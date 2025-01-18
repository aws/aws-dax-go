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
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestBuildDocumentPath(t *testing.T) {
	cases := []struct {
		projectionExpression     string
		expressionAttributeNames map[string]string
		documentPath             documentPath
	}{
		{
			"a", nil,
			documentPath{[]documentPathElement{{name: "a", index: -1}}},
		},
		{
			"a.b", nil,
			documentPath{[]documentPathElement{
				{name: "a", index: -1},
				{name: "b", index: -1},
			}},
		},
		{
			"a[3]", nil,
			documentPath{[]documentPathElement{
				{name: "a", index: -1},
				{name: "", index: 3},
			}},
		},
		{
			"a.#s.c", map[string]string{"#s": "b"},
			documentPath{[]documentPathElement{
				{name: "a", index: -1},
				{name: "b", index: -1},
				{name: "c", index: -1},
			}},
		},
		{
			"#a[1].#b", map[string]string{"#a": "with.dot", "#b": "sub.field"},
			documentPath{[]documentPathElement{
				{name: "with.dot", index: -1},
				{name: "", index: 1},
				{name: "sub.field", index: -1},
			}},
		},
	}

	for _, c := range cases {
		actual, err := buildDocumentPath(c.projectionExpression, c.expressionAttributeNames)
		if err != nil {
			t.Errorf(fmt.Sprintf("unexpected error %v", err))
		}
		if !reflect.DeepEqual(c.documentPath, actual) {
			t.Errorf(fmt.Sprintf("expected %v, got %v for %s", c.documentPath, actual, c.projectionExpression))
		}
	}
}

func TestBuildProjectionOrdinals(t *testing.T) {
	cases := []struct {
		projectionExpression     string
		expressionAttributeNames map[string]string
		documentPaths            []documentPath
	}{
		{
			"#1",
			map[string]string{"#1": "a"},
			[]documentPath{{[]documentPathElement{{name: "a", index: -1}}}},
		},
		{
			"#1, #2",
			map[string]string{"#1": "a", "#2": "b"},
			[]documentPath{{[]documentPathElement{{name: "a", index: -1}}}, {[]documentPathElement{{name: "b", index: -1}}}},
		},
	}

	for _, c := range cases {
		actual, err := buildProjectionOrdinals(&c.projectionExpression, c.expressionAttributeNames)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if !reflect.DeepEqual(c.documentPaths, actual) {
			t.Errorf("expected %v, got %v for %s", c.documentPaths, actual, c.projectionExpression)
		}
	}

}

