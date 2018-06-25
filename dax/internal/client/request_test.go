package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"testing"
)

func TestHasDuplicatesWriteRequests(t *testing.T) {
	hk := "hk"
	d := []dynamodb.AttributeDefinition{
		{AttributeName: aws.String(hk), AttributeType: aws.String(dynamodb.ScalarAttributeTypeS)},
	}
	cases := []struct {
		w []*dynamodb.WriteRequest
		e bool
	}{
		{
			w: nil,
			e: false,
		},
		{
			w: []*dynamodb.WriteRequest{},
			e: false,
		},
		{
			w: []*dynamodb.WriteRequest{nil},
			e: false,
		},
		{
			w: []*dynamodb.WriteRequest{nil, nil, nil},
			e: false, // continue with request processing
		},
		{
			w: []*dynamodb.WriteRequest{
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("abc")}}}},
			},
			e: false,
		},
		{
			w: []*dynamodb.WriteRequest{
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("abc")}}}},
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("abc")}}}},
			},
			e: true,
		},
		{
			w: []*dynamodb.WriteRequest{
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("abc")}}}},
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("def")}}}},
			},
			e: false,
		},
		{
			w: []*dynamodb.WriteRequest{
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("abc")}}}},
				{DeleteRequest: &dynamodb.DeleteRequest{Key: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("abc")}}}},
			},
			e: true,
		},
		{
			w: []*dynamodb.WriteRequest{
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("abc")}}}},
				{DeleteRequest: &dynamodb.DeleteRequest{Key: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("def")}}}},
			},
			e: false,
		},
		{
			w: []*dynamodb.WriteRequest{
				{DeleteRequest: &dynamodb.DeleteRequest{Key: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("abc")}}}},
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("def")}}}},
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("xyz")}}}},
				{DeleteRequest: &dynamodb.DeleteRequest{Key: map[string]*dynamodb.AttributeValue{hk: {S: aws.String("def")}}}},
			},
			e: true,
		},
	}

	for _, c := range cases {
		a := hasDuplicatesWriteRequests(c.w, d)
		if a != c.e {
			t.Errorf("expected TestHasDuplicatesWriteRequests(%v)=%v, got %v", c.w, c.e, a)
		}
	}
}

func TestHasDuplicateKeysAndAttributes(t *testing.T) {
	hk := "hk"
	d := []dynamodb.AttributeDefinition{
		{AttributeName: aws.String(hk), AttributeType: aws.String(dynamodb.ScalarAttributeTypeS)},
	}
	cases := []struct {
		kaas *dynamodb.KeysAndAttributes
		e    bool
	}{
		{
			kaas: nil,
			e:    false,
		},
		{
			kaas: &dynamodb.KeysAndAttributes{},
			e:    false,
		},
		{
			kaas: &dynamodb.KeysAndAttributes{Keys: []map[string]*dynamodb.AttributeValue{}},
			e:    false,
		},
		{
			kaas: &dynamodb.KeysAndAttributes{Keys: []map[string]*dynamodb.AttributeValue{nil}},
			e:    false,
		},
		{
			kaas: &dynamodb.KeysAndAttributes{Keys: []map[string]*dynamodb.AttributeValue{nil, nil, nil}},
			e:    false, // continue with request processing
		},
		{
			kaas: &dynamodb.KeysAndAttributes{Keys: []map[string]*dynamodb.AttributeValue{
				{hk: {S: aws.String("abc")}},
			}},
			e: false,
		},
		{
			kaas: &dynamodb.KeysAndAttributes{Keys: []map[string]*dynamodb.AttributeValue{
				{hk: {S: aws.String("abc")}},
				{hk: {S: aws.String("def")}},
			}},
			e: false,
		},
		{
			kaas: &dynamodb.KeysAndAttributes{Keys: []map[string]*dynamodb.AttributeValue{
				{hk: {S: aws.String("abc")}},
				{hk: {S: aws.String("abc")}},
			}},
			e: true,
		},
		{
			kaas: &dynamodb.KeysAndAttributes{Keys: []map[string]*dynamodb.AttributeValue{
				{hk: {S: aws.String("abc")}},
				{hk: {S: aws.String("def")}},
				{hk: {S: aws.String("abc")}},
			}},
			e: true,
		},
	}
	for _, c := range cases {
		a := hasDuplicateKeysAndAttributes(c.kaas, d)
		if a != c.e {
			t.Errorf("expected hasDuplicateKeysAndAttributes(%v)=%v, got %v", c.kaas, c.e, a)
		}
	}
}

func reverse(a []interface{}) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}
