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
	"net"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/aws/smithy-go/logging"

	aws2 "github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testRetryer struct {
	Attemps     int
	Duration    time.Duration
	CalledCount int
}

var _ aws2.Retryer = (*testRetryer)(nil)

func (r *testRetryer) IsErrorRetryable(_ error) bool {
	return true
}
func (r *testRetryer) MaxAttempts() int {
	return r.Attemps
}

func (r *testRetryer) RetryDelay(_ int, _ error) (time.Duration, error) {
	r.CalledCount++
	return r.Duration, nil
}
func (r *testRetryer) GetRetryToken(ctx context.Context, opErr error) (releaseToken func(error) error, err error) {
	return releaseToken, nil
}
func (r *testRetryer) GetInitialToken() (releaseToken func(error) error) {
	return releaseToken
}

func testTaskExecutor(t *testing.T) { // disabled as test is time sensitive
	executor := newExecutor()

	var cnt1, cnt2, cnt3 int32
	executor.start(10*time.Millisecond, func() error {
		atomic.AddInt32(&cnt1, 1)
		return nil
	})
	executor.start(20*time.Millisecond, func() error {
		atomic.AddInt32(&cnt2, 1)
		return nil
	})
	executor.start(50*time.Millisecond, func() error {
		atomic.AddInt32(&cnt3, 1)
		return nil
	})
	<-time.After(105 * time.Millisecond)
	if c := atomic.LoadInt32(&cnt1); c != 10 {
		t.Errorf("expected 10, got %d", c)
	}
	if c := atomic.LoadInt32(&cnt2); c != 5 {
		t.Errorf("expected 10, got %d", c)
	}
	if c := atomic.LoadInt32(&cnt3); c != 2 {
		t.Errorf("expected 10, got %d", c)
	}
	if executor.numTasks() != 3 {
		t.Errorf("expected 3, got %d", executor.numTasks())
	}
	executor.stopAll()
	<-time.After(105 * time.Millisecond)
	if c := atomic.LoadInt32(&cnt1); c != 10 {
		t.Errorf("expected 10, got %d", c)
	}
	if c := atomic.LoadInt32(&cnt2); c != 5 {
		t.Errorf("expected 10, got %d", c)
	}
	if c := atomic.LoadInt32(&cnt3); c != 2 {
		t.Errorf("expected 10, got %d", c)
	}
	if executor.numTasks() != 0 {
		t.Errorf("expected 0, got %d", executor.numTasks())
	}
}

