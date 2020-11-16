package dax

import (
	"reflect"
	"testing"

	"github.com/aws/aws-dax-go/dax/internal/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewWithInternalClient(c client.DaxAPI) *Dax {
	return &Dax{client: c, config: DefaultConfig()}
}

func TestPaginationBatchGetItemPage(t *testing.T) {
	pages, numPages, gotToEnd := map[string][]map[string]*dynamodb.AttributeValue{}, 0, false

	resps := []*dynamodb.BatchGetItemOutput{
		{
			Responses: map[string][]map[string]*dynamodb.AttributeValue{
				"tablename": {
					{
						"key":  {S: aws.String("key1")},
						"attr": {S: aws.String("attr1")},
					},
					{
						"key":  {S: aws.String("key2")},
						"attr": {S: aws.String("attr2")},
					},
				},
			},
			UnprocessedKeys: map[string]*dynamodb.KeysAndAttributes{
				"tablename": {
					Keys: []map[string]*dynamodb.AttributeValue{
						{"key": {S: aws.String("key3")}},
						{"key": {S: aws.String("key4")}},
						{"key": {S: aws.String("key5")}},
					},
				},
			},
		},
		{
			Responses: map[string][]map[string]*dynamodb.AttributeValue{
				"tablename": {
					{
						"key":  {S: aws.String("key3")},
						"attr": {S: aws.String("attr3")},
					},
					{
						"key":  {S: aws.String("key4")},
						"attr": {S: aws.String("attr4")},
					},
				},
			},
			UnprocessedKeys: map[string]*dynamodb.KeysAndAttributes{
				"tablename": {
					Keys: []map[string]*dynamodb.AttributeValue{
						{"key": {S: aws.String("key5")}},
					},
				},
			},
		},
		{
			Responses: map[string][]map[string]*dynamodb.AttributeValue{
				"tablename": {
					{
						"key":  {S: aws.String("key5")},
						"attr": {S: aws.String("attr5")},
					},
				},
			},
		},
	}

	stub := client.NewClientStub(resps, nil, nil)
	db := NewWithInternalClient(stub)
	params := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			"tablename": {
				Keys: []map[string]*dynamodb.AttributeValue{
					{"key": {S: aws.String("key1")}},
					{"key": {S: aws.String("key2")}},
					{"key": {S: aws.String("key3")}},
					{"key": {S: aws.String("key4")}},
					{"key": {S: aws.String("key5")}},
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
		map[string][]map[string]*dynamodb.AttributeValue{
			"tablename": {
				{"key": {S: aws.String("key1")}, "attr": {S: aws.String("attr1")}},
				{"key": {S: aws.String("key2")}, "attr": {S: aws.String("attr2")}},
				{"key": {S: aws.String("key3")}, "attr": {S: aws.String("attr3")}},
				{"key": {S: aws.String("key4")}, "attr": {S: aws.String("attr4")}},
				{"key": {S: aws.String("key5")}, "attr": {S: aws.String("attr5")}},
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
	if e, a := "key5", *stub.GetBatchGetItemRequests()[2].RequestItems["tablename"].Keys[0]["key"].S; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestPaginationQueryPage(t *testing.T) {
	pages, numPages, gotToEnd := []map[string]*dynamodb.AttributeValue{}, 0, false

	resps := []*dynamodb.QueryOutput{
		{
			LastEvaluatedKey: map[string]*dynamodb.AttributeValue{"key": {S: aws.String("key1")}},
			Count:            aws.Int64(1),
			Items: []map[string]*dynamodb.AttributeValue{
				{
					"key": {S: aws.String("key1")},
				},
			},
		},
		{
			LastEvaluatedKey: map[string]*dynamodb.AttributeValue{"key": {S: aws.String("key2")}},
			Count:            aws.Int64(1),
			Items: []map[string]*dynamodb.AttributeValue{
				{
					"key": {S: aws.String("key2")},
				},
			},
		},
		{
			LastEvaluatedKey: map[string]*dynamodb.AttributeValue{},
			Count:            aws.Int64(1),
			Items: []map[string]*dynamodb.AttributeValue{
				{
					"key": {S: aws.String("key3")},
				},
			},
		},
	}

	stub := client.NewClientStub(nil, resps, nil)
	db := NewWithInternalClient(stub)
	params := &dynamodb.QueryInput{
		Limit:     aws.Int64(1),
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
		[]map[string]*dynamodb.AttributeValue{
			{"key": {S: aws.String("key1")}},
			{"key": {S: aws.String("key2")}},
			{"key": {S: aws.String("key3")}},
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
		if a := *stub.GetQueryRequests()[i+1].ExclusiveStartKey["key"].S; e != a {
			t.Errorf("expect %s, got %s at index %d", e, a, i+1)
		}
	}
}

func TestPaginationScanPage(t *testing.T) {
	pages, numPages, gotToEnd := []map[string]*dynamodb.AttributeValue{}, 0, false

	resps := []*dynamodb.ScanOutput{
		{
			LastEvaluatedKey: map[string]*dynamodb.AttributeValue{"key": {S: aws.String("key1")}},
			Count:            aws.Int64(1),
			Items: []map[string]*dynamodb.AttributeValue{
				{
					"key": {S: aws.String("key1")},
				},
			},
		},
		{
			LastEvaluatedKey: map[string]*dynamodb.AttributeValue{"key": {S: aws.String("key2")}},
			Count:            aws.Int64(1),
			Items: []map[string]*dynamodb.AttributeValue{
				{
					"key": {S: aws.String("key2")},
				},
			},
		},
		{
			LastEvaluatedKey: map[string]*dynamodb.AttributeValue{},
			Count:            aws.Int64(1),
			Items: []map[string]*dynamodb.AttributeValue{
				{
					"key": {S: aws.String("key3")},
				},
			},
		},
	}

	stub := client.NewClientStub(nil, nil, resps)
	db := NewWithInternalClient(stub)
	params := &dynamodb.ScanInput{
		Limit:     aws.Int64(1),
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
		[]map[string]*dynamodb.AttributeValue{
			{"key": {S: aws.String("key1")}},
			{"key": {S: aws.String("key2")}},
			{"key": {S: aws.String("key3")}},
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
		if a := *stub.GetScanRequests()[i+1].ExclusiveStartKey["key"].S; e != a {
			t.Errorf("expect %s, got %s at index %d", e, a, i+1)
		}
	}
}
