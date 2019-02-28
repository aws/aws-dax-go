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
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	attributeNamesKeyPrefix  = "#key"
	attributeValuesKeyPrefix = ":val"
)

func translateLegacyGetItemInput(input *dynamodb.GetItemInput) (*dynamodb.GetItemInput, error) {
	f, err := hasAttributesToGet(input.AttributesToGet, input.ProjectionExpression)
	if err != nil || !f {
		return input, err
	}

	output := input
	output.ProjectionExpression, output.ExpressionAttributeNames, err = translateAttributesToGet(input.AttributesToGet, input.ExpressionAttributeNames)
	output.AttributesToGet = nil
	if err != nil {
		return input, err
	}
	return output, err
}

func translateLegacyPutItemInput(input *dynamodb.PutItemInput) (*dynamodb.PutItemInput, error) {
	f, err := hasExpected(input.Expected, input.ConditionExpression)
	if err != nil || !f {
		return input, err
	}

	output := input
	output.ConditionExpression, output.ExpressionAttributeNames, output.ExpressionAttributeValues, err =
		translateExpected(output.ConditionalOperator, output.Expected, output.ExpressionAttributeNames, output.ExpressionAttributeValues)
	if err != nil {
		return input, err
	}
	output.ConditionalOperator = nil
	output.Expected = nil
	return output, err
}

func translateLegacyDeleteItemInput(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemInput, error) {
	f, err := hasExpected(input.Expected, input.ConditionExpression)
	if err != nil || !f {
		return input, err
	}

	output := input
	output.ConditionExpression, output.ExpressionAttributeNames, output.ExpressionAttributeValues, err =
		translateExpected(output.ConditionalOperator, output.Expected, input.ExpressionAttributeNames, output.ExpressionAttributeValues)
	output.ConditionalOperator = nil
	output.Expected = nil
	if err != nil {
		return input, err
	}
	return output, err
}

func translateLegacyUpdateItemInput(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemInput, error) {
	cf, err := hasExpected(input.Expected, input.ConditionExpression)
	if err != nil {
		return input, err
	}
	uf, err := hasAttributeUpdates(input.AttributeUpdates, input.UpdateExpression)
	if err != nil {
		return input, err
	}
	if !uf && !cf {
		return input, nil
	}

	output := input
	if cf {
		output.ConditionExpression, output.ExpressionAttributeNames, output.ExpressionAttributeValues, err =
			translateExpected(output.ConditionalOperator, output.Expected, output.ExpressionAttributeNames, output.ExpressionAttributeValues)
		if err != nil {
			return input, err
		}
		output.ConditionalOperator = nil
		output.Expected = nil
	}
	if uf {
		output.UpdateExpression, output.ExpressionAttributeNames, output.ExpressionAttributeValues, err =
			translateAttributeUpdates(output.AttributeUpdates, output.ExpressionAttributeNames, output.ExpressionAttributeValues)
		if err != nil {
			return input, err
		}
		output.AttributeUpdates = nil
	}
	return output, nil
}

func translateLegacyScanInput(input *dynamodb.ScanInput) (*dynamodb.ScanInput, error) {
	pf, err := hasAttributesToGet(input.AttributesToGet, input.ProjectionExpression)
	if err != nil {
		return input, err
	}
	cf, err := hasFilter(input.ScanFilter, input.FilterExpression)
	if err != nil {
		return input, err
	}
	if !pf && !cf {
		return input, nil
	}

	output := input
	if pf {
		output.ProjectionExpression, output.ExpressionAttributeNames, err =
			translateAttributesToGet(output.AttributesToGet, output.ExpressionAttributeNames)
		if err != nil {
			return input, err
		}
		output.AttributesToGet = nil
	}
	if cf {
		output.FilterExpression, output.ExpressionAttributeNames, output.ExpressionAttributeValues, err =
			translateFilter(output.ConditionalOperator, output.ScanFilter, output.ExpressionAttributeNames, output.ExpressionAttributeValues, false)
		if err != nil {
			return input, err
		}
		output.ConditionalOperator = nil
		output.ScanFilter = nil
	}

	return output, nil
}

