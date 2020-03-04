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

package dax

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dmartin1/aws-dax-go/dax/internal/client"
)

func (d *Dax) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return d.PutItemWithContext(nil, input)
}

func (d *Dax) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.PutItemWithOptions(input, &dynamodb.PutItemOutput{}, o)
}

func (d *Dax) PutItemRequest(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	op := &request.Operation{Name: client.OpPutItem}
	if input == nil {
		input = &dynamodb.PutItemInput{}
	}
	output := &dynamodb.PutItemOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return d.DeleteItemWithContext(nil, input)
}

func (d *Dax) DeleteItemWithContext(ctx aws.Context, input *dynamodb.DeleteItemInput, opts ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.DeleteItemWithOptions(input, &dynamodb.DeleteItemOutput{}, o)
}

func (d *Dax) DeleteItemRequest(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	op := &request.Operation{Name: client.OpDeleteItem}
	if input == nil {
		input = &dynamodb.DeleteItemInput{}
	}
	output := &dynamodb.DeleteItemOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return d.UpdateItemWithContext(nil, input)
}

func (d *Dax) UpdateItemWithContext(ctx aws.Context, input *dynamodb.UpdateItemInput, opts ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.UpdateItemWithOptions(input, &dynamodb.UpdateItemOutput{}, o)
}

func (d *Dax) UpdateItemRequest(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	op := &request.Operation{Name: client.OpUpdateItem}
	if input == nil {
		input = &dynamodb.UpdateItemInput{}
	}
	output := &dynamodb.UpdateItemOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return d.GetItemWithContext(nil, input)
}

func (d *Dax) GetItemWithContext(ctx aws.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.GetItemWithOptions(input, &dynamodb.GetItemOutput{}, o)
}

func (d *Dax) GetItemRequest(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	op := &request.Operation{Name: client.OpGetItem}
	if input == nil {
		input = &dynamodb.GetItemInput{}
	}
	output := &dynamodb.GetItemOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return d.ScanWithContext(nil, input)
}

func (d *Dax) ScanWithContext(ctx aws.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.ScanWithOptions(input, &dynamodb.ScanOutput{}, o)
}

func (d *Dax) ScanRequest(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	op := &request.Operation{Name: client.OpScan}
	if input == nil {
		input = &dynamodb.ScanInput{}
	}
	output := &dynamodb.ScanOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return d.QueryWithContext(nil, input)
}

func (d *Dax) QueryWithContext(ctx aws.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.QueryWithOptions(input, &dynamodb.QueryOutput{}, o)
}

func (d *Dax) QueryRequest(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	op := &request.Operation{Name: client.OpQuery}
	if input == nil {
		input = &dynamodb.QueryInput{}
	}
	output := &dynamodb.QueryOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return d.BatchWriteItemWithContext(nil, input)
}

func (d *Dax) BatchWriteItemWithContext(ctx aws.Context, input *dynamodb.BatchWriteItemInput, opts ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.BatchWriteItemWithOptions(input, &dynamodb.BatchWriteItemOutput{}, o)
}

func (d *Dax) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	op := &request.Operation{Name: client.OpBatchWriteItem}
	if input == nil {
		input = &dynamodb.BatchWriteItemInput{}
	}
	output := &dynamodb.BatchWriteItemOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return d.BatchGetItemWithContext(nil, input)
}

func (d *Dax) BatchGetItemWithContext(ctx aws.Context, input *dynamodb.BatchGetItemInput, opts ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.BatchGetItemWithOptions(input, &dynamodb.BatchGetItemOutput{}, o)
}

func (d *Dax) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	op := &request.Operation{Name: client.OpBatchGetItem}
	if input == nil {
		input = &dynamodb.BatchGetItemInput{}
	}
	output := &dynamodb.BatchGetItemOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	return d.TransactWriteItemsWithContext(nil, input)
}

