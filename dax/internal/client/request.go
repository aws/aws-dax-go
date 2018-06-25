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
	"github.com/aws/aws-dax-go/dax/internal/parser"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"sort"
	"strings"
)

const daxServiceId = 1

const (
	// Dax Control
	authorizeConnection_1489122155_1_Id    = 1489122155
	defineAttributeList_670678385_1_Id     = 670678385
	defineAttributeListId_N1230579644_1_Id = -1230579644
	defineKeySchema_N742646399_1_Id        = -742646399
	endpoints_455855874_1_Id               = 455855874
	methods_785068263_1_Id                 = 785068263
	services_N1016793520_1_Id              = -1016793520

	// DynamoDB Data
	batchGetItem_N697851100_1_Id  = -697851100
	batchWriteItem_116217951_1_Id = 116217951
	getItem_263244906_1_Id        = 263244906
	putItem_N2106490455_1_Id      = -2106490455
	deleteItem_1013539361_1_Id    = 1013539361
	updateItem_1425579023_1_Id    = 1425579023
	query_N931250863_1_Id         = -931250863
	scan_N1875390620_1_Id         = -1875390620

	// DynamoDB Control
	createTable_N313431286_1_Id    = -313431286
	deleteTable_2120496185_1_Id    = 2120496185
	describeTable_N819330193_1_Id  = -819330193
	updateTable_383747477_1_Id     = 383747477
	listTables_1874119219_1_Id     = 1874119219
	describeLimits_N475661135_1_Id = -475661135
)

const (
	requestParamProjectionExpression = iota
	requestParamExpressionAttributeNames
	requestParamConsistentRead
	requestParamReturnConsumedCapacity
	requestParamConditionExpression
	requestParamExpressionAttributeValues
	requestParamReturnItemCollectionMetrics
	requestParamReturnValues
	requestParamUpdateExpression
	requestParamExclusiveStartKey
	requestParamFilterExpression
	requestParamIndexName
	requestParamKeyConditionExpression
	requestParamLimit
	requestParamScanIndexForward
	requestParamSelect
	requestParamSegment
	requestParamTotalSegments
	requestParamRequestItems
)

const (
	returnValueNone = 1 + iota
	returnValueAllOld
	returnValueUpdatedOld
	returnValueAllNew
	returnValueUpdatedNew
)

const (
	returnConsumedCapacityNone = iota
	returnConsumedCapacityTotal
	returnConsumedCapacityIndexes
)

const (
	returnItemCollectionMetricsNone = iota
	returnItemCollectionMetricsSize
)

const (
	selectAllAttributes = 1 + iota
	selectAllProjectedAttributes
	selectCount
	selectSpecificAttributes
)

const maxWriteBatchSize = 25

func encodeEndpointsInput(writer *cbor.Writer) error {
	if err := encodeServiceAndMethod(endpoints_455855874_1_Id, writer); err != nil {
		return err
	}
	return nil
}

func encodeAuthInput(accessKey, sessionToken, stringToSign, signature, userAgent string, writer *cbor.Writer) error {
	if err := encodeServiceAndMethod(authorizeConnection_1489122155_1_Id, writer); err != nil {
		return err
	}
	if err := writer.WriteString(accessKey); err != nil {
		return err
	}
	if err := writer.WriteString(signature); err != nil {
		return err
	}
	if err := writer.WriteBytes([]byte(stringToSign)); err != nil {
		return err
	}
	if len(sessionToken) == 0 {
		if err := writer.WriteNull(); err != nil {
			return err
		}
	} else {
		if err := writer.WriteString(sessionToken); err != nil {
			return err
		}
	}
	if len(userAgent) == 0 {
		if err := writer.WriteNull(); err != nil {
			return err
		}
	} else {
		if err := writer.WriteString(userAgent); err != nil {
			return err
		}
	}
	return nil
}

func encodeDefineAttributeListIdInput(attrNames []string, writer *cbor.Writer) error {
	if err := encodeServiceAndMethod(defineAttributeListId_N1230579644_1_Id, writer); err != nil {
		return err
	}
	if err := writer.WriteArrayHeader(len(attrNames)); err != nil {
		return err
	}
	for _, an := range attrNames {
		if err := writer.WriteString(an); err != nil {
			return err
		}
	}
	return nil
}

