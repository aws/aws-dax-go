package dax

import (
	"reflect"
	"testing"

	"github.com/aws/aws-dax-go/dax/internal/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func NewWithInternalClient(c client.DaxAPI) *Dax {
	return &Dax{client: c, config: DefaultConfig()}
}

func TestPaginationBatchGetItemPage(t *testing.T) {
	pages, numPages, gotToEnd := map[string][]map[string]types.AttributeValue{}, 0, false

	resps := []*dynamodb.BatchGetItemOutput{
		{
			Responses: map[string][]map[string]types.AttributeValue{
				"tablename": {
					{
						"key":  &types.AttributeValueMemberS{Value: "key1"},
						"attr": &types.AttributeValueMemberS{Value: "attr1"},
					},
					{
						"key":  &types.AttributeValueMemberS{Value: "key2"},
						"attr": &types.AttributeValueMemberS{Value: "attr2"},
					},
				},
			},
			UnprocessedKeys: map[string]types.KeysAndAttributes{
				"tablename": {
					Keys: []map[string]types.AttributeValue{
						{"key": &types.AttributeValueMemberS{Value: "key3"}},
						{"key": &types.AttributeValueMemberS{Value: "key4"}},
						{"key": &types.AttributeValueMemberS{Value: "key5"}},
					},
				},
			},
		},
		{
			Responses: map[string][]map[string]types.AttributeValue{
				"tablename": {
					{
						"key":  &types.AttributeValueMemberS{Value: "key3"},
						"attr": &types.AttributeValueMemberS{Value: "attr3"},
					},
					{
						"key":  &types.AttributeValueMemberS{Value: "key4"},
						"attr": &types.AttributeValueMemberS{Value: "attr4"},
					},
				},
			},
			UnprocessedKeys: map[string]types.KeysAndAttributes{
				"tablename": {
					Keys: []map[string]types.AttributeValue{
						{"key": &types.AttributeValueMemberS{Value: "key5"}},
					},
				},
			},
		},
		{
			Responses: map[string][]map[string]types.AttributeValue{
				"tablename": {
					{
						"key":  &types.AttributeValueMemberS{Value: "key5"},
						"attr": &types.AttributeValueMemberS{Value: "attr5"},
					},
				},
			},
		},
	}

	stub := client.NewClientStub(resps, nil, nil)
	db := NewWithInternalClient(stub)
	params := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			"tablename": {
				Keys: []map[string]types.AttributeValue{
					{"key": &types.AttributeValueMemberS{Value: "key1"}},
					{"key": &types.AttributeValueMemberS{Value: "key2"}},
					{"key": &types.AttributeValueMemberS{Value: "key3"}},
					{"key": &types.AttributeValueMemberS{Value: "key4"}},
					{"key": &types.AttributeValueMemberS{Value: "key5"}},
				},
			},
		},
	}
	err := db.BatchGetItemPages(params, func(p *dynamodb.BatchGetItemOutput, last bool) bool {
		numPages++
		for k, v := range p.Responses {
			pages[k] = append(pages[k], v...)
		}
		if last {
			if gotToEnd {
				t.Errorf("last=true happened twice")
			}
			gotToEnd = true
		}
		return true
	})

	// There was no error
	if err != nil {
		t.Errorf("expect nil, %v", err)
	}

	// The items were all returned
	if e, a :=
		map[string][]map[string]types.AttributeValue{
			"tablename": {
				{"key": &types.AttributeValueMemberS{Value: "key1"}, "attr": &types.AttributeValueMemberS{Value: "attr1"}},
				{"key": &types.AttributeValueMemberS{Value: "key2"}, "attr": &types.AttributeValueMemberS{Value: "attr2"}},
				{"key": &types.AttributeValueMemberS{Value: "key3"}, "attr": &types.AttributeValueMemberS{Value: "attr3"}},
				{"key": &types.AttributeValueMemberS{Value: "key4"}, "attr": &types.AttributeValueMemberS{Value: "attr4"}},
				{"key": &types.AttributeValueMemberS{Value: "key5"}, "attr": &types.AttributeValueMemberS{Value: "attr5"}},
			}}, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expect %v, got %v", e, a)
	}

	// The results were returned in the expected number of pages
	if e, a := 3, numPages; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	// The last page was signaled
	if !gotToEnd {
		t.Errorf("expect true")
	}

	// Each request had the correct number of keys
	for i, e := range []int{5, 3, 1} {
		a := len(stub.GetBatchGetItemRequests()[i].RequestItems["tablename"].Keys)
		if e != a {
			t.Errorf("expect %v, got %v at index %d", e, a, i)
		}
	}

	// The last request had the correct key
	if e, a := "key5", stub.GetBatchGetItemRequests()[2].RequestItems["tablename"].Keys[0]["key"].(*types.AttributeValueMemberS).Value; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestPaginationQueryPage(t *testing.T) {
	var pages []map[string]types.AttributeValue
	numPages, gotToEnd := 0, false

	resps := []*dynamodb.QueryOutput{
		{
			LastEvaluatedKey: map[string]types.AttributeValue{"key": &types.AttributeValueMemberS{Value: ("key1")}},
			Count:            1,
			Items: []map[string]types.AttributeValue{
				{
					"key": &types.AttributeValueMemberS{Value: "key1"},
				},
			},
		},
		{
			LastEvaluatedKey: map[string]types.AttributeValue{"key": &types.AttributeValueMemberS{Value: "key2"}},
			Count:            1,
			Items: []map[string]types.AttributeValue{
				{
					"key": &types.AttributeValueMemberS{Value: "key2"},
				},
			},
		},
		{
			LastEvaluatedKey: map[string]types.AttributeValue{},
			Count:            1,
			Items: []map[string]types.AttributeValue{
				{
					"key": &types.AttributeValueMemberS{Value: "key3"},
				},
			},
		},
	}

	stub := client.NewClientStub(nil, resps, nil)
	db := NewWithInternalClient(stub)
	params := &dynamodb.QueryInput{
		Limit:     aws.Int32(1),
		TableName: aws.String("tablename"),
	}
	err := db.QueryPages(params, func(p *dynamodb.QueryOutput, last bool) bool {
		numPages++
		pages = append(pages, p.Items...)
		if last {
			if gotToEnd {
				t.Errorf("last=true happened twice")
			}
			gotToEnd = true
		}
		return true
	})

	// There was no error
	if err != nil {
		t.Errorf("expect nil, %v", err)
	}

	// The correct items were returned
	if e, a :=
		[]map[string]types.AttributeValue{
			{"key": &types.AttributeValueMemberS{Value: "key1"}},
			{"key": &types.AttributeValueMemberS{Value: "key2"}},
			{"key": &types.AttributeValueMemberS{Value: "key3"}},
		}, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expect %v, got %v", e, a)
	}

	// Items were returned in the correct number of pages
	if e, a := 3, numPages; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	// The last page was signaled
	if !gotToEnd {
		t.Errorf("expect true")
	}

	// Each request had the correct start key
	if a := stub.GetQueryRequests()[0].ExclusiveStartKey; a != nil {
		t.Errorf("expect nil, %v", a)
	}
	for i, e := range []string{"key1", "key2"} {
		if a := stub.GetQueryRequests()[i+1].ExclusiveStartKey["key"].(*types.AttributeValueMemberS).Value; e != a {
			t.Errorf("expect %s, got %s at index %d", e, a, i+1)
		}
	}
}

