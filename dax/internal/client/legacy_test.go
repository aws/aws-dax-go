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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
	"testing"
)

var functions = map[reflect.Type]interface{}{
	reflect.TypeOf(&dynamodb.GetItemInput{}):      translateLegacyGetItemInput,
	reflect.TypeOf(&dynamodb.PutItemInput{}):      translateLegacyPutItemInput,
	reflect.TypeOf(&dynamodb.DeleteItemInput{}):   translateLegacyDeleteItemInput,
	reflect.TypeOf(&dynamodb.UpdateItemInput{}):   translateLegacyUpdateItemInput,
	reflect.TypeOf(&dynamodb.ScanInput{}):         translateLegacyScanInput,
	reflect.TypeOf(&dynamodb.QueryInput{}):        translateLegacyQueryInput,
	reflect.TypeOf(&dynamodb.BatchGetItemInput{}): translateLegacyBatchGetItemInput,
}

func TestTranslateLegacyPositive(t *testing.T) {
	cases := []struct {
		inp interface{}
		exp interface{}
	}{
		{
			&dynamodb.GetItemInput{
				AttributesToGet: []*string{aws.String("a1"), aws.String("a2")},
			},
			&dynamodb.GetItemInput{
				ProjectionExpression:     aws.String("#key0,#key1"),
				ExpressionAttributeNames: map[string]*string{"#key0": aws.String("a1"), "#key1": aws.String("a2")},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(true), Value: &dynamodb.AttributeValue{N: aws.String("5")}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("#key0 = :val0"),
				ExpressionAttributeNames:  map[string]*string{"#key0": aws.String("a")},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":val0": {N: aws.String("5")}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(false)},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("attribute_not_exists(#key0)"),
				ExpressionAttributeNames: map[string]*string{"#key0": aws.String("a")},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorBetween),
						AttributeValueList: []*dynamodb.AttributeValue{{N: aws.String("5")}, {N: aws.String("6")}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("#key0 between :val0 and :val1"),
				ExpressionAttributeNames:  map[string]*string{"#key0": aws.String("a")},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":val0": {N: aws.String("5")}, ":val1": {N: aws.String("6")}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorBeginsWith),
						AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String("abc")}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("begins_with(#key0,:val0)"),
				ExpressionAttributeNames:  map[string]*string{"#key0": aws.String("a")},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":val0": {S: aws.String("abc")}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorContains),
						AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String("abc")}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("contains(#key0,:val0)"),
				ExpressionAttributeNames:  map[string]*string{"#key0": aws.String("a")},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":val0": {S: aws.String("abc")}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorNotContains),
						AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String("abc")}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("not contains(#key0,:val0)"),
				ExpressionAttributeNames:  map[string]*string{"#key0": aws.String("a")},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":val0": {S: aws.String("abc")}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorNull)},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("attribute_not_exists(#key0)"),
				ExpressionAttributeNames: map[string]*string{"#key0": aws.String("a")},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorNotNull)},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("attribute_exists(#key0)"),
				ExpressionAttributeNames: map[string]*string{"#key0": aws.String("a")},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorIn),
						AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String("abc")}, {S: aws.String("def")}, {S: aws.String("ghi")}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("#key0 in (:val0,:val1,:val2)"),
				ExpressionAttributeNames:  map[string]*string{"#key0": aws.String("a")},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":val0": {S: aws.String("abc")}, ":val1": {S: aws.String("def")}, ":val2": {S: aws.String("ghi")}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorNe),
						AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String("abc")}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("#key0 <> :val0"),
				ExpressionAttributeNames:  map[string]*string{"#key0": aws.String("a")},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":val0": {S: aws.String("abc")}},
			},
		},
		{
			&dynamodb.DeleteItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq),
						AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String("abc")}}},
				},
			},
			&dynamodb.DeleteItemInput{
				ConditionExpression:       aws.String("#key0 = :val0"),
				ExpressionAttributeNames:  map[string]*string{"#key0": aws.String("a")},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":val0": {S: aws.String("abc")}},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorLe),
						AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String("abc")}}},
				},
				AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
					"b": {Value: &dynamodb.AttributeValue{S: aws.String("def")}},
				},
			},
			&dynamodb.UpdateItemInput{
				ConditionExpression: aws.String("#key0 <= :val0"),
				UpdateExpression:    aws.String("set #key1=:val1"),
				ExpressionAttributeNames: map[string]*string{
					"#key0": aws.String("a"),
					"#key1": aws.String("b"),
				},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":val0": {S: aws.String("abc")},
					":val1": {S: aws.String("def")},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
					"a": {Action: aws.String(dynamodb.AttributeActionPut), Value: &dynamodb.AttributeValue{S: aws.String("def")}},
				},
			},
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("set #key0=:val0"),
				ExpressionAttributeNames: map[string]*string{
					"#key0": aws.String("a"),
				},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":val0": {S: aws.String("def")},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
					"a": {Action: aws.String(dynamodb.AttributeActionAdd), Value: &dynamodb.AttributeValue{S: aws.String("def")}},
				},
			},
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("add #key0 :val0"),
				ExpressionAttributeNames: map[string]*string{
					"#key0": aws.String("a"),
				},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":val0": {S: aws.String("def")},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
					"a": {Action: aws.String(dynamodb.AttributeActionDelete), Value: &dynamodb.AttributeValue{S: aws.String("def")}},
				},
			},
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("delete #key0 :val0"),
				ExpressionAttributeNames: map[string]*string{
					"#key0": aws.String("a"),
				},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":val0": {S: aws.String("def")},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
					"a": {Action: aws.String(dynamodb.AttributeActionDelete)},
				},
			},
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("remove #key0"),
				ExpressionAttributeNames: map[string]*string{
					"#key0": aws.String("a"),
				},
			},
		},
		{
			&dynamodb.ScanInput{
				AttributesToGet:     []*string{aws.String("a1"), aws.String("a2")},
				ConditionalOperator: aws.String(dynamodb.ConditionalOperatorOr),
				ScanFilter: map[string]*dynamodb.Condition{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorGe), AttributeValueList: []*dynamodb.AttributeValue{{N: aws.String("5")}}},
				},
			},
			&dynamodb.ScanInput{
				ProjectionExpression: aws.String("#key0,#key1"),
				FilterExpression:     aws.String("#key2 >= :val0"),
				ExpressionAttributeNames: map[string]*string{
					"#key0": aws.String("a1"),
					"#key1": aws.String("a2"),
					"#key2": aws.String("a"),
				},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":val0": {N: aws.String("5")},
				},
			},
		},
		{
			&dynamodb.QueryInput{
				AttributesToGet:     []*string{aws.String("a1"), aws.String("a2")},
				ConditionalOperator: aws.String(dynamodb.ConditionalOperatorOr),
				QueryFilter: map[string]*dynamodb.Condition{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorGe), AttributeValueList: []*dynamodb.AttributeValue{{N: aws.String("5")}}},
				},
				KeyConditions: map[string]*dynamodb.Condition{
					"k": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String("abc")}}},
				},
			},
			&dynamodb.QueryInput{
				ProjectionExpression:   aws.String("#key0,#key1"),
				FilterExpression:       aws.String("#key2 >= :val0"),
				KeyConditionExpression: aws.String("#key3 = :val1"),
				ExpressionAttributeNames: map[string]*string{
					"#key0": aws.String("a1"),
					"#key1": aws.String("a2"),
					"#key2": aws.String("a"),
					"#key3": aws.String("k"),
				},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":val0": {N: aws.String("5")},
					":val1": {S: aws.String("abc")},
				},
			},
		},
		{
			&dynamodb.BatchGetItemInput{
				RequestItems: map[string]*dynamodb.KeysAndAttributes{
					"table1": {AttributesToGet: []*string{aws.String("a1"), aws.String("a2")}},
					"table2": {AttributesToGet: []*string{aws.String("a3"), aws.String("a4")}},
				},
			},
			&dynamodb.BatchGetItemInput{
				RequestItems: map[string]*dynamodb.KeysAndAttributes{
					"table1": {
						AttributesToGet:      []*string{aws.String("a1"), aws.String("a2")},
						ProjectionExpression: aws.String("#key0,#key1"),
						ExpressionAttributeNames: map[string]*string{
							"#key0": aws.String("a1"),
							"#key1": aws.String("a2"),
						},
					},
					"table2": {
						AttributesToGet:      []*string{aws.String("a3"), aws.String("a4")},
						ProjectionExpression: aws.String("#key0,#key1"),
						ExpressionAttributeNames: map[string]*string{
							"#key0": aws.String("a3"),
							"#key1": aws.String("a4"),
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		fn := reflect.ValueOf(functions[reflect.TypeOf(c.inp)])
		out := fn.Call([]reflect.Value{reflect.ValueOf(c.inp)})
		act, err := out[0].Interface(), out[1].Interface()
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if !reflect.DeepEqual(c.exp, act) {
			t.Errorf("expected %v, got %v", c.exp, act)
		}
	}
}

func TestTranslateLegacyNegative(t *testing.T) {
	cases := []struct {
		inp interface{}
		err awserr.Error
	}{
		{
			&dynamodb.GetItemInput{
				AttributesToGet:      []*string{aws.String("a1"), aws.String("a2")},
				ProjectionExpression: aws.String("a1, a2"),
			},
			awserr.New(ErrCodeValidationException, "Cannot specify both AttributesToGet and ProjectionExpression", nil),
		},
		{
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("a < :v"),
				Expected:                  map[string]*dynamodb.ExpectedAttributeValue{"a": {Exists: aws.Bool(true), Value: &dynamodb.AttributeValue{N: aws.String("5")}}},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":v": {N: aws.String("5")}},
			},
			awserr.New(ErrCodeValidationException, "Cannot specify both Expected and ConditionExpression", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {
						Value:              &dynamodb.AttributeValue{N: aws.String("5")},
						AttributeValueList: []*dynamodb.AttributeValue{{N: aws.String("5")}},
					},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: Value and AttributeValueList cannot be used together for Attribute: a", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {AttributeValueList: []*dynamodb.AttributeValue{{N: aws.String("5")}}},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: AttributeValueList can only be used with a ComparisonOperator for Attribute: a", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(true)},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: Value must be provided when Exists is true for Attribute: a", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(false), Value: &dynamodb.AttributeValue{N: aws.String("5")}},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: Value cannot be used when Exists is false for Attribute: a", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]*dynamodb.ExpectedAttributeValue{
					"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorBetween), AttributeValueList: []*dynamodb.AttributeValue{{N: aws.String("5")}, {NULL: aws.Bool(true)}}},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: ComparisonOperator BETWEEN is not valid for NULL AttributeValue type", nil),
		},
		{
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("a < :v"),
				AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
					"a": {Action: aws.String(dynamodb.AttributeActionDelete)},
				},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":v": {N: aws.String("5")}},
			},
			awserr.New(ErrCodeValidationException, "Cannot specify both AttributeUpdates and UpdateExpression", nil),
		},
		{
			&dynamodb.UpdateItemInput{
				ConditionExpression:       aws.String("a < :v"),
				Expected:                  map[string]*dynamodb.ExpectedAttributeValue{"a": {Exists: aws.Bool(true), Value: &dynamodb.AttributeValue{N: aws.String("5")}}},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":v": {N: aws.String("5")}},
			},
			awserr.New(ErrCodeValidationException, "Cannot specify both Expected and ConditionExpression", nil),
		},
		{
			&dynamodb.ScanInput{
				ScanFilter: map[string]*dynamodb.Condition{"a": {}},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: AttributeValueList can only be used with a ComparisonOperator for Attribute: a", nil),
		},
		{
			&dynamodb.QueryInput{
				KeyConditions: map[string]*dynamodb.Condition{"a": {ComparisonOperator: aws.String(dynamodb.ComparisonOperatorContains), AttributeValueList: []*dynamodb.AttributeValue{{N: aws.String("5")}}}},
			},
			awserr.New(ErrCodeValidationException, "Unsupported operator on KeyCondition: CONTAINS", nil),
		},
	}

	for _, c := range cases {
		fn := reflect.ValueOf(functions[reflect.TypeOf(c.inp)])
		out := fn.Call([]reflect.Value{reflect.ValueOf(c.inp)})
		_, err := out[0].Interface(), out[1].Interface()
		if !reflect.DeepEqual(c.err, err) {
			t.Errorf("expected %v, got %v", c.err, err)
		}
	}
}