func TestItemBuilder(t *testing.T) {
	cases := []struct {
		projectionExpression     string
		expressionAttributeNames map[string]string
		values                   map[int]types.AttributeValue
		item                     map[string]types.AttributeValue
	}{
		{
			projectionExpression:     "a",
			expressionAttributeNames: nil,
			values: map[int]types.AttributeValue{
				0: &types.AttributeValueMemberS{Value: "av"},
			},
			item: map[string]types.AttributeValue{
				"a": &types.AttributeValueMemberS{Value: "av"},
			},
		},
		{
			projectionExpression:     "a,b[2],c.d",
			expressionAttributeNames: nil,
			values:                   map[int]types.AttributeValue{},
			item:                     map[string]types.AttributeValue{},
		},
		{
			projectionExpression:     "a.b",
			expressionAttributeNames: nil,
			values: map[int]types.AttributeValue{
				0: &types.AttributeValueMemberS{Value: "av"},
			},
			item: map[string]types.AttributeValue{
				"a": &types.AttributeValueMemberM{
					Value: map[string]types.AttributeValue{
						"b": &types.AttributeValueMemberS{Value: "av"},
					},
				},
			},
		},
		{
			projectionExpression:     "a[3]",
			expressionAttributeNames: nil,
			values: map[int]types.AttributeValue{
				0: &types.AttributeValueMemberS{Value: "av"},
			},
			item: map[string]types.AttributeValue{
				"a": &types.AttributeValueMemberL{
					Value: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: "av"},
					},
				},
			},
		},
		{
			projectionExpression:     "a[3],a[2]",
			expressionAttributeNames: nil,
			values: map[int]types.AttributeValue{
				0: &types.AttributeValueMemberS{Value: "av3"},
				1: &types.AttributeValueMemberS{Value: "av2"},
			},
			item: map[string]types.AttributeValue{
				"a": &types.AttributeValueMemberL{
					Value: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: "av2"},
						&types.AttributeValueMemberS{Value: "av3"},
					},
				},
			},
		},
		{
			projectionExpression:     "a[2],a[3]",
			expressionAttributeNames: nil,
			values: map[int]types.AttributeValue{
				0: &types.AttributeValueMemberS{Value: "av2"},
				1: &types.AttributeValueMemberS{Value: "av3"},
			},
			item: map[string]types.AttributeValue{
				"a": &types.AttributeValueMemberL{
					Value: []types.AttributeValue{
						&types.AttributeValueMemberS{Value: "av2"},
						&types.AttributeValueMemberS{Value: "av3"},
					},
				},
			},
		},
		{
			projectionExpression:     "a[2].b.c,a[2].b.d,a[1].b.e",
			expressionAttributeNames: nil,
			values: map[int]types.AttributeValue{
				2: &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
					"field": &types.AttributeValueMemberS{Value: "value"},
				}},
				0: &types.AttributeValueMemberN{Value: "4"},
				1: &types.AttributeValueMemberN{Value: "2"},
			},
			item: map[string]types.AttributeValue{
				"a": &types.AttributeValueMemberL{
					Value: []types.AttributeValue{
						&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
							"b": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
								"e": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
									"field": &types.AttributeValueMemberS{Value: "value"},
								}},
							}},
						}},
						&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
							"b": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
								"c": &types.AttributeValueMemberN{Value: "4"},
								"d": &types.AttributeValueMemberN{Value: "2"},
							}},
						}},
					},
				},
			},
		},
		{
			projectionExpression:     "a[4],a[2],b.c[12]",
			expressionAttributeNames: nil,
			values: map[int]types.AttributeValue{
				2: &types.AttributeValueMemberL{Value: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: "elem"},
				}},
				0: &types.AttributeValueMemberN{Value: "4"},
				1: &types.AttributeValueMemberN{Value: "2"},
			},
			item: map[string]types.AttributeValue{
				"a": &types.AttributeValueMemberL{
					Value: []types.AttributeValue{
						&types.AttributeValueMemberN{Value: "2"},
						&types.AttributeValueMemberN{Value: "4"},
					},
				},
				"b": &types.AttributeValueMemberM{
					Value: map[string]types.AttributeValue{
						"c": &types.AttributeValueMemberL{
							Value: []types.AttributeValue{
								&types.AttributeValueMemberL{
									Value: []types.AttributeValue{
										&types.AttributeValueMemberS{Value: "elem"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			projectionExpression: "#a[1].#b",
			expressionAttributeNames: map[string]string{
				"#a": "with.dot",
				"#b": "sub.field",
			},
			values: map[int]types.AttributeValue{
				0: &types.AttributeValueMemberN{Value: "4"},
			},
			item: map[string]types.AttributeValue{
				"with.dot": &types.AttributeValueMemberL{
					Value: []types.AttributeValue{
						&types.AttributeValueMemberM{
							Value: map[string]types.AttributeValue{
								"sub.field": &types.AttributeValueMemberN{Value: "4"},
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		dps, err := buildProjectionOrdinals(aws.String(c.projectionExpression), c.expressionAttributeNames)
		if err != nil {
			t.Errorf(fmt.Sprintf("unexpected error %v", err))
		}

		ib := &itemBuilder{}
		for k, v := range c.values {
			ib.insert(dps[k], v)
		}

		actual := ib.toItem()
		if !reflect.DeepEqual(c.item, actual) {
			t.Errorf(fmt.Sprintf("expected %v, got %v for %s", c.item, actual, c.projectionExpression))
		}
	}
}
