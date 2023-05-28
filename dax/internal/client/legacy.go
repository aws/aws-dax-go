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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
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
		output.QueryFilter = nil
	}
	if kf {
		output.KeyConditionExpression, output.ExpressionAttributeNames, output.ExpressionAttributeValues, err =
			translateFilter(types.ConditionalOperatorAnd, output.KeyConditions, output.ExpressionAttributeNames, output.ExpressionAttributeValues, true)
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

func hasAttributesToGet(a []string, p *string) (bool, error) {
	af := len(a) != 0
	pf := p != nil
	if af && pf {
		return false, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: "Cannot specify both AttributesToGet and ProjectionExpression",
			Fault:   smithy.FaultClient,
		}
	}
	return af, nil
}

func hasExpected(e map[string]types.ExpectedAttributeValue, c *string) (bool, error) {
	ef := len(e) != 0
	cf := c != nil
	if ef && cf {
		return false, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: "Cannot specify both Expected and ConditionExpression",
			Fault:   smithy.FaultClient,
		}
	}
	return ef, nil
}

func hasAttributeUpdates(u map[string]types.AttributeValueUpdate, e *string) (bool, error) {
	uf := len(u) > 0
	ef := e != nil
	if uf && ef {
		return false, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: "Cannot specify both AttributeUpdates and UpdateExpression",
			Fault:   smithy.FaultClient,
		}
	}
	return uf, nil
}

func hasFilter(c map[string]types.Condition, e *string) (bool, error) {
	cf := len(c) > 0
	ef := e != nil
	if cf && ef {
		return false, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: "Cannot specify both [Scan|Query]Filter and [Scan|Query]FilterExpression",
			Fault:   smithy.FaultClient,
		}
	}
	return cf, nil
}

func translateAttributesToGet(attrs []string, subs map[string]string) (*string, map[string]string, error) {
	out, sub := appendAttributeNames(nil, attrs, subs)
	return aws.String(string(out)), sub, nil
}

func translateExpected(o types.ConditionalOperator, e map[string]types.ExpectedAttributeValue, subs map[string]string, vars map[string]types.AttributeValue) (*string, map[string]string, map[string]types.AttributeValue, error) {
	op := types.ConditionalOperatorAnd
	if len(o) > 0 {
		op = o
	}
	ops := fmt.Sprintf(" %s ", strings.TrimSpace(string(op)))

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

func translateFilter(o types.ConditionalOperator, c map[string]types.Condition, subs map[string]string, vars map[string]types.AttributeValue, keyCondition bool) (*string, map[string]string, map[string]types.AttributeValue, error) {
	op := types.ConditionalOperatorAnd
	if len(o) > 0 {
		op = o
	}
	ops := fmt.Sprintf(" %s ", strings.TrimSpace(string(op)))

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

func translateAttributeUpdates(avus map[string]types.AttributeValueUpdate, subs map[string]string, vars map[string]types.AttributeValue) (*string, map[string]string, map[string]types.AttributeValue, error) {
	var sets, adds, dels, rems []string

	for a, avu := range avus {
		act := avu.Action
		if avu.Value == nil && act != types.AttributeActionDelete {
			return nil, subs, vars, &smithy.GenericAPIError{
				Code:    ErrCodeValidationException,
				Message: "only DELETE action is allowed when no attribute value is specified",
				Fault:   smithy.FaultClient,
			}
		}

		var an, av string
		subs, an = appendAttributeName(subs, a)
		if avu.Value != nil {
			vars, av = appendAttributeValue(vars, avu.Value)
		}

		switch act {
		case types.AttributeActionPut:
			sets = append(sets, fmt.Sprintf("%s=%s", an, av))
		case types.AttributeActionAdd:
			adds = append(adds, fmt.Sprintf("%s %s", an, av))
		case types.AttributeActionDelete:
			if len(av) == 0 {
				rems = append(rems, an)
			} else {
				dels = append(dels, fmt.Sprintf("%s %s", an, av))
			}
		default:
			return nil, subs, vars, &smithy.GenericAPIError{
				Code:    ErrCodeValidationException,
				Message: fmt.Sprintf("unknown AttributeValueUpdate Action: %s", act),
				Fault:   smithy.FaultClient,
			}
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

func appendAttributeNames(in []byte, attrs []string, sub map[string]string) ([]byte, map[string]string) {
	f := true
	for _, a := range attrs {
		if len(a) == 0 {
			continue
		}
		if f {
			f = false
		} else {
			in = append(in, []byte(",")...)
		}
		var an string
		sub, an = appendAttributeName(sub, a)
		in = append(in, []byte(an)...)
	}
	return in, sub
}

func appendAttributeValues(in []byte, avl []types.AttributeValue, vars map[string]types.AttributeValue) ([]byte, map[string]types.AttributeValue) {
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
		vars, av = appendAttributeValue(vars, v)
		in = append(in, []byte(av)...)
	}
	return in, vars
}

func appendCondition(in []byte, a string, eav types.ExpectedAttributeValue, subs map[string]string, vars map[string]types.AttributeValue) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	if eav.Value != nil && len(eav.AttributeValueList) > 0 {
		return in, subs, vars, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("One or more parameter values were invalid: Value and AttributeValueList cannot be used together for Attribute: %s", a),
			Fault:   smithy.FaultClient,
		}
	}

	op := eav.ComparisonOperator
	if len(op) == 0 {
		return appendExistsCondition(in, a, eav, subs, vars)
	} else if eav.Exists != nil && *eav.Exists {
		return in, subs, vars, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("One or more parameter values were invalid: Exists and ComparisonOperator cannot be used together for Attribute: %s", a),
			Fault:   smithy.FaultClient,
		}
	}
	avl := eav.AttributeValueList
	if len(avl) == 0 && eav.Value != nil {
		avl = []types.AttributeValue{eav.Value}
	}
	return appendComparisonCondition(in, a, op, avl, subs, vars, false)
}

func appendFilterCondition(in []byte, a string, c types.Condition, subs map[string]string, vars map[string]types.AttributeValue, keyCondition bool) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	op := c.ComparisonOperator
	if len(op) == 0 {
		return in, subs, vars, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("One or more parameter values were invalid: AttributeValueList can only be used with a ComparisonOperator for Attribute: %s", a),
			Fault:   smithy.FaultClient,
		}
	}
	return appendComparisonCondition(in, a, op, c.AttributeValueList, subs, vars, keyCondition)
}