func encodeDefineAttributeListInput(id int64, writer *cbor.Writer) error {
	if err := encodeServiceAndMethod(defineAttributeList_670678385_1_Id, writer); err != nil {
		return err
	}
	return writer.WriteInt64(id)
}

func encodeDefineKeySchemaInput(table string, writer *cbor.Writer) error {
	if err := encodeServiceAndMethod(defineKeySchema_N742646399_1_Id, writer); err != nil {
		return err
	}
	return writer.WriteBytes([]byte(table))
}

func encodePutItemInput(ctx aws.Context, input *dynamodb.PutItemInput, keySchema *lru.Lru, attrNamesListToId *lru.Lru, writer *cbor.Writer) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, fmt.Sprintf("input cannot be nil"), nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}
	if input, err = translateLegacyPutItemInput(input); err != nil {
		return err
	}
	table := *input.TableName
	keys, err := getKeySchema(ctx, keySchema, table)
	if err != nil {
		return err
	}

	if err := encodeServiceAndMethod(putItem_N2106490455_1_Id, writer); err != nil {
		return err
	}
	if err := writer.WriteBytes([]byte(table)); err != nil {
		return err
	}

	if err := cbor.EncodeItemKey(input.Item, keys, writer); err != nil {
		return err
	}
	if err := encodeNonKeyAttributes(ctx, input.Item, keys, attrNamesListToId, writer); err != nil {
		return err
	}

	return encodeItemOperationOptionalParams(input.ReturnValues, input.ReturnConsumedCapacity, input.ReturnItemCollectionMetrics, nil,
		nil, input.ConditionExpression, nil, input.ExpressionAttributeNames, input.ExpressionAttributeValues, writer)
}

func encodeDeleteItemInput(ctx aws.Context, input *dynamodb.DeleteItemInput, keySchema *lru.Lru, writer *cbor.Writer) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, fmt.Sprintf("input cannot be nil"), nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}
	if input, err = translateLegacyDeleteItemInput(input); err != nil {
		return err
	}
	table := *input.TableName
	keys, err := getKeySchema(ctx, keySchema, *input.TableName)
	if err != nil {
		return nil
	}

	if err := encodeServiceAndMethod(deleteItem_1013539361_1_Id, writer); err != nil {
		return err
	}
	if err := writer.WriteBytes([]byte(table)); err != nil {
		return err
	}

	if err := cbor.EncodeItemKey(input.Key, keys, writer); err != nil {
		return err
	}

	return encodeItemOperationOptionalParams(input.ReturnValues, input.ReturnConsumedCapacity, input.ReturnItemCollectionMetrics, nil,
		nil, input.ConditionExpression, nil, input.ExpressionAttributeNames, input.ExpressionAttributeValues, writer)
}

func encodeUpdateItemInput(ctx aws.Context, input *dynamodb.UpdateItemInput, keySchema *lru.Lru, writer *cbor.Writer) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, fmt.Sprintf("input cannot be nil"), nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}
	if input, err = translateLegacyUpdateItemInput(input); err != nil {
		return err
	}
	table := *input.TableName
	keys, err := getKeySchema(ctx, keySchema, *input.TableName)
	if err != nil {
		return nil
	}

	if err := encodeServiceAndMethod(updateItem_1425579023_1_Id, writer); err != nil {
		return err
	}
	if err := writer.WriteBytes([]byte(table)); err != nil {
		return err
	}

	if err := cbor.EncodeItemKey(input.Key, keys, writer); err != nil {
		return err
	}

	return encodeItemOperationOptionalParams(input.ReturnValues, input.ReturnConsumedCapacity, input.ReturnItemCollectionMetrics, nil,
		nil, input.ConditionExpression, input.UpdateExpression, input.ExpressionAttributeNames, input.ExpressionAttributeValues, writer)
}