func translateLegacyQueryInput(input *dynamodb.QueryInput) (*dynamodb.QueryInput, error) {
	pf, err := hasAttributesToGet(input.AttributesToGet, input.ProjectionExpression)
	if err != nil {
		return input, err
	}
	ff, err := hasFilter(input.QueryFilter, input.FilterExpression)
	if err != nil {
		return input, err
	}
	kf, err := hasFilter(input.KeyConditions, input.KeyConditionExpression)
	if err != nil {
		return input, err
	}
	if !pf && !ff && !kf {
		return input, nil
	}

	output := input
	if pf {
		output.ProjectionExpression, output.ExpressionAttributeNames, err =
			translateAttributesToGet(output.AttributesToGet, output.ExpressionAttributeNames)
		if err != nil {
			return input, err
		}
		output.AttributesToGet = nil
	}
	if ff {
		output.FilterExpression, output.ExpressionAttributeNames, output.ExpressionAttributeValues, err =
			translateFilter(output.ConditionalOperator, output.QueryFilter, output.ExpressionAttributeNames, output.ExpressionAttributeValues, false)
		if err != nil {
			return input, err
		}
		output.ConditionalOperator = nil
		output.QueryFilter = nil
	}
	if kf {
		output.KeyConditionExpression, output.ExpressionAttributeNames, output.ExpressionAttributeValues, err =
			translateFilter(aws.String(dynamodb.ConditionalOperatorAnd), output.KeyConditions, output.ExpressionAttributeNames, output.ExpressionAttributeValues, true)
		if err != nil {
			return input, err
		}
		output.KeyConditions = nil
	}

	return output, nil
}

func translateLegacyBatchGetItemInput(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemInput, error) {
	if len(input.RequestItems) == 0 {
		return input, nil
	}

	for _, kaas := range input.RequestItems {
		f, err := hasAttributesToGet(kaas.AttributesToGet, kaas.ProjectionExpression)
		if err != nil {
			return input, err
		}
		if !f {
			continue
		}
		kaas.ProjectionExpression, kaas.ExpressionAttributeNames, err = translateAttributesToGet(kaas.AttributesToGet, kaas.ExpressionAttributeNames)
		if err != nil {
			return input, err
		}
	}
	return input, nil
}

func hasAttributesToGet(a []*string, p *string) (bool, error) {
	af := len(a) != 0
	pf := p != nil
	if af && pf {
		return false, awserr.New(ErrCodeValidationException, "Cannot specify both AttributesToGet and ProjectionExpression", nil)
	}
	return af, nil
}

func hasExpected(e map[string]*dynamodb.ExpectedAttributeValue, c *string) (bool, error) {
	ef := len(e) != 0
	cf := c != nil
	if ef && cf {
		return false, awserr.New(ErrCodeValidationException, "Cannot specify both Expected and ConditionExpression", nil)
	}
	return ef, nil
}

func hasAttributeUpdates(u map[string]*dynamodb.AttributeValueUpdate, e *string) (bool, error) {
	uf := len(u) > 0
	ef := e != nil
	if uf && ef {
		return false, awserr.New(ErrCodeValidationException, "Cannot specify both AttributeUpdates and UpdateExpression", nil)
	}
	return uf, nil
}

func hasFilter(c map[string]*dynamodb.Condition, e *string) (bool, error) {
	cf := len(c) > 0
	ef := e != nil
	if cf && ef {
		return false, awserr.New(ErrCodeValidationException, "Cannot specify both [Scan|Query]Filter and [Scan|Query]FilterExpression", nil)
	}
	return cf, nil
}

func translateAttributesToGet(attrs []*string, subs map[string]*string) (*string, map[string]*string, error) {
	out, sub := appendAttributeNames(nil, attrs, subs)
	return aws.String(string(out)), sub, nil
}

