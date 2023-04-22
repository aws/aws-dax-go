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
	"sort"
	"strings"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-dax-go/dax/internal/parser"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/gofrs/uuid"
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
	transactWriteItems_N1160037738_1_Id = -1160037738
	transactGetItems_1866287579_1_Id    = 1866287579
	batchGetItem_N697851100_1_Id        = -697851100
	batchWriteItem_116217951_1_Id       = 116217951
	getItem_263244906_1_Id              = 263244906
	putItem_N2106490455_1_Id            = -2106490455
	deleteItem_1013539361_1_Id          = 1013539361
	updateItem_1425579023_1_Id          = 1425579023
	query_N931250863_1_Id               = -931250863
	scan_N1875390620_1_Id               = -1875390620

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
	requestParamRequestItemsClientRequestToken
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

const (
	getOperation = iota + 1
	putOperation
	exchangeOperation
	insertOperation
	replaceOperation
	updateOperation
	deleteOperation
	removeOperation
	partialUpdateOperation
	batchGetOperation
	batchOperation
	checkOperation
	transactWriteOperation
	transactGetOperation
	scanOperation
	queryOperation
	createTableOperation
	deleteTableOperation
	describeTableOperation
	listTablesOperation
	updateTableOperation
)

const (
	returnValueOnConditionCheckFailureNone = iota + 1
	returnValueOnConditionCheckFailureAllOld
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

	return encodeItemOperationOptionalParams(
		input.ReturnValues,
		input.ReturnConsumedCapacity,
		input.ReturnItemCollectionMetrics,
		nil,
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

	return encodeItemOperationOptionalParams(
		input.ReturnValues,
		input.ReturnConsumedCapacity,
		input.ReturnItemCollectionMetrics,
		nil,
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

	return encodeItemOperationOptionalParams(
		input.ReturnValues,
		input.ReturnConsumedCapacity,
		input.ReturnItemCollectionMetrics,
		nil,
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
	return encodeItemOperationOptionalParams(
		types.ReturnValueNone,
		input.ReturnConsumedCapacity,
		types.ReturnItemCollectionMetricsNone,
		input.ConsistentRead,
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
	return encodeScanQueryOptionalParams(
		ctx,
		input.IndexName,
		input.Select,
		input.ReturnConsumedCapacity,
		input.ConsistentRead,
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
	return encodeScanQueryOptionalParams(
		ctx,
		input.IndexName,
		input.Select,
		input.ReturnConsumedCapacity,
		input.ConsistentRead,
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
	return encodeItemOperationOptionalParams(
		types.ReturnValueNone,
		input.ReturnConsumedCapacity, input.ReturnItemCollectionMetrics, nil, nil, nil, nil, nil, nil, writer)
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

	return encodeItemOperationOptionalParams(
		types.ReturnValueNone,
		input.ReturnConsumedCapacity,
		types.ReturnItemCollectionMetricsNone,
		nil, nil, nil, nil, nil, nil, writer)
}

func encodeTransactWriteItemsInput(ctx aws.Context, input *dynamodb.TransactWriteItemsInput, keySchema *lru.Lru, attrNamesListToId *lru.Lru, writer *cbor.Writer, extractedKeys []map[string]*dynamodb.AttributeValue) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, "input cannot be nil", nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}

	if err = encodeServiceAndMethod(transactWriteItems_N1160037738_1_Id, writer); err != nil {
		return err
	}

	var operationsBuf, tableNamesBuf, keysBuf, valuesBuf, conditionExpressionsBuf,
		updateExpressionsBuf, rvOnConditionCheckFailureBuf bytes.Buffer
	operationWriter := cbor.NewWriter(&operationsBuf)
	tableNamesWriter := cbor.NewWriter(&tableNamesBuf)
	keysWriter := cbor.NewWriter(&keysBuf)
	valuesWriter := cbor.NewWriter(&valuesBuf)
	conditionExpressionsWriter := cbor.NewWriter(&conditionExpressionsBuf)
	updateExpressionsWriter := cbor.NewWriter(&updateExpressionsBuf)
	rvOnConditionCheckFailureWriter := cbor.NewWriter(&rvOnConditionCheckFailureBuf)

	l := len(input.TransactItems)

	if err = operationWriter.WriteArrayHeader(l); err != nil {
		return err
	}
	if err = tableNamesWriter.WriteArrayHeader(l); err != nil {
		return err
	}
	if err = keysWriter.WriteArrayHeader(l); err != nil {
		return err
	}
	if err = valuesWriter.WriteArrayHeader(l); err != nil {
		return err
	}
	if err = conditionExpressionsWriter.WriteArrayHeader(l); err != nil {
		return err
	}
	if err = updateExpressionsWriter.WriteArrayHeader(l); err != nil {
		return err
	}
	if err = rvOnConditionCheckFailureWriter.WriteArrayHeader(l); err != nil {
		return err
	}

	defer func() {
		operationWriter.Close()
		tableNamesWriter.Close()
		keysWriter.Close()
		valuesWriter.Close()
		conditionExpressionsWriter.Close()
		updateExpressionsWriter.Close()
		rvOnConditionCheckFailureWriter.Close()
	}()

	tableKeySet := make(map[string]bool)
	for i, twi := range input.TransactItems {
		if twi == nil {
			return awserr.New(request.ParamRequiredErrCode, "TransactWriteItem cannot be nil", nil)
		}
		var operation int
		var tableName *string
		var key map[string]*dynamodb.AttributeValue
		var item map[string]*dynamodb.AttributeValue
		var isItem bool = false
		var conditionExpression *string
		var updateExpression *string
		var expressionAttributeNames map[string]*string
		var expressionAttributeValues map[string]*dynamodb.AttributeValue
		var rvOnConditionCheckFailure *string
		opCount := 0
		if check := twi.ConditionCheck; check != nil {
			opCount++
			operation = checkOperation
			conditionExpression = check.ConditionExpression
			expressionAttributeNames = check.ExpressionAttributeNames
			expressionAttributeValues = check.ExpressionAttributeValues
			tableName = check.TableName
			key = check.Key
			rvOnConditionCheckFailure = check.ReturnValuesOnConditionCheckFailure
		}
		if delete := twi.Delete; delete != nil {
			opCount++
			operation = deleteOperation
			conditionExpression = delete.ConditionExpression
			expressionAttributeNames = delete.ExpressionAttributeNames
			expressionAttributeValues = delete.ExpressionAttributeValues
			tableName = delete.TableName
			key = delete.Key
			rvOnConditionCheckFailure = delete.ReturnValuesOnConditionCheckFailure
		}
		if put := twi.Put; put != nil {
			opCount++
			operation = putOperation
			conditionExpression = put.ConditionExpression
			expressionAttributeNames = put.ExpressionAttributeNames
			expressionAttributeValues = put.ExpressionAttributeValues
			tableName = put.TableName
			item = put.Item
			isItem = true
			rvOnConditionCheckFailure = put.ReturnValuesOnConditionCheckFailure
		}
		if update := twi.Update; update != nil {
			opCount++
			operation = partialUpdateOperation
			conditionExpression = update.ConditionExpression
			expressionAttributeNames = update.ExpressionAttributeNames
			expressionAttributeValues = update.ExpressionAttributeValues
			tableName = update.TableName
			key = update.Key
			updateExpression = update.UpdateExpression
			rvOnConditionCheckFailure = update.ReturnValuesOnConditionCheckFailure
		}

		if opCount == 0 {
			return awserr.New(request.ParamRequiredErrCode, "Invalid Request: TransactWriteItemsInput should contain Delete or Put or Update request", nil)
		}
		if opCount > 1 {
			return awserr.New(request.ParamRequiredErrCode, "TransactItems can only contain one of ConditionalCheck, Put, Update or Delete", nil)
		}

		if err := operationWriter.WriteInt(operation); err != nil {
			return err
		}
		if err := tableNamesWriter.WriteBytes([]byte(*tableName)); err != nil {
			return err
		}

		keydef, err := getKeySchema(ctx, keySchema, *tableName)
		if err != nil {
			return nil
		}

		// Check if duplicate [key, tableName] pair exists
		var keyBytes []byte
		if isItem {
			keyBytes, err = cbor.GetEncodedItemKey(item, keydef)
		} else {
			keyBytes, err = cbor.GetEncodedItemKey(key, keydef)
		}
		if err != nil {
			return err
		}
		keysWriter.WriteBytes(keyBytes)
		keyBytes = append(keyBytes, []byte(*tableName)...)
		tableKey := string(keyBytes)
		_, ok := tableKeySet[tableKey]
		if ok {
			return awserr.New(request.ParamRequiredErrCode, "Transaction request cannot include multiple operations on one item", nil)
		} else {
			tableKeySet[tableKey] = true
		}

		switch operation {
		case checkOperation, deleteOperation, partialUpdateOperation:
			if err := valuesWriter.WriteNull(); err != nil {
				return err
			}
		case putOperation:
			if err := encodeNonKeyAttributes(ctx, item, keydef, attrNamesListToId, valuesWriter); err != nil {
				return err
			}
			key = map[string]*dynamodb.AttributeValue{}
			for _, attrDef := range keydef {
				key[*attrDef.AttributeName] = item[*attrDef.AttributeName]
			}
		}

		extractedKeys[i] = key

		encoded, err := parseExpressions(conditionExpression, updateExpression, nil, expressionAttributeNames, expressionAttributeValues)
		if err != nil {
			return err
		}
		if parsedConditionExpr := encoded[parser.ConditionExpr]; parsedConditionExpr != nil {
			if err := conditionExpressionsWriter.WriteBytes(parsedConditionExpr); err != nil {
				return err
			}
		} else {
			if err := conditionExpressionsWriter.WriteNull(); err != nil {
				return err
			}
		}

		if parsedUpdateExpr := encoded[parser.UpdateExpr]; parsedUpdateExpr != nil {
			if err := updateExpressionsWriter.WriteBytes(parsedUpdateExpr); err != nil {
				return err
			}
		} else {
			if err := updateExpressionsWriter.WriteNull(); err != nil {
				return err
			}
		}

		if rvOnConditionCheckFailure != nil && *rvOnConditionCheckFailure == dynamodb.ReturnValuesOnConditionCheckFailureAllOld {
			if err := rvOnConditionCheckFailureWriter.WriteInt(returnValueOnConditionCheckFailureAllOld); err != nil {
				return err
			}
		} else {
			if err := rvOnConditionCheckFailureWriter.WriteInt(returnValueOnConditionCheckFailureNone); err != nil {
				return err
			}
		}
	}

	if err := operationWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(operationsBuf.Bytes()); err != nil {
		return err
	}
	if err := tableNamesWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(tableNamesBuf.Bytes()); err != nil {
		return err
	}
	if err := keysWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(keysBuf.Bytes()); err != nil {
		return err
	}
	if err := valuesWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(valuesBuf.Bytes()); err != nil {
		return err
	}
	// Write null for returnValues
	if err := writer.WriteNull(); err != nil {
		return err
	}
	if err := rvOnConditionCheckFailureWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(rvOnConditionCheckFailureBuf.Bytes()); err != nil {
		return err
	}
	if err := conditionExpressionsWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(conditionExpressionsBuf.Bytes()); err != nil {
		return err
	}
	if err := updateExpressionsWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(updateExpressionsBuf.Bytes()); err != nil {
		return err
	}

	if input.ClientRequestToken == nil {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}
		input.ClientRequestToken = aws.String(id.String())
	}
	return encodeItemOperationOptionalParamsWithToken(
		types.ReturnValueNone,
		input.ReturnConsumedCapacity,
		input.ReturnItemCollectionMetrics,
		nil, nil, nil, nil, nil, nil, input.ClientRequestToken, writer)
}

func encodeTransactGetItemsInput(ctx aws.Context, input *dynamodb.TransactGetItemsInput, keySchema *lru.Lru, writer *cbor.Writer, extractedKeys []map[string]*dynamodb.AttributeValue) error {
	if input == nil {
		return awserr.New(request.ParamRequiredErrCode, "input cannot be nil", nil)
	}
	var err error
	if err = input.Validate(); err != nil {
		return err
	}

	if err = encodeServiceAndMethod(transactGetItems_1866287579_1_Id, writer); err != nil {
		return err
	}

	var tableNamesBuf, keysBuf, projectionExpressionsBuf bytes.Buffer
	tableNamesWriter := cbor.NewWriter(&tableNamesBuf)
	keysWriter := cbor.NewWriter(&keysBuf)
	projectionExpressionsWriter := cbor.NewWriter(&projectionExpressionsBuf)

	len := len(input.TransactItems)

	if err = tableNamesWriter.WriteArrayHeader(len); err != nil {
		return err
	}
	if err = keysWriter.WriteArrayHeader(len); err != nil {
		return err
	}
	if err = projectionExpressionsWriter.WriteArrayHeader(len); err != nil {
		return err
	}

	defer func() {
		tableNamesWriter.Close()
		keysWriter.Close()
		projectionExpressionsWriter.Close()
	}()

	for i, tgi := range input.TransactItems {
		if tgi == nil {
			return awserr.New(request.ParamRequiredErrCode, "TransactGetItem cannot be nil", nil)
		}
		var tableName *string
		var key map[string]*dynamodb.AttributeValue
		var projectionExpression *string
		var expressionAttributeNames map[string]*string
		get := tgi.Get
		tableName = get.TableName
		key = get.Key
		expressionAttributeNames = get.ExpressionAttributeNames
		projectionExpression = get.ProjectionExpression

		extractedKeys[i] = key

		if err := tableNamesWriter.WriteBytes([]byte(*tableName)); err != nil {
			return err
		}

		keydef, err := getKeySchema(ctx, keySchema, *tableName)
		if err != nil {
			return err
		}

		if err := cbor.EncodeItemKey(key, keydef, keysWriter); err != nil {
			return err
		}

		encoded, err := parseExpressions(nil, nil, projectionExpression, expressionAttributeNames, nil)
		if err != nil {
			return err
		}

		if parsedProjectionExpr := encoded[parser.ProjectionExpr]; parsedProjectionExpr != nil {
			if err := projectionExpressionsWriter.WriteBytes(parsedProjectionExpr); err != nil {
				return err
			}
		} else {
			if err := projectionExpressionsWriter.WriteNull(); err != nil {
				return err
			}
		}
	}

	if err := tableNamesWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(tableNamesBuf.Bytes()); err != nil {
		return err
	}
	if err := keysWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(keysBuf.Bytes()); err != nil {
		return err
	}
	if err := projectionExpressionsWriter.Flush(); err != nil {
		return err
	}
	if err := writer.Write(projectionExpressionsBuf.Bytes()); err != nil {
		return err
	}

	return encodeItemOperationOptionalParams(
		types.ReturnValueNone,
		input.ReturnConsumedCapacity,
		types.ReturnItemCollectionMetricsNone,
		nil, nil, nil, nil, nil, nil, writer)
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

func encodeScanQueryOptionalParams(
	ctx aws.Context,
	index *string,
	selection types.Select,
	returnConsumedCapacity types.ReturnConsumedCapacity,
	consistentRead *bool,
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
	if selection != types.SelectAllAttributes {
		if err = writer.WriteInt(requestParamSelect); err != nil {
			return err
		}
		if err = writer.WriteInt(translateSelect(selection)); err != nil {
			return err
		}
	}
	if returnConsumedCapacity != types.ReturnConsumedCapacityNone {
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

func encodeItemOperationOptionalParamsWithToken(
	returnValues types.ReturnValue,
	returnConsumedCapacity types.ReturnConsumedCapacity,
	returnItemCollectionMetrics types.ReturnItemCollectionMetrics,
	consistentRead *bool,
	projectionExp, conditionalExpr, updateExpr *string, exprAttrNames map[string]*string, exprAttrValues map[string]*dynamodb.AttributeValue, clientRequestToken *string, writer *cbor.Writer) error {
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
		encoded, err := parseExpressions(conditionalExpr, updateExpr, projectionExp, exprAttrNames, exprAttrValues)
		if err != nil {
			return err
		}
		for k := range encoded {
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
			if err := writer.WriteInt(e); err != nil {
				return err
			}
			if err := writer.WriteBytes(encoded[k]); err != nil {
				return err
			}
		}
	}

	if clientRequestToken != nil {
		if err := writer.WriteInt(requestParamRequestItemsClientRequestToken); err != nil {
			return err
		}
		if err := writer.WriteString(*clientRequestToken); err != nil {
			return err
		}
	}

	return writer.WriteStreamBreak()
}

func encodeItemOperationOptionalParams(
	returnValues types.ReturnValue,
	returnConsumedCapacity types.ReturnConsumedCapacity,
	returnItemCollectionMetrics types.ReturnItemCollectionMetrics,
	consistentRead *bool,
	projectionExp, conditionalExpr, updateExpr *string, exprAttrNames map[string]*string, exprAttrValues map[string]*dynamodb.AttributeValue, writer *cbor.Writer) error {
	return encodeItemOperationOptionalParamsWithToken(
		returnValues,
		returnConsumedCapacity,
		returnItemCollectionMetrics,
		consistentRead,
		projectionExp, conditionalExpr, updateExpr, exprAttrNames, exprAttrValues, nil, writer)
}

func parseExpressions(conditionalExpr, updateExpr, projectionExp *string, exprAttrNames map[string]*string, exprAttrValues map[string]*dynamodb.AttributeValue) (map[int][]byte, error) {
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
	encoded, err := encoder.Parse()
	if err != nil {
		return nil, err
	}
	return encoded, nil
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

func translateReturnValues(returnValues types.ReturnValue) int {
	switch returnValues {
	case types.ReturnValueAllOld:
		return returnValueAllOld
	case types.ReturnValueUpdatedOld:
		return returnValueUpdatedOld
	case types.ReturnValueAllNew:
		return returnValueAllNew
	case types.ReturnValueUpdatedNew:
		return returnValueUpdatedNew
	default:
		return returnValueNone
	}
}

func translateReturnConsumedCapacity(returnConsumedCapacity types.ReturnConsumedCapacity) int {
	switch returnConsumedCapacity {
	case types.ReturnConsumedCapacityTotal:
		return returnConsumedCapacityTotal
	case types.ReturnConsumedCapacityIndexes:
		return returnConsumedCapacityIndexes
	default:
		return returnConsumedCapacityNone
	}
}

func translateReturnItemCollectionMetrics(returnItemCollectionMetrics types.ReturnItemCollectionMetrics) int {
	if types.ReturnItemCollectionMetricsSize == returnItemCollectionMetrics {
		return returnItemCollectionMetricsSize
	}
	return returnItemCollectionMetricsNone
}

func translateSelect(selection types.Select) int {
	switch selection {
	case types.SelectAllProjectedAttributes:
		return selectAllProjectedAttributes
	case types.SelectCount:
		return selectCount
	case types.SelectSpecificAttributes:
		return selectSpecificAttributes
	case types.SelectAllAttributes:
		return selectAllAttributes
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
