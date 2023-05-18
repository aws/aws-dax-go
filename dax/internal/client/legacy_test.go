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
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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
				AttributesToGet: []string{"a1", "a2"},
			},
			&dynamodb.GetItemInput{
				ProjectionExpression:     aws.String("#key0,#key1"),
				ExpressionAttributeNames: map[string]string{"#key0": "a1", "#key1": "a2"},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(true), Value: &types.AttributeValueMemberN{Value: "5"}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:       aws.String("#key0 = :val0"),
				ExpressionAttributeNames:  map[string]string{"#key0": "a"},
				ExpressionAttributeValues: map[string]types.AttributeValue{":val0": &types.AttributeValueMemberN{Value: "5"}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(false)},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("attribute_not_exists(#key0)"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorBetween,
						AttributeValueList: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "5"},
							&types.AttributeValueMemberN{Value: "6"},
						}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("#key0 between :val0 and :val1"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberN{Value: "5"},
					":val1": &types.AttributeValueMemberN{Value: "6"}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorBeginsWith,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberS{Value: "abc"}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("begins_with(#key0,:val0)"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "abc"},
				},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorContains,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberS{Value: "abc"}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("contains(#key0,:val0)"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "abc"},
				},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorNotContains,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberS{Value: "abc"}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("not contains(#key0,:val0)"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "abc"},
				},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorNull},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("attribute_not_exists(#key0)"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorNotNull},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("attribute_exists(#key0)"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorIn,
						AttributeValueList: []types.AttributeValue{
							&types.AttributeValueMemberS{Value: "abc"},
							&types.AttributeValueMemberS{Value: "def"},
							&types.AttributeValueMemberS{Value: "ghi"},
						},
					},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("#key0 in (:val0,:val1,:val2)"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "abc"},
					":val1": &types.AttributeValueMemberS{Value: "def"},
					":val2": &types.AttributeValueMemberS{Value: "ghi"}},
			},
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorNe,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberS{Value: "abc"}}},
				},
			},
			&dynamodb.PutItemInput{
				ConditionExpression:      aws.String("#key0 <> :val0"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "abc"},
				},
			},
		},
		{
			&dynamodb.DeleteItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorEq,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberS{Value: "abc"}}},
				},
			},
			&dynamodb.DeleteItemInput{
				ConditionExpression:      aws.String("#key0 = :val0"),
				ExpressionAttributeNames: map[string]string{"#key0": "a"},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "abc"},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {ComparisonOperator: types.ComparisonOperatorLe,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberS{Value: "abc"}}},
				},
				AttributeUpdates: map[string]types.AttributeValueUpdate{
					"b": {Value: &types.AttributeValueMemberS{Value: "def"}},
				},
			},
			&dynamodb.UpdateItemInput{
				ConditionExpression: aws.String("#key0 <= :val0"),
				UpdateExpression:    aws.String("set #key1=:val1"),
				ExpressionAttributeNames: map[string]string{
					"#key0": "a",
					"#key1": "b",
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "abc"},
					":val1": &types.AttributeValueMemberS{Value: "def"},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				AttributeUpdates: map[string]types.AttributeValueUpdate{
					"a": {Action: types.AttributeActionPut, Value: &types.AttributeValueMemberS{Value: "def"}},
				},
			},
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("set #key0=:val0"),
				ExpressionAttributeNames: map[string]string{
					"#key0": "a",
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "def"},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				AttributeUpdates: map[string]types.AttributeValueUpdate{
					"a": {
						Action: types.AttributeActionAdd,
						Value:  &types.AttributeValueMemberS{Value: "def"}},
				},
			},
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("add #key0 :val0"),
				ExpressionAttributeNames: map[string]string{
					"#key0": "a",
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "def"},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				AttributeUpdates: map[string]types.AttributeValueUpdate{
					"a": {
						Action: types.AttributeActionDelete,
						Value:  &types.AttributeValueMemberS{Value: "def"}},
				},
			},
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("delete #key0 :val0"),
				ExpressionAttributeNames: map[string]string{
					"#key0": "a",
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberS{Value: "def"},
				},
			},
		},
		{
			&dynamodb.UpdateItemInput{
				AttributeUpdates: map[string]types.AttributeValueUpdate{
					"a": {Action: types.AttributeActionDelete},
				},
			},
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("remove #key0"),
				ExpressionAttributeNames: map[string]string{
					"#key0": "a",
				},
			},
		},
		{
			&dynamodb.ScanInput{
				AttributesToGet:     []string{"a1", "a2"},
				ConditionalOperator: types.ConditionalOperatorOr,
				ScanFilter: map[string]types.Condition{
					"a": {
						ComparisonOperator: types.ComparisonOperatorGe,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberN{Value: "5"}}},
				},
			},
			&dynamodb.ScanInput{
				ProjectionExpression: aws.String("#key0,#key1"),
				FilterExpression:     aws.String("#key2 >= :val0"),
				ExpressionAttributeNames: map[string]string{
					"#key0": "a1",
					"#key1": "a2",
					"#key2": "a",
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberN{Value: "5"},
				},
			},
		},
		{
			&dynamodb.QueryInput{
				AttributesToGet:     []string{"a1", "a2"},
				ConditionalOperator: types.ConditionalOperatorOr,
				QueryFilter: map[string]types.Condition{
					"a": {
						ComparisonOperator: types.ComparisonOperatorGe,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberN{Value: "5"}},
					},
				},
				KeyConditions: map[string]types.Condition{
					"k": {
						ComparisonOperator: types.ComparisonOperatorEq,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberS{Value: "abc"}},
					},
				},
			},
			&dynamodb.QueryInput{
				ProjectionExpression:   aws.String("#key0,#key1"),
				FilterExpression:       aws.String("#key2 >= :val0"),
				KeyConditionExpression: aws.String("#key3 = :val1"),
				ExpressionAttributeNames: map[string]string{
					"#key0": "a1",
					"#key1": "a2",
					"#key2": "a",
					"#key3": "k",
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":val0": &types.AttributeValueMemberN{Value: "5"},
					":val1": &types.AttributeValueMemberS{Value: "abc"},
				},
			},
		},
		{
			inp: &dynamodb.BatchGetItemInput{
				RequestItems: map[string]types.KeysAndAttributes{
					"table1": {AttributesToGet: []string{"a1", "a2"}},
					"table2": {AttributesToGet: []string{"a3", "a4"}},
				},
			},
			exp: &dynamodb.BatchGetItemInput{
				RequestItems: map[string]types.KeysAndAttributes{
					"table1": {
						AttributesToGet:      []string{"a1", "a2"},
						ProjectionExpression: aws.String("#key0,#key1"),
						ExpressionAttributeNames: map[string]string{
							"#key0": "a1",
							"#key1": "a2",
						},
					},
					"table2": {
						AttributesToGet:      []string{"a3", "a4"},
						ProjectionExpression: aws.String("#key0,#key1"),
						ExpressionAttributeNames: map[string]string{
							"#key0": "a3",
							"#key1": "a4",
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
				AttributesToGet:      []string{"a1", "a2"},
				ProjectionExpression: aws.String("a1, a2"),
			},
			awserr.New(ErrCodeValidationException, "Cannot specify both AttributesToGet and ProjectionExpression", nil),
		},
		{
			&dynamodb.PutItemInput{
				ConditionExpression: aws.String("a < :v"),
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(true), Value: &types.AttributeValueMemberN{Value: "5"}},
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{":v": &types.AttributeValueMemberN{Value: "5"}},
			},
			awserr.New(ErrCodeValidationException, "Cannot specify both Expected and ConditionExpression", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {
						Value:              &types.AttributeValueMemberN{Value: "5"},
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberN{Value: "5"}},
					},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: Value and AttributeValueList cannot be used together for Attribute: a", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberN{Value: "5"}}},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: AttributeValueList can only be used with a ComparisonOperator for Attribute: a", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(true)},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: Value must be provided when Exists is true for Attribute: a", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {Exists: aws.Bool(false), Value: &types.AttributeValueMemberN{Value: "5"}},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: Value cannot be used when Exists is false for Attribute: a", nil),
		},
		{
			&dynamodb.PutItemInput{
				Expected: map[string]types.ExpectedAttributeValue{
					"a": {
						ComparisonOperator: types.ComparisonOperatorBetween,
						AttributeValueList: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "5"},
							&types.AttributeValueMemberNULL{Value: true}}},
				},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: ComparisonOperator BETWEEN is not valid for NULL AttributeValue type", nil),
		},
		{
			&dynamodb.UpdateItemInput{
				UpdateExpression: aws.String("a < :v"),
				AttributeUpdates: map[string]types.AttributeValueUpdate{
					"a": {Action: types.AttributeActionDelete},
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{":v": &types.AttributeValueMemberN{Value: "5"}},
			},
			awserr.New(ErrCodeValidationException, "Cannot specify both AttributeUpdates and UpdateExpression", nil),
		},
		{
			&dynamodb.UpdateItemInput{
				ConditionExpression:       aws.String("a < :v"),
				Expected:                  map[string]types.ExpectedAttributeValue{"a": {Exists: aws.Bool(true), Value: &types.AttributeValueMemberN{Value: "5"}}},
				ExpressionAttributeValues: map[string]types.AttributeValue{":v": &types.AttributeValueMemberN{Value: "5"}},
			},
			awserr.New(ErrCodeValidationException, "Cannot specify both Expected and ConditionExpression", nil),
		},
		{
			&dynamodb.ScanInput{
				ScanFilter: map[string]types.Condition{"a": {}},
			},
			awserr.New(ErrCodeValidationException, "One or more parameter values were invalid: AttributeValueList can only be used with a ComparisonOperator for Attribute: a", nil),
		},
		{
			&dynamodb.QueryInput{
				KeyConditions: map[string]types.Condition{
					"a": {
						ComparisonOperator: types.ComparisonOperatorContains,
						AttributeValueList: []types.AttributeValue{&types.AttributeValueMemberN{Value: "5"}}}},
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