func translateExpected(o *string, e map[string]*dynamodb.ExpectedAttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) (*string, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	op := dynamodb.ConditionalOperatorAnd
	if o != nil && len(*o) > 0 {
		op = *o
	}
	ops := fmt.Sprintf(" %s ", strings.TrimSpace(op))

	var out []byte
	var err error
	f := true
	for k, v := range e {
		if f {
			f = false
		} else {
			out = append(out, []byte(ops)...)
		}
		out, subs, vars, err = appendCondition(out, k, v, subs, vars)
		if err != nil {
			return nil, subs, vars, err
		}
	}
	return aws.String(string(out)), subs, vars, err
}

func translateFilter(o *string, c map[string]*dynamodb.Condition, subs map[string]*string, vars map[string]*dynamodb.AttributeValue, keyCondition bool) (*string, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	op := dynamodb.ConditionalOperatorAnd
	if o != nil && len(*o) > 0 {
		op = *o
	}
	ops := fmt.Sprintf(" %s ", strings.TrimSpace(op))

	var out []byte
	var err error
	f := true
	for k, v := range c {
		if f {
			f = false
		} else {
			out = append(out, []byte(ops)...)
		}
		out, subs, vars, err = appendFilterCondition(out, k, v, subs, vars, keyCondition)
		if err != nil {
			return nil, subs, vars, err
		}
	}
	return aws.String(string(out)), subs, vars, err
}

func translateAttributeUpdates(avus map[string]*dynamodb.AttributeValueUpdate, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) (*string, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	var sets, adds, dels, rems []string

	for a, avu := range avus {
		if avu == nil {
			continue
		}
		act := dynamodb.AttributeActionPut
		if avu.Action != nil {
			act = *avu.Action
		}
		if avu.Value == nil && act != dynamodb.AttributeActionDelete {
			return nil, subs, vars, awserr.New(ErrCodeValidationException, "only DELETE action is allowed when no attribute value is specified", nil)
		}

		var an, av string
		subs, an = appendAttributeName(subs, a)
		if avu.Value != nil {
			vars, av = appendAttributeValue(vars, *avu.Value)
		}

		switch act {
		case dynamodb.AttributeActionPut:
			sets = append(sets, fmt.Sprintf("%s=%s", an, av))
		case dynamodb.AttributeActionAdd:
			adds = append(adds, fmt.Sprintf("%s %s", an, av))
		case dynamodb.AttributeActionDelete:
			if len(av) == 0 {
				rems = append(rems, an)
			} else {
				dels = append(dels, fmt.Sprintf("%s %s", an, av))
			}
		default:
			return nil, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("unknown AttributeValueUpdate Action: %s", act), nil)
		}
	}

	var all []string
	if len(sets) > 0 {
		all = append(all, fmt.Sprintf("set %s", strings.Join(sets, ",")))
	}
	if len(adds) > 0 {
		all = append(all, fmt.Sprintf("add %s", strings.Join(adds, ",")))
	}
	if len(dels) > 0 {
		all = append(all, fmt.Sprintf("delete %s", strings.Join(dels, ",")))
	}
	if len(rems) > 0 {
		all = append(all, fmt.Sprintf("remove %s", strings.Join(rems, ",")))
	}
	return aws.String(strings.Join(all, " ")), subs, vars, nil
}

func appendAttributeNames(in []byte, attrs []*string, sub map[string]*string) ([]byte, map[string]*string) {
	f := true
	for _, a := range attrs {
		if a == nil || len(*a) == 0 {
			continue
		}
		if f {
			f = false
		} else {
			in = append(in, []byte(",")...)
		}
		var an string
		sub, an = appendAttributeName(sub, *a)
		in = append(in, []byte(an)...)
	}
	return in, sub
}

