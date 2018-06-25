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
	"fmt"
	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"time"
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
	credentials        *credentials.Credentials
	tubeAuthWindowSecs int64

	handlers          *request.Handlers
	pool              *tubePool
	keySchema         *lru.Lru
	attrNamesListToId *lru.Lru
	attrListIdToNames *lru.Lru
}

func NewSingleClient(endpoint, region string, credentials *credentials.Credentials) (*SingleDaxClient, error) {
	return newSingleClientWithOptions(endpoint, region, credentials, -1)

}

func newSingleClientWithOptions(endpoint, region string, credentials *credentials.Credentials, maxPendingConnections int) (*SingleDaxClient, error) {
	po := defaultTubePoolOptions
	if maxPendingConnections > 0 {
		po.maxConcurrentConnAttempts = maxPendingConnections
	}
	client := &SingleDaxClient{
		region:             region,
		credentials:        credentials,
		tubeAuthWindowSecs: authTtlSecs * tubeAuthWindowScalar,
		pool:               newTubePoolWithOptions(endpoint, po),
	}

	client.handlers = client.buildHandlers()
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
	if client.pool != nil {
		return client.pool.Close()
	}
	return nil
}

func (client *SingleDaxClient) endpoints(opt RequestOptions) ([]serviceEndpoint, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeEndpointsInput(writer)
	}
	var out []serviceEndpoint
	var err error
	decoder := func(reader *cbor.Reader) error {
		out, err = decodeEndpointsOutput(reader)
		return err
	}
	if err = client.executeWithRetries(opEndpoints, opt, encoder, decoder); err != nil {
		return nil, err
	}
	return out, nil
}

func (client *SingleDaxClient) defineAttributeListId(ctx aws.Context, attrNames []string) (int64, error) {
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
	opt := RequestOptions{Context: ctx}
	if err = client.executeWithRetries(opDefineAttributeListId, opt, encoder, decoder); err != nil {
		return 0, err
	}
	return out, nil
}

func (client *SingleDaxClient) defineAttributeList(ctx aws.Context, id int64) ([]string, error) {
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
	opt := RequestOptions{Context: ctx}
	if err = client.executeWithRetries(opDefineAttributeList, opt, encoder, decoder); err != nil {
		return nil, err
	}
	return out, nil
}

func (client *SingleDaxClient) defineKeySchema(ctx aws.Context, table string) ([]dynamodb.AttributeDefinition, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeDefineKeySchemaInput(table, writer)
	}
	var out []dynamodb.AttributeDefinition
	var err error
	decoder := func(reader *cbor.Reader) error {
		out, err = decodeDefineKeySchemaOutput(reader)
		return err
	}
	opt := RequestOptions{Context: ctx}
	if err = client.executeWithRetries(opDefineKeySchema, opt, encoder, decoder); err != nil {
		return nil, err
	}
	return out, nil
}

