/*
  Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

  Licensed under the Apache License, Version 2.0 (the "License").
  You may not use this file except in compliance with the License.
  A copy of the License is located at

      http://www.apache.org/licenses/LICENSE-2.0

  or in the "license" file accompanying this file. This file is distributed
  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
  express or implied. See the License for the specific language governing
  permissions and limitations under the License.
*/

package client

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/aws/smithy-go"

	aws2 "github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
)

const (
	userAgent  = "DaxGoClient-1.0.0"
	daxAddress = "https://dax.amazonaws.com"

	authTtlSecs          = 5 * 60
	tubeAuthWindowScalar = 0.75

	emptyAttributeListId = 1
)

const (
	serviceName             = "dax"
	opDefineAttributeList   = "DefineAttributeList"
	opDefineAttributeListId = "DefineAttributeListId"
	opDefineKeySchema       = "DefineKeySchema"
	opEndpoints             = "Endpoints"
	OpGetItem               = "GetItem"
	OpPutItem               = "PutItem"
	OpUpdateItem            = "UpdateItem"
	OpDeleteItem            = "DeleteItem"
	OpBatchGetItem          = "BatchGetItem"
	OpBatchWriteItem        = "BatchWriteItem"
	OpTransactGetItems      = "TransactGetItems"
	OpTransactWriteItems    = "TransactWriteItems"
	OpQuery                 = "Query"
	OpScan                  = "Scan"
)

var clientInfo = metadata.ClientInfo{ServiceName: serviceName}

const (
	keySchemaLruCacheSize     = 100
	attributeListLruCacheSize = 1000
)

type SingleDaxClient struct {
	region             string
	credentials        aws2.CredentialsProvider
	tubeAuthWindowSecs int64
	executor           *taskExecutor

	pool              *tubePool
	keySchema         *lru.Lru
	attrNamesListToId *lru.Lru
	attrListIdToNames *lru.Lru
}

var _ DaxAPI = (*SingleDaxClient)(nil)

func NewSingleClient(endpoint string, connConfigData connConfig, region string, credentials aws2.CredentialsProvider) (*SingleDaxClient, error) {
	return newSingleClientWithOptions(endpoint, connConfigData, region, credentials, -1, defaultDialer.DialContext)
}

func newSingleClientWithOptions(endpoint string, connConfigData connConfig, region string, credentials aws2.CredentialsProvider, maxPendingConnections int, dialContextFn dialContext) (*SingleDaxClient, error) {
	po := defaultTubePoolOptions
	if maxPendingConnections > 0 {
		po.maxConcurrentConnAttempts = maxPendingConnections
	}

	po.dialContext = dialContextFn

	client := &SingleDaxClient{
		region:             region,
		credentials:        credentials,
		tubeAuthWindowSecs: authTtlSecs * tubeAuthWindowScalar,
		pool:               newTubePoolWithOptions(endpoint, po, connConfigData),
		executor:           newExecutor(),
	}

	client.keySchema = &lru.Lru{
		MaxEntries: keySchemaLruCacheSize,
		LoadFunc: func(ctx aws.Context, key lru.Key) (interface{}, error) {
			table, ok := key.(string)
			if !ok {
				return nil, awserr.New(request.ErrCodeSerialization, "unexpected type for table name", nil)
			}
			if ctx == nil {
				ctx = aws.BackgroundContext()
			}
			return client.defineKeySchema(ctx, table)
		},
	}

	client.attrNamesListToId = &lru.Lru{
		MaxEntries: attributeListLruCacheSize,
		LoadFunc: func(ctx aws.Context, key lru.Key) (interface{}, error) {
			attrNames, ok := key.([]string)
			if !ok {
				return nil, awserr.New(request.ErrCodeSerialization, "unexpected type for attribute list", nil)
			}
			if ctx == nil {
				ctx = aws.BackgroundContext()
			}
			return client.defineAttributeListId(ctx, attrNames)
		},
		KeyMarshaller: func(key lru.Key) lru.Key {
			var buf bytes.Buffer
			w := cbor.NewWriter(&buf)
			defer w.Close()
			for _, v := range key.([]string) {
				w.WriteString(v)
			}
			w.Flush()
			return string(buf.Bytes())
		},
	}

	client.attrListIdToNames = &lru.Lru{
		MaxEntries: attributeListLruCacheSize,
		LoadFunc: func(ctx aws.Context, key lru.Key) (interface{}, error) {
			id, ok := key.(int64)
			if !ok {
				return nil, awserr.New(request.ErrCodeSerialization, "unexpected type for attribute list id", nil)
			}
			if ctx == nil {
				ctx = aws.BackgroundContext()
			}
			return client.defineAttributeList(ctx, id)
		},
	}

	return client, nil
}

