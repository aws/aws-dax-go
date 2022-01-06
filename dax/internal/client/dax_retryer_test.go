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
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
)

func TestRetryThrottleCodes(t *testing.T) {

	req := request.Request{}
	retryer := DaxRetryer{}
	attempt := 2
	req.RetryCount = attempt
	baseThrottleDelay := 70 * time.Millisecond
	//for throttling exception
	req.Error = newDaxRequestFailure([]int{0}, "ThrottlingException", "", "", 400)

	if !retryer.ShouldRetry(&req) {
		t.Errorf("expected retry on throttling")
	}
	delay := retryer.RetryRules(&req)
	maxDelay := time.Duration(1<<uint64(attempt)) * baseThrottleDelay
	if delay > maxDelay {
		t.Errorf("delay more than expected, expected upto %d, got %d ", maxDelay, delay)
	}
	if delay <= 0 {
		t.Errorf("delay for throttled error should be greater than 0, got %d", delay)
	}

	//for non throttling exception
	req.Error = newDaxRequestFailure([]int{0}, "AccessDeniedException", "", "", 400)
	if retryer.ShouldRetry(&req) || retryer.RetryRules(&req) != 0 {
		t.Errorf("no retry expected")
	}
}

func TestRetryOnThrottlingException(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	flag := 0
	action := func(client DaxAPI, o RequestOptions) error {
		if flag == 0 {
			flag = 1
			return newDaxRequestFailure([]int{0}, "ThrottlingException", "", "", 400)
		}
		return nil
	}

	opt := RequestOptions{
		MaxRetries: 2,
	}

	err := cc.retry("op", action, opt)

	if err != nil {
		t.Errorf("error %v", err)
	}
}

func TestRetryOnAuthenticationRequiredException(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	flag := 0
	codes := []int{4, 23, 31, 33}
	action := func(client DaxAPI, o RequestOptions) error {
		if flag == 0 {
			flag = 1
			return newDaxRequestFailure(codes, "AuthenticationRequiredException", "", "", 400)
		}
		return nil
	}

	opt := RequestOptions{
		MaxRetries: 2,
	}

	err := cc.retry("op", action, opt)

	if err != nil {
		t.Errorf("error %v", err)
	}

}