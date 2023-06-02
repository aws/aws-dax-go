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
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type documentPathElement struct {
	index int
	name  string
}

type documentPath struct {
	elements []documentPathElement
}

func documentPathElementFromIndex(idx int) documentPathElement {
	return documentPathElement{index: idx, name: ""}
}

func documentPathElementFromName(nm string) documentPathElement {
	return documentPathElement{index: -1, name: nm}
}

func buildProjectionOrdinals(projectionExpression *string, expressionAttributeNames map[string]string) ([]documentPath, error) {
	if projectionExpression == nil || *projectionExpression == "" {
		return nil, nil
	}
	terms := strings.Split(*projectionExpression, ",")
	dps := make([]documentPath, 0, len(terms))
	for _, t := range terms {
		dp, err := buildDocumentPath(strings.TrimSpace(t), expressionAttributeNames)
		if err != nil {
			return nil, err
		}
		dps = append(dps, dp)
	}
	return dps, nil
}

func buildDocumentPath(path string, expressionAttributeNames map[string]string) (documentPath, error) {
	var substitutes map[string]string
	if expressionAttributeNames != nil {
		substitutes = expressionAttributeNames
	}

	res := strings.Split(path, ".")
	var elements []documentPathElement

	for _, re := range res {
		idx := strings.Index(re, "[")
		if idx == -1 {
			elements = append(elements, documentPathElementFromName(getOrDefault(substitutes, re, re)))
			continue
		}

		if idx == 0 {
			return documentPath{}, errors.New("invalid path: " + path)
		}

		pre := re[0:idx]
		elements = append(elements, documentPathElementFromName(getOrDefault(substitutes, pre, pre)))

		for idx != -1 {
			re = re[idx+1:]
			idx = strings.Index(re, "]")

			if idx == -1 {
				return documentPath{}, errors.New("invalid path: " + path)
			}

			lidx, err := strconv.Atoi(re[:idx])
			if err != nil {
				return documentPath{}, err
			}
			elements = append(elements, documentPathElementFromIndex(lidx))

			re = re[idx+1:]
			idx = strings.Index(re, "[")
			if idx > 0 {
				return documentPath{}, errors.New("invalid path: " + path)
			}
		}

		if len(elements) == 0 {
			return documentPath{}, errors.New("invalid path: " + path)
		}
	}

	return documentPath{elements: elements}, nil
}

func getOrDefault(m map[string]string, key, value string) string {
	v, ok := m[key]
	if ok {
		return v
	} else {
		return value
	}
}

type itemNode struct {
	children map[documentPathElement]*itemNode
	value    types.AttributeValue
}

func (in *itemNode) toAttribute() types.AttributeValue {
	if in == nil {
		return nil
	}
	if in.value != nil {
		return in.value
	}
	var val types.AttributeValue
	if len(in.children) != 0 {
		m := in.children
		keys := make([]documentPathElement, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		// it is safe to assume keys do not have a mix of array and map indexes as the request succeeded in DynamoDB at this point
		if keys[0].index < 0 {
			attr := types.AttributeValueMemberM{
				Value: make(map[string]types.AttributeValue, len(keys)),
			}
			for k, v := range m {
				attr.Value[k.name] = v.toAttribute()
			}
			val = &attr
		} else {
			// order in the response item should be same as order in actual item, not the one in projection expression
			// eg: projection expression "list[1],list[0]" returns "list[valAt(0),valAt(1)]
			sort.Slice(keys, func(i, j int) bool { return keys[i].index < keys[j].index })
			attr := types.AttributeValueMemberL{
				Value: make([]types.AttributeValue, 0, len(keys)),
			}
			for _, k := range keys {
				attr.Value = append(attr.Value, m[k].toAttribute())
			}
			val = &attr
		}
	}
	return val
}

type itemBuilder struct {
	root *itemNode
}

func (ib *itemBuilder) insert(path documentPath, value types.AttributeValue) {
	if ib.root == nil {
		var children map[documentPathElement]*itemNode
		ib.root = &itemNode{children: children}
	}
	ib.insertNode(ib.root, path.elements, value)
}

func (ib *itemBuilder) toItem() map[string]types.AttributeValue {
	items := make(map[string]types.AttributeValue)
	if ib.root == nil {
		return items
	}

	c := ib.root.children
	for k, v := range c { // top level attribute names are strings
		items[k.name] = v.toAttribute()
	}
	return items
}

func (ib *itemBuilder) insertNode(node *itemNode, elements []documentPathElement, value types.AttributeValue) {
	if len(elements) == 0 {
		node.value = value
		return
	}

	e := elements[0]
	if node.children == nil {
		node.children = make(map[documentPathElement]*itemNode)
	}

	n, ok := node.children[e]
	if !ok {
		n = &itemNode{}
		node.children[e] = n
	}
	ib.insertNode(n, elements[1:], value)
}