func appendAttributeValues(in []byte, avl []*dynamodb.AttributeValue, vars map[string]*dynamodb.AttributeValue) ([]byte, map[string]*dynamodb.AttributeValue) {
	f := true
	for _, v := range avl {
		if v == nil {
			continue
		}
		if f {
			f = false
		} else {
			in = append(in, []byte(",")...)
		}
		var av string
		vars, av = appendAttributeValue(vars, *v)
		in = append(in, []byte(av)...)
	}
	return in, vars
}

func appendCondition(in []byte, a string, eav *dynamodb.ExpectedAttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	if eav == nil {
		return in, subs, vars, nil
	}
	if eav.Value != nil && len(eav.AttributeValueList) > 0 {
		return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: Value and AttributeValueList cannot be used together for Attribute: %s", a), nil)
	}

	op := eav.ComparisonOperator
	if op == nil || len(*op) == 0 {
		return appendExistsCondition(in, a, eav, subs, vars)
	} else if eav.Exists != nil && *eav.Exists {
		return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: Exists and ComparisonOperator cannot be used together for Attribute: %s", a), nil)
	}
	avl := eav.AttributeValueList
	if len(avl) == 0 && eav.Value != nil {
		avl = []*dynamodb.AttributeValue{eav.Value}
	}
	return appendComparisonCondition(in, a, *op, avl, subs, vars, false)
}

func appendFilterCondition(in []byte, a string, c *dynamodb.Condition, subs map[string]*string, vars map[string]*dynamodb.AttributeValue, keyCondition bool) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	if c == nil {
		if keyCondition {
			return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("KeyCondition cannot be nil for key: %s", a), nil)
		}
		return in, subs, vars, nil
	}
	op := c.ComparisonOperator
	if op == nil || len(*op) == 0 {
		return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: AttributeValueList can only be used with a ComparisonOperator for Attribute: %s", a), nil)
	}
	return appendComparisonCondition(in, a, *op, c.AttributeValueList, subs, vars, keyCondition)
}

func appendExistsCondition(in []byte, a string, eav *dynamodb.ExpectedAttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	if len(eav.AttributeValueList) != 0 {
		return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: AttributeValueList can only be used with a ComparisonOperator for Attribute: %s", a), nil)
	}
	if eav.Exists == nil || *eav.Exists {
		if eav.Value == nil {
			var s string
			if eav.Exists == nil {
				s = "nil"
			} else {
				s = fmt.Sprintf("%v", *eav.Exists)
			}
			return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: Value must be provided when Exists is %s for Attribute: %s", s, a), nil)
		}

		var an, av string
		subs, an = appendAttributeName(subs, a)
		vars, av = appendAttributeValue(vars, *eav.Value)
		in = append(in, []byte(an)...)
		in = append(in, []byte(" = ")...)
		in = append(in, []byte(av)...)
	} else {
		if eav.Value != nil {
			return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: Value cannot be used when Exists is false for Attribute: %s", a), nil)
		}

		var an string
		subs, an = appendAttributeName(subs, a)
		in = append(in, []byte("attribute_not_exists(")...)
		in = append(in, []byte(an)...)
		in = append(in, []byte(")")...)
	}
	return in, subs, vars, nil
}

