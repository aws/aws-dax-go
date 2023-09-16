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
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DaxAPI interface {
	PutItemWithOptions(ctx context.Context, input *dynamodb.PutItemInput, opt RequestOptions) (*dynamodb.PutItemOutput, error)
	DeleteItemWithOptions(ctx context.Context, input *dynamodb.DeleteItemInput, opt RequestOptions) (*dynamodb.DeleteItemOutput, error)
	UpdateItemWithOptions(ctx context.Context, input *dynamodb.UpdateItemInput, opt RequestOptions) (*dynamodb.UpdateItemOutput, error)

	GetItemWithOptions(ctx context.Context, input *dynamodb.GetItemInput, opt RequestOptions) (*dynamodb.GetItemOutput, error)
	ScanWithOptions(ctx context.Context, input *dynamodb.ScanInput, opt RequestOptions) (*dynamodb.ScanOutput, error)
	QueryWithOptions(ctx context.Context, input *dynamodb.QueryInput, opt RequestOptions) (*dynamodb.QueryOutput, error)

	BatchWriteItemWithOptions(ctx context.Context, input *dynamodb.BatchWriteItemInput, opt RequestOptions) (*dynamodb.BatchWriteItemOutput, error)
	BatchGetItemWithOptions(ctx context.Context, input *dynamodb.BatchGetItemInput, opt RequestOptions) (*dynamodb.BatchGetItemOutput, error)

	TransactWriteItemsWithOptions(ctx context.Context, input *dynamodb.TransactWriteItemsInput, opt RequestOptions) (*dynamodb.TransactWriteItemsOutput, error)
	TransactGetItemsWithOptions(ctx context.Context, input *dynamodb.TransactGetItemsInput, opt RequestOptions) (*dynamodb.TransactGetItemsOutput, error)

	endpoints(ctx context.Context, opt RequestOptions) ([]serviceEndpoint, error)
}
