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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	retries := 3
	for successfulAttempt := 1; successfulAttempt < retries+5; successfulAttempt++ {
		calls := 0
		action := func(client DaxAPI, o RequestOptions) error {
			if o.MaxRetries != 0 {
				t.Errorf("expected 0 retries, found %v", o.MaxRetries)
			}
			calls++
			if calls == successfulAttempt {
				return nil
			}
			return errors.New("error")
		}

		opt := RequestOptions{MaxRetries: retries}
		err := cc.retry("op	", action, opt)
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
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	action := func(client DaxAPI, o RequestOptions) error {
		return errors.New("error")
	}

	var sleepCallCount int
	opt := RequestOptions{
		MaxRetries:   0,
		RetryDelay:   0,
		SleepDelayFn: func(d time.Duration) { sleepCallCount++ },
	}

	cc.retry("op", action, opt)

	if sleepCallCount != 0 {
		t.Fatalf("Sleep was called %d times, but expected none", sleepCallCount)
	}

	opt.MaxRetries = 3
	opt.RetryDelay = 1

	cc.retry("op", action, opt)

	if sleepCallCount != opt.MaxRetries {
		t.Fatalf("Sleep was called %d times, but expected %d", sleepCallCount, opt.MaxRetries)
	}
}

func TestClusterDaxClient_retryReturnsLastError(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	callCount := 0
	action := func(client DaxAPI, o RequestOptions) error {
		callCount++
		return fmt.Errorf("Error_%d", callCount)
	}

	opt := RequestOptions{
		MaxRetries: 2,
		RetryDelay: 1,
	}

	err := cc.retry("op", action, opt)
	expectedError := fmt.Errorf("Error_%d", callCount)
	if err.Error() != expectedError.Error() {
		t.Fatalf("Wrong error. Expected %v, but got %v", expectedError, err)
	}
}