func encodeGetItemInput(ctx aws.Context, input *dynamodb.GetItemInput, keySchema *lru.Lru, writer *cbor.Writer) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, fmt.Sprintf("input cannot be nil"), nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}
	if input, err = translateLegacyGetItemInput(input); err != nil {
		return err
	}
	table := *input.TableName
	keys, err := getKeySchema(ctx, keySchema, table)
	if err != nil {
		return err
	}

	if err := encodeServiceAndMethod(getItem_263244906_1_Id, writer); err != nil {
		return err
	}
	if err := writer.WriteBytes([]byte(table)); err != nil {
		return err
	}
	if err := cbor.EncodeItemKey(input.Key, keys, writer); err != nil {
		return err
	}
	return encodeItemOperationOptionalParams(nil, input.ReturnConsumedCapacity, nil, input.ConsistentRead,
		input.ProjectionExpression, nil, nil, input.ExpressionAttributeNames, nil, writer)
}

func encodeScanInput(ctx aws.Context, input *dynamodb.ScanInput, keySchema *lru.Lru, writer *cbor.Writer) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, fmt.Sprintf("input cannot be nil"), nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}
	if input, err = translateLegacyScanInput(input); err != nil {
		return err
	}
	if err := encodeServiceAndMethod(scan_N1875390620_1_Id, writer); err != nil {
		return err
	}
	if err := writer.WriteBytes([]byte(*input.TableName)); err != nil {
		return err
	}
	expressions, err := encodeExpressions(input.ProjectionExpression, input.FilterExpression, nil, input.ExpressionAttributeNames, input.ExpressionAttributeValues)
	if err != nil {
		return err
	}
	return encodeScanQueryOptionalParams(ctx, input.IndexName, input.Select, input.ReturnConsumedCapacity, input.ConsistentRead,
		expressions, input.Segment, input.TotalSegments, input.Limit, nil, input.ExclusiveStartKey, keySchema, *input.TableName, writer)
}

func encodeQueryInput(ctx aws.Context, input *dynamodb.QueryInput, keySchema *lru.Lru, writer *cbor.Writer) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, fmt.Sprintf("input cannot be nil"), nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}
	if input, err = translateLegacyQueryInput(input); err != nil {
		return err
	}
	if input.KeyConditionExpression == nil {
		return awserr.New(request.ParamRequiredErrCode, "KeyConditionExpression cannot be nil", nil)
	}
	if err := encodeServiceAndMethod(query_N931250863_1_Id, writer); err != nil {
		return err
	}
	if err := writer.WriteBytes([]byte(*input.TableName)); err != nil {
		return err
	}
	expressions, err := encodeExpressions(input.ProjectionExpression, input.FilterExpression, input.KeyConditionExpression, input.ExpressionAttributeNames, input.ExpressionAttributeValues)
	if err != nil {
		return err
	}
	if err = writer.WriteBytes(expressions[parser.KeyConditionExpr]); err != nil {
		return err
	}
	return encodeScanQueryOptionalParams(ctx, input.IndexName, input.Select, input.ReturnConsumedCapacity, input.ConsistentRead,
		expressions, nil, nil, input.Limit, input.ScanIndexForward, input.ExclusiveStartKey, keySchema, *input.TableName, writer)
}