func (client *SingleDaxClient) Close() error {
	client.executor.stopAll()
	if client.pool != nil {
		return client.pool.Close()
	}
	return nil
}

func (client *SingleDaxClient) startHealthChecks(ctx context.Context, cc *cluster, host hostPort) {
	cc.debugLog("Starting health checks for :: " + host.host)
	client.executor.start(cc.config.ClientHealthCheckInterval, func() error {
		ctx, cfn := context.WithTimeout(ctx, 1*time.Second)
		defer cfn()
		var err error
		opts := RequestOptions{}
		opts.RetryMaxAttempts = 2
		_, err = client.endpoints(ctx, opts)
		if err != nil {
			cc.debugLog("Health checks failed with error " + err.Error() + " for host :: " + host.host)
			cc.onHealthCheckFailed(host)
		} else {
			cc.debugLog("Health checks succeeded for host:: " + host.host)
		}
		return nil
	})
}

func (client *SingleDaxClient) endpoints(ctx context.Context, opt RequestOptions) ([]serviceEndpoint, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeEndpointsInput(writer)
	}
	var out []serviceEndpoint
	var err error
	decoder := func(reader *cbor.Reader) error {
		out, err = decodeEndpointsOutput(reader)
		return err
	}
	if err = client.executeWithRetries(ctx, opEndpoints, opt, encoder, decoder); err != nil {
		return nil, err
	}
	return out, nil
}

func (client *SingleDaxClient) defineAttributeListId(ctx context.Context, attrNames []string) (int64, error) {
	if len(attrNames) == 0 {
		return emptyAttributeListId, nil
	}
	encoder := func(writer *cbor.Writer) error {
		return encodeDefineAttributeListIdInput(attrNames, writer)
	}
	var out int64
	var err error
	decoder := func(reader *cbor.Reader) error {
		out, err = decodeDefineAttributeListIdOutput(reader)
		return err
	}
	opt := RequestOptions{}
	if err = client.executeWithRetries(ctx, opDefineAttributeListId, opt, encoder, decoder); err != nil {
		return 0, err
	}
	return out, nil
}

func (client *SingleDaxClient) defineAttributeList(ctx context.Context, id int64) ([]string, error) {
	if id == emptyAttributeListId {
		return []string{}, nil
	}
	encoder := func(writer *cbor.Writer) error {
		return encodeDefineAttributeListInput(id, writer)
	}
	var out []string
	var err error
	decoder := func(reader *cbor.Reader) error {
		out, err = decodeDefineAttributeListOutput(reader)
		return err
	}
	opt := RequestOptions{}
	if err = client.executeWithRetries(ctx, opDefineAttributeList, opt, encoder, decoder); err != nil {
		return nil, err
	}
	return out, nil
}

func (client *SingleDaxClient) defineKeySchema(ctx context.Context, table string) ([]types.AttributeDefinition, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeDefineKeySchemaInput(table, writer)
	}
	var out []types.AttributeDefinition
	var err error
	decoder := func(reader *cbor.Reader) error {
		out, err = decodeDefineKeySchemaOutput(reader)
		return err
	}
	opt := RequestOptions{}
	if err = client.executeWithRetries(ctx, opDefineKeySchema, opt, encoder, decoder); err != nil {
		return nil, err
	}
	return out, nil
}

