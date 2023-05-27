/*
A stub implementation of the DaxAPI interface that can be configured to
return a series of responses to BatchGetItem, Query, and Scan requests. It
is used to test pagination logic.
*/
package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ClientStub struct {
	batchGetItemRequests  []*dynamodb.BatchGetItemInput
	batchGetItemResponses []*dynamodb.BatchGetItemOutput
	queryRequests         []*dynamodb.QueryInput
	queryResponses        []*dynamodb.QueryOutput
	scanRequests          []*dynamodb.ScanInput
	scanResponses         []*dynamodb.ScanOutput
}

var _ DaxAPI = (*ClientStub)(nil)

// NewClientStub creates ClientStub.
func NewClientStub(batchGetItemResponses []*dynamodb.BatchGetItemOutput, queryResponses []*dynamodb.QueryOutput, scanResponses []*dynamodb.ScanOutput) *ClientStub {
	return &ClientStub{batchGetItemResponses: batchGetItemResponses, queryResponses: queryResponses, scanResponses: scanResponses}
}

// Stub methods

func (stub *ClientStub) GetBatchGetItemRequests() []*dynamodb.BatchGetItemInput {
	return stub.batchGetItemRequests
}

func (stub *ClientStub) GetQueryRequests() []*dynamodb.QueryInput {
	return stub.queryRequests
}

func (stub *ClientStub) GetScanRequests() []*dynamodb.ScanInput {
	return stub.scanRequests
}

// DaxAPI methods

func (stub *ClientStub) PutItemWithOptions(_ context.Context, _ *dynamodb.PutItemInput, _ RequestOptions) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) DeleteItemWithOptions(_ context.Context, _ *dynamodb.DeleteItemInput, _ RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) UpdateItemWithOptions(_ context.Context, _ *dynamodb.UpdateItemInput, _ RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) GetItemWithOptions(_ context.Context, _ *dynamodb.GetItemInput, _ RequestOptions) (*dynamodb.GetItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) ScanWithOptions(_ context.Context, _ *dynamodb.ScanInput, _ RequestOptions) (*dynamodb.ScanOutput, error) {
	output := stub.scanResponses[0]
	stub.scanResponses = stub.scanResponses[1:]
	return output, nil
}

func (stub *ClientStub) QueryWithOptions(_ context.Context, _ *dynamodb.QueryInput, _ RequestOptions) (*dynamodb.QueryOutput, error) {
	output := stub.queryResponses[0]
	stub.queryResponses = stub.queryResponses[1:]
	return output, nil
}

func (stub *ClientStub) BatchWriteItemWithOptions(_ context.Context, _ *dynamodb.BatchWriteItemInput, _ RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) BatchGetItemWithOptions(_ context.Context, _ *dynamodb.BatchGetItemInput, _ RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	output := stub.batchGetItemResponses[0]
	stub.batchGetItemResponses = stub.batchGetItemResponses[1:]
	return output, nil
}

func (stub *ClientStub) TransactWriteItemsWithOptions(_ context.Context, _ *dynamodb.TransactWriteItemsInput, _ RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, nil
}

func (stub *ClientStub) TransactGetItemsWithOptions(_ context.Context, _ *dynamodb.TransactGetItemsInput, _ RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	return nil, nil
}

func (stub *ClientStub) endpoints(_ context.Context, _ RequestOptions) ([]serviceEndpoint, error) {
	return nil, nil
}
