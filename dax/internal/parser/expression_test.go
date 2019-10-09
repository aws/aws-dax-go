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

package parser

import (
	"bytes"
	"encoding/hex"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestExpressionEncoder(t *testing.T) {
	cases := []struct {
		typ  int
		in   string
		subs map[string]*string
		vars map[string]*dynamodb.AttributeValue
		out  []byte
	}{
		{
			typ: ProjectionExpr,
			in:  "a1",
			out: fromHex("0x8201818212626131"),
		},
		{
			typ: ProjectionExpr,
			in:  "a1,a2",
			out: fromHex("0x82018282126261318212626132"),
		},
		{
			typ: ProjectionExpr,
			in:  "a1,a2.k1",
			out: fromHex("0x82018282126261318312626132626B31"),
		},
		{
			typ: ProjectionExpr,
			in:  "a1,a4[0]",
			out: fromHex("0x82018282126261318312626134D90CFC00"),
		},
		{
			typ:  ProjectionExpr,
			in:   "a1,a3.#s1",
			subs: map[string]*string{"#s1": aws.String("k2")},
			out:  fromHex("0x82018282126261318312626133626B32"),
		},
		{
			typ: FilterExpr,
			in:  "a1 = a2",
			out: fromHex("0x830183008212626131821262613280"),
		},
		{
			typ:  FilterExpr,
			in:   "a1 <> :v1",
			vars: map[string]*dynamodb.AttributeValue{":v1": &dynamodb.AttributeValue{N: aws.String("5")}},
			out:  fromHex("0x8301830182126261318211008105"),
		},
		{
			typ: FilterExpr,
			in:  "a1 < :v1 and a2 >= :v2",
			vars: map[string]*dynamodb.AttributeValue{
				":v1": &dynamodb.AttributeValue{N: aws.String("5")},
				":v2": &dynamodb.AttributeValue{N: aws.String("10")},
			},
			out: fromHex("0x83018306830282126261318211008303821262613282110182050A"),
		},
		{
			typ: FilterExpr,
			in:  "a1 > :v1 or a2 <= :v2",
			vars: map[string]*dynamodb.AttributeValue{
				":v1": &dynamodb.AttributeValue{N: aws.String("5")},
				":v2": &dynamodb.AttributeValue{N: aws.String("10")},
			},
			out: fromHex("0x83018307830482126261318211008305821262613282110182050A"),
		},
		{
			typ: FilterExpr,
			in:  "not (a1 <> a2)",
			out: fromHex("0x8301820883018212626131821262613280"),
		},
		{
			typ: FilterExpr,
			in:  "a in (b,c,d)",
			out: fromHex("0x8301830A821261618382126162821261638212616480"),
		},
		{
			typ: FilterExpr,
			in:  "a between :v1 and :v2",
			vars: map[string]*dynamodb.AttributeValue{
				":v1": &dynamodb.AttributeValue{N: aws.String("5")},
				":v2": &dynamodb.AttributeValue{N: aws.String("10")},
			},
			out: fromHex("0x830184098212616182110082110182050A"),
		},
		{
			typ: ConditionExpr,
			in:  "attribute_exists(a)",
			out: fromHex("0x8301820B8212616180"),
		},
		{
			typ:  ConditionExpr,
			in:   "attribute_not_exists(#a.k1)",
			subs: map[string]*string{"#a": aws.String("a1")},
			out:  fromHex("0x8301820C8312626131626B3180"),
		},
		{
			typ: ConditionExpr,
			in:  "Attribute_type(a, S)",
			out: fromHex("0x8301830D821261618212615380"),
		},
		{
			typ: ConditionExpr,
			in:  "begins_With(a, substr)",
			out: fromHex("0x8301830E8212616182126673756273747280"),
		},
		{
			typ:  ConditionExpr,
			in:   "CONTAINS(a, :v)",
			vars: map[string]*dynamodb.AttributeValue{":v": &dynamodb.AttributeValue{N: aws.String("5")}},
			out:  fromHex("0x8301830F821261618211008105"),
		},
		{
			typ: ConditionExpr,
			in:  "a > size(c)",
			out: fromHex("0x830183048212616182108212616380"),
		},
		{
			typ: UpdateExpr,
			in:  "SET #pr.#5star[1] = :r5, #pr.#3star = :r3",
			subs: map[string]*string{
				"#pr":    aws.String("a1"),
				"#5star": aws.String("k5"),
				"#3star": aws.String("k3"),
			},
			vars: map[string]*dynamodb.AttributeValue{
				":r3": &dynamodb.AttributeValue{N: aws.String("3")},
				":r5": &dynamodb.AttributeValue{N: aws.String("5")},
			},
			out: fromHex("0x83018283138412626131626B35D90CFC0182110083138312626131626B33821101820503"),
		},
		{
			typ:  UpdateExpr,
			in:   "SET Price = Price - :p",
			vars: map[string]*dynamodb.AttributeValue{":p": &dynamodb.AttributeValue{N: aws.String("5")}},
			out:  fromHex("0x8301818313821265507269636583181A82126550726963658211008105"),
		},
		{
			typ:  UpdateExpr,
			in:   "SET #ri = list_append(#ri, :vals)",
			subs: map[string]*string{"#ri": aws.String("RelatedItems")},
			vars: map[string]*dynamodb.AttributeValue{":vals": &dynamodb.AttributeValue{N: aws.String("5")}},
			out:  fromHex("0x830181831382126C52656C617465644974656D7383181882126C52656C617465644974656D738211008105"),
		},
		{
			typ:  UpdateExpr,
			in:   "SET Price = if_not_exists(Price, :p)",
			vars: map[string]*dynamodb.AttributeValue{":p": &dynamodb.AttributeValue{N: aws.String("10")}},
			out:  fromHex("0x8301818313821265507269636583178212655072696365821100810A"),
		},
		{
			typ: UpdateExpr,
			in:  "REMOVE RelatedItems[1], RelatedItems[2]",
			out: fromHex("0x830182821683126C52656C617465644974656D73D90CFC01821683126C52656C617465644974656D73D90CFC0280"),
		},
		{
			typ:  UpdateExpr,
			in:   "ADD QuantityOnHand :q",
			vars: map[string]*dynamodb.AttributeValue{":q": &dynamodb.AttributeValue{N: aws.String("5")}},
			out:  fromHex("0x830181831482126E5175616E746974794F6E48616E648211008105"),
		},
		{
			typ:  UpdateExpr,
			in:   "DELETE Color :p",
			vars: map[string]*dynamodb.AttributeValue{":p": &dynamodb.AttributeValue{N: aws.String("5")}},
			out:  fromHex("0x8301818315821265436F6C6F728211008105"),
		},
		{
			typ:  UpdateExpr,
			in:   "DELETE Color :p, Color_2 :p",
			vars: map[string]*dynamodb.AttributeValue{":p": &dynamodb.AttributeValue{N: aws.String("5")}},
			out:  fromHex("0x8301828315821265436F6C6F728211008315821267436F6C6F725F328211008105"),
		},
	}

	for _, c := range cases {
		var buf bytes.Buffer
		r := make(map[int]string)
		r[c.typ] = c.in

		encoder := NewExpressionEncoder(r, c.subs, c.vars)
		if _, err := encoder.Parse(); err != nil {
			t.Errorf("unexpected error %v for %s", err, c.in)
			continue
		}
		if err := encoder.Write(c.typ, &buf); err != nil {
			t.Errorf("unexpected error %v for %s", err, c.in)
			continue
		}

		actual := buf.Bytes()
		if !reflect.DeepEqual(c.out, actual) {
			t.Errorf("expected %v, actual %v for %s", c.out, actual, c.in)
		}
	}
}

func TestExpressionEncoderErrors(t *testing.T) {
	cases := []struct {
		typ  int
		in   string
		subs map[string]*string
		vars map[string]*dynamodb.AttributeValue
		err  error
	}{
		{
			typ: ProjectionExpr,
			in:  "(a1)",
			err: newInvalidParameterError("Invalid Projection: Syntax error; token: (, near line 1 char 0"),
		},
		{
			typ:  ProjectionExpr,
			in:   "a",
			subs: map[string]*string{"#b": aws.String("c")},
			err:  newInvalidParameterError("Value provided in ExpressionAttributeNames unused in expressions: keys: {#b}"),
		},
		{
			typ: ProjectionExpr,
			in:  "#a",
			err: newInvalidParameterError("Invalid ProjectionExpression. Substitution value not provided for #a"),
		},
		{
			typ: FilterExpr,
			in:  "a < ",
			err: newInvalidParameterError("Invalid Filter: Syntax error; token: <EOF>, near line 1 char 4"),
		},
		{
			typ: ConditionExpr,
			in:  "a < :v",
			vars: map[string]*dynamodb.AttributeValue{
				":v": &dynamodb.AttributeValue{N: aws.String("10")},
				":z": &dynamodb.AttributeValue{N: aws.String("5")},
			},
			err: newInvalidParameterError("Value provided in ExpressionAttributeValues unused in expressions: keys: {:z}"),
		},
		{
			typ: FilterExpr,
			in:  "a < :v",
			err: newInvalidParameterError("Invalid FilterExpression: An expression attribute value used in expression is not defined: attribute value :v"),
		},
		{
			typ: KeyConditionExpr,
			in:  "a < b[-25]",
			err: newInvalidParameterError("Invalid KeyCondition: Syntax error; token: -, near line 1 char 6"),
		},
		{
			typ: ConditionExpr,
			in:  "a < if_not_exists(c)",
			err: newInvalidParameterError("Invalid ConditionExpression: The function is not allowed in a condition expression"),
		},
		{
			typ: ConditionExpr,
			in:  "a < size(attribute_exists(c))",
			err: newInvalidParameterError("Only size() function is allowed to be nested"),
		},
		{
			typ: UpdateExpr,
			in:  "set a = size(b)",
			err: newInvalidParameterError("Invalid UpdateExpression: The function is not allowed in a update expression"),
		},
		{
			typ: UpdateExpr,
			in:  "set a = if_not_exists(list_append(a, b))",
			err: newInvalidParameterError("Only if_not_exists() function can be nested"),
		},
	}

	for _, c := range cases {
		r := make(map[int]string)
		r[c.typ] = c.in

		encoder := NewExpressionEncoder(r, c.subs, c.vars)
		_, actual := encoder.Parse()
		if !reflect.DeepEqual(actual, c.err) {
			t.Errorf("expected error '%v', got '%v' for %s", c.err, actual, c.in)
		}
	}
}

func BenchmarkProjection(b *testing.B) {
	var buf bytes.Buffer
	expr := map[int]string{
		ProjectionExpr: "a1.k1,a2.l[0],a3",
	}
	expected := []byte{130, 1, 131, 131, 18, 98, 97, 49, 98, 107, 49, 132, 18, 98, 97, 50, 97, 108, 217, 12, 252, 0, 130, 18, 98, 97, 51}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		encoder := NewExpressionEncoder(expr, nil, nil)
		if _, err := encoder.Parse(); err != nil {
			b.Fatalf("unexpected error %v", err)
		}
		if err := encoder.Write(ProjectionExpr, &buf); err != nil {
			b.Fatalf("unexpected error %v", err)
		}

		actual := buf.Bytes()
		if !reflect.DeepEqual(expected, actual) {
			b.Fatalf("expected %v, actual %v for %s", expected, actual, expr[ProjectionExpr])
		}
	}
}

func BenchmarkFilter(b *testing.B) {
	var buf bytes.Buffer
	expr := map[int]string{
		KeyConditionExpr: "pk = :v1 and hk < :v2",
	}
	vars := map[string]*dynamodb.AttributeValue{
		":v1": &dynamodb.AttributeValue{S: aws.String("pkval")},
		":v2": &dynamodb.AttributeValue{N: aws.String("5")},
	}
	expected := []byte{131, 1, 131, 6, 131, 0, 130, 18, 98, 112, 107, 130, 17, 0, 131, 2, 130, 18, 98, 104, 107, 130, 17, 1, 130, 101, 112, 107, 118, 97, 108, 5}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		encoder := NewExpressionEncoder(expr, nil, vars)
		if _, err := encoder.Parse(); err != nil {
			b.Fatalf("unexpected error %v", err)
		}
		if err := encoder.Write(KeyConditionExpr, &buf); err != nil {
			b.Fatalf("unexpected error %v", err)
		}

		actual := buf.Bytes()
		if !reflect.DeepEqual(expected, actual) {
			b.Fatalf("expected %v, actual %v for %s", expected, actual, expr[KeyConditionExpr])
		}
	}
}

func BenchmarkFunction(b *testing.B) {
	var buf bytes.Buffer
	expr := map[int]string{
		FilterExpr: "attribute_exists(a)",
	}
	expected := []byte{131, 1, 130, 11, 130, 18, 97, 97, 128}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		encoder := NewExpressionEncoder(expr, nil, nil)
		if _, err := encoder.Parse(); err != nil {
			b.Fatalf("unexpected error %v", err)
		}
		if err := encoder.Write(FilterExpr, &buf); err != nil {
			b.Fatalf("unexpected error %v", err)
		}

		actual := buf.Bytes()
		if !reflect.DeepEqual(expected, actual) {
			b.Fatalf("expected %v, actual %v for %s", expected, actual, expr[FilterExpr])
		}
	}

}

func fromHex(s string) []byte {
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return dst[:n]
}