func TestClusterDaxClient_retry(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update(context.Background(), []serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	retries := 3
	for successfulAttempt := 1; successfulAttempt < retries+5; successfulAttempt++ {
		calls := 0
		action := func(client DaxAPI, o RequestOptions) error {
			if o.RetryMaxAttempts != 0 {
				t.Errorf("expected 0 retries, found %v", o.RetryMaxAttempts)
			}
			calls++
			if calls == successfulAttempt {
				return nil
			}
			return errors.New("error")
		}

		opt := RequestOptions{}
		opt.RetryMaxAttempts = retries
		err := cc.retry(context.Background(), "op	", action, opt)
		maxAttempts := retries + 1
		if successfulAttempt <= maxAttempts {
			if calls != successfulAttempt {
				t.Errorf("expected success on %d call, but made %d calls", successfulAttempt, calls)
			}
			if err != nil {
				t.Errorf("unexpected error")
			}
		} else {
			if calls != retries+1 {
				t.Errorf("expected %d retries, but made %d", retries+1, calls)
			}
			if err == nil {
				t.Errorf("unexpected success %d %d", successfulAttempt, maxAttempts)
			}
		}
	}
}

func TestClusterDaxClient_retrySleepCycleCount(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update(context.Background(), []serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	action := func(client DaxAPI, o RequestOptions) error {
		return errors.New("error")
	}

	tr := &testRetryer{
		Duration: time.Millisecond * 100,
	}
	opt := RequestOptions{}
	opt.Retryer = tr

	cc.retry(context.Background(), "op", action, opt)

	if tr.CalledCount != 0 {
		t.Fatalf("Sleep was called %d times, but expected none", tr.CalledCount)
	}

	opt.RetryMaxAttempts = 3
	tr.Attemps = opt.RetryMaxAttempts

	cc.retry(context.Background(), "op", action, opt)

	if tr.CalledCount != opt.RetryMaxAttempts {
		t.Fatalf("Sleep was called %d times, but expected %d", tr.CalledCount, opt.RetryMaxAttempts)
	}
}

func TestClusterDaxClient_retryReturnsLastError(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update(context.Background(), []serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	callCount := 0
	action := func(client DaxAPI, o RequestOptions) error {
		callCount++
		return fmt.Errorf("Error_%d", callCount)
	}

	opt := RequestOptions{}
	opt.RetryMaxAttempts = 2

	err := cc.retry(context.Background(), "op", action, opt)
	expectedError := fmt.Errorf("Error_%d", callCount)
	if err.Error() != expectedError.Error() {
		t.Fatalf("Wrong error. Expected %v, but got %v", expectedError, err)
	}
}

func TestClusterDaxClient_retryReturnsCorrectErrorType(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update(context.Background(), []serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	message := "Message"
	statusCode := 0
	requestID := "RequestID"
	defaultErrCode := "empty"

	cases := []struct {
		// input
		codes []int

		// output
		errCode string
		class   reflect.Type
	}{
		{
			codes:   []int{4, 23, 24},
			errCode: (&types.ResourceNotFoundException{}).ErrorCode(),
			class:   reflect.TypeOf(types.ResourceNotFoundException{}),
		},
		{
			codes:   []int{4, 23, 35},
			errCode: (&types.ResourceInUseException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ResourceInUseException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 40},
			errCode: (&types.ProvisionedThroughputExceededException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ProvisionedThroughputExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 40},
			errCode: (&types.ProvisionedThroughputExceededException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ProvisionedThroughputExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 41},
			errCode: (&types.ResourceNotFoundException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ResourceNotFoundException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 43},
			errCode: (&types.ConditionalCheckFailedException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ConditionalCheckFailedException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 45},
			errCode: (&types.ResourceInUseException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ResourceInUseException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 46},
			errCode: ErrCodeValidationException,
			class:   reflect.TypeOf(awserr.NewRequestFailure(nil, 0, "")),
		},
		{
			codes:   []int{4, 37, 38, 39, 47},
			errCode: (&types.InternalServerError{}).ErrorCode(),
			class:   reflect.TypeOf(&types.InternalServerError{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 48},
			errCode: (&types.ItemCollectionSizeLimitExceededException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ItemCollectionSizeLimitExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 49},
			errCode: (&types.LimitExceededException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.LimitExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 50},
			errCode: ErrCodeThrottlingException,
			class:   reflect.TypeOf(awserr.NewRequestFailure(nil, 0, "")),
		},
		{
			codes:   []int{4, 37, 38, 39, 57},
			errCode: (&types.TransactionConflictException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.TransactionConflictException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 58},
			errCode: (&types.TransactionCanceledException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.TransactionCanceledException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 59},
			errCode: (&types.TransactionInProgressException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.TransactionInProgressException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 60},
			errCode: (&types.IdempotentParameterMismatchException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.IdempotentParameterMismatchException{}),
		},
		{
			codes:   []int{4, 37, 38, 44},
			errCode: ErrCodeNotImplemented,
			class:   reflect.TypeOf(awserr.NewRequestFailure(nil, 0, "")),
		},
	}

	for _, c := range cases {
		action := func(client DaxAPI, o RequestOptions) error {
			if c.errCode == (&types.TransactionCanceledException{}).ErrorCode() {
				return newDaxTransactionCanceledFailure(c.codes, defaultErrCode, message, requestID, statusCode, nil, nil, nil)
			}
			return newDaxRequestFailure(c.codes, defaultErrCode, message, requestID, statusCode)
		}

		opt := RequestOptions{}

		err := cc.retry(context.Background(), "op", action, opt)
		actualClass := reflect.TypeOf(err)
		if actualClass != c.class {
			t.Errorf("conversion of code sequence %v failed: expected %s, but got %s", c.codes, c.class.String(), actualClass.String())
		}
		f, _ := err.(awserr.RequestFailure)
		require.NotNilf(t, f, "conversion of code sequence %v failed: expected implement awserr.Error", c.codes)
		require.Equal(t, c.errCode, f.Code())
		require.Equal(t, statusCode, f.StatusCode())
		require.Equal(t, requestID, f.RequestID())
		require.Equal(t, message, f.Message())
	}
}

func TestCluster_parseHostPorts(t *testing.T) {
	endpoints := []string{"dax.us-east-1.amazonaws.com:8111"}
	hostPorts, _, _, err := getHostPorts(endpoints, "op")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(hostPorts) != len(endpoints) {
		t.Errorf("expected %v, got %v", len(endpoints), len(hostPorts))
	}
	if hostPorts[0].host != "dax.us-east-1.amazonaws.com" {
		t.Errorf("expected %v, got %v", "dax.us-east-1.amazonaws.com", hostPorts[0].host)
	}
	if hostPorts[0].port != 8111 {
		t.Errorf("expected %v, got %v", 8111, hostPorts[0].port)
	}
}

func TestCluster_pullFromNextSeed(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"non-existent-host:8888", "127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", port: 8121}})

	if err := cluster.refresh(context.Background(), false); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}
	client := clientBuilder.clients[0]
	assertDiscoveryClient(client, t)
	assertActiveClient(clientBuilder.clients[1], t)
	expected := hostPort{"127.0.0.1", 8111}
	if expected != client.hp {
		t.Errorf("expected %v, got %v", expected, client.hp)
	}
}

func TestCluster_refreshEmpty(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{})

	if err := cluster.refresh(context.Background(), false); err != nil {
		t.Errorf("unexpected error %v", err)
	}

	assertNumRoutes(cluster, 0, t)
	if _, err := cluster.client(nil, "op"); err == nil {
		t.Errorf("expected err, got nil")
	}
	if len(clientBuilder.clients) != 1 {
		t.Errorf("expected 1, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)
}

func TestCluster_refreshThreshold(t *testing.T) {
	cfg := DefaultConfig()
	cfg.ClusterUpdateThreshold = time.Millisecond * 100
	cfg.HostPorts = []string{"127.0.0.1:8111"}
	cfg.Region = "us-west-2"

	cluster, clientBuilder := newTestClusterWithConfig(cfg)
	for i := 0; i < 10; i++ {
		if err := cluster.refresh(context.Background(), false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if 1 != len(clientBuilder.clients) {
		t.Errorf("expected 1, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)

	<-time.After(cfg.ClusterUpdateThreshold)
	for i := 0; i < 10; i++ {
		if err := cluster.refresh(context.Background(), false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if 2 != len(clientBuilder.clients) {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[1], t)
}

func TestCluster_refreshDup(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", port: 8121}})

	if err := cluster.refreshNow(context.Background()); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %v", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)
	assertActiveClient(clientBuilder.clients[1], t)

	oldActive := cluster.active
	oldRoutes := cluster.routes
	if err := cluster.refreshNow(context.Background()); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if fmt.Sprintf("%p", cluster.active) != fmt.Sprintf("%p", oldActive) {
		t.Errorf("unexpected updation to active")
	}
	if fmt.Sprintf("%p", cluster.routes) != fmt.Sprintf("%p", oldRoutes) {
		t.Errorf("unexpected updation to routes")
	}
	if len(clientBuilder.clients) != 3 {
		t.Errorf("expected 3, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[2], t)
}

func TestCluster_refreshUpdate(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", port: 8121}})

	if err := cluster.refreshNow(context.Background()); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)
	assertActiveClient(clientBuilder.clients[1], t)

	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}})
	if err := cluster.refreshNow(context.Background()); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 2, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if len(clientBuilder.clients) != 4 {
		t.Errorf("expected 3, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[2], t)
	assertActiveClient(clientBuilder.clients[3], t)
}

func TestCluster_update(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8888"})

	first := []serviceEndpoint{{hostname: "localhost", port: 8121}}
	if !cluster.hasChanged(first) {
		t.Errorf("expected config change")
	}
	cluster.update(context.Background(), first)
	assertNumRoutes(cluster, 1, t)
	assertConnections(cluster, first, t)
	assertHealthCheckCalls(cluster, t)

	// add new hosts
	second := []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8123}}
	if !cluster.hasChanged(second) {
		t.Errorf("expected config change")
	}
	cluster.update(context.Background(), second)
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, second, t)
	assertHealthCheckCalls(cluster, t)

	// replace host
	third := []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if !cluster.hasChanged(third) {
		t.Errorf("expected config change")
	}
	cluster.update(context.Background(), third)
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, third, t)
	assertHealthCheckCalls(cluster, t)

	// remove host
	fourth := []serviceEndpoint{{hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if !cluster.hasChanged(fourth) {
		t.Errorf("expected config change")
	}
	cluster.update(context.Background(), fourth)
	assertNumRoutes(cluster, 2, t)
	assertConnections(cluster, fourth, t)
	assertHealthCheckCalls(cluster, t)

	// no change
	fifth := []serviceEndpoint{{hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if cluster.hasChanged(fifth) {
		t.Errorf("unexpected config change")
	}
	cluster.update(context.Background(), fifth)
	assertNumRoutes(cluster, 2, t)
	assertConnections(cluster, fifth, t)
	assertHealthCheckCalls(cluster, t)
}

func TestCluster_onHealthCheckFailed(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8888"})
	endpoint := serviceEndpoint{hostname: "localhost", port: 8123}
	first := []serviceEndpoint{endpoint, {hostname: "localhost", port: 8124}, {hostname: "localhost", port: 8125}}
	cluster.update(context.Background(), first)

	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, first, t)
	assertHealthCheckCalls(cluster, t)
	// Replace old instance of client with new one. Total client instances: 3 + 0
	assert.Equal(t, 3, len(clientBuilder.clients))
	assertCloseCalls(cluster, 0, t)

	cluster.onHealthCheckFailed(context.Background(), endpoint.hostPort())
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, first, t)
	assertHealthCheckCalls(cluster, t)
	// Replace old instance of client with new one. Total client instances: 3 + 1
	assert.Equal(t, 4, len(clientBuilder.clients))
	assertCloseCalls(cluster, 1, t)

	// Another failure
	cluster.onHealthCheckFailed(context.Background(), endpoint.hostPort())
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, first, t)
	assertHealthCheckCalls(cluster, t)
	// Replace old instance of client with new one. Total client instances: 3 + 2
	assert.Equal(t, 5, len(clientBuilder.clients))
	assertCloseCalls(cluster, 2, t)
}

func TestCluster_client(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8888"})
	endpoints := []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8123}}

	cluster.update(context.Background(), endpoints)
	assertNumRoutes(cluster, 3, t)
	prev, err := cluster.client(nil, "op")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	for i := 0; i < 100; i++ {
		next, err := cluster.client(prev, "op")
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if next == prev {
			t.Errorf("expected next != prev")
		}
		prev = next
	}
}

func TestCluster_Close(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", port: 8121}})

	if err := cluster.refreshNow(context.Background()); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}

	cluster.Close()
	for _, c := range clientBuilder.clients {
		if c.closeCalls != 1 {
			t.Errorf("expected 1, got %d", c.closeCalls)
		}
	}
}

func Test_CorrectHostPortUrlFormat(t *testing.T) {
	hostPort := "dax://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com:1234"
	host, port, scheme, _ := parseHostPort(hostPort, "op")
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 1234, port, "")
	assertEqual(t, "dax", scheme, "")
}

func Test_MissingScheme(t *testing.T) {
	hostPort := "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com:8111"
	host, port, scheme, _ := parseHostPort(hostPort, "op")
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 8111, port, "")
	assertEqual(t, "dax", scheme, "")
}