func encodeBatchWriteItemInput(ctx aws.Context, input *dynamodb.BatchWriteItemInput, keySchema *lru.Lru, attrNamesListToId *lru.Lru, writer *cbor.Writer) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, fmt.Sprintf("input cannot be nil"), nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}
	if err = encodeServiceAndMethod(batchWriteItem_116217951_1_Id, writer); err != nil {
		return err
	}
	if err = writer.WriteMapHeader(len(input.RequestItems)); err != nil {
		return err
	}
	totalRequests := 0
	for table, wrs := range input.RequestItems {
		keys, err := getKeySchema(ctx, keySchema, table)
		if err != nil {
			return err
		}

		l := len(wrs)
		if l == 0 {
			return awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("1 validation error detected: Value '{%s=%d}' at 'requestItems' failed to satisfy constraint:"+
				" Map value must satisfy constraint: [Member must have length less than or equal to 25, Member must have length greater than or equal to 1", table, l), nil)
		}
		totalRequests = totalRequests + l
		if totalRequests > maxWriteBatchSize {
			return awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("1 validation error detected: Value '{%s=%d}' at 'requestItems' failed to satisfy constraint:"+
				" Map value must satisfy constraint: [Member must have length less than or equal to 25, Member must have length greater than or equal to 1", table, totalRequests), nil)
		}

		if err = writer.WriteString(table); err != nil {
			return err
		}
		if err = writer.WriteArrayHeader(2 * l); err != nil {
			return err
		}

		if hasDuplicatesWriteRequests(wrs, keys) {
			return awserr.New(request.InvalidParameterErrCode, "Provided list of item keys contains duplicates", nil)
		}
		for _, wr := range wrs {
			if pr := wr.PutRequest; pr != nil {
				attrs := pr.Item
				if err = cbor.EncodeItemKey(attrs, keys, writer); err != nil {
					return err
				}
				if err = encodeNonKeyAttributes(ctx, attrs, keys, attrNamesListToId, writer); err != nil {
					return err
				}
			} else if dr := wr.DeleteRequest; dr != nil {
				if err = cbor.EncodeItemKey(dr.Key, keys, writer); err != nil {
					return err
				}
				if err = writer.WriteNull(); err != nil {
					return err
				}
			} else {
				return awserr.New(request.ParamRequiredErrCode, "Both PutRequest and DeleteRequest cannot be empty", nil)
			}
		}
	}
	return encodeItemOperationOptionalParams(nil, input.ReturnConsumedCapacity, input.ReturnItemCollectionMetrics, nil, nil, nil, nil, nil, nil, writer)
}

func encodeBatchGetItemInput(ctx aws.Context, input *dynamodb.BatchGetItemInput, keySchema *lru.Lru, writer *cbor.Writer) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, fmt.Sprintf("input cannot be nil"), nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}
	if input, err = translateLegacyBatchGetItemInput(input); err != nil {
		return err
	}
	if err = encodeServiceAndMethod(batchGetItem_N697851100_1_Id, writer); err != nil {
		return err
	}
	if err = writer.WriteMapHeader(len(input.RequestItems)); err != nil {
		return err
	}
	for table, kaas := range input.RequestItems {
		if err = writer.WriteString(table); err != nil {
			return err
		}

		if err = writer.WriteArrayHeader(3); err != nil {
			return err
		}

		cr := false
		if kaas.ConsistentRead != nil {
			cr = *kaas.ConsistentRead
		}
		if err = writer.WriteBoolean(cr); err != nil {
			return err
		}
		if kaas.ProjectionExpression != nil {
			expressions := make(map[int]string)
			expressions[parser.ProjectionExpr] = *kaas.ProjectionExpression
			encoder := parser.NewExpressionEncoder(expressions, kaas.ExpressionAttributeNames, nil)
			if _, err = encoder.Parse(); err != nil {
				return err
			}
			var buf bytes.Buffer
			if err = encoder.Write(parser.ProjectionExpr, &buf); err != nil {
				return err
			}
			if err = writer.WriteBytes(buf.Bytes()); err != nil {
				return err
			}
		} else {
			if err = writer.WriteNull(); err != nil {
				return err
			}
		}

		tableKeys, err := getKeySchema(ctx, keySchema, table)
		if err != nil {
			return err
		}
		if err = writer.WriteArrayHeader(len(kaas.Keys)); err != nil {
			return err
		}
		if hasDuplicateKeysAndAttributes(kaas, tableKeys) {
			return awserr.New(request.InvalidParameterErrCode, "Provided list of item keys contains duplicates", nil)
		}
		for _, keys := range kaas.Keys {
			if err = cbor.EncodeItemKey(keys, tableKeys, writer); err != nil {
				return err
			}
		}
	}

	return encodeItemOperationOptionalParams(nil, input.ReturnConsumedCapacity, nil, nil, nil, nil, nil, nil, nil, writer)
}

func encodeCompoundKey(key map[string]*dynamodb.AttributeValue, writer *cbor.Writer) error {
	var buf bytes.Buffer
	w := cbor.NewWriter(&buf)
	defer w.Close()
	if err := w.WriteMapStreamHeader(); err != nil {
		return err
	}
	if len(key) > 0 {
		for k, v := range key {
			if err := w.WriteString(k); err != nil {
				return err
			}
			if err := cbor.EncodeAttributeValue(v, w); err != nil {
				return err
			}
		}
	}
	if err := w.WriteStreamBreak(); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return writer.WriteBytes(buf.Bytes())
}

func encodeNonKeyAttributes(ctx aws.Context, item map[string]*dynamodb.AttributeValue, keys []dynamodb.AttributeDefinition,
	attrNamesListToId *lru.Lru, writer *cbor.Writer) error {
	var buf bytes.Buffer
	w := cbor.NewWriter(&buf)
	defer w.Close()
	if err := cbor.EncodeItemNonKeyAttributes(ctx, item, keys, attrNamesListToId, w); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return writer.WriteBytes(buf.Bytes())
}

func encodeScanQueryOptionalParams(ctx aws.Context, index, selection, returnConsumedCapacity *string, consistentRead *bool,
	encodedExpressions map[int][]byte, segment, totalSegment, limit *int64, forward *bool,
	startKey map[string]*dynamodb.AttributeValue, keySchema *lru.Lru, table string, writer *cbor.Writer) error {

	var err error
	if err = writer.WriteMapStreamHeader(); err != nil {
		return err
	}
	if index != nil {
		if err = writer.WriteInt(requestParamIndexName); err != nil {
			return err
		}
		if err = writer.WriteBytes([]byte(*index)); err != nil {
			return err
		}
	}
	if selection != nil {
		if err = writer.WriteInt(requestParamSelect); err != nil {
			return err
		}
		if err = writer.WriteInt(translateSelect(selection)); err != nil {
			return err
		}
	}
	if returnConsumedCapacity != nil {
		if err = writer.WriteInt(requestParamReturnConsumedCapacity); err != nil {
			return err
		}
		if err = writer.WriteInt(translateReturnConsumedCapacity(returnConsumedCapacity)); err != nil {
			return err
		}
	}
	if consistentRead != nil {
		if err = writer.WriteInt(requestParamConsistentRead); err != nil {
			return err
		}
		cr := 0
		if *consistentRead {
			cr = 1
		}
		if err = writer.WriteInt(cr); err != nil {
			return err
		}
	}

	if len(startKey) != 0 {
		if err = writer.WriteInt(requestParamExclusiveStartKey); err != nil {
			return err
		}
		if index == nil {
			tableKeys, err := getKeySchema(ctx, keySchema, table)
			if err != nil {
				return nil
			}
			if err = cbor.EncodeItemKey(startKey, tableKeys, writer); err != nil {
				return err
			}
		} else {
			if err = encodeCompoundKey(startKey, writer); err != nil {
				return err
			}
		}
	}
	if segment != nil {
		if err = writer.WriteInt(requestParamSegment); err != nil {
			return err
		}
		if err = writer.WriteInt64(*segment); err != nil {
			return err
		}
	}
	if totalSegment != nil {
		if err = writer.WriteInt(requestParamTotalSegments); err != nil {
			return err
		}
		if err = writer.WriteInt64(*totalSegment); err != nil {
			return err
		}
	}
	if limit != nil {
		if err = writer.WriteInt(requestParamLimit); err != nil {
			return err
		}
		if err = writer.WriteInt64(*limit); err != nil {
			return err
		}
	}
	if forward != nil {
		if err = writer.WriteInt(requestParamScanIndexForward); err != nil {
			return err
		}
		if err = writer.WriteInt(translateScanIndexForward(forward)); err != nil {
			return err
		}
	}

	if len(encodedExpressions) > 0 {
		for k, v := range encodedExpressions {
			var e int
			switch k {
			case parser.ProjectionExpr:
				e = requestParamProjectionExpression
			case parser.FilterExpr:
				e = requestParamFilterExpression
			default:
				continue
			}
			if err = writer.WriteInt(e); err != nil {
				return err
			}
			if err = writer.WriteBytes(v); err != nil {
				return err
			}
		}
	}

	return writer.WriteStreamBreak()
}

