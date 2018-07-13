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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
	"testing"
)

func TestBuildDocumentPath(t *testing.T) {
	cases := []struct {
		projectionExpression     string
		expressionAttributeNames map[string]*string
		documentPath             documentPath
	}{
		{
			"a", nil,
			documentPath{[]documentPathElement{documentPathElement{name: "a", index: -1}}},
		},
		{
			"a.b", nil,
			documentPath{[]documentPathElement{
				documentPathElement{name: "a", index: -1},
				documentPathElement{name: "b", index: -1},
			}},
		},
		{
			"a[3]", nil,
			documentPath{[]documentPathElement{
				documentPathElement{name: "a", index: -1},
				documentPathElement{name: "", index: 3},
			}},
		},
		{
			"a.#s.c", map[string]*string{"#s": aws.String("b")},
			documentPath{[]documentPathElement{
				documentPathElement{name: "a", index: -1},
				documentPathElement{name: "b", index: -1},
				documentPathElement{name: "c", index: -1},
			}},
		},
		{
			"#a[1].#b", map[string]*string{"#a": aws.String("with.dot"), "#b": aws.String("sub.field")},
			documentPath{[]documentPathElement{
				documentPathElement{name: "with.dot", index: -1},
				documentPathElement{name: "", index: 1},
				documentPathElement{name: "sub.field", index: -1},
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
		expressionAttributeNames map[string]*string
		documentPaths            []documentPath
	}{
		{
			"#1",
			map[string]*string{"#1": aws.String("a")},
			[]documentPath{{[]documentPathElement{{name: "a", index: -1}}}},
		},
		{
			"#1, #2",
			map[string]*string{"#1": aws.String("a"), "#2": aws.String("b")},
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
		expressionAttributeNames map[string]*string
		values                   map[int]*dynamodb.AttributeValue
		item                     map[string]*dynamodb.AttributeValue
	}{
		{
			"a", nil,
			map[int]*dynamodb.AttributeValue{
				0: &dynamodb.AttributeValue{S: aws.String("av")},
			},
			map[string]*dynamodb.AttributeValue{
				"a": &dynamodb.AttributeValue{S: aws.String("av")},
			},
		},
		{
			"a,b[2],c.d", nil,
			map[int]*dynamodb.AttributeValue{},
			map[string]*dynamodb.AttributeValue{},
		},
		{
			"a.b", nil,
			map[int]*dynamodb.AttributeValue{
				0: &dynamodb.AttributeValue{S: aws.String("av")},
			},
			map[string]*dynamodb.AttributeValue{
				"a": &dynamodb.AttributeValue{
					M: map[string]*dynamodb.AttributeValue{
						"b": &dynamodb.AttributeValue{S: aws.String("av")},
					},
				},
			},
		},
		{
			"a[3]", nil,
			map[int]*dynamodb.AttributeValue{
				0: &dynamodb.AttributeValue{S: aws.String("av")},
			},
			map[string]*dynamodb.AttributeValue{
				"a": &dynamodb.AttributeValue{
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{S: aws.String("av")},
					},
				},
			},
		},
		{
			"a[3],a[2]", nil,
			map[int]*dynamodb.AttributeValue{
				0: &dynamodb.AttributeValue{S: aws.String("av3")},
				1: &dynamodb.AttributeValue{S: aws.String("av2")},
			},
			map[string]*dynamodb.AttributeValue{
				"a": &dynamodb.AttributeValue{
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{S: aws.String("av2")},
						&dynamodb.AttributeValue{S: aws.String("av3")},
					},
				},
			},
		},
		{
			"a[2],a[3]", nil,
			map[int]*dynamodb.AttributeValue{
				0: &dynamodb.AttributeValue{S: aws.String("av2")},
				1: &dynamodb.AttributeValue{S: aws.String("av3")},
			},
			map[string]*dynamodb.AttributeValue{
				"a": &dynamodb.AttributeValue{
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{S: aws.String("av2")},
						&dynamodb.AttributeValue{S: aws.String("av3")},
					},
				},
			},
		},
		{
			"a[2].b.c,a[2].b.d,a[1].b.e", nil,
			map[int]*dynamodb.AttributeValue{
				2: &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
					"field": &dynamodb.AttributeValue{S: aws.String("value")},
				}},
				0: &dynamodb.AttributeValue{N: aws.String("4")},
				1: &dynamodb.AttributeValue{N: aws.String("2")},
			},
			map[string]*dynamodb.AttributeValue{
				"a": &dynamodb.AttributeValue{
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
							"b": &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
								"e": &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
									"field": &dynamodb.AttributeValue{S: aws.String("value")},
								}},
							}},
						}},
						&dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
							"b": &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
								"c": &dynamodb.AttributeValue{N: aws.String("4")},
								"d": &dynamodb.AttributeValue{N: aws.String("2")},
							}},
						}},
					},
				},
			},
		},
		{
			"a[4],a[2],b.c[12]", nil,
			map[int]*dynamodb.AttributeValue{
				2: &dynamodb.AttributeValue{L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{S: aws.String("elem")},
				}},
				0: &dynamodb.AttributeValue{N: aws.String("4")},
				1: &dynamodb.AttributeValue{N: aws.String("2")},
			},
			map[string]*dynamodb.AttributeValue{
				"a": &dynamodb.AttributeValue{
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{N: aws.String("2")},
						&dynamodb.AttributeValue{N: aws.String("4")},
					},
				},
				"b": &dynamodb.AttributeValue{
					M: map[string]*dynamodb.AttributeValue{
						"c": &dynamodb.AttributeValue{
							L: []*dynamodb.AttributeValue{
								&dynamodb.AttributeValue{
									L: []*dynamodb.AttributeValue{
										&dynamodb.AttributeValue{S: aws.String("elem")},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"#a[1].#b",
			map[string]*string{
				"#a": aws.String("with.dot"),
				"#b": aws.String("sub.field"),
			},
			map[int]*dynamodb.AttributeValue{
				0: &dynamodb.AttributeValue{N: aws.String("4")},
			},
			map[string]*dynamodb.AttributeValue{
				"with.dot": &dynamodb.AttributeValue{
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{
							M: map[string]*dynamodb.AttributeValue{
								"sub.field": &dynamodb.AttributeValue{N: aws.String("4")},
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