func Test_MissingPortForDax(t *testing.T) {
	hostPort := "dax://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com"
	host, port, scheme, _ := parseHostPort(hostPort, "op")
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 8111, port, "")
	assertEqual(t, "dax", scheme, "")
}

func Test_MissingPortForDaxs(t *testing.T) {
	hostPort := "daxs://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com"
	host, port, scheme, _ := parseHostPort(hostPort, "op")
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 9111, port, "")
	assertEqual(t, "daxs", scheme, "")
}

func Test_UnsupportedScheme(t *testing.T) {
	hostPort := "sample://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com"
	_, _, _, err := parseHostPort(hostPort, "op")
	assertEqual(t, reflect.TypeOf(err), reflect.TypeOf(awserr.New(request.ErrCodeRequestError, "", nil)), "")
}

func Test_DaxsCorrectUrlFormat(t *testing.T) {
	hostPort := "daxs://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com:1234"
	host, port, scheme, _ := parseHostPort(hostPort, "op")
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 1234, port, "")
	assertEqual(t, "daxs", scheme, "")
}

var nonEncEp = "dax://cluster.random.alpha-dax-clusters.us-east-1.amazonaws.com"
var nonEncNodeEp = "cluster-a.random.nodes.alpha-dax-clusters.us-east-1.amazonaws.com:8111"
var encEp = "daxs://cluster2.random.alpha-dax-clusters.us-east-1.amazonaws.com"
var encNodeEp = "daxs://cluster2-a.random.nodes.alpha-dax-clusters.us-east-1.amazonaws.com:9111"

func Test_InconsistentScheme(t *testing.T) {
	_, _, _, err := getHostPorts([]string{nonEncEp, encEp}, "op")
	assertEqual(t, reflect.TypeOf(err), reflect.TypeOf(awserr.New(request.ErrCodeRequestError, "", nil)), "")
}

func Test_MultipleUnEncryptedEndpoints(t *testing.T) {
	hps, _, _, _ := getHostPorts([]string{nonEncEp, nonEncNodeEp}, "op")
	assert.Contains(t, hps, hostPort{"cluster.random.alpha-dax-clusters.us-east-1.amazonaws.com", 8111})
	assert.Contains(t, hps, hostPort{"cluster-a.random.nodes.alpha-dax-clusters.us-east-1.amazonaws.com", 8111})
}

func Test_MultipleEncryptedEndpoints(t *testing.T) {
	_, _, _, err := getHostPorts([]string{encEp, encNodeEp}, "op")
	assertEqual(t, reflect.TypeOf(err), reflect.TypeOf(awserr.New(request.ErrCodeRequestError, "", nil)), "")
}

func assertConnections(cluster *cluster, endpoints []serviceEndpoint, t *testing.T) {
	if len(cluster.active) != len(endpoints) {
		t.Errorf("expected %d, got %d", len(cluster.active), len(endpoints))
	}
	for _, ep := range endpoints {
		hp := ep.hostPort()
		c, ok := cluster.active[hp]
		if !ok {
			t.Errorf("missing client %v", hp)
		}
		if tc, ok := c.client.(*testClient); ok {
			if tc.hp != hp {
				t.Errorf("expected %v, got %v", hp, tc.hp)
			}
		}
	}
	return
}