func appendExistsCondition(in []byte, a string, eav types.ExpectedAttributeValue, subs map[string]string, vars map[string]types.AttributeValue) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	if len(eav.AttributeValueList) != 0 {
		return in, subs, vars, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("One or more parameter values were invalid: AttributeValueList can only be used with a ComparisonOperator for Attribute: %s", a),
			Fault:   smithy.FaultClient,
		}
	}
	if eav.Exists == nil || *eav.Exists {
		if eav.Value == nil {
			var s string
			if eav.Exists == nil {
				s = "nil"
			} else {
				s = fmt.Sprintf("%v", *eav.Exists)
			}
			return in, subs, vars, &smithy.GenericAPIError{
				Code:    ErrCodeValidationException,
				Message: fmt.Sprintf("One or more parameter values were invalid: Value must be provided when Exists is %s for Attribute: %s", s, a),
				Fault:   smithy.FaultClient,
			}
		}

		var an, av string
		subs, an = appendAttributeName(subs, a)
		vars, av = appendAttributeValue(vars, eav.Value)
		in = append(in, []byte(an)...)
		in = append(in, []byte(" = ")...)
		in = append(in, []byte(av)...)
	} else {
		if eav.Value != nil {
			return in, subs, vars, &smithy.GenericAPIError{
				Code:    ErrCodeValidationException,
				Message: fmt.Sprintf("One or more parameter values were invalid: Value cannot be used when Exists is false for Attribute: %s", a),
				Fault:   smithy.FaultClient,
			}
		}

		var an string
		subs, an = appendAttributeName(subs, a)
		in = append(in, []byte("attribute_not_exists(")...)
		in = append(in, []byte(an)...)
		in = append(in, []byte(")")...)
	}
	return in, subs, vars, nil
}

