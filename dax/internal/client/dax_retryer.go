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

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/aws"
)

////DaxRetryer implements EqualJitterBackoffStratergy for throttled requests
//type DaxRetryer struct {
//	BaseThrottleDelay time.Duration
//	MaxBackoffDelay   time.Duration
//}
//
//const (
//	//DefaultBaseRetryDelay is base delay for throttled requests
//	DefaultBaseRetryDelay = 70 * time.Millisecond
//	//DefaultMaxBackoffDelay is max backoff delay for throttled requests
//	DefaultMaxBackoffDelay = 20 * time.Second
//)
//
//func (r *DaxRetryer) setRetryerDefaults() {
//	if r.BaseThrottleDelay == 0 {
//		r.BaseThrottleDelay = DefaultBaseRetryDelay
//	}
//	if r.MaxBackoffDelay == 0 {
//		r.MaxBackoffDelay = DefaultMaxBackoffDelay
//	}
//}
//
////RetryRules returns the delay duration before retrying this request again
//func (r DaxRetryer) RetryRules(req *request.Request) time.Duration {
//	// ??? ProvisionedThroughputExceededException
//	if req.IsErrorThrottle() {
//		r.setRetryerDefaults()
//		attempt := req.RetryCount
//		minDelay := time.Duration(1<<uint64(attempt)) * r.BaseThrottleDelay
//		if minDelay > r.MaxBackoffDelay {
//			minDelay = r.MaxBackoffDelay
//		}
//		jitter := time.Duration(rand.Intn(int(minDelay)/2 + 1))
//
//		return minDelay/2 + jitter
//	}
//	return 0
//}
//
////ShouldRetry returns true if the request should be retried.
//func (r DaxRetryer) ShouldRetry(req *request.Request) bool {
//	daxErr := req.Error.(daxError)
//	codes := daxErr.CodeSequence()
//	return len(codes) > 0 && (codes[0] == 1 || codes[0] == 2) || req.IsErrorThrottle() || isAuthCRequiredException(codes)
//}
//
//// Error code [4.23.31.33] is for AuthenticationRequiredException
//func isAuthCRequiredException(codes []int) bool {
//	return len(codes) == 4 && codes[0] == 4 && codes[1] == 23 && codes[2] == 31 && codes[3] == 33
//}
//
//// MaxRetries returns the number of maximum retries the service will use to make
//// an individual API request.
//func (r DaxRetryer) MaxRetries() int {
//	return 0
//}

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