func assertNumRoutes(cluster *cluster, num int, t *testing.T) {
	if len(cluster.active) != num {
		t.Errorf("expected %d, got %d", num, len(cluster.active))
	}
	if len(cluster.routes) != num {
		t.Errorf("expected %d, got %d", num, len(cluster.routes))
	}
}

func assertHealthCheckCalls(cluster *cluster, t *testing.T) {
	for _, cliAndCfg := range cluster.active {
		healtCheckCalls := cliAndCfg.client.(*testClient).healthCheckCalls
		if healtCheckCalls != 1 {
			t.Errorf("expected 1 healthcheck call, got %d", healtCheckCalls)
		}
	}
}

func assertCloseCalls(cluster *cluster, num int, t *testing.T) {
	cnt := 0
	for _, client := range cluster.clientBuilder.(*testClientBuilder).clients {
		if client.closeCalls == 1 {
			cnt++
		}
	}
	assert.Equal(t, num, cnt)
}

func assertDiscoveryClient(client *testClient, t *testing.T) {
	if client.endpointsCalls != 1 {
		t.Errorf("expected 1, got %d", client.endpointsCalls)
	}
	if client.closeCalls != 1 {
		t.Errorf("expected 1, got %d", client.closeCalls)
	}
}

func assertActiveClient(client *testClient, t *testing.T) {
	if client.endpointsCalls != 0 {
		t.Errorf("expected 0, got %d", client.endpointsCalls)
	}
	if client.closeCalls != 0 {
		t.Errorf("expected 0, got %d", client.closeCalls)
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func newTestCluster(seeds []string) (*cluster, *testClientBuilder) {
	cfg := DefaultConfig()
	cfg.HostPorts = seeds
	cfg.Region = "us-west-2"
	return newTestClusterWithConfig(cfg)
}

func newTestClusterWithConfig(config Config) (*cluster, *testClientBuilder) {
	cluster, _ := newCluster(config)
	b := &testClientBuilder{}
	cluster.clientBuilder = b
	return cluster, b
}

func setExpectation(cluster *cluster, ep []serviceEndpoint) {
	cluster.clientBuilder.(*testClientBuilder).ep = ep
}

func TestCluster_customDialer(t *testing.T) {
	ours, theirs := net.Pipe()
	var wg sync.WaitGroup
	var result []byte
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			buf := make([]byte, 4096)
			n, _ := ours.Read(buf)
			result = buf[:n]
			ours.Close()
			return
		}
	}()

	var dialContextFn dialContext = func(ctx context.Context, address string, network string) (net.Conn, error) {
		return theirs, nil
	}
	cfg := Config{
		MaxPendingConnectionsPerHost: 1,
		ClusterUpdateInterval:        1 * time.Second,
		Credentials:                  nil,
		DialContext:                  dialContextFn,
		Region:                       "us-west-2",
		HostPorts:                    []string{"localhost:9121"},
		logger:                       &logging.Nop{},
		IdleConnectionReapDelay:      30 * time.Second,
	}
	cc, err := New(context.Background(), cfg)
	require.NoError(t, err)
	cc.GetItemWithOptions(context.Background(), &dynamodb.GetItemInput{TableName: aws.String("MyTable")}, RequestOptions{})

	wg.Wait()

	assert.Equal(t, magic, string(result[1:8]), "expected the ClusterClient to write to the connection provided by the custom dialer")
}