func appendComparisonCondition(in []byte, a string, op types.ComparisonOperator, avl []types.AttributeValue, subs map[string]string, vars map[string]types.AttributeValue, keyCondition bool) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	switch op {
	case types.ComparisonOperatorBetween:
		return appendBetweenCondition(in, a, op, avl, subs, vars)
	case types.ComparisonOperatorBeginsWith:
		return appendBeginsWithCondition(in, a, op, avl, subs, vars)
	case types.ComparisonOperatorContains:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		return appendContainsCondition(in, a, op, avl, subs, vars, true)
	case types.ComparisonOperatorNotContains:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		return appendContainsCondition(in, a, op, avl, subs, vars, false)
	case types.ComparisonOperatorNull:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		var err error
		in, subs, err = appendNullCondition(in, a, op, avl, subs, true)
		return in, subs, vars, err
	case types.ComparisonOperatorNotNull:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		var err error
		in, subs, err = appendNullCondition(in, a, op, avl, subs, false)
		return in, subs, vars, err
	case types.ComparisonOperatorIn:
		if err := validateNotKeyCondition(keyCondition, op); err != nil {
			return in, subs, vars, err
		}
		return appendInCondition(in, a, op, avl, subs, vars)
	default:
		switch op {
		case types.ComparisonOperatorEq:
			// do nothing
		case types.ComparisonOperatorNe:
			if err := validateNotKeyCondition(keyCondition, op); err != nil {
				return in, subs, vars, err
			}
		case types.ComparisonOperatorLe:
			// do nothing
		case types.ComparisonOperatorGe:
			// do nothing
		case types.ComparisonOperatorLt:
			// do nothing
		case types.ComparisonOperatorGt:
			// do nothing
		default:
			return in, subs, vars, &smithy.GenericAPIError{
				Code:    ErrCodeValidationException,
				Message: fmt.Sprintf("Unknown comparison operator: %s", op),
				Fault:   smithy.FaultClient,
			}
		}
		return appendArithmeticComparisonCondition(in, a, op, avl, subs, vars)
	}
}

func appendBetweenCondition(in []byte, a string, op types.ComparisonOperator, avl []types.AttributeValue, subs map[string]string, vars map[string]types.AttributeValue) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	if err := validateArgCount(2, avl, op, a); err != nil {
		return in, subs, vars, err
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an, av0, av1 string
	subs, an = appendAttributeName(subs, a)
	vars, av0 = appendAttributeValue(vars, avl[0])
	vars, av1 = appendAttributeValue(vars, avl[1])
	in = append(in, []byte(an)...)
	in = append(in, []byte(" between ")...)
	in = append(in, []byte(av0)...)
	in = append(in, []byte(" and ")...)
	in = append(in, []byte(av1)...)
	return in, subs, vars, nil
}

func appendBeginsWithCondition(in []byte, a string, op types.ComparisonOperator, avl []types.AttributeValue, subs map[string]string, vars map[string]types.AttributeValue) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	if err := validateArgCount(1, avl, op, a); err != nil {
		return in, subs, vars, err
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an, av0 string
	subs, an = appendAttributeName(subs, a)
	vars, av0 = appendAttributeValue(vars, avl[0])
	in = append(in, []byte("begins_with(")...)
	in = append(in, []byte(an)...)
	in = append(in, []byte(",")...)
	in = append(in, []byte(av0)...)
	in = append(in, []byte(")")...)
	return in, subs, vars, nil
}

