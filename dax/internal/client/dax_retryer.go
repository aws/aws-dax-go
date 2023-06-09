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
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go"
)

// IsErrorRetryable returns if the error is daxError
// if code sequences correct any condition return a value other than unknown.
func IsErrorRetryable(err error) aws.Ternary {
	de, ok := err.(daxError)
	if !ok {
		return aws.UnknownTernary
	}
	codes := de.CodeSequence()
	if len(codes) > 0 && (codes[0] == 1 || codes[0] == 2) {
		return aws.TrueTernary
	}
	// Error code [4.23.31.33] is for AuthenticationRequiredException
	if len(codes) == 4 && codes[0] == 4 && codes[1] == 23 && codes[2] == 31 && codes[3] == 33 {
		return aws.TrueTernary
	}
	return aws.FalseTernary
}

// Sleep will wait for the timer duration to expire, or the context
// is canceled. Which ever happens first. If the context is canceled the Context's
// error will be returned.
//
// Expects Context to always return a non-nil error if the Done channel is closed.
func Sleep(ctx context.Context, op string, dur time.Duration) error {
	t := time.NewTimer(dur)
	defer t.Stop()

	select {
	case <-t.C:
		break
	case <-ctx.Done():
		err := ctx.Err()
		if errors.Is(err, context.Canceled) {
			return &smithy.CanceledError{Err: err}
		}
		return &smithy.OperationError{Err: err, OperationName: op}
	}

	return nil
}

func isRetryable(o RequestOptions, attempt int, err error) bool {
	if o.Retryer == nil {
		return false
	}
	if attempt > o.Retryer.MaxAttempts() {
		return false
	}
	return o.Retryer.IsErrorRetryable(err)
}