func TestClusterDaxClient_retryReturnsCorrectErrorType(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
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
			errCode: dynamodb.ErrCodeResourceNotFoundException,
			class:   reflect.TypeOf(&dynamodb.ResourceNotFoundException{}),
		},
		{
			codes:   []int{4, 23, 35},
			errCode: dynamodb.ErrCodeResourceInUseException,
			class:   reflect.TypeOf(&dynamodb.ResourceInUseException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 40},
			errCode: dynamodb.ErrCodeProvisionedThroughputExceededException,
			class:   reflect.TypeOf(&dynamodb.ProvisionedThroughputExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 40},
			errCode: dynamodb.ErrCodeProvisionedThroughputExceededException,
			class:   reflect.TypeOf(&dynamodb.ProvisionedThroughputExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 41},
			errCode: dynamodb.ErrCodeResourceNotFoundException,
			class:   reflect.TypeOf(&dynamodb.ResourceNotFoundException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 43},
			errCode: dynamodb.ErrCodeConditionalCheckFailedException,
			class:   reflect.TypeOf(&dynamodb.ConditionalCheckFailedException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 45},
			errCode: dynamodb.ErrCodeResourceInUseException,
			class:   reflect.TypeOf(&dynamodb.ResourceInUseException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 46},
			errCode: ErrCodeValidationException,
			class:   reflect.TypeOf(awserr.NewRequestFailure(nil, 0, "")),
		},
		{
			codes:   []int{4, 37, 38, 39, 47},
			errCode: dynamodb.ErrCodeInternalServerError,
			class:   reflect.TypeOf(&dynamodb.InternalServerError{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 48},
			errCode: dynamodb.ErrCodeItemCollectionSizeLimitExceededException,
			class:   reflect.TypeOf(&dynamodb.ItemCollectionSizeLimitExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 49},
			errCode: dynamodb.ErrCodeLimitExceededException,
			class:   reflect.TypeOf(&dynamodb.LimitExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 50},
			errCode: ErrCodeThrottlingException,
			class:   reflect.TypeOf(awserr.NewRequestFailure(nil, 0, "")),
		},
		{
			codes:   []int{4, 37, 38, 39, 57},
			errCode: dynamodb.ErrCodeTransactionConflictException,
			class:   reflect.TypeOf(&dynamodb.TransactionConflictException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 58},
			errCode: dynamodb.ErrCodeTransactionCanceledException,
			class:   reflect.TypeOf(&dynamodb.TransactionCanceledException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 59},
			errCode: dynamodb.ErrCodeTransactionInProgressException,
			class:   reflect.TypeOf(&dynamodb.TransactionInProgressException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 60},
			errCode: dynamodb.ErrCodeIdempotentParameterMismatchException,
			class:   reflect.TypeOf(&dynamodb.IdempotentParameterMismatchException{}),
		},
		{
			codes:   []int{4, 37, 38, 44},
			errCode: ErrCodeNotImplemented,
			class:   reflect.TypeOf(awserr.NewRequestFailure(nil, 0, "")),
		},
	}

	for _, c := range cases {
		action := func(client DaxAPI, o RequestOptions) error {
			if c.errCode == dynamodb.ErrCodeTransactionCanceledException {
				return newDaxTransactionCanceledFailure(c.codes, defaultErrCode, message, requestID, statusCode, nil, nil, nil)
			}
			return newDaxRequestFailure(c.codes, defaultErrCode, message, requestID, statusCode)
		}

		opt := RequestOptions{
			MaxRetries: 0,
		}

		err := cc.retry("op", action, opt)
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
	hostPorts, err := parseHostPorts(endpoints)
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

	if err := cluster.refresh(false); err != nil {
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

	if err := cluster.refresh(false); err != nil {
		t.Errorf("unexpected error %v", err)
	}

	assertNumRoutes(cluster, 0, t)
	if _, err := cluster.client(nil); err == nil {
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
		cluster.refresh(false)
	}
	if 1 != len(clientBuilder.clients) {
		t.Errorf("expected 1, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)

	<-time.After(cfg.ClusterUpdateThreshold)
	for i := 0; i < 10; i++ {
		cluster.refresh(false)
	}
	if 2 != len(clientBuilder.clients) {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[1], t)
}

func TestCluster_refreshDup(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", port: 8121}})

	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %v", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)
	assertActiveClient(clientBuilder.clients[1], t)

	oldActive := cluster.active
	oldRoutes := cluster.routes
	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil); err != nil {
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

	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)
	assertActiveClient(clientBuilder.clients[1], t)

	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}})
	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 2, t)
	if _, err := cluster.client(nil); err != nil {
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
	if err := cluster.update(first); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	assertConnections(cluster, first, t)

	// add new hosts
	second := []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8123}}
	if !cluster.hasChanged(second) {
		t.Errorf("expected config change")
	}
	if err := cluster.update(second); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, second, t)

	// replace host
	third := []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if !cluster.hasChanged(third) {
		t.Errorf("expected config change")
	}
	if err := cluster.update(third); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, third, t)

	// remove host
	fourth := []serviceEndpoint{{hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if !cluster.hasChanged(fourth) {
		t.Errorf("expected config change")
	}
	if err := cluster.update(fourth); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	assertNumRoutes(cluster, 2, t)
	assertConnections(cluster, fourth, t)

	// no change
	fifth := []serviceEndpoint{{hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if cluster.hasChanged(fifth) {
		t.Errorf("unexpected config change")
	}
	assertNumRoutes(cluster, 2, t)
	assertConnections(cluster, fifth, t)
}

func TestCluster_client(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8888"})
	endpoints := []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8123}}

	if err := cluster.update(endpoints); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	assertNumRoutes(cluster, 3, t)
	prev, err := cluster.client(nil)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	for i := 0; i < 100; i++ {
		next, err := cluster.client(prev)
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

	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil); err != nil {
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
		if tc, ok := c.(*testClient); ok {
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
	wg := &sync.WaitGroup{}
	var result []byte
	go func() {
		wg.Add(1)
		defer wg.Done()

		for {
			buf := make([]byte, 4096)
			n, _ := ours.Read(buf)
			result = buf[:n]
			ours.Close()
			return
		}
	}()

	var dialFn contextDialer = func(ctx context.Context, address string, network string) (net.Conn, error) {
		return theirs, nil
	}
	cfg := Config{
		MaxPendingConnectionsPerHost: 1,
		ClusterUpdateInterval:        1 * time.Second,
		Credentials:                  credentials.NewStaticCredentials("id", "secret", "tok"),
		DialFn:                       dialFn,
		Region:                       "us-west-2",
		HostPorts:                    []string{"localhost:9121"},
	}
	cc, err := New(cfg)
	require.NoError(t, err)
	cc.GetItemWithOptions(&dynamodb.GetItemInput{TableName: aws.String("MyTable")}, &dynamodb.GetItemOutput{}, RequestOptions{})

	wg.Wait()

	assert.Equal(t, magic, string(result[1:8]), "expected the ClusterClient to write to the connection provided by the custom dialer")
}

type testClientBuilder struct {
	ep      []serviceEndpoint
	clients []*testClient
}

func (b *testClientBuilder) newClient(ip net.IP, port int, region string, credentials *credentials.Credentials, maxConns int) (DaxAPI, error) {
	t := &testClient{ep: b.ep, hp: hostPort{ip.String(), port}}
	b.clients = append(b.clients, []*testClient{t}...)
	return t, nil
}

type testClient struct {
	hp                         hostPort
	ep                         []serviceEndpoint
	endpointsCalls, closeCalls int
}

func (c *testClient) endpoints(opt RequestOptions) ([]serviceEndpoint, error) {
	c.endpointsCalls++
	return c.ep, nil
}

func (c *testClient) Close() error {
	c.closeCalls++
	return nil
}

func (c *testClient) PutItemWithOptions(input *dynamodb.PutItemInput, output *dynamodb.PutItemOutput, opt RequestOptions) (*dynamodb.PutItemOutput, error) {
	panic("unimpl")
}
func (c *testClient) DeleteItemWithOptions(input *dynamodb.DeleteItemInput, output *dynamodb.DeleteItemOutput, opt RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	panic("unimpl")
}
func (c *testClient) UpdateItemWithOptions(input *dynamodb.UpdateItemInput, output *dynamodb.UpdateItemOutput, opt RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	panic("unimpl")
}

func (c *testClient) GetItemWithOptions(input *dynamodb.GetItemInput, output *dynamodb.GetItemOutput, opt RequestOptions) (*dynamodb.GetItemOutput, error) {
	panic("unimpl")
}
func (c *testClient) ScanWithOptions(input *dynamodb.ScanInput, output *dynamodb.ScanOutput, opt RequestOptions) (*dynamodb.ScanOutput, error) {
	panic("unimpl")
}
func (c *testClient) QueryWithOptions(input *dynamodb.QueryInput, output *dynamodb.QueryOutput, opt RequestOptions) (*dynamodb.QueryOutput, error) {
	panic("unimpl")
}

func (c *testClient) BatchWriteItemWithOptions(input *dynamodb.BatchWriteItemInput, output *dynamodb.BatchWriteItemOutput, opt RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	panic("unimpl")
}
func (c *testClient) BatchGetItemWithOptions(input *dynamodb.BatchGetItemInput, output *dynamodb.BatchGetItemOutput, opt RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	panic("unimpl")
}
func (c *testClient) NewDaxRequest(op *request.Operation, input, output interface{}, opt RequestOptions) *request.Request {
	panic("unimpl")
}

func (c *testClient) TransactWriteItemsWithOptions(input *dynamodb.TransactWriteItemsInput, output *dynamodb.TransactWriteItemsOutput, opt RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("unimpl")
}
func (c *testClient) TransactGetItemsWithOptions(input *dynamodb.TransactGetItemsInput, output *dynamodb.TransactGetItemsOutput, opt RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	panic("unimpl")
}

func (c *testClient) build(req *request.Request) { panic("unimpl") }
func (c *testClient) send(req *request.Request)  { panic("unimpl") }