func encodeItemOperationOptionalParams(returnValues, returnConsumedCapacity, returnItemCollectionMetrics *string, consistentRead *bool,
	projectionExp, conditionalExpr, updateExpr *string, exprAttrNames map[string]*string, exprAttrValues map[string]*dynamodb.AttributeValue, writer *cbor.Writer) error {
	if err := writer.WriteMapStreamHeader(); err != nil {
		return err
	}

	if consistentRead != nil {
		if err := writer.WriteInt(requestParamConsistentRead); err != nil {
			return err
		}
		if err := writer.WriteBoolean(*consistentRead); err != nil {
			return err
		}
	}

	if dv := translateReturnValues(returnValues); dv != returnValueNone {
		if err := writer.WriteInt(requestParamReturnValues); err != nil {
			return err
		}
		if err := writer.WriteInt(dv); err != nil {
			return err
		}
	}

	if dv := translateReturnConsumedCapacity(returnConsumedCapacity); dv != returnConsumedCapacityNone {
		if err := writer.WriteInt(requestParamReturnConsumedCapacity); err != nil {
			return err
		}
		if err := writer.WriteInt(dv); err != nil {
			return err
		}
	}

	if dv := translateReturnItemCollectionMetrics(returnItemCollectionMetrics); dv != returnItemCollectionMetricsNone {
		if err := writer.WriteInt(requestParamReturnItemCollectionMetrics); err != nil {
			return err
		}
		if err := writer.WriteInt(dv); err != nil {
			return err
		}
	}

	if conditionalExpr != nil || updateExpr != nil || projectionExp != nil {
		expressions := make(map[int]string)
		if conditionalExpr != nil {
			expressions[parser.ConditionExpr] = *conditionalExpr
		}
		if updateExpr != nil {
			expressions[parser.UpdateExpr] = *updateExpr
		}
		if projectionExp != nil {
			expressions[parser.ProjectionExpr] = *projectionExp
		}
		encoder := parser.NewExpressionEncoder(expressions, exprAttrNames, exprAttrValues)
		if _, err := encoder.Parse(); err != nil {
			return err
		}
		for k := range expressions {
			var e int
			switch k {
			case parser.ConditionExpr:
				e = requestParamConditionExpression
			case parser.UpdateExpr:
				e = requestParamUpdateExpression
			case parser.ProjectionExpr:
				e = requestParamProjectionExpression
			default:
				continue
			}
			var buf bytes.Buffer
			if err := encoder.Write(k, &buf); err != nil {
				return nil
			}
			if err := writer.WriteInt(e); err != nil {
				return err
			}
			if err := writer.WriteBytes(buf.Bytes()); err != nil {
				return err
			}
		}
	}

	return writer.WriteStreamBreak()
}

func encodeServiceAndMethod(method int, writer *cbor.Writer) error {
	if err := writer.WriteInt(daxServiceId); err != nil {
		return err
	}
	return writer.WriteInt(method)
}

func encodeExpressions(projection, filter, keyCondition *string, exprAttrNames map[string]*string, exprAttrValues map[string]*dynamodb.AttributeValue) (map[int][]byte, error) {
	expressions := make(map[int]string)
	if projection != nil {
		expressions[parser.ProjectionExpr] = *projection
	}
	if filter != nil {
		expressions[parser.FilterExpr] = *filter
	}
	if keyCondition != nil {
		expressions[parser.KeyConditionExpr] = *keyCondition
	}
	encoder := parser.NewExpressionEncoder(expressions, exprAttrNames, exprAttrValues)
	return encoder.Parse()
}

func translateReturnValues(returnValues *string) int {
	if returnValues == nil {
		return returnValueNone
	}
	switch *returnValues {
	case dynamodb.ReturnValueAllOld:
		return returnValueAllOld
	case dynamodb.ReturnValueUpdatedOld:
		return returnValueUpdatedOld
	case dynamodb.ReturnValueAllNew:
		return returnValueAllNew
	case dynamodb.ReturnValueUpdatedNew:
		return returnValueUpdatedNew
	default:
		return returnValueNone
	}
}

func translateReturnConsumedCapacity(returnConsumedCapacity *string) int {
	if returnConsumedCapacity == nil {
		return returnConsumedCapacityNone
	}
	switch *returnConsumedCapacity {
	case dynamodb.ReturnConsumedCapacityTotal:
		return returnConsumedCapacityTotal
	case dynamodb.ReturnConsumedCapacityIndexes:
		return returnConsumedCapacityIndexes
	default:
		return returnItemCollectionMetricsNone
	}
}