type testClientBuilder struct {
	ep      []serviceEndpoint
	clients []*testClient
}

var _ clientBuilder = (*testClientBuilder)(nil)

func (b *testClientBuilder) newClient(ip net.IP, port int, _ connConfig, _ string, _ aws2.CredentialsProvider, _ int, _ dialContext) (DaxAPI, error) {
	t := &testClient{ep: b.ep, hp: hostPort{ip.String(), port}}
	b.clients = append(b.clients, []*testClient{t}...)
	return t, nil
}

type testClient struct {
	hp                                           hostPort
	ep                                           []serviceEndpoint
	endpointsCalls, closeCalls, healthCheckCalls int
}

var _ DaxAPI = (*testClient)(nil)

func (c *testClient) startHealthChecks(_ *cluster, _ hostPort) {
	c.healthCheckCalls++
}

func (c *testClient) endpoints(_ context.Context, _ RequestOptions) ([]serviceEndpoint, error) {
	c.endpointsCalls++
	return c.ep, nil
}

func (c *testClient) Close() error {
	c.closeCalls++
	return nil
}

func (c *testClient) PutItemWithOptions(_ context.Context, _ *dynamodb.PutItemInput, _ RequestOptions) (*dynamodb.PutItemOutput, error) {
	panic("not implemented")
}
func (c *testClient) DeleteItemWithOptions(_ context.Context, _ *dynamodb.DeleteItemInput, _ RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	panic("not implemented")
}
func (c *testClient) UpdateItemWithOptions(_ context.Context, _ *dynamodb.UpdateItemInput, _ RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	panic("not implemented")
}

func (c *testClient) GetItemWithOptions(_ context.Context, _ *dynamodb.GetItemInput, _ RequestOptions) (*dynamodb.GetItemOutput, error) {
	panic("not implemented")
}
func (c *testClient) ScanWithOptions(_ context.Context, _ *dynamodb.ScanInput, _ RequestOptions) (*dynamodb.ScanOutput, error) {
	panic("not implemented")
}
func (c *testClient) QueryWithOptions(_ context.Context, _ *dynamodb.QueryInput, _ RequestOptions) (*dynamodb.QueryOutput, error) {
	panic("not implemented")
}

func (c *testClient) BatchWriteItemWithOptions(_ context.Context, _ *dynamodb.BatchWriteItemInput, _ RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	panic("not implemented")
}
func (c *testClient) BatchGetItemWithOptions(_ context.Context, _ *dynamodb.BatchGetItemInput, _ RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	panic("not implemented")
}
func (c *testClient) NewDaxRequest(op *request.Operation, input, output interface{}, _ RequestOptions) *request.Request {
	panic("not implemented")
}

func (c *testClient) TransactWriteItemsWithOptions(_ context.Context, _ *dynamodb.TransactWriteItemsInput, _ RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("not implemented")
}
func (c *testClient) TransactGetItemsWithOptions(_ context.Context, _ *dynamodb.TransactGetItemsInput, _ RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	panic("not implemented")
}

func (c *testClient) build(req *request.Request) { panic("not implemented") }
func (c *testClient) send(req *request.Request)  { panic("not implemented") }
