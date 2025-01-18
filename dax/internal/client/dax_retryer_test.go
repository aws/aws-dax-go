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
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go"
)

func TestIsErrorRetryable(t *testing.T) {
	err := fmt.Errorf("fmt error")
	te := IsErrorRetryable(err)
	if te != aws.UnknownTernary {
		t.Fatalf("ternary is wrong: %v", te)
	}

	for _, code := range []int{1, 2} {
		err = newDaxRequestFailure([]int{code}, "", "", "", 500)
		te = IsErrorRetryable(err)
		if te != aws.TrueTernary {
			t.Fatalf("ternary is wrong: %v", te)
		}
	}

	err = newDaxRequestFailure([]int{4, 23, 31, 33}, "", "", "", 500)
	te = IsErrorRetryable(err)
	if te != aws.TrueTernary {
		t.Fatalf("ternary is wrong: %v", te)
	}

	err = newDaxRequestFailure([]int{0}, "", "", "", 500)
	te = IsErrorRetryable(err)
	if te != aws.FalseTernary {
		t.Fatalf("ternary is wrong: %v", te)
	}
}

func TestSleep(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := Sleep(ctx, "op", time.Millisecond*100)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error is not context.Canceled, but %T", err)
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Millisecond*100)
	err = Sleep(ctx, "op", time.Millisecond*200)

	var opErr *smithy.OperationError
	if !errors.As(err, &opErr) {
		t.Fatalf("error is not OperationError, but %T", err)
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Millisecond*200)
	err = Sleep(ctx, "op", time.Millisecond*100)
	if err != nil {
		t.Fatalf("err must be nil, but %T", err)
	}
}