func translateReturnItemCollectionMetrics(returnItemCollectionMetrics *string) int {
	if returnItemCollectionMetrics == nil {
		return returnItemCollectionMetricsNone
	}
	if dynamodb.ReturnItemCollectionMetricsSize == *returnItemCollectionMetrics {
		return returnItemCollectionMetricsSize
	}
	return returnItemCollectionMetricsNone
}

func translateSelect(selection *string) int {
	if selection == nil {
		return selectAllAttributes
	}
	switch *selection {
	case dynamodb.SelectAllAttributes:
		return selectAllAttributes
	case dynamodb.SelectAllProjectedAttributes:
		return selectAllProjectedAttributes
	case dynamodb.SelectCount:
		return selectCount
	case dynamodb.SelectSpecificAttributes:
		return selectSpecificAttributes
	default:
		return selectAllAttributes
	}
}

func translateScanIndexForward(b *bool) int {
	if b == nil {
		return 1
	}
	if *b {
		return 1
	}
	return 0
}

func hasDuplicatesWriteRequests(wrs []*dynamodb.WriteRequest, d []dynamodb.AttributeDefinition) bool {
	if len(wrs) <= 1 {
		return false
	}
	face := make([]item, len(wrs))
	for i, v := range wrs {
		if v == nil {
			return false // continue with request processing, will fail later with proper error msg
		}
		face[i] = (*writeItem)(v)
	}

	var err error
	sort.Sort(dupKeys{d, face, func(a, b item) int {
		if err != nil {
			return 0
		}
		for _, k := range d {
			r := strings.Compare(a.key(k), b.key(k))
			if r != 0 {
				return r
			}
		}
		err = fmt.Errorf("dup %v %v", a, b)
		return 0
	}})
	return err != nil
}

func hasDuplicateKeysAndAttributes(kaas *dynamodb.KeysAndAttributes, d []dynamodb.AttributeDefinition) bool {
	if kaas == nil || len(kaas.Keys) <= 1 {
		return false
	}
	face := make([]item, len(kaas.Keys))
	for i, v := range kaas.Keys {
		if v == nil {
			return false // continue with request processing, will fail later with proper error msg
		}
		face[i] = (attrItem)(v)
	}

	var err error
	sort.Sort(dupKeys{d, face, func(a, b item) int {
		if err != nil {
			return 0
		}
		for _, k := range d {
			r := strings.Compare(a.key(k), b.key(k))
			if r != 0 {
				return r
			}
		}
		err = fmt.Errorf("dup %v %v", a, b)
		return 0
	}})
	return err != nil
}

type item interface {
	key(def dynamodb.AttributeDefinition) string
}

type itemKey dynamodb.AttributeDefinition

func (i itemKey) extract(v *dynamodb.AttributeValue) string {
	if v == nil {
		return ""
	}
	switch *i.AttributeType {
	case dynamodb.ScalarAttributeTypeS:
		if v.S != nil {
			return *v.S
		}
	case dynamodb.ScalarAttributeTypeN:
		if v.N != nil {
			return *v.N
		}
	case dynamodb.ScalarAttributeTypeB:
		return string(v.B)
	}
	return ""
}

type writeItem dynamodb.WriteRequest

func (w writeItem) key(def dynamodb.AttributeDefinition) string {
	var v *dynamodb.AttributeValue
	if w.PutRequest != nil && w.PutRequest.Item != nil {
		v = w.PutRequest.Item[*def.AttributeName]
	} else if w.DeleteRequest != nil && w.DeleteRequest.Key != nil {
		v = w.DeleteRequest.Key[*def.AttributeName]
	}
	return itemKey(def).extract(v)
}

type attrItem map[string]*dynamodb.AttributeValue

func (w attrItem) key(def dynamodb.AttributeDefinition) string {
	v := w[*def.AttributeName]
	return itemKey(def).extract(v)
}

type dupKeys struct {
	defs  []dynamodb.AttributeDefinition
	items []item
	eq    func(a, b item) int
}

// Implements sort.Interface
func (d dupKeys) Len() int           { return len(d.items) }
func (d dupKeys) Swap(i, j int)      { d.items[i], d.items[j] = d.items[j], d.items[i] }
func (d dupKeys) Less(i, j int) bool { return d.eq(d.items[i], d.items[j]) <= 0 }
