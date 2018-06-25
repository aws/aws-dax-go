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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"reflect"
	"testing"
	"time"
)

func TestRequestOptions(t *testing.T) {
	e := RequestOptions{
		Logger:     aws.NewDefaultLogger(),
		LogLevel:   aws.LogDebug,
		RetryDelay: 1 * time.Second,
		MaxRetries: 5,
		Context:    aws.BackgroundContext(),
	}

	r := request.New(aws.Config{}, metadata.ClientInfo{}, request.Handlers{}, nil, &request.Operation{Name: OpPutItem}, nil, nil)
	e.applyTo(r)

	a := RequestOptions{}
	if err := a.mergeFromRequest(r, true); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, got %v", e, a)
	}

}

func TestRequestOptions_MergeFromRequestOptions(t *testing.T) {
	in := request.WithLogLevel(aws.LogDebugWithHTTPBody)
	out := RequestOptions{}
	if err := out.MergeFromRequestOptions(aws.BackgroundContext(), in); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if aws.LogDebugWithHTTPBody != out.LogLevel {
		t.Errorf("expected %v, got %v", aws.LogDebugWithHTTPBody, out.LogLevel)
	}
	if aws.BackgroundContext() != out.Context {
		t.Errorf("expected %v, got %v", aws.BackgroundContext(), out.Context)
	}
}