func (d *Dax) TransactWriteItemsWithContext(ctx aws.Context, input *dynamodb.TransactWriteItemsInput, opts ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	o, cfn, err := d.config.requestOptions(false, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.TransactWriteItemsWithOptions(input, &dynamodb.TransactWriteItemsOutput{}, o)
}

func (d *Dax) TransactWriteItemsRequest(input *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	op := &request.Operation{Name: client.OpTransactWriteItems}
	if input == nil {
		input = &dynamodb.TransactWriteItemsInput{}
	}
	output := &dynamodb.TransactWriteItemsOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) TransactGetItems(input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	return d.TransactGetItemsWithContext(nil, input)
}

func (d *Dax) TransactGetItemsWithContext(ctx aws.Context, input *dynamodb.TransactGetItemsInput, opts ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	o, cfn, err := d.config.requestOptions(true, ctx, opts...)
	if err != nil {
		return nil, err
	}
	if cfn != nil {
		defer cfn()
	}
	return d.client.TransactGetItemsWithOptions(input, &dynamodb.TransactGetItemsOutput{}, o)
}

func (d *Dax) TransactGetItemsRequest(input *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	op := &request.Operation{Name: client.OpTransactGetItems}
	if input == nil {
		input = &dynamodb.TransactGetItemsInput{}
	}
	output := &dynamodb.TransactGetItemsOutput{}
	opt := client.RequestOptions{Context: aws.BackgroundContext()}
	req := d.client.NewDaxRequest(op, input, output, opt)
	return req, output
}

func (d *Dax) BatchGetItemPages(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	return d.unImpl()
}

func (d *Dax) BatchGetItemPagesWithContext(aws.Context, *dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool, ...request.Option) error {
	return d.unImpl()
}

func (d *Dax) QueryPages(input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error {
	return d.QueryPagesWithContext(aws.BackgroundContext(), input, fn)
}

func (d *Dax) QueryPagesWithContext(ctx aws.Context, input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *dynamodb.QueryInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := d.QueryRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	for p.Next() {
		if !fn(p.Page().(*dynamodb.QueryOutput), !p.HasNextPage()) {
			break
		}
	}
	return p.Err()
}

func (d *Dax) ScanPages(input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool) error {
	return d.ScanPagesWithContext(aws.BackgroundContext(), input, fn)
}

func (d *Dax) ScanPagesWithContext(ctx aws.Context, input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *dynamodb.ScanInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := d.ScanRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	for p.Next() {
		if !fn(p.Page().(*dynamodb.ScanOutput), !p.HasNextPage()) {
			break
		}
	}
	return p.Err()
}

func (d *Dax) CreateBackup(*dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateBackupWithContext(aws.Context, *dynamodb.CreateBackupInput, ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateBackupRequest(*dynamodb.CreateBackupInput) (r *request.Request, o *dynamodb.CreateBackupOutput) {
	d.unImpl()
	return
}

func (d *Dax) CreateGlobalTable(*dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateGlobalTableWithContext(aws.Context, *dynamodb.CreateGlobalTableInput, ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateGlobalTableRequest(*dynamodb.CreateGlobalTableInput) (r *request.Request, o *dynamodb.CreateGlobalTableOutput) {
	d.unImpl()
	return
}

func (d *Dax) CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateTableWithContext(aws.Context, *dynamodb.CreateTableInput, ...request.Option) (*dynamodb.CreateTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) CreateTableRequest(*dynamodb.CreateTableInput) (r *request.Request, o *dynamodb.CreateTableOutput) {
	d.unImpl()
	return
}

func (d *Dax) DeleteBackup(*dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteBackupWithContext(aws.Context, *dynamodb.DeleteBackupInput, ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteBackupRequest(*dynamodb.DeleteBackupInput) (r *request.Request, o *dynamodb.DeleteBackupOutput) {
	d.unImpl()
	return
}

func (d *Dax) DeleteTable(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteTableWithContext(aws.Context, *dynamodb.DeleteTableInput, ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DeleteTableRequest(*dynamodb.DeleteTableInput) (r *request.Request, o *dynamodb.DeleteTableOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeBackup(*dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeBackupWithContext(aws.Context, *dynamodb.DescribeBackupInput, ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeBackupRequest(*dynamodb.DescribeBackupInput) (r *request.Request, o *dynamodb.DescribeBackupOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeContinuousBackups(*dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeContinuousBackupsWithContext(aws.Context, *dynamodb.DescribeContinuousBackupsInput, ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeContinuousBackupsRequest(*dynamodb.DescribeContinuousBackupsInput) (r *request.Request, o *dynamodb.DescribeContinuousBackupsOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeContributorInsights(*dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeContributorInsightsWithContext(aws.Context, *dynamodb.DescribeContributorInsightsInput, ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeContributorInsightsRequest(*dynamodb.DescribeContributorInsightsInput) (r *request.Request, o *dynamodb.DescribeContributorInsightsOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeEndpoints(*dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeEndpointsWithContext(aws.Context, *dynamodb.DescribeEndpointsInput, ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeEndpointsRequest(*dynamodb.DescribeEndpointsInput) (r *request.Request, o *dynamodb.DescribeEndpointsOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeGlobalTable(*dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeGlobalTableWithContext(aws.Context, *dynamodb.DescribeGlobalTableInput, ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeGlobalTableRequest(*dynamodb.DescribeGlobalTableInput) (r *request.Request, o *dynamodb.DescribeGlobalTableOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeGlobalTableSettings(*dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeGlobalTableSettingsWithContext(aws.Context, *dynamodb.DescribeGlobalTableSettingsInput, ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeGlobalTableSettingsRequest(*dynamodb.DescribeGlobalTableSettingsInput) (r *request.Request, o *dynamodb.DescribeGlobalTableSettingsOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeLimits(*dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeLimitsWithContext(aws.Context, *dynamodb.DescribeLimitsInput, ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeLimitsRequest(*dynamodb.DescribeLimitsInput) (r *request.Request, o *dynamodb.DescribeLimitsOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTableWithContext(aws.Context, *dynamodb.DescribeTableInput, ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTableRequest(*dynamodb.DescribeTableInput) (r *request.Request, o *dynamodb.DescribeTableOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeTableReplicaAutoScaling(*dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTableReplicaAutoScalingWithContext(aws.Context, *dynamodb.DescribeTableReplicaAutoScalingInput, ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTableReplicaAutoScalingRequest(*dynamodb.DescribeTableReplicaAutoScalingInput) (r *request.Request, o *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	d.unImpl()
	return
}

func (d *Dax) DescribeTimeToLive(*dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTimeToLiveWithContext(aws.Context, *dynamodb.DescribeTimeToLiveInput, ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) DescribeTimeToLiveRequest(*dynamodb.DescribeTimeToLiveInput) (r *request.Request, o *dynamodb.DescribeTimeToLiveOutput) {
	d.unImpl()
	return
}

func (d *Dax) ListBackups(*dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListBackupsWithContext(aws.Context, *dynamodb.ListBackupsInput, ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListBackupsRequest(*dynamodb.ListBackupsInput) (r *request.Request, o *dynamodb.ListBackupsOutput) {
	d.unImpl()
	return
}

func (d *Dax) ListContributorInsights(*dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListContributorInsightsWithContext(aws.Context, *dynamodb.ListContributorInsightsInput, ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListContributorInsightsRequest(*dynamodb.ListContributorInsightsInput) (r *request.Request, o *dynamodb.ListContributorInsightsOutput) {
	d.unImpl()
	return
}

func (d *Dax) ListContributorInsightsPages(*dynamodb.ListContributorInsightsInput, func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	return d.unImpl()
}

func (d *Dax) ListContributorInsightsPagesWithContext(aws.Context, *dynamodb.ListContributorInsightsInput, func(*dynamodb.ListContributorInsightsOutput, bool) bool, ...request.Option) error {
	return d.unImpl()
}

func (d *Dax) ListGlobalTables(*dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListGlobalTablesWithContext(aws.Context, *dynamodb.ListGlobalTablesInput, ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListGlobalTablesRequest(*dynamodb.ListGlobalTablesInput) (r *request.Request, o *dynamodb.ListGlobalTablesOutput) {
	d.unImpl()
	return
}

func (d *Dax) ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListTablesWithContext(aws.Context, *dynamodb.ListTablesInput, ...request.Option) (*dynamodb.ListTablesOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListTablesRequest(*dynamodb.ListTablesInput) (r *request.Request, o *dynamodb.ListTablesOutput) {
	d.unImpl()
	return
}

func (d *Dax) ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error {
	return d.unImpl()
}

func (d *Dax) ListTablesPagesWithContext(aws.Context, *dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool, ...request.Option) error {
	return d.unImpl()
}

func (d *Dax) ListTagsOfResource(*dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListTagsOfResourceWithContext(aws.Context, *dynamodb.ListTagsOfResourceInput, ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) ListTagsOfResourceRequest(*dynamodb.ListTagsOfResourceInput) (r *request.Request, o *dynamodb.ListTagsOfResourceOutput) {
	d.unImpl()
	return
}

func (d *Dax) RestoreTableFromBackup(*dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) RestoreTableFromBackupWithContext(aws.Context, *dynamodb.RestoreTableFromBackupInput, ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) RestoreTableFromBackupRequest(*dynamodb.RestoreTableFromBackupInput) (r *request.Request, o *dynamodb.RestoreTableFromBackupOutput) {
	d.unImpl()
	return
}

func (d *Dax) RestoreTableToPointInTime(*dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) RestoreTableToPointInTimeWithContext(aws.Context, *dynamodb.RestoreTableToPointInTimeInput, ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) RestoreTableToPointInTimeRequest(*dynamodb.RestoreTableToPointInTimeInput) (r *request.Request, o *dynamodb.RestoreTableToPointInTimeOutput) {
	d.unImpl()
	return
}

func (d *Dax) TagResource(*dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) TagResourceWithContext(aws.Context, *dynamodb.TagResourceInput, ...request.Option) (*dynamodb.TagResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) TagResourceRequest(*dynamodb.TagResourceInput) (r *request.Request, o *dynamodb.TagResourceOutput) {
	d.unImpl()
	return
}

func (d *Dax) UntagResource(*dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UntagResourceWithContext(aws.Context, *dynamodb.UntagResourceInput, ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UntagResourceRequest(*dynamodb.UntagResourceInput) (r *request.Request, o *dynamodb.UntagResourceOutput) {
	d.unImpl()
	return
}

func (d *Dax) UpdateContinuousBackups(*dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateContinuousBackupsWithContext(aws.Context, *dynamodb.UpdateContinuousBackupsInput, ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateContinuousBackupsRequest(*dynamodb.UpdateContinuousBackupsInput) (r *request.Request, o *dynamodb.UpdateContinuousBackupsOutput) {
	d.unImpl()
	return
}

func (d *Dax) UpdateContributorInsights(*dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateContributorInsightsWithContext(aws.Context, *dynamodb.UpdateContributorInsightsInput, ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateContributorInsightsRequest(*dynamodb.UpdateContributorInsightsInput) (r *request.Request, o *dynamodb.UpdateContributorInsightsOutput) {
	d.unImpl()
	return
}

func (d *Dax) UpdateGlobalTable(*dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateGlobalTableWithContext(aws.Context, *dynamodb.UpdateGlobalTableInput, ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateGlobalTableRequest(*dynamodb.UpdateGlobalTableInput) (r *request.Request, o *dynamodb.UpdateGlobalTableOutput) {
	d.unImpl()
	return
}

func (d *Dax) UpdateGlobalTableSettings(*dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateGlobalTableSettingsWithContext(aws.Context, *dynamodb.UpdateGlobalTableSettingsInput, ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateGlobalTableSettingsRequest(*dynamodb.UpdateGlobalTableSettingsInput) (r *request.Request, o *dynamodb.UpdateGlobalTableSettingsOutput) {
	d.unImpl()
	return
}

func (d *Dax) UpdateTable(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTableWithContext(aws.Context, *dynamodb.UpdateTableInput, ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTableRequest(*dynamodb.UpdateTableInput) (r *request.Request, o *dynamodb.UpdateTableOutput) {
	d.unImpl()
	return
}

func (d *Dax) UpdateTableReplicaAutoScaling(*dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTableReplicaAutoScalingWithContext(aws.Context, *dynamodb.UpdateTableReplicaAutoScalingInput, ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTableReplicaAutoScalingRequest(*dynamodb.UpdateTableReplicaAutoScalingInput) (r *request.Request, o *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	d.unImpl()
	return
}

func (d *Dax) UpdateTimeToLive(*dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTimeToLiveWithContext(aws.Context, *dynamodb.UpdateTimeToLiveInput, ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return nil, d.unImpl()
}

func (d *Dax) UpdateTimeToLiveRequest(*dynamodb.UpdateTimeToLiveInput) (r *request.Request, o *dynamodb.UpdateTimeToLiveOutput) {
	d.unImpl()
	return
}

func (d *Dax) WaitUntilTableExists(*dynamodb.DescribeTableInput) error {
	return d.unImpl()
}

func (d *Dax) WaitUntilTableExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error {
	return d.unImpl()
}

func (d *Dax) WaitUntilTableNotExists(*dynamodb.DescribeTableInput) error {
	return d.unImpl()
}

func (d *Dax) WaitUntilTableNotExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error {
	return d.unImpl()
}

func (d *Dax) unImpl() error {
	panic("unimpl")
}

func (d *Dax) Close() error {
	if c, ok := d.client.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