func (client *SingleDaxClient) PutItemWithOptions(input *dynamodb.PutItemInput, output *dynamodb.PutItemOutput, opt RequestOptions) (*dynamodb.PutItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodePutItemInput(opt.Context, input, client.keySchema, client.attrNamesListToId, writer)
	}
	var err error
	decoder := func(reader *cbor.Reader) error {
		output, err = decodePutItemOutput(opt.Context, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err = client.executeWithRetries(OpPutItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) DeleteItemWithOptions(input *dynamodb.DeleteItemInput, output *dynamodb.DeleteItemOutput, opt RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeDeleteItemInput(opt.Context, input, client.keySchema, writer)
	}
	var err error
	decoder := func(reader *cbor.Reader) error {
		output, err = decodeDeleteItemOutput(opt.Context, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err = client.executeWithRetries(OpDeleteItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) UpdateItemWithOptions(input *dynamodb.UpdateItemInput, output *dynamodb.UpdateItemOutput, opt RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeUpdateItemInput(opt.Context, input, client.keySchema, writer)
	}
	var err error
	decoder := func(reader *cbor.Reader) error {
		output, err = decodeUpdateItemOutput(opt.Context, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err = client.executeWithRetries(OpUpdateItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) GetItemWithOptions(input *dynamodb.GetItemInput, output *dynamodb.GetItemOutput, opt RequestOptions) (*dynamodb.GetItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeGetItemInput(opt.Context, input, client.keySchema, writer)
	}
	var err error
	decoder := func(reader *cbor.Reader) error {
		output, err = decodeGetItemOutput(opt.Context, reader, input, client.attrListIdToNames, output)
		return err
	}
	if err = client.executeWithRetries(OpGetItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) ScanWithOptions(input *dynamodb.ScanInput, output *dynamodb.ScanOutput, opt RequestOptions) (*dynamodb.ScanOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeScanInput(opt.Context, input, client.keySchema, writer)
	}
	var err error
	decoder := func(reader *cbor.Reader) error {
		output, err = decodeScanOutput(opt.Context, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err = client.executeWithRetries(OpScan, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) QueryWithOptions(input *dynamodb.QueryInput, output *dynamodb.QueryOutput, opt RequestOptions) (*dynamodb.QueryOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeQueryInput(opt.Context, input, client.keySchema, writer)
	}
	var err error
	decoder := func(reader *cbor.Reader) error {
		output, err = decodeQueryOutput(opt.Context, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err = client.executeWithRetries(OpQuery, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) BatchWriteItemWithOptions(input *dynamodb.BatchWriteItemInput, output *dynamodb.BatchWriteItemOutput, opt RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeBatchWriteItemInput(opt.Context, input, client.keySchema, client.attrNamesListToId, writer)
	}
	var err error
	decoder := func(reader *cbor.Reader) error {
		output, err = decodeBatchWriteItemOutput(opt.Context, reader, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err = client.executeWithRetries(OpBatchWriteItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) BatchGetItemWithOptions(input *dynamodb.BatchGetItemInput, output *dynamodb.BatchGetItemOutput, opt RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	encoder := func(writer *cbor.Writer) error {
		return encodeBatchGetItemInput(opt.Context, input, client.keySchema, writer)
	}
	var err error
	decoder := func(reader *cbor.Reader) error {
		output, err = decodeBatchGetItemOutput(opt.Context, reader, input, client.keySchema, client.attrListIdToNames, output)
		return err
	}
	if err = client.executeWithRetries(OpBatchGetItem, opt, encoder, decoder); err != nil {
		return output, err
	}
	return output, nil
}

func (client *SingleDaxClient) NewDaxRequest(op *request.Operation, input, output interface{}, opt RequestOptions) *request.Request {
	req := request.New(aws.Config{}, clientInfo, *client.handlers, nil, op, input, output)
	opt.applyTo(req)
	return req
}

func (client *SingleDaxClient) buildHandlers() *request.Handlers {
	h := &request.Handlers{}
	h.Build.PushFrontNamed(request.NamedHandler{Name: "dax.BuildHandler", Fn: client.build})
	h.Send.PushFrontNamed(request.NamedHandler{Name: "dax.SendHandler", Fn: client.send})
	return h
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
	default:
		req.Error = awserr.New(request.InvalidParameterErrCode, "unknown op "+req.Operation.Name, nil)
		return
	}
	req.SetBufferBody(buf.Bytes())
}

func (client *SingleDaxClient) send(req *request.Request) {
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
		req.Data, req.Error = client.GetItemWithOptions(input, output, opt)
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
		req.Data, req.Error = client.ScanWithOptions(input, output, opt)
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
		req.Data, req.Error = client.QueryWithOptions(input, output, opt)
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
		req.Data, req.Error = client.BatchGetItemWithOptions(input, output, opt)
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
		req.Data, req.Error = client.PutItemWithOptions(input, output, opt)
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
		req.Data, req.Error = client.DeleteItemWithOptions(input, output, opt)
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
		req.Data, req.Error = client.UpdateItemWithOptions(input, output, opt)
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
		req.Data, req.Error = client.BatchWriteItemWithOptions(input, output, opt)
	default:
		req.Error = awserr.New(request.InvalidParameterErrCode, "unknown op "+req.Operation.Name, nil)
		return
	}
}

func (client *SingleDaxClient) newContext(o RequestOptions) aws.Context {
	if o.Context != nil {
		return o.Context
	}
	return aws.BackgroundContext()
}

func (client *SingleDaxClient) executeWithRetries(op string, o RequestOptions, encoder func(writer *cbor.Writer) error, decoder func(reader *cbor.Reader) error) error {
	ctx := client.newContext(o)
	attempts := o.MaxRetries + 1

	var err error
	for i := 0; i < attempts; i++ {
		if i > 0 && o.Logger != nil && o.LogLevel.Matches(aws.LogDebugWithRequestRetries) {
			o.Logger.Log(fmt.Sprintf("DEBUG: Retrying Request %s/%s, attempt %d", service, op, i+1))
		}
		if err = client.executeWithContext(ctx, op, encoder, decoder); err == nil {
			return nil
		} else if ctx != nil && err == ctx.Err() {
			return awserr.New(request.CanceledErrorCode, "request context canceled", err)
		}
		d := o.RetryDelay
		if d > 0 {
			if s := o.SleepDelayFn; s != nil {
				s(d)
			} else if err = aws.SleepWithContext(ctx, d); err != nil {
				return awserr.New(request.CanceledErrorCode, "request context canceled", err)
			}
		}
	}
	if err != nil {
		return translateError(err)
	} else {
		return awserr.New("UnknownError", "ran out of retries", nil)
	}
}

func (client *SingleDaxClient) executeWithContext(ctx aws.Context, op string, encoder func(writer *cbor.Writer) error, decoder func(reader *cbor.Reader) error) error {
	tube, err := client.pool.getWithContext(ctx, client.isHighPriority(op))
	if err != nil {
		return err
	}
	if err = client.pool.setDeadline(tube, ctx); err != nil {
		return err
	}

	if err = client.auth(tube); err != nil {
		client.recycleTube(tube, err)
		return err
	}

	writer := tube.cborWriter
	if err = encoder(tube.cborWriter); err != nil {
		client.recycleTube(tube, err)
		return err
	}
	if err := writer.Flush(); err != nil {
		client.recycleTube(tube, err)
		return err
	}

	reader := tube.cborReader
	ex, err := decodeError(reader)
	if err != nil { // decode or network error
		client.recycleTube(tube, err)
		return err
	}
	if ex != nil { // user or server error
		client.recycleTube(tube, nil) // do not close conn
		return ex
	}

	err = decoder(reader)
	client.recycleTube(tube, err)
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

func (client *SingleDaxClient) recycleTube(tube *tube, err error) {
	if tube == nil {
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
			tube.setAuthExpiryUnix(time.Now().Unix())
		}
	}
	if recycle {
		client.pool.put(tube)
	} else {
		client.pool.discard(tube)
	}
}
func (client *SingleDaxClient) auth(tube *tube) error {
	// TODO credentials.Get() cause a throughput drop of ~25 with 250 goroutines with DefaultCredentialChain (only instance profile credentials available)
	creds, err := client.credentials.Get()
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	if tube.compareAndSwapAuthId(creds.AccessKeyID) || tube.getAuthExpiryUnix() <= now.Unix() {
		stringToSign, signature := generateSigV4WithTime(creds, daxAddress, client.region, "", now)
		writer := tube.cborWriter
		if err := encodeAuthInput(creds.AccessKeyID, creds.SessionToken, stringToSign, signature, userAgent, writer); err != nil {
			return err
		}
		if err := writer.Flush(); err != nil {
			return err
		}
		tube.setAuthExpiryUnix(now.Unix() + client.tubeAuthWindowSecs)
	}
	return nil
}

func (client *SingleDaxClient) reapIdleConnections() {
	client.pool.reapIdleConnections()
}
