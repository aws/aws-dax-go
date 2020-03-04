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
	"github.com/dmartin1/aws-dax-go/dax/daxerr"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
)

//DaxRetryer implements EqualJitterBackoffStratergy for throttled requests
type DaxRetryer struct {
	BaseThrottleDelay time.Duration
	MaxBackoffDelay   time.Duration
}

const (
	//DefaultBaseRetryDelay is base delay for throttled requests
	DefaultBaseRetryDelay = 70 * time.Millisecond
	//DefaultMaxBackoffDelay is max backoff delay for throttled requests
	DefaultMaxBackoffDelay = 20 * time.Second
)

func (r *DaxRetryer) setRetryerDefaults() {
	if r.BaseThrottleDelay == 0 {
		r.BaseThrottleDelay = DefaultBaseRetryDelay
	}
	if r.MaxBackoffDelay == 0 {
		r.MaxBackoffDelay = DefaultMaxBackoffDelay
	}
}

//RetryRules returns the delay duration before retrying this request again
func (r DaxRetryer) RetryRules(req *request.Request) time.Duration {
	if req.IsErrorThrottle() {
		r.setRetryerDefaults()
		attempt := req.RetryCount
		minDelay := time.Duration(1<<uint64(attempt)) * r.BaseThrottleDelay
		if minDelay > r.MaxBackoffDelay {
			minDelay = r.MaxBackoffDelay
		}
		jitter := time.Duration(rand.Intn(int(minDelay)/2 + 1))

		return minDelay/2 + jitter
	}
	return 0
}

//ShouldRetry returns true if the request should be retried.
func (r DaxRetryer) ShouldRetry(req *request.Request) bool {
	daxErr := req.Error.(*daxerr.DaxRequestFailure)
	return len(daxErr.Codes) > 0 && (daxErr.Codes[0] == 1 || daxErr.Codes[0] == 2) || req.IsErrorThrottle()
}

// MaxRetries returns the number of maximum retries the service will use to make
// an individual API request.
func (r DaxRetryer) MaxRetries() int {
	return 0
}