func appendComparisonCondition(in []byte, a string, op string, avl []*dynamodb.AttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue, keyCondition bool) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	switch op {
	case dynamodb.ComparisonOperatorBetween:
		return appendBetweenCondition(in, a, op, avl, subs, vars)
	case dynamodb.ComparisonOperatorBeginsWith:
		return appendBeginsWithCondition(in, a, op, avl, subs, vars)
	case dynamodb.ComparisonOperatorContains:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		return appendContainsCondition(in, a, op, avl, subs, vars, true)
	case dynamodb.ComparisonOperatorNotContains:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		return appendContainsCondition(in, a, op, avl, subs, vars, false)
	case dynamodb.ComparisonOperatorNull:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		var err error
		in, subs, err = appendNullCondition(in, a, op, avl, subs, true)
		return in, subs, vars, err
	case dynamodb.ComparisonOperatorNotNull:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		var err error
		in, subs, err = appendNullCondition(in, a, op, avl, subs, false)
		return in, subs, vars, err
	case dynamodb.ComparisonOperatorIn:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		return appendInCondition(in, a, op, avl, subs, vars)
	default:
		var eop string
		switch op {
		case dynamodb.ComparisonOperatorEq:
			eop = "="
		case dynamodb.ComparisonOperatorNe:
			if err := validateNotKeyCondition(keyCondition, op); err != nil {
				return in, subs, vars, err
			}
			eop = "<>"
		case dynamodb.ComparisonOperatorLe:
			eop = "<="
		case dynamodb.ComparisonOperatorGe:
			eop = ">="
		case dynamodb.ComparisonOperatorLt:
			eop = "<"
		case dynamodb.ComparisonOperatorGt:
			eop = ">"
		default:
			return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("Unknown comparison operator: %s", op), nil)
		}
		return appendArithmeticComparisonCondition(in, a, eop, avl, subs, vars)
	}
}