func TestPaginationScanPage(t *testing.T) {
	var pages []map[string]types.AttributeValue
	numPages, gotToEnd := 0, false

	resps := []*dynamodb.ScanOutput{
		{
			LastEvaluatedKey: map[string]types.AttributeValue{"key": &types.AttributeValueMemberS{Value: "key1"}},
			Count:            1,
			Items: []map[string]types.AttributeValue{
				{
					"key": &types.AttributeValueMemberS{Value: "key1"},
				},
			},
		},
		{
			LastEvaluatedKey: map[string]types.AttributeValue{"key": &types.AttributeValueMemberS{Value: "key2"}},
			Count:            1,
			Items: []map[string]types.AttributeValue{
				{
					"key": &types.AttributeValueMemberS{Value: "key2"},
				},
			},
		},
		{
			LastEvaluatedKey: map[string]types.AttributeValue{},
			Count:            1,
			Items: []map[string]types.AttributeValue{
				{
					"key": &types.AttributeValueMemberS{Value: "key3"},
				},
			},
		},
	}

	stub := client.NewClientStub(nil, nil, resps)
	db := NewWithInternalClient(stub)
	params := &dynamodb.ScanInput{
		Limit:     aws.Int32(1),
		TableName: aws.String("tablename"),
	}
	err := db.ScanPages(params, func(p *dynamodb.ScanOutput, last bool) bool {
		numPages++
		pages = append(pages, p.Items...)
		if last {
			if gotToEnd {
				t.Errorf("last=true happened twice")
			}
			gotToEnd = true
		}
		return true
	})

	// There was no error
	if err != nil {
		t.Errorf("expect nil, %v", err)
	}

	// The correct items were returned
	if e, a :=
		[]map[string]types.AttributeValue{
			{"key": &types.AttributeValueMemberS{Value: "key1"}},
			{"key": &types.AttributeValueMemberS{Value: "key2"}},
			{"key": &types.AttributeValueMemberS{Value: "key3"}},
		}, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expect %v, got %v", e, a)
	}

	// Items were returned in the correct number of pages
	if e, a := 3, numPages; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	// Last page was signaled
	if !gotToEnd {
		t.Errorf("expect true")
	}
	// Each request had the correct start key
	if a := stub.GetScanRequests()[0].ExclusiveStartKey; a != nil {
		t.Errorf("expect nil, %v", a)
	}
	for i, e := range []string{"key1", "key2"} {
		if a := stub.GetScanRequests()[i+1].ExclusiveStartKey["key"].(*types.AttributeValueMemberS).Value; e != a {
			t.Errorf("expect %s, got %s at index %d", e, a, i+1)
		}
	}
}
