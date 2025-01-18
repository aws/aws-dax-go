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
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
)

const (
	responseParamItem = iota
	responseParamConsumedCapacity
	responseParamAttributes
	responseParamItemCollectionMetrics
	responseParamResponses
	responseParamUnprocessedKeys
	responseParamUnprocessedItems
	responseParamItems
	responseParamCount
	responseParamLastEvaluatedKey
	responseParamScannedCount
	responseParamTableDescription
)

const (
	roleLeader  = 1
	roleReplica = 2
)

const (
	keyNodeId = iota
	keyHostname
	keyAddress
	keyPort
	keyRole
	keyAvailablityZone
	keyLeaderSessionId
)

const (
	capacityUnits = iota + 1
	readCapacityUnits
	writeCapacityUnits
	tableName
	table
	globalSecondaryIndexes
	localSecondaryIndexes
)

func decodeEndpointsOutput(reader *cbor.Reader) ([]serviceEndpoint, error) {
	length, err := reader.ReadArrayLength()
	if err != nil {
		return nil, err
	}
	if length <= 0 {
		return []serviceEndpoint{}, nil
	}
	o := make([]serviceEndpoint, length)
	for i := 0; i < length; i++ {
		o[i], err = decodeEndpoint(reader)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func decodeEndpoint(reader *cbor.Reader) (serviceEndpoint, error) {
	se := serviceEndpoint{}
	err := consumeMap(reader, func(key int, r *cbor.Reader) error {
		var err error
		switch key {
		case keyNodeId:
			if se.nodeId, err = r.ReadInt64(); err != nil {
				return err
			}
		case keyHostname:
			if se.hostname, err = r.ReadString(); err != nil {
				return err
			}
		case keyAddress:
			if se.address, err = r.ReadBytes(); err != nil {
				return err
			}
		case keyPort:
			if se.port, err = r.ReadInt(); err != nil {
				return err
			}
		case keyRole:
			if role, err := r.ReadInt(); err != nil {
				return err
			} else {
				if role != roleLeader && role != roleReplica {
					return &smithy.SerializationError{Err: fmt.Errorf("unknown role %d", role)}
				}
				se.role = role
			}
		case keyAvailablityZone:
			if se.availabilityZone, err = r.ReadString(); err != nil {
				return err
			}
		case keyLeaderSessionId:
			if se.leaderSessionId, err = r.ReadInt64(); err != nil {
				return err
			}
		default:
			// inorder to ensure backward compatibility on future field additions, new/unknown fields are ignored
		}
		return nil
	})
	return se, err
}

func decodeDefineAttributeListIdOutput(reader *cbor.Reader) (int64, error) {
	id, err := reader.ReadInt64()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func decodeDefineAttributeListOutput(reader *cbor.Reader) ([]string, error) {
	length, err := reader.ReadArrayLength()
	if err != nil {
		return nil, err
	}
	attrNames := make([]string, length)
	for i := 0; i < length; i++ {
		an, err := reader.ReadString()
		if err != nil {
			return nil, err
		}
		attrNames[i] = an
	}
	return attrNames, nil
}

func decodeDefineKeySchemaOutput(reader *cbor.Reader) ([]types.AttributeDefinition, error) {
	length, err := reader.ReadMapLength()
	if err != nil {
		return nil, err
	}
	keys := make([]types.AttributeDefinition, length)
	for i := 0; i < length; i++ {
		name, err := reader.ReadString()
		if err != nil {
			return nil, err
		}
		typ, err := reader.ReadString()
		if err != nil {
			return nil, err
		}
		keys[i] = types.AttributeDefinition{AttributeName: &name, AttributeType: types.ScalarAttributeType(typ)}
	}
	return keys, nil
}

func decodePutItemOutput(ctx context.Context, reader *cbor.Reader, input *dynamodb.PutItemInput, keySchemaCache *lru.Lru, attrListIdToNames *lru.Lru) (*dynamodb.PutItemOutput, error) {
	output := &dynamodb.PutItemOutput{}
	if consumed, err := consumeNil(reader); err != nil {
		return output, err
	} else if consumed {
		return output, nil
	}

	tableName := *input.TableName

	err := consumeMap(reader, func(key int, reader *cbor.Reader) error {
		switch key {
		case responseParamConsumedCapacity:
			var err error
			if output.ConsumedCapacity, err = decodeConsumedCapacity(reader); err != nil {
				return err
			}
		case responseParamItemCollectionMetrics:
			keys, err := getKeySchema(ctx, keySchemaCache, tableName)
			if err != nil {
				return err
			}
			if output.ItemCollectionMetrics, err = decodeItemCollectionMetrics(reader, *keys[0].AttributeName); err != nil {
				return err
			}
		case responseParamAttributes:
			attrs, err := decodeNonKeyAttributes(ctx, reader, attrListIdToNames, nil)
			if err != nil {
				return err
			}
			keys, err := getKeySchema(ctx, keySchemaCache, tableName)
			if err != nil {
				return err
			}
			for _, ad := range keys {
				k := *ad.AttributeName
				attrs[k] = input.Item[k]
			}
			output.Attributes = attrs
		default:
			return &smithy.SerializationError{Err: fmt.Errorf("unknown response param key %d", key)}
		}
		return nil
	})
	if err != nil {
		return output, err
	}

	return output, nil
}

func decodeDeleteItemOutput(ctx context.Context, reader *cbor.Reader, input *dynamodb.DeleteItemInput, keySchemaCache *lru.Lru, attrListIdToNames *lru.Lru) (*dynamodb.DeleteItemOutput, error) {
	output := &dynamodb.DeleteItemOutput{}
	if consumed, err := consumeNil(reader); err != nil {
		return output, err
	} else if consumed {
		return output, nil
	}

	tableName := *input.TableName

	err := consumeMap(reader, func(key int, reader *cbor.Reader) error {
		switch key {
		case responseParamConsumedCapacity:
			var err error
			if output.ConsumedCapacity, err = decodeConsumedCapacity(reader); err != nil {
				return err
			}
		case responseParamItemCollectionMetrics:
			keys, err := getKeySchema(ctx, keySchemaCache, tableName)
			if err != nil {
				return err
			}
			if output.ItemCollectionMetrics, err = decodeItemCollectionMetrics(reader, *keys[0].AttributeName); err != nil {
				return err
			}
		case responseParamAttributes:
			attrs, err := decodeNonKeyAttributes(ctx, reader, attrListIdToNames, nil)
			if err != nil {
				return err
			}
			for k, v := range input.Key {
				attrs[k] = v
			}
			output.Attributes = attrs
		default:
			return &smithy.SerializationError{Err: fmt.Errorf("unknown response param key %d", key)}
		}
		return nil
	})
	if err != nil {
		return output, err
	}

	return output, nil
}

func decodeUpdateItemOutput(
	ctx context.Context, reader *cbor.Reader, input *dynamodb.UpdateItemInput,
	keySchemaCache *lru.Lru, attrListIdToNames *lru.Lru,
) (*dynamodb.UpdateItemOutput, error) {
	output := &dynamodb.UpdateItemOutput{}
	if consumed, err := consumeNil(reader); err != nil {
		return output, err
	} else if consumed {
		return output, nil
	}

	tableName := *input.TableName

	err := consumeMap(reader, func(key int, reader *cbor.Reader) error {
		switch key {
		case responseParamConsumedCapacity:
			var err error
			if output.ConsumedCapacity, err = decodeConsumedCapacity(reader); err != nil {
				return err
			}
		case responseParamItemCollectionMetrics:
			keys, err := getKeySchema(ctx, keySchemaCache, tableName)
			if err != nil {
				return err
			}
			if output.ItemCollectionMetrics, err = decodeItemCollectionMetrics(reader, *keys[0].AttributeName); err != nil {
				return err
			}
		case responseParamAttributes:
			rv := input.ReturnValues
			switch rv {
			case types.ReturnValueAllNew, types.ReturnValueAllOld:
				attrs, err := decodeNonKeyAttributes(ctx, reader, attrListIdToNames, nil)
				if err != nil {
					return err
				}
				for k, v := range input.Key {
					attrs[k] = v
				}
				output.Attributes = attrs
			case types.ReturnValueUpdatedNew, types.ReturnValueUpdatedOld:
				var err error
				if output.Attributes, err = decodeAttributeProjection(ctx, reader, attrListIdToNames); err != nil {
					return err
				}
			default:
				return &smithy.SerializationError{Err: fmt.Errorf("unexpected return value %s", rv)}
			}
		default:
			return &smithy.SerializationError{Err: fmt.Errorf("unknown response param key %d", key)}
		}
		return nil
	})
	if err != nil {
		return output, err
	}

	return output, nil
}

func decodeGetItemOutput(ctx context.Context, reader *cbor.Reader, input *dynamodb.GetItemInput, attrListIdToNames *lru.Lru) (*dynamodb.GetItemOutput, error) {
	output := &dynamodb.GetItemOutput{}
	if consumed, err := consumeNil(reader); err != nil {
		return output, err
	} else if consumed {
		return output, nil
	}

	projectionOrdinals, err := buildProjectionOrdinals(input.ProjectionExpression, input.ExpressionAttributeNames)
	if err != nil {
		return output, err
	}

	err = consumeMap(reader, func(key int, reader *cbor.Reader) error {
		switch key {
		case responseParamConsumedCapacity:
			if output.ConsumedCapacity, err = decodeConsumedCapacity(reader); err != nil {
				return err
			}
		case responseParamItem:
			item, err := decodeNonKeyAttributes(ctx, reader, attrListIdToNames, projectionOrdinals)
			if err != nil {
				return err
			}
			if len(projectionOrdinals) == 0 {
				for k, v := range input.Key {
					item[k] = v
				}
			}
			output.Item = item
		default:
			return &smithy.SerializationError{Err: fmt.Errorf("unknown response param key %d", key)}
		}
		return nil
	})
	if err != nil {
		return output, err
	}

	return output, nil
}

func decodeScanOutput(ctx context.Context, reader *cbor.Reader, input *dynamodb.ScanInput, keySchemaCache *lru.Lru, attrNamesListToId *lru.Lru) (*dynamodb.ScanOutput, error) {
	output := &dynamodb.ScanOutput{}
	out, err := decodeScanQueryOutput(ctx, reader, *input.TableName, input.IndexName != nil, input.ProjectionExpression, input.ExpressionAttributeNames, keySchemaCache, attrNamesListToId)
	if err != nil {
		return output, err
	}
	if out == nil {
		return output, nil
	}
	return out.scanOutput(output), nil
}

func decodeQueryOutput(ctx context.Context, reader *cbor.Reader, input *dynamodb.QueryInput, keySchemaCache *lru.Lru, attrNamesListToId *lru.Lru) (*dynamodb.QueryOutput, error) {
	output := &dynamodb.QueryOutput{}
	out, err := decodeScanQueryOutput(ctx, reader, *input.TableName, input.IndexName != nil, input.ProjectionExpression, input.ExpressionAttributeNames, keySchemaCache, attrNamesListToId)
	if err != nil {
		return output, err
	}
	if out == nil {
		return output, nil
	}
	return out.queryOutput(output), nil
}

type scanQueryOutput struct {
	dynamodb.ScanOutput
}

func (o *scanQueryOutput) scanOutput(output *dynamodb.ScanOutput) *dynamodb.ScanOutput {
	if output != nil {
		output.Items = o.Items
		output.ConsumedCapacity = o.ConsumedCapacity
		output.Count = o.Count
		output.ScannedCount = o.ScannedCount
		output.LastEvaluatedKey = o.LastEvaluatedKey
		return output
	}
	return &o.ScanOutput
}

func (o *scanQueryOutput) queryOutput(output *dynamodb.QueryOutput) *dynamodb.QueryOutput {
	if output != nil {
		output.Items = o.Items
		output.ConsumedCapacity = o.ConsumedCapacity
		output.Count = o.Count
		output.ScannedCount = o.ScannedCount
		output.LastEvaluatedKey = o.LastEvaluatedKey
		return output
	}
	return &dynamodb.QueryOutput{
		Items:            o.Items,
		ConsumedCapacity: o.ConsumedCapacity,
		Count:            o.Count,
		ScannedCount:     o.ScannedCount,
		LastEvaluatedKey: o.LastEvaluatedKey,
	}
}

func decodeScanQueryOutput(ctx context.Context, reader *cbor.Reader, table string, indexed bool, projection *string, exprAttrNames map[string]string, keySchemaCache *lru.Lru, attrNamesListToId *lru.Lru) (*scanQueryOutput, error) {
	if consumed, err := consumeNil(reader); err != nil {
		return nil, err
	} else if consumed {
		return nil, nil
	}

	out := &scanQueryOutput{}
	out.Items = []map[string]types.AttributeValue{}
	var err error
	err = consumeMap(reader, func(key int, reader *cbor.Reader) error {
		switch key {
		case responseParamItems:
			projectionOrdinals, err := buildProjectionOrdinals(projection, exprAttrNames)
			if err != nil {
				return err
			}
			if out.Items, err = decodeScanQueryItems(ctx, reader, table, keySchemaCache, attrNamesListToId, projectionOrdinals); err != nil {
				return err
			}
		case responseParamConsumedCapacity:
			if out.ConsumedCapacity, err = decodeConsumedCapacity(reader); err != nil {
				return err
			}
		case responseParamCount:
			c, err := reader.ReadInt64()
			if err != nil {
				return err
			}
			out.Count = int32(c)
		case responseParamScannedCount:
			c, err := reader.ReadInt64()
			if err != nil {
				return err
			}
			out.ScannedCount = int32(c)
		case responseParamLastEvaluatedKey:
			k, err := decodeLastEvaluatedKey(ctx, reader, table, indexed, keySchemaCache)
			if err != nil {
				return err
			}
			if len(k) > 0 {
				out.LastEvaluatedKey = k
			}
		default:
			return &smithy.SerializationError{Err: fmt.Errorf("unknown response param key %d", key)}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func decodeBatchWriteItemOutput(ctx context.Context, reader *cbor.Reader, keySchemaCache *lru.Lru, attrNamesListToId *lru.Lru, output *dynamodb.BatchWriteItemOutput) (*dynamodb.BatchWriteItemOutput, error) {
	if output != nil {
		output.UnprocessedItems = map[string][]types.WriteRequest{}
	}
	if consumed, err := consumeNil(reader); err != nil {
		return output, err
	} else if consumed {
		return output, nil
	}
	numTables, err := reader.ReadMapLength()
	if err != nil {
		return output, err
	}
	if output == nil {
		output = &dynamodb.BatchWriteItemOutput{UnprocessedItems: map[string][]types.WriteRequest{}}
	}
	if numTables > 0 {
		unprocessed := make(map[string][]types.WriteRequest, numTables)
		for i := 0; i < numTables; i++ {
			table, err := reader.ReadString()
			if err != nil {
				return output, err
			}
			tableKeys, err := getKeySchema(ctx, keySchemaCache, table)
			if err != nil {
				return output, err
			}
			numObjs, err := reader.ReadArrayLength()
			if err != nil {
				return output, err
			}
			numItems := numObjs / 2
			wrs := make([]types.WriteRequest, numItems)
			for j := 0; j < numItems; j++ {
				keys, err := decodeKey(reader, tableKeys)
				if err != nil {
					return output, err
				}
				item, err := decodeNonKeyAttributes(ctx, reader, attrNamesListToId, nil)
				if err != nil {
					return output, err
				}
				wr := types.WriteRequest{}
				if item == nil {
					wr.DeleteRequest = &types.DeleteRequest{Key: keys}
				} else {
					for k, v := range keys {
						item[k] = v
					}
					wr.PutRequest = &types.PutRequest{Item: item}
				}
				wrs[j] = wr
			}
			unprocessed[table] = wrs
		}
		if len(unprocessed) != 0 {
			output.UnprocessedItems = unprocessed
		}
	}

	numCC, err := reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if numCC > 0 {
		output.ConsumedCapacity = make([]types.ConsumedCapacity, numCC)
		for i := 0; i < numCC; i++ {
			capacity, err := decodeConsumedCapacity(reader)
			if err != nil {
				return output, err
			}
			output.ConsumedCapacity[i] = *capacity
		}
	}

	icmLen, err := reader.ReadMapLength()
	if err != nil {
		return output, err
	}
	if icmLen > 0 {
		output.ItemCollectionMetrics = make(map[string][]types.ItemCollectionMetrics, icmLen)
		for i := 0; i < icmLen; i++ {
			table, err := reader.ReadString()
			if err != nil {
				return output, err
			}
			keys, err := getKeySchema(ctx, keySchemaCache, table)
			if err != nil {
				return output, err
			}
			pkey := *keys[0].AttributeName
			numMetrics, err := reader.ReadArrayLength()
			if err != nil {
				return output, err
			}
			metrics := make([]types.ItemCollectionMetrics, numMetrics)
			for j := 0; j < numMetrics; j++ {
				m, err := decodeItemCollectionMetrics(reader, pkey)
				if err != nil {
					return output, err
				}
				metrics[j] = *m
			}
			output.ItemCollectionMetrics[table] = metrics
		}
	}

	return output, nil
}

func decodeBatchGetItemOutput(
	ctx context.Context, reader *cbor.Reader, input *dynamodb.BatchGetItemInput,
	keySchemaCache *lru.Lru, attrNamesListToId *lru.Lru, output *dynamodb.BatchGetItemOutput,
) (*dynamodb.BatchGetItemOutput, error) {
	if consumed, err := consumeNil(reader); err != nil {
		return output, err
	} else if consumed {
		return output, nil
	}

	l, err := reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if l != 2 {
		return output, &smithy.SerializationError{Err: fmt.Errorf("Unexpected number of objects %d in BatchGetItemOutput", l)}
	}

	projectionsByTable := make(map[string][]documentPath, len(input.RequestItems))
	for table, kaas := range input.RequestItems {
		if kaas.ProjectionExpression != nil {
			dp, err := buildProjectionOrdinals(kaas.ProjectionExpression, kaas.ExpressionAttributeNames)
			if err != nil {
				return output, err
			}
			projectionsByTable[table] = dp
		}
	}

	numTables, err := reader.ReadMapLength()
	if err != nil {
		return output, err
	}

	if output == nil {
		output = &dynamodb.BatchGetItemOutput{}
	}
	if numTables > 0 {
		output.Responses = make(map[string][]map[string]types.AttributeValue, numTables)
		for i := 0; i < numTables; i++ {
			table, err := reader.ReadString()
			if err != nil {
				return output, err
			}

			projections, hasProjections := projectionsByTable[table]
			if hasProjections {
				numItems, err := reader.ReadArrayLength()
				if err != nil {
					return output, err
				}
				items := make([]map[string]types.AttributeValue, numItems)
				for j := 0; j < numItems; j++ {
					if items[j], err = decodeNonKeyAttributes(ctx, reader, attrNamesListToId, projections); err != nil {
						return output, err
					}
				}
				output.Responses[table] = items
			} else {
				tableKeys, err := getKeySchema(ctx, keySchemaCache, table)
				if err != nil {
					return output, err
				}
				numObjs, err := reader.ReadArrayLength()
				if err != nil {
					return output, err
				}
				numItems := numObjs / 2
				items := make([]map[string]types.AttributeValue, numItems)
				for j := 0; j < numItems; j++ {
					keys, err := decodeKey(reader, tableKeys)
					if err != nil {
						return output, err
					}
					item, err := decodeNonKeyAttributes(ctx, reader, attrNamesListToId, projections)
					if err != nil {
						return output, err
					}
					for k, v := range keys {
						item[k] = v
					}
					items[j] = item
				}
				output.Responses[table] = items
			}
		}
	}

	numUnprocessed, err := reader.ReadMapLength()
	if err != nil {
		return output, err
	}
	if numUnprocessed > 0 {
		unprocessed := make(map[string]types.KeysAndAttributes, numUnprocessed)
		for i := 0; i < numUnprocessed; i++ {
			table, err := reader.ReadString()
			if err != nil {
				return output, err
			}
			tableKeys, err := getKeySchema(ctx, keySchemaCache, table)
			if err != nil {
				return output, err
			}
			numKeys, err := reader.ReadArrayLength()
			if err != nil {
				return output, err
			}
			if numKeys <= 0 {
				continue
			}
			keys := make([]map[string]types.AttributeValue, numKeys)
			for j := 0; j < numKeys; j++ {
				if keys[j], err = decodeKey(reader, tableKeys); err != nil {
					return output, err
				}
			}
			outKaas := types.KeysAndAttributes{Keys: keys}
			if inKaas, ok := input.RequestItems[table]; ok {
				outKaas.AttributesToGet = inKaas.AttributesToGet
				outKaas.ConsistentRead = inKaas.ConsistentRead
				outKaas.ExpressionAttributeNames = inKaas.ExpressionAttributeNames
				outKaas.Keys = keys
				outKaas.ProjectionExpression = inKaas.ProjectionExpression
			}
			unprocessed[table] = outKaas
		}
		if len(unprocessed) != 0 {
			output.UnprocessedKeys = unprocessed
		}
	}

	numCC, err := reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if numCC > 0 {
		output.ConsumedCapacity = make([]types.ConsumedCapacity, numCC)
		for i := 0; i < numCC; i++ {
			capacity, err := decodeConsumedCapacity(reader)
			if err != nil {
				return output, err
			}
			output.ConsumedCapacity[i] = *capacity
		}
	}

	return output, nil
}

func decodeTransactWriteItemsOutput(ctx context.Context, reader *cbor.Reader, input *dynamodb.TransactWriteItemsInput, keySchemaCache *lru.Lru, attrListIdToNames *lru.Lru, output *dynamodb.TransactWriteItemsOutput) (*dynamodb.TransactWriteItemsOutput, error) {
	len, err := reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if len != 3 {
		// returnValues still in the tube even though it's not being returned
		// But user shouldn't be able to see it.
		return output, &smithy.SerializationError{Err: fmt.Errorf("TransactWriteResponse needs to have 2 elements, instead had: %d", len)}
	}
	_, err = reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if output == nil {
		output = &dynamodb.TransactWriteItemsOutput{}
	}

	numCC, err := reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if numCC > 0 {
		output.ConsumedCapacity = make([]types.ConsumedCapacity, numCC)
		for i := 0; i < numCC; i++ {
			capacity, err := decodeConsumedCapacityExtended(reader)
			if err != nil {
				return output, err
			}
			output.ConsumedCapacity[i] = *capacity
		}
	}

	icmLen, err := reader.ReadMapLength()
	if err != nil {
		return output, err
	}
	if icmLen > 0 {
		output.ItemCollectionMetrics = make(map[string][]types.ItemCollectionMetrics, icmLen)
		for i := 0; i < icmLen; i++ {
			table, err := reader.ReadString()
			if err != nil {
				return output, err
			}
			keys, err := getKeySchema(ctx, keySchemaCache, table)
			if err != nil {
				return output, err
			}
			pkey := *keys[0].AttributeName
			numMetrics, err := reader.ReadArrayLength()
			if err != nil {
				return output, err
			}
			metrics := make([]types.ItemCollectionMetrics, numMetrics)
			for j := 0; j < numMetrics; j++ {
				m, err := decodeItemCollectionMetrics(reader, pkey)
				if err != nil {
					return output, err
				}
				metrics[j] = *m
			}
			output.ItemCollectionMetrics[table] = metrics
		}
	}

	return output, nil
}

func decodeTransactGetItemsOutput(ctx context.Context, reader *cbor.Reader, input *dynamodb.TransactGetItemsInput, keySchemaCache *lru.Lru, attrListIdToNames *lru.Lru, output *dynamodb.TransactGetItemsOutput) (*dynamodb.TransactGetItemsOutput, error) {
	length, err := reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if length != 2 {
		return output, &smithy.SerializationError{Err: fmt.Errorf("TransactGetResponse needs to have 2 elements, instead had: %d", length)}
	}

	if output == nil {
		output = &dynamodb.TransactGetItemsOutput{}
	}

	numR, err := reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if numR != len(input.TransactItems) {
		return output, &smithy.SerializationError{Err: fmt.Errorf("TransactGetResponse need to have the same number of Responses "+
			"as the length of TransactItems in the input: %d, instead had: %d", len(input.TransactItems), numR)}
	}

	responses := make([]types.ItemResponse, numR)
	for i := 0; i < numR; i++ {
		get := input.TransactItems[i].Get
		projectionOrdinals, err := buildProjectionOrdinals(get.ProjectionExpression, get.ExpressionAttributeNames)
		if err != nil {
			return output, err
		}
		item, err := decodeNonKeyAttributes(ctx, reader, attrListIdToNames, projectionOrdinals)
		if err != nil {
			return output, err
		}
		// The key attributes are only added if it's NOT a projection
		if item != nil && len(projectionOrdinals) == 0 {
			for k, v := range get.Key {
				item[k] = v
			}
		}
		responses[i] = types.ItemResponse{Item: item}
	}
	output.Responses = responses

	numCC, err := reader.ReadArrayLength()
	if err != nil {
		return output, err
	}
	if numCC > 0 {
		output.ConsumedCapacity = make([]types.ConsumedCapacity, numCC)
		for i := 0; i < numCC; i++ {
			capacity, err := decodeConsumedCapacityExtended(reader)
			if err != nil {
				return output, err
			}
			output.ConsumedCapacity[i] = *capacity
		}
	}

	return output, nil
}

func decodeScanQueryItems(ctx context.Context, reader *cbor.Reader, table string, keySchemaCache *lru.Lru, attrNamesListToId *lru.Lru, projectionOrdinals []documentPath) ([]map[string]types.AttributeValue, error) {
	consumed, err := consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if consumed {
		return nil, nil
	}

	items := []map[string]types.AttributeValue{}
	if len(projectionOrdinals) > 0 {
		err := consumeArray(reader, func(reader *cbor.Reader) error {
			i, err := decodeProjection(reader, projectionOrdinals)
			if err != nil {
				return err
			}
			items = append(items, i)
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		tableKeys, err := getKeySchema(ctx, keySchemaCache, table)
		if err != nil {
			return nil, err
		}
		err = consumeArray(reader, func(reader *cbor.Reader) error {
			length, err := reader.ReadArrayLength()
			if err != nil {
				return err
			}
			if length != 2 {
				return &smithy.SerializationError{Err: fmt.Errorf("expected array of size 2 containing key and value, got %d", length)}
			}
			key, err := decodeKey(reader, tableKeys)
			if err != nil {
				return err
			}
			item, err := decodeNonKeyAttributes(ctx, reader, attrNamesListToId, projectionOrdinals)
			if err != nil {
				return err
			}
			for k, v := range key {
				item[k] = v
			}
			items = append(items, item)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}

func decodeLastEvaluatedKey(ctx context.Context, reader *cbor.Reader, table string, indexed bool, keySchemaCache *lru.Lru) (map[string]types.AttributeValue, error) {
	if indexed {
		key, err := decodeCompoundKey(reader)
		if err != nil {
			return nil, err
		}
		return key, nil
	} else {
		tableKeys, err := getKeySchema(ctx, keySchemaCache, table)
		if err != nil {
			return nil, err
		}
		key, err := decodeKey(reader, tableKeys)
		if err != nil {
			return nil, err
		}
		return key, nil
	}
}

func consumeArray(reader *cbor.Reader, consumer func(reader *cbor.Reader) error) error {
	hdr, err := reader.PeekHeader()
	if err != nil {
		return err
	}
	len, err := reader.ReadArrayLength()
	if err != nil {
		return err
	}
	if hdr == cbor.ArrayStream {
		len = -1
	}
	for i := 0; len < 0 || i < len; i++ {
		if len < 0 {
			consumed, err := consumeBreak(reader)
			if err != nil {
				return err
			}
			if consumed {
				break
			}
		}
		if err := consumer(reader); err != nil {
			return err
		}
	}
	return nil
}

func consumeMap(reader *cbor.Reader, consumer func(int, *cbor.Reader) error) error {
	hdr, err := reader.PeekHeader()
	if err != nil {
		return err
	}
	len, err := reader.ReadMapLength()
	if err != nil {
		return err
	}
	if hdr == cbor.MapStream {
		len = -1
	}
	for i := 0; len < 0 || i < len; i++ {
		if len < 0 {
			consumed, err := consumeBreak(reader)
			if err != nil {
				return err
			}
			if consumed {
				break
			}
		}

		id, err := reader.ReadInt()
		if err != nil {
			return err
		}
		if err := consumer(id, reader); err != nil {
			return err
		}
	}
	return nil
}

func consumeNil(reader *cbor.Reader) (bool, error) {
	hdr, err := reader.PeekHeader()
	if err != nil {
		return false, err
	}
	if hdr != cbor.Nil {
		return false, nil
	}
	if err := reader.ReadNil(); err != nil {
		return false, err
	}
	return true, nil
}

func consumeBreak(reader *cbor.Reader) (bool, error) {
	hdr, err := reader.PeekHeader()
	if err != nil {
		return false, err
	}
	if hdr != cbor.Break {
		return false, nil
	}
	if err = reader.ReadBreak(); err != nil {
		return false, err
	}
	return true, nil
}

func decodeKey(reader *cbor.Reader, keys []types.AttributeDefinition) (map[string]types.AttributeValue, error) {
	consumed, err := consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if consumed {
		return nil, nil
	}
	k, err := cbor.DecodeItemKey(reader, keys)
	if err != nil {
		return nil, err
	}
	return k, nil
}

func decodeCompoundKey(reader *cbor.Reader) (map[string]types.AttributeValue, error) {
	consumed, err := consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if consumed {
		return nil, nil
	}

	r, err := reader.BytesReader()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	hdr, err := r.PeekHeader()
	if err != nil {
		return nil, err
	}
	if hdr != cbor.MapStream {
		return nil, &smithy.SerializationError{Err: errors.New("bad compound key")}
	}
	_, err = r.ReadMapLength()
	if err != nil {
		return nil, err
	}
	key := make(map[string]types.AttributeValue, 4)
	for {
		consumed, err := consumeBreak(r)
		if err != nil {
			return nil, err
		}
		if consumed {
			break
		}

		k, err := r.ReadString()
		if err != nil {
			return nil, err
		}
		v, err := cbor.DecodeAttributeValue(r)
		if err != nil {
			return nil, err
		}
		key[k] = v
	}
	return key, nil
}

func decodeNonKeyAttributes(ctx context.Context, reader *cbor.Reader, attrNamesListToId *lru.Lru, projectionOrdinals []documentPath) (map[string]types.AttributeValue, error) {
	hdr, err := reader.PeekHeader()
	if err != nil {
		return nil, err
	}
	if hdr == cbor.Nil {
		if err = reader.ReadNil(); err != nil {
			return nil, err
		}
		return nil, nil
	}

	switch hdr & cbor.MajorTypeMask {
	case cbor.Bytes:
		r, err := reader.BytesReader()
		if err != nil {
			return nil, err
		}
		defer r.Close()
		item, err := cbor.DecodeItemNonKeyAttributes(ctx, r, attrNamesListToId)
		if err != nil {
			return nil, err
		}
		return item, nil
	case cbor.Map:
		return decodeProjection(reader, projectionOrdinals)
	}
	return nil, &smithy.SerializationError{Err: fmt.Errorf("unexpected cbor type %v", hdr)}

}

func decodeProjection(reader *cbor.Reader, projectionOrdinals []documentPath) (map[string]types.AttributeValue, error) {
	ib := &itemBuilder{}
	err := consumeMap(reader, func(ord int, r *cbor.Reader) error {
		if ord > len(projectionOrdinals) {
			return &smithy.SerializationError{Err: fmt.Errorf("unexpected ordinal %v", ord)}
		}
		p := projectionOrdinals[ord]
		v, err := cbor.DecodeAttributeValue(r)
		if err != nil {
			return err
		}
		ib.insert(p, v)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ib.toItem(), nil
}

func decodeAttributeProjection(ctx context.Context, reader *cbor.Reader, attrListIdToNames *lru.Lru) (map[string]types.AttributeValue, error) {
	r, err := reader.BytesReader()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	attrListId, err := r.ReadInt64()
	if err != nil {
		return nil, err
	}
	attrNames, err := attrListIdToNames.GetWithContext(ctx, attrListId)
	if err != nil {
		return nil, err
	}
	ans, ok := attrNames.([]string)
	if !ok {
		return nil, &smithy.SerializationError{Err: errors.New("invalid type for attribute names list")}
	}
	attrs := make(map[string]types.AttributeValue)
	err = consumeMap(r, func(ord int, reader *cbor.Reader) error {
		if ord > len(ans) {
			return &smithy.SerializationError{Err: errors.New("invalid ordinal")}
		}
		av, err := cbor.DecodeAttributeValue(reader)
		if err != nil {
			return err
		}
		attrs[ans[ord]] = av
		return nil
	})
	return attrs, nil
}

func decodeConsumedCapacity(reader *cbor.Reader) (*types.ConsumedCapacity, error) {
	consumed, err := consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if consumed {
		return nil, nil
	}

	if _, err := reader.ReadBytesLength(); err != nil {
		return nil, err
	}

	cc := &types.ConsumedCapacity{}

	t, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	cc.TableName = aws.String(t)

	c, err := reader.ReadFloat64()
	if err != nil {
		return nil, err
	}
	cc.CapacityUnits = aws.Float64(c)

	consumed, err = consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if !consumed {
		c, err = reader.ReadFloat64()
		if err != nil {
			return nil, err
		}
		cc.Table = &types.Capacity{
			CapacityUnits: aws.Float64(c),
		}
	}

	cc.GlobalSecondaryIndexes, err = decodeIndexConsumedCapacity(reader, false)
	if err != nil {
		return nil, err
	}
	cc.LocalSecondaryIndexes, err = decodeIndexConsumedCapacity(reader, false)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func decodeConsumedCapacityExtended(reader *cbor.Reader) (*types.ConsumedCapacity, error) {
	consumed, err := consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if consumed {
		return nil, nil
	}

	cc := &types.ConsumedCapacity{}
	err = consumeMap(reader, func(key int, reader *cbor.Reader) error {
		switch key {
		case tableName:
			s, err := reader.ReadString()
			if err != nil {
				return err
			}
			cc.TableName = &s
		case capacityUnits:
			f, err := reader.ReadFloat64()
			if err != nil {
				return err
			}
			cc.CapacityUnits = &f
		case readCapacityUnits:
			f, err := reader.ReadFloat64()
			if err != nil {
				return err
			}
			cc.ReadCapacityUnits = &f
		case writeCapacityUnits:
			f, err := reader.ReadFloat64()
			if err != nil {
				return err
			}
			cc.WriteCapacityUnits = &f
		case table:
			c, err := decodeCapacity(reader)
			if err != nil {
				return err
			}
			cc.Table = c
		case globalSecondaryIndexes:
			c, err := decodeIndexConsumedCapacity(reader, true)
			if err != nil {
				return err
			}
			cc.GlobalSecondaryIndexes = c
		case localSecondaryIndexes:
			c, err := decodeIndexConsumedCapacity(reader, true)
			if err != nil {
				return err
			}
			cc.LocalSecondaryIndexes = c
		default:
			return &smithy.SerializationError{Err: fmt.Errorf("unknown response param key %d", key)}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func decodeCapacity(reader *cbor.Reader) (*types.Capacity, error) {
	consumed, err := consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if consumed {
		return nil, nil
	}

	c := &types.Capacity{}
	err = consumeMap(reader, func(key int, reader *cbor.Reader) error {
		switch key {
		case capacityUnits:
			f, err := reader.ReadFloat64()
			if err != nil {
				return err
			}
			c.CapacityUnits = &f
		case readCapacityUnits:
			f, err := reader.ReadFloat64()
			if err != nil {
				return err
			}
			c.ReadCapacityUnits = &f
		case writeCapacityUnits:
			f, err := reader.ReadFloat64()
			if err != nil {
				return err
			}
			c.WriteCapacityUnits = &f
		default:
			return &smithy.SerializationError{Err: fmt.Errorf("unknown response param key %d", key)}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}

func decodeIndexConsumedCapacity(reader *cbor.Reader, extended bool) (map[string]types.Capacity, error) {
	consumed, err := consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if consumed {
		return nil, nil
	}

	len, err := reader.ReadMapLength()
	if err != nil {
		return nil, err
	}
	index := make(map[string]types.Capacity, len)
	for len > 0 {
		len--
		i, err := reader.ReadString()
		if err != nil {
			return nil, err
		}
		var c *types.Capacity
		if extended {
			c, err = decodeCapacity(reader)
			if err != nil {
				return nil, err
			}
		} else {
			f, err := reader.ReadFloat64()
			if err != nil {
				return nil, err
			}
			c = &types.Capacity{
				CapacityUnits: aws.Float64(f),
			}
		}
		index[i] = *c
	}
	return index, nil
}

func decodeItemCollectionMetrics(reader *cbor.Reader, partitionKey string) (*types.ItemCollectionMetrics, error) {
	consumed, err := consumeNil(reader)
	if err != nil {
		return nil, err
	}
	if consumed {
		return nil, nil
	}

	if _, err := reader.ReadBytesLength(); err != nil {
		return nil, err
	}

	av, err := cbor.DecodeAttributeValue(reader)
	if err != nil {
		return nil, err
	}
	l, err := reader.ReadFloat64()
	if err != nil {
		return nil, err
	}
	r, err := reader.ReadFloat64()
	if err != nil {
		return nil, err
	}

	icm := types.ItemCollectionMetrics{
		ItemCollectionKey: map[string]types.AttributeValue{
			partitionKey: av,
		},
		SizeEstimateRangeGB: []float64{l, r},
	}
	return &icm, nil
}

func getKeySchema(ctx context.Context, keySchemaCache *lru.Lru, table string) ([]types.AttributeDefinition, error) {
	k, err := keySchemaCache.GetWithContext(ctx, table)
	if err != nil {
		return nil, err
	}
	keys, ok := k.([]types.AttributeDefinition)
	if !ok {
		return nil, &smithy.SerializationError{Err: errors.New("invalid type for keyschema")}
	}
	return keys, nil
}