func (client *SingleDaxClient) PutItemWithOptions(ctx context.Context, input *dynamodb.PutItemInput, opt RequestOptions) (*dynamodb.PutItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodePutItemInput(ctx, input, client.keySchema, client.attrNamesListToId, writer)
	}
	var output *dynamodb.PutItemOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodePutItemOutput(ctx, reader, input, client.keySchema, client.attrListIdToNames)
		return err
	}
	if err := client.executeWithRetries(ctx, OpPutItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) DeleteItemWithOptions(ctx context.Context, input *dynamodb.DeleteItemInput, opt RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeDeleteItemInput(ctx, input, client.keySchema, writer)
	}
	var output *dynamodb.DeleteItemOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeDeleteItemOutput(ctx, reader, input, client.keySchema, client.attrListIdToNames)
		return err
	}
	if err := client.executeWithRetries(ctx, OpDeleteItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) UpdateItemWithOptions(ctx context.Context, input *dynamodb.UpdateItemInput, opt RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeUpdateItemInput(ctx, input, client.keySchema, writer)
	}
	var output *dynamodb.UpdateItemOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeUpdateItemOutput(ctx, reader, input, client.keySchema, client.attrListIdToNames)
		return err
	}
	if err := client.executeWithRetries(ctx, OpUpdateItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) GetItemWithOptions(ctx context.Context, input *dynamodb.GetItemInput, opt RequestOptions) (*dynamodb.GetItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeGetItemInput(ctx, input, client.keySchema, writer)
	}
	var output *dynamodb.GetItemOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeGetItemOutput(ctx, reader, input, client.attrListIdToNames)
		return err
	}
	if err := client.executeWithRetries(ctx, OpGetItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) ScanWithOptions(ctx context.Context, input *dynamodb.ScanInput, opt RequestOptions) (*dynamodb.ScanOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeScanInput(ctx, input, client.keySchema, writer)
	}
	var output *dynamodb.ScanOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeScanOutput(ctx, reader, input, client.keySchema, client.attrListIdToNames)
		return err
	}
	if err := client.executeWithRetries(ctx, OpScan, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) QueryWithOptions(ctx context.Context, input *dynamodb.QueryInput, opt RequestOptions) (*dynamodb.QueryOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeQueryInput(ctx, input, client.keySchema, writer)
	}
	var output *dynamodb.QueryOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeQueryOutput(ctx, reader, input, client.keySchema, client.attrListIdToNames)
		return err
	}
	if err := client.executeWithRetries(ctx, OpQuery, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) BatchWriteItemWithOptions(ctx context.Context, input *dynamodb.BatchWriteItemInput, opt RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeBatchWriteItemInput(ctx, input, client.keySchema, client.attrNamesListToId, writer)
	}
	var output *dynamodb.BatchWriteItemOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeBatchWriteItemOutput(ctx, reader, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err := client.executeWithRetries(ctx, OpBatchWriteItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) BatchGetItemWithOptions(ctx context.Context, input *dynamodb.BatchGetItemInput, opt RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeBatchGetItemInput(ctx, input, client.keySchema, writer)
	}
	var output *dynamodb.BatchGetItemOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeBatchGetItemOutput(ctx, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err := client.executeWithRetries(ctx, OpBatchGetItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) TransactWriteItemsWithOptions(ctx context.Context, input *dynamodb.TransactWriteItemsInput, opt RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	extractedKeys := make([]map[string]types.AttributeValue, len(input.TransactItems))
	encoder := func(writer *cbor.Writer) error {
		return encodeTransactWriteItemsInput(ctx, input, client.keySchema, client.attrNamesListToId, writer, extractedKeys)
	}
	var output *dynamodb.TransactWriteItemsOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeTransactWriteItemsOutput(ctx, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err := client.executeWithRetries(ctx, OpBatchWriteItem, opt, encoder, decoder); err != nil {
		if failure, ok := err.(*daxTransactionCanceledFailure); ok {
			var cancellationReasons []types.CancellationReason
			if cancellationReasons, err = decodeTransactionCancellationReasons(ctx, failure, extractedKeys, client.attrListIdToNames); err != nil {
				return output, err
			}
			failure.cancellationReasons = cancellationReasons
			return output, failure
		}
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) TransactGetItemsWithOptions(ctx context.Context, input *dynamodb.TransactGetItemsInput, opt RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	extractedKeys := make([]map[string]types.AttributeValue, len(input.TransactItems))
	encoder := func(writer *cbor.Writer) error {
		return encodeTransactGetItemsInput(ctx, input, client.keySchema, writer, extractedKeys)
	}
	var output *dynamodb.TransactGetItemsOutput
	decoder := func(reader *cbor.Reader) error {
		var err error
		output, err = decodeTransactGetItemsOutput(ctx, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err := client.executeWithRetries(ctx, OpBatchWriteItem, opt, encoder, decoder); err != nil {
		if failure, ok := err.(*daxTransactionCanceledFailure); ok {
			var cancellationReasons []types.CancellationReason
			if cancellationReasons, err = decodeTransactionCancellationReasons(ctx, failure, extractedKeys, client.attrListIdToNames); err != nil {
				return output, err
			}
			failure.cancellationReasons = cancellationReasons
			return output, failure
		}
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) build(req *request.Request) {
	var buf bytes.Buffer
	w := cbor.NewWriter(&buf)
	defer w.Close()
	switch req.Operation.Name {
	case OpGetItem:
		input, ok := req.Params.(*dynamodb.GetItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *GetItemInput", nil)
			return
		}
		if err := encodeGetItemInput(req.Context(), input, client.keySchema, w); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpScan:
		input, ok := req.Params.(*dynamodb.ScanInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *ScanInput", nil)
			return
		}
		if err := encodeScanInput(req.Context(), input, client.keySchema, w); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpQuery:
		input, ok := req.Params.(*dynamodb.QueryInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *QueryInput", nil)
			return
		}
		if err := encodeQueryInput(req.Context(), input, client.keySchema, w); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpBatchGetItem:
		input, ok := req.Params.(*dynamodb.BatchGetItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *BatchGetItemInput", nil)
			return
		}
		if err := encodeBatchGetItemInput(req.Context(), input, client.keySchema, w); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpPutItem:
		input, ok := req.Params.(*dynamodb.PutItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *PutItemInput", nil)
			return
		}
		if err := encodePutItemInput(req.Context(), input, client.keySchema, client.attrNamesListToId, w); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpDeleteItem:
		input, ok := req.Params.(*dynamodb.DeleteItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *DeleteItemInput", nil)
			return
		}
		if err := encodeDeleteItemInput(req.Context(), input, client.keySchema, w); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpUpdateItem:
		input, ok := req.Params.(*dynamodb.UpdateItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *UpdateItemInput", nil)
			return
		}
		if err := encodeUpdateItemInput(req.Context(), input, client.keySchema, w); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpBatchWriteItem:
		input, ok := req.Params.(*dynamodb.BatchWriteItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *BatchWriteItemInput", nil)
			return
		}
		if err := encodeBatchWriteItemInput(req.Context(), input, client.keySchema, client.attrNamesListToId, w); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpTransactGetItems:
		input, ok := req.Params.(*dynamodb.TransactGetItemsInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *TransactGetItemsInput", nil)
			return
		}
		extractedKeys := make([]map[string]types.AttributeValue, len(input.TransactItems))
		if err := encodeTransactGetItemsInput(req.Context(), input, client.keySchema, w, extractedKeys); err != nil {
			req.Error = translateError(err)
			return
		}
	case OpTransactWriteItems:
		input, ok := req.Params.(*dynamodb.TransactWriteItemsInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *TransactWriteItemsInput", nil)
			return
		}
		extractedKeys := make([]map[string]types.AttributeValue, len(input.TransactItems))
		if err := encodeTransactWriteItemsInput(req.Context(), input, client.keySchema, client.attrNamesListToId, w, extractedKeys); err != nil {
			req.Error = translateError(err)
			return
		}
	default:
		req.Error = awserr.New(request.InvalidParameterErrCode, "unknown op "+req.Operation.Name, nil)
		return
	}
	req.SetBufferBody(buf.Bytes())
}

func (client *SingleDaxClient) send(ctx context.Context, req *request.Request) {
	opt := RequestOptions{}
	if err := opt.mergeFromRequest(req, true); err != nil {
		req.Error = err
		return
	}
	switch req.Operation.Name {
	case OpGetItem:
		input, ok := req.Params.(*dynamodb.GetItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *GetItemInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.GetItemOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *GetItemOutput", nil)
			return
		}
		req.Data, req.Error = client.GetItemWithOptions(ctx, input, opt)
	case OpScan:
		input, ok := req.Params.(*dynamodb.ScanInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *ScanInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.ScanOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *ScanOutput", nil)
			return
		}
		req.Data, req.Error = client.ScanWithOptions(ctx, input, opt)
	case OpQuery:
		input, ok := req.Params.(*dynamodb.QueryInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *QueryInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.QueryOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *QueryOutput", nil)
			return
		}
		req.Data, req.Error = client.QueryWithOptions(ctx, input, opt)
	case OpBatchGetItem:
		input, ok := req.Params.(*dynamodb.BatchGetItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *BatchGetItemInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.BatchGetItemOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *BatchGetItemOutput", nil)
			return
		}
		req.Data, req.Error = client.BatchGetItemWithOptions(ctx, input, opt)
	case OpPutItem:
		input, ok := req.Params.(*dynamodb.PutItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *PutItemInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.PutItemOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *PutItemOutput", nil)
			return
		}
		req.Data, req.Error = client.PutItemWithOptions(ctx, input, opt)
	case OpDeleteItem:
		input, ok := req.Params.(*dynamodb.DeleteItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *DeleteItemInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.DeleteItemOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *DeleteItemOutput", nil)
			return
		}
		req.Data, req.Error = client.DeleteItemWithOptions(ctx, input, opt)
	case OpUpdateItem:
		input, ok := req.Params.(*dynamodb.UpdateItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *UpdateItemInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.UpdateItemOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *UpdateItemOutput", nil)
			return
		}
		req.Data, req.Error = client.UpdateItemWithOptions(ctx, input, opt)
	case OpBatchWriteItem:
		input, ok := req.Params.(*dynamodb.BatchWriteItemInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *BatchWriteItemInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.BatchWriteItemOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *BatchWriteItemOutput", nil)
			return
		}
		req.Data, req.Error = client.BatchWriteItemWithOptions(ctx, input, opt)
	case OpTransactGetItems:
		input, ok := req.Params.(*dynamodb.TransactGetItemsInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *TransactGetItemsInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.TransactGetItemsOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *TransactGetItemsOutput", nil)
			return
		}
		req.Data, req.Error = client.TransactGetItemsWithOptions(ctx, input, opt)
	case OpTransactWriteItems:
		input, ok := req.Params.(*dynamodb.TransactWriteItemsInput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *TransactWriteItemsInput", nil)
			return
		}
		output, ok := req.Data.(*dynamodb.TransactWriteItemsOutput)
		if !ok {
			req.Error = awserr.New(request.ErrCodeSerialization, "expected *TransactWriteItemsOutput", nil)
			return
		}
		req.Data, req.Error = client.TransactWriteItemsWithOptions(ctx, input, opt)
	default:
		req.Error = awserr.New(request.InvalidParameterErrCode, "unknown op "+req.Operation.Name, nil)
		return
	}
}

func (client *SingleDaxClient) executeWithRetries(ctx context.Context, op string, o RequestOptions, encoder func(writer *cbor.Writer) error, decoder func(reader *cbor.Reader) error) error {

	var err error
	attempts := o.RetryMaxAttempts
	// Start from 0 to accommodate for the initial request
	for i := 0; i <= attempts; i++ {
		if i > 0 && o.Logger != nil {
			o.Logger.Logf(ClassificationDebug, "Retrying Request %s/%s, attempt %d", service, op, i)
		}

		err = client.executeWithContext(ctx, op, encoder, decoder, o)
		if err == nil {
			return nil
		}
		if errors.Is(err, context.Canceled) {
			return &smithy.CanceledError{Err: err}
		}

		if !isRetryable(o, i+1, err) {
			return &smithy.OperationError{Err: err, OperationName: op}
		}

		d, err := o.Retryer.RetryDelay(i+1, err)
		if err != nil {
			return &smithy.OperationError{Err: err, OperationName: op}
		}
		if err = Sleep(ctx, op, d); err != nil {
			return err
		}

		if o.Logger != nil {
			o.Logger.Logf(ClassificationDebug, "Error in executing %s%s : %s", service, op, err)
		}
	}
	// Return the last error occurred
	return translateError(err)
}

func (client *SingleDaxClient) executeWithContext(ctx context.Context, op string, encoder func(writer *cbor.Writer) error, decoder func(reader *cbor.Reader) error, opt RequestOptions) error {
	t, err := client.pool.getWithContext(ctx, client.isHighPriority(op), opt)
	if err != nil {
		return err
	}
	if err = client.pool.setDeadline(ctx, t); err != nil {
		client.pool.discard(t)
		return err
	}

	if err = client.auth(ctx, t); err != nil {
		client.pool.discard(t)
		return err
	}

	writer := t.CborWriter()
	if err = encoder(writer); err != nil {
		// Validation errors will cause pool to be discarded as there is no guarantee
		// that the validation was performed before any data was written into tube
		client.pool.discard(t)
		return err
	}
	if err := writer.Flush(); err != nil {
		client.pool.discard(t)
		return err
	}

	reader := t.CborReader()
	ex, err := decodeError(reader)
	if err != nil { // decode or network error
		client.pool.discard(t)
		return err
	}
	if ex != nil { // user or server error
		client.recycleTube(t, ex)
		return ex
	}

	err = decoder(reader)
	if err != nil {
		client.pool.discard(t)
	} else {
		client.pool.put(t)
	}
	return err
}

func (client *SingleDaxClient) isHighPriority(op string) bool {
	switch op {
	case opDefineAttributeListId, opDefineAttributeList, opDefineKeySchema:
		return true
	default:
		return false
	}
}

func (client *SingleDaxClient) recycleTube(t tube, err error) {
	if t == nil {
		return
	}

	var recycle bool
	if err == nil {
		recycle = true
	} else {
		// IO streams are guaranteed to be completely drained only on daxRequestException
		d, ok := err.(*daxRequestFailure)
		recycle = ok
		if ok && d.authError() {
			t.SetAuthExpiryUnix(time.Now().Unix())
		}
	}
	if recycle {
		client.pool.put(t)
	} else {
		client.pool.discard(t)
	}
}
func (client *SingleDaxClient) auth(ctx context.Context, t tube) error {
	// TODO credentials.Get() cause a throughput drop of ~25 with 250 goroutines with DefaultCredentialChain (only instance profile credentials available)
	creds, err := client.credentials.Retrieve(ctx)
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	if t.CompareAndSwapAuthID(creds.AccessKeyID) || t.AuthExpiryUnix() <= now.Unix() {
		stringToSign, signature := generateSigV4WithTime(creds, daxAddress, client.region, "", now)
		writer := t.CborWriter()
		if err := encodeAuthInput(creds.AccessKeyID, creds.SessionToken, stringToSign, signature, userAgent, writer); err != nil {
			return err
		}
		if err := writer.Flush(); err != nil {
			return err
		}
		t.SetAuthExpiryUnix(now.Unix() + client.tubeAuthWindowSecs)
	}
	return nil
}

func (client *SingleDaxClient) reapIdleConnections() {
	client.pool.reapIdleConnections()
}

type HealthCheckDaxAPI interface {
	startHealthChecks(cc *cluster, host hostPort)
}