func appendContainsCondition(in []byte, a string, op types.ComparisonOperator, avl []types.AttributeValue, subs map[string]string, vars map[string]types.AttributeValue, p bool) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	if err := validateArgCount(1, avl, op, a); err != nil {
		return in, subs, vars, err
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an, av0 string
	subs, an = appendAttributeName(subs, a)
	vars, av0 = appendAttributeValue(vars, avl[0])

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

func appendNullCondition(in []byte, a string, op types.ComparisonOperator, avl []types.AttributeValue, subs map[string]string, p bool) ([]byte, map[string]string, error) {
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

func appendInCondition(in []byte, a string, op types.ComparisonOperator, avl []types.AttributeValue, subs map[string]string, vars map[string]types.AttributeValue) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	if avl == nil {
		return in, subs, vars, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("One or more parameter values were invalid: AttributeValueList must be used with ComparisonOperator: %s for Attribute: %s", op, a),
			Fault:   smithy.FaultClient,
		}
	}
	if len(avl) == 0 {
		return in, subs, vars, &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("One or more parameter values were invalid: Invalid number of argument(s)0 for the %s ComparisonOperator", op),
			Fault:   smithy.FaultClient,
		}
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

func appendArithmeticComparisonCondition(in []byte, a string, op types.ComparisonOperator, avl []types.AttributeValue, subs map[string]string, vars map[string]types.AttributeValue) ([]byte, map[string]string, map[string]types.AttributeValue, error) {
	if err := validateArgCount(1, avl, op, a); err != nil {
		return in, subs, vars, err
	}
	if err := validateScalarAttribute(avl, op); err != nil {
		return in, subs, vars, err
	}

	var an, av0 string
	subs, an = appendAttributeName(subs, a)
	vars, av0 = appendAttributeValue(vars, avl[0])

	in = append(in, []byte(an)...)
	in = append(in, []byte(" ")...)
	in = append(in, []byte(op)...)
	in = append(in, []byte(" ")...)
	in = append(in, []byte(av0)...)
	return in, subs, vars, nil
}

func appendAttributeName(subs map[string]string, a string) (map[string]string, string) {
	if len(a) == 0 {
		return subs, ""
	}
	if subs == nil {
		subs = make(map[string]string)
	}
	l := len(subs)
	k := fmt.Sprintf("%s%d", attributeNamesKeyPrefix, l)
	for _, ok := subs[k]; ok; {
		l++
		k = fmt.Sprintf("%s%d", attributeNamesKeyPrefix, l)
	}
	subs[k] = a
	return subs, k
}

func appendAttributeValue(vars map[string]types.AttributeValue, av types.AttributeValue) (map[string]types.AttributeValue, string) {
	if vars == nil {
		vars = make(map[string]types.AttributeValue)
	}
	l := len(vars)
	k := fmt.Sprintf("%s%d", attributeValuesKeyPrefix, l)
	for _, ok := vars[k]; ok; {
		l++
		k = fmt.Sprintf("%s%d", attributeValuesKeyPrefix, l)
	}
	vars[k] = av
	return vars, k
}

func validateArgCount(e int, a []types.AttributeValue, op types.ComparisonOperator, n string) error {
	if a == nil && e > 0 {
		return &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("One or more parameter values were invalid: Value or AttributeValueList must be used with ComparisonOperator: %s for Attribute %s", op, n),
			Fault:   smithy.FaultClient,
		}
	}
	if len(a) != e {
		return &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("One or more parameter values were invalid: Invalid number of argument(s) for the %s ComparisonOperator", op),
			Fault:   smithy.FaultClient,
		}
	}
	for _, i := range a {
		if i == nil {
			return &smithy.GenericAPIError{
				Code:    ErrCodeValidationException,
				Message: fmt.Sprintf("One or more parameter values were invalid: Invalid number of argument(s) for the %s ComparisonOperator", op),
				Fault:   smithy.FaultClient,
			}
		}
	}
	return nil
}

func validateScalarAttribute(avl []types.AttributeValue, op types.ComparisonOperator) error {
	if op == types.ComparisonOperatorEq || op == types.ComparisonOperatorNe {
		return nil
	}
	for _, v := range avl {
		if v != nil {
			switch v.(type) {
			case *types.AttributeValueMemberS, *types.AttributeValueMemberN, *types.AttributeValueMemberB:
			// ok
			default:
				return &smithy.GenericAPIError{
					Code:    ErrCodeValidationException,
					Message: fmt.Sprintf("One or more parameter values were invalid: ComparisonOperator %s is not valid for %s AttributeValue type", op, attributeTypeName(v)),
					Fault:   smithy.FaultClient,
				}
			}
		}
	}
	return nil
}

func validateNotKeyCondition(kc bool, op types.ComparisonOperator) error {
	if kc {
		return &smithy.GenericAPIError{
			Code:    ErrCodeValidationException,
			Message: fmt.Sprintf("Unsupported operator on KeyCondition: %s", op),
			Fault:   smithy.FaultClient,
		}
	}
	return nil
}

func attributeTypeName(v types.AttributeValue) string {
	switch v.(type) {
	case *types.AttributeValueMemberS:
		return string(types.ScalarAttributeTypeS)
	case *types.AttributeValueMemberN:
		return string(types.ScalarAttributeTypeN)
	case *types.AttributeValueMemberB:
		return string(types.ScalarAttributeTypeB)
	case *types.AttributeValueMemberSS:
		return "SS"
	case *types.AttributeValueMemberNS:
		return "NS"
	case *types.AttributeValueMemberBS:
		return "BS"
	case *types.AttributeValueMemberM:
		return "M"
	case *types.AttributeValueMemberL:
		return "L"
	case *types.AttributeValueMemberBOOL:
		return "BOOL"
	case *types.AttributeValueMemberNULL:
		return "NULL"
	default:
		return ""
	}
}