func appendBetweenCondition(in []byte, a string, op string, avl []*dynamodb.AttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	if err := validateArgCount(2, avl, op, a); err != nil {
		return in, subs, vars, err
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an, av0, av1 string
	subs, an = appendAttributeName(subs, a)
	vars, av0 = appendAttributeValue(vars, *avl[0])
	vars, av1 = appendAttributeValue(vars, *avl[1])
	in = append(in, []byte(an)...)
	in = append(in, []byte(" between ")...)
	in = append(in, []byte(av0)...)
	in = append(in, []byte(" and ")...)
	in = append(in, []byte(av1)...)
	return in, subs, vars, nil
}

func appendBeginsWithCondition(in []byte, a string, op string, avl []*dynamodb.AttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	if err := validateArgCount(1, avl, op, a); err != nil {
		return in, subs, vars, err
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an, av0 string
	subs, an = appendAttributeName(subs, a)
	vars, av0 = appendAttributeValue(vars, *avl[0])
	in = append(in, []byte("begins_with(")...)
	in = append(in, []byte(an)...)
	in = append(in, []byte(",")...)
	in = append(in, []byte(av0)...)
	in = append(in, []byte(")")...)
	return in, subs, vars, nil
}

func appendContainsCondition(in []byte, a string, op string, avl []*dynamodb.AttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue, p bool) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	if err := validateArgCount(1, avl, op, a); err != nil {
		return in, subs, vars, err
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an, av0 string
	subs, an = appendAttributeName(subs, a)
	vars, av0 = appendAttributeValue(vars, *avl[0])

	if !p {
		in = append(in, []byte("not ")...)
	}
	in = append(in, []byte("contains(")...)
	in = append(in, []byte(an)...)
	in = append(in, []byte(",")...)
	in = append(in, []byte(av0)...)
	in = append(in, []byte(")")...)
	return in, subs, vars, nil
}

func appendNullCondition(in []byte, a string, op string, avl []*dynamodb.AttributeValue, subs map[string]*string, p bool) ([]byte, map[string]*string, error) {
	if err := validateArgCount(0, avl, op, a); err != nil {
		return in, subs, err
	}

	var an string
	subs, an = appendAttributeName(subs, a)

	if p {
		in = append(in, []byte("attribute_not_exists(")...)
	} else {
		in = append(in, []byte("attribute_exists(")...)
	}
	in = append(in, []byte(an)...)
	in = append(in, []byte(")")...)
	return in, subs, nil
}

func appendInCondition(in []byte, a string, op string, avl []*dynamodb.AttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	if avl == nil {
		return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: AttributeValueList must be used with ComparisonOperator: %s for Attribute: %s", op, a), nil)
	}
	if len(avl) == 0 {
		return in, subs, vars, awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: Invalid number of argument(s)0 for the %s ComparisonOperator", op), nil)
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an string
	subs, an = appendAttributeName(subs, a)
	in = append(in, []byte(an)...)
	in = append(in, []byte(" in (")...)
	in, vars = appendAttributeValues(in, avl, vars)
	in = append(in, []byte(")")...)
	return in, subs, vars, nil
}

func appendArithmeticComparisonCondition(in []byte, a string, op string, avl []*dynamodb.AttributeValue, subs map[string]*string, vars map[string]*dynamodb.AttributeValue) ([]byte, map[string]*string, map[string]*dynamodb.AttributeValue, error) {
	if err := validateArgCount(1, avl, op, a); err != nil {
		return in, subs, vars, err
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an, av0 string
	subs, an = appendAttributeName(subs, a)
	vars, av0 = appendAttributeValue(vars, *avl[0])

	in = append(in, []byte(an)...)
	in = append(in, []byte(" ")...)
	in = append(in, []byte(op)...)
	in = append(in, []byte(" ")...)
	in = append(in, []byte(av0)...)
	return in, subs, vars, nil
}

func appendAttributeName(subs map[string]*string, a string) (map[string]*string, string) {
	if len(a) == 0 {
		return subs, ""
	}
	if subs == nil {
		subs = make(map[string]*string)
	}
	l := len(subs)
	k := fmt.Sprintf("%s%d", attributeNamesKeyPrefix, l)
	for _, ok := subs[k]; ok; {
		l++
		k = fmt.Sprintf("%s%d", attributeNamesKeyPrefix, l)
	}
	subs[k] = &a
	return subs, k
}

func appendAttributeValue(vars map[string]*dynamodb.AttributeValue, av dynamodb.AttributeValue) (map[string]*dynamodb.AttributeValue, string) {
	if vars == nil {
		vars = make(map[string]*dynamodb.AttributeValue)
	}
	l := len(vars)
	k := fmt.Sprintf("%s%d", attributeValuesKeyPrefix, l)
	for _, ok := vars[k]; ok; {
		l++
		k = fmt.Sprintf("%s%d", attributeValuesKeyPrefix, l)
	}
	vars[k] = &av
	return vars, k
}

func validateArgCount(e int, a []*dynamodb.AttributeValue, op, n string) error {
	if a == nil && e > 0 {
		return awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: Value or AttributeValueList must be used with ComparisonOperator: %s for Attribute %s", op, n), nil)
	}
	if len(a) != e {
		return awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: Invalid number of argument(s) for the %s ComparisonOperator", op), nil)
	}
	for _, i := range a {
		if i == nil {
			return awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: Invalid number of argument(s) for the %s ComparisonOperator", op), nil)
		}
	}
	return nil
}

func validateScalarAttribute(avl []*dynamodb.AttributeValue, op string) error {
	if op == "=" || op == "<>" {
		return nil
	}
	for _, v := range avl {
		if v != nil {
			if v.S == nil && v.N == nil && v.B == nil {
				return awserr.New(ErrCodeValidationException, fmt.Sprintf("One or more parameter values were invalid: ComparisonOperator %s is not valid for %s AttributeValue type", op, attributeTypeName(*v)), nil)
			}
		}
	}
	return nil
}

func validateNotKeyCondition(kc bool, op string) error {
	if kc {
		return awserr.New(ErrCodeValidationException, fmt.Sprintf("Unsupported operator on KeyCondition: %s", op), nil)
	}
	return nil
}

func attributeTypeName(v dynamodb.AttributeValue) string {
	switch {
	case v.S != nil:
		return dynamodb.ScalarAttributeTypeS
	case v.N != nil:
		return dynamodb.ScalarAttributeTypeN
	case v.B != nil:
		return dynamodb.ScalarAttributeTypeB
	case len(v.SS) > 0:
		return "SS"
	case len(v.NS) > 0:
		return "NS"
	case len(v.BS) > 0:
		return "BS"
	case len(v.M) > 0:
		return "M"
	case len(v.L) > 0:
		return "L"
	case v.BOOL != nil:
		return "BOOL"
	case v.NULL != nil:
		return "NULL"
	default:
		return ""
	}
}
