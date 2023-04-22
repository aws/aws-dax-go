/*
A stub implementation of the DaxAPI interface that can be configured to
return a series of responses to BatchGetItem, Query, and Scan requests. It
is used to test pagination logic.
*/
package client

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

type ClientStub struct {
	batchGetItemRequests  []*dynamodb.BatchGetItemInput
	batchGetItemResponses []*dynamodb.BatchGetItemOutput
	queryRequests         []*dynamodb.QueryInput
	queryResponses        []*dynamodb.QueryOutput
	scanRequests          []*dynamodb.ScanInput
	scanResponses         []*dynamodb.ScanOutput
}

// Constructor
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
func (stub *ClientStub) PutItemWithOptions(input *dynamodb.PutItemInput, output *dynamodb.PutItemOutput, opt RequestOptions) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) DeleteItemWithOptions(input *dynamodb.DeleteItemInput, output *dynamodb.DeleteItemOutput, opt RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) UpdateItemWithOptions(input *dynamodb.UpdateItemInput, output *dynamodb.UpdateItemOutput, opt RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) GetItemWithOptions(input *dynamodb.GetItemInput, output *dynamodb.GetItemOutput, opt RequestOptions) (*dynamodb.GetItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) ScanWithOptions(input *dynamodb.ScanInput, output *dynamodb.ScanOutput, opt RequestOptions) (*dynamodb.ScanOutput, error) {
	output, stub.scanResponses = stub.scanResponses[0], stub.scanResponses[1:]
	return output, nil
}

func (stub *ClientStub) QueryWithOptions(input *dynamodb.QueryInput, output *dynamodb.QueryOutput, opt RequestOptions) (*dynamodb.QueryOutput, error) {
	output, stub.queryResponses = stub.queryResponses[0], stub.queryResponses[1:]
	return output, nil
}

func (stub *ClientStub) BatchWriteItemWithOptions(input *dynamodb.BatchWriteItemInput, output *dynamodb.BatchWriteItemOutput, opt RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, nil
}

func (stub *ClientStub) BatchGetItemWithOptions(input *dynamodb.BatchGetItemInput, output *dynamodb.BatchGetItemOutput, opt RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	output, stub.batchGetItemResponses = stub.batchGetItemResponses[0], stub.batchGetItemResponses[1:]
	return output, nil
}

func (stub *ClientStub) TransactWriteItemsWithOptions(input *dynamodb.TransactWriteItemsInput, output *dynamodb.TransactWriteItemsOutput, opt RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, nil
}

func (stub *ClientStub) TransactGetItemsWithOptions(input *dynamodb.TransactGetItemsInput, output *dynamodb.TransactGetItemsOutput, opt RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	return nil, nil
}

func (stub *ClientStub) NewDaxRequest(op *request.Operation, input, output interface{}, opt RequestOptions) *request.Request {
	h := request.Handlers{}
	h.Build.PushFrontNamed(request.NamedHandler{Name: "dax.BuildHandler", Fn: stub.build})
	h.Send.PushFrontNamed(request.NamedHandler{Name: "dax.SendHandler", Fn: stub.send})

	req := request.New(aws.Config{}, clientInfo, h, nil, op, input, output)
	opt.applyTo(req)
	return req
}

func (stub *ClientStub) build(req *request.Request) {
}

func (stub *ClientStub) send(req *request.Request) {
	opt := RequestOptions{}
	switch req.Operation.Name {
	case OpBatchGetItem:
		input, _ := req.Params.(*dynamodb.BatchGetItemInput)
		stub.batchGetItemRequests = append(stub.batchGetItemRequests, input)
		output, _ := req.Data.(*dynamodb.BatchGetItemOutput)
		req.Data, req.Error = stub.BatchGetItemWithOptions(input, output, opt)
	case OpQuery:
		input, _ := req.Params.(*dynamodb.QueryInput)
		stub.queryRequests = append(stub.queryRequests, input)
		output, _ := req.Data.(*dynamodb.QueryOutput)
		req.Data, req.Error = stub.QueryWithOptions(input, output, opt)
	case OpScan:
		input, _ := req.Params.(*dynamodb.ScanInput)
		stub.scanRequests = append(stub.scanRequests, input)
		output, _ := req.Data.(*dynamodb.ScanOutput)
		req.Data, req.Error = stub.ScanWithOptions(input, output, opt)
	}
}

func (stub *ClientStub) endpoints(opt RequestOptions) ([]serviceEndpoint, error) {
	return nil, nil
}
