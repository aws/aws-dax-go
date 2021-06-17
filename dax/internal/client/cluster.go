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
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type serviceEndpoint struct {
	nodeId           int64
	hostname         string
	address          []byte
	port             int
	role             int
	availabilityZone string
	leaderSessionId  int64
}

func (e *serviceEndpoint) hostPort() hostPort {
	return hostPort{net.IP(e.address).String(), e.port}
}

type hostPort struct {
	host string
	port int
}

const idleConnectionReapDelay = 30 * time.Second

type Config struct {
	MaxPendingConnectionsPerHost int
	ClusterUpdateThreshold       time.Duration
	ClusterUpdateInterval        time.Duration

	HostPorts   []string
	Region      string
	Credentials *credentials.Credentials
	DialContext func(ctx context.Context, network string, address string) (net.Conn, error)
	connConfig  connConfig

	SkipHostnameVerification bool
	logger                   aws.Logger
}

type connConfig struct {
	isEncrypted              bool
	hostname                 string
	skipHostnameVerification bool
}

func (cfg *Config) validate() error {
	if cfg.HostPorts == nil || len(cfg.HostPorts) == 0 {
		return awserr.New(request.ParamRequiredErrCode, "HostPorts is required", nil)
	}
	if len(cfg.Region) == 0 {
		return awserr.New(request.ParamRequiredErrCode, "Region is required", nil)
	}
	if cfg.Credentials == nil {
		return awserr.New(request.ParamRequiredErrCode, "Credentials is required", nil)
	}
	if cfg.MaxPendingConnectionsPerHost < 0 {
		return awserr.New(request.InvalidParameterErrCode, "MaxPendingConnectionsPerHost cannot be negative", nil)
	}
	return nil
}

func (cfg *Config) validateConnConfig() {
	if cfg.connConfig.isEncrypted && cfg.SkipHostnameVerification {
		cfg.logger.Log(fmt.Sprintf("WARN: Skip hostname verification of TLS connections. The default is to perform hostname verification, setting this to True will skip verification. Be sure you understand the implication of doing so, which is the inability to authenticate the cluster that you are connecting to."))
	}
}

var defaultConfig = Config{
	MaxPendingConnectionsPerHost: 10,
	ClusterUpdateInterval:        time.Second * 4,
	ClusterUpdateThreshold:       time.Millisecond * 125,

	Credentials: defaults.CredChain(defaults.Config(), defaults.Handlers()),

	connConfig:               connConfig{},
	SkipHostnameVerification: false,
	logger:                   aws.NewDefaultLogger(),
}

var defaultPorts = map[string]int{
	"dax":  8111,
	"daxs": 9111,
}

func DefaultConfig() Config {
	return defaultConfig
}

type ClusterDaxClient struct {
	config  Config
	cluster *cluster

	handlers *request.Handlers
}

func New(config Config) (*ClusterDaxClient, error) {
	cluster, err := newCluster(config)
	if err != nil {
		return nil, err
	}
	err = cluster.start()
	if err != nil {
		return nil, err
	}
	client := &ClusterDaxClient{config: config, cluster: cluster}
	client.handlers = client.buildHandlers()
	return client, nil
}

func (cc *ClusterDaxClient) Close() error {
	return cc.cluster.Close()
}

func (cc *ClusterDaxClient) endpoints(opt RequestOptions) ([]serviceEndpoint, error) {
	var out []serviceEndpoint
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		out, err = client.endpoints(o)
		return err
	}
	if err = cc.retry(opEndpoints, action, opt); err != nil {
		return nil, err
	}
	return out, nil
}

func (cc *ClusterDaxClient) PutItemWithOptions(input *dynamodb.PutItemInput, output *dynamodb.PutItemOutput, opt RequestOptions) (*dynamodb.PutItemOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.PutItemWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpPutItem, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) DeleteItemWithOptions(input *dynamodb.DeleteItemInput, output *dynamodb.DeleteItemOutput, opt RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.DeleteItemWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpDeleteItem, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) UpdateItemWithOptions(input *dynamodb.UpdateItemInput, output *dynamodb.UpdateItemOutput, opt RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.UpdateItemWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpUpdateItem, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) BatchWriteItemWithOptions(input *dynamodb.BatchWriteItemInput, output *dynamodb.BatchWriteItemOutput, opt RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.BatchWriteItemWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpBatchWriteItem, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) TransactWriteItemsWithOptions(input *dynamodb.TransactWriteItemsInput, output *dynamodb.TransactWriteItemsOutput, opt RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.TransactWriteItemsWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpTransactWriteItems, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) TransactGetItemsWithOptions(input *dynamodb.TransactGetItemsInput, output *dynamodb.TransactGetItemsOutput, opt RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.TransactGetItemsWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpTransactGetItems, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) GetItemWithOptions(input *dynamodb.GetItemInput, output *dynamodb.GetItemOutput, opt RequestOptions) (*dynamodb.GetItemOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.GetItemWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpGetItem, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) QueryWithOptions(input *dynamodb.QueryInput, output *dynamodb.QueryOutput, opt RequestOptions) (*dynamodb.QueryOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.QueryWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpQuery, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) ScanWithOptions(input *dynamodb.ScanInput, output *dynamodb.ScanOutput, opt RequestOptions) (*dynamodb.ScanOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.ScanWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpScan, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) BatchGetItemWithOptions(input *dynamodb.BatchGetItemInput, output *dynamodb.BatchGetItemOutput, opt RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		output, err = client.BatchGetItemWithOptions(input, output, o)
		return err
	}
	if err = cc.retry(OpBatchGetItem, action, opt); err != nil {
		return output, err
	}
	return output, nil
}

func (cc *ClusterDaxClient) NewDaxRequest(op *request.Operation, input, output interface{}, opt RequestOptions) *request.Request {
	req := request.New(aws.Config{}, clientInfo, *cc.handlers, nil, op, input, output)
	opt.applyTo(req)
	return req
}

func (cc *ClusterDaxClient) buildHandlers() *request.Handlers {
	h := &request.Handlers{}
	h.Build.PushFrontNamed(request.NamedHandler{Name: "dax.BuildHandler", Fn: cc.build})
	h.Send.PushFrontNamed(request.NamedHandler{Name: "dax.SendHandler", Fn: cc.send})
	return h
}

func (cc *ClusterDaxClient) build(req *request.Request) {
	// Do not involve IO. Safe to retry on same client
	c, err := cc.cluster.client(nil)
	if err != nil {
		req.Error = err
	} else {
		c.build(req)
	}
}

func (cc *ClusterDaxClient) send(req *request.Request) {
	opt := RequestOptions{}
	if err := opt.mergeFromRequest(req, true); err != nil {
		req.Error = err
		return
	}
	action := func(client DaxAPI, o RequestOptions) error {
		o.applyTo(req)
		client.send(req)
		return req.Error
	}
	if err := cc.retry(req.Operation.Name, action, opt); err != nil {
		req.Error = err
	}
}

func (cc *ClusterDaxClient) retry(op string, action func(client DaxAPI, o RequestOptions) error, opt RequestOptions) (err error) {
	defer func() {
		if daxErr, ok := err.(daxError); ok {
			err = convertDaxError(daxErr)
		}
	}()

	ctx := cc.newContext(opt)

	var sleepFun func() error
	if opt.RetryDelay > 0 {
		retryDelay := opt.RetryDelay
		if opt.SleepDelayFn == nil {
			sleepFun = func() error {
				return aws.SleepWithContext(ctx, retryDelay)
			}
		} else {
			sleepFun = func() error {
				opt.SleepDelayFn(retryDelay)
				return nil
			}
		}
	}

	attempts := opt.MaxRetries
	opt.MaxRetries = 0 // disable retries on single node client

	var req request.Request
	var ok bool
	var client DaxAPI
	// Start from 0 to accomodate for the initial request
	for i := 0; i <= attempts; i++ {
		if i > 0 && opt.Logger != nil && opt.LogLevel.Matches(aws.LogDebugWithRequestRetries) {
			opt.Logger.Log(fmt.Sprintf("DEBUG: Retrying Request %s/%s, attempt %d", service, op, i))
		}
		client, err = cc.cluster.client(client)
		if err != nil {
			if req, ok = cc.shouldRetry(opt, err); !ok {
				return err
			}
		}

		if err == nil {
			if err = action(client, opt); err == nil {
				return nil
			} else if req, ok = cc.shouldRetry(opt, err); !ok {
				return err
			}
		}

		if i != attempts {
			req.RetryCount = i + 1
			delay := opt.Retryer.RetryRules(&req)
			if delay != 0 {
				if opt.SleepDelayFn == nil {
					aws.SleepWithContext(ctx, delay)
				} else {
					opt.SleepDelayFn(delay)
				}
			} else if sleepFun != nil {
				if err := sleepFun(); err != nil {
					return awserr.New(request.CanceledErrorCode, "request context canceled", err)
				}
			}
		}
	}
	return err
}

func (cc *ClusterDaxClient) newContext(o RequestOptions) aws.Context {
	if o.Context != nil {
		return o.Context
	}
	return aws.BackgroundContext()
}

func (cc *ClusterDaxClient) shouldRetry(o RequestOptions, err error) (request.Request, bool) {
	req := request.Request{}
	req.Error = err
	if _, ok := err.(daxError); ok {
		retry := o.Retryer.ShouldRetry(&req)
		return req, retry
	}
	return req, true
}

type cluster struct {
	lock           sync.RWMutex
	active         map[hostPort]DaxAPI // protected by lock
	routes         []DaxAPI            // protected by lock
	closed         bool                // protected by lock
	lastRefreshErr error               // protected by lock

	lastUpdateNs int64
	executor     *taskExecutor

	seeds         []hostPort
	config        Config
	clientBuilder clientBuilder
}

func newCluster(cfg Config) (*cluster, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	seeds, hostname, isEncrypted, err := getHostPorts(cfg.HostPorts)
	if err != nil {
		return nil, err
	}
	cfg.connConfig.isEncrypted = isEncrypted
	cfg.connConfig.skipHostnameVerification = cfg.SkipHostnameVerification
	cfg.connConfig.hostname = hostname
	cfg.validateConnConfig()
	return &cluster{seeds: seeds, config: cfg, executor: newExecutor(), clientBuilder: &singleClientBuilder{}}, nil
}

func getHostPorts(hosts []string) (hostPorts []hostPort, hostname string, isEncrypted bool, err error) {
	out := make([]hostPort, len(hosts))

	handle := func(e error) (hostPorts []hostPort, hostname string, isEncrypted bool, err error) {
		return nil, "", false, e
	}

	for i, hp := range hosts {
		host, port, scheme, err := parseHostPort(hp)
		if err != nil {
			return handle(err)
		}

		if isEncrypted != (scheme == "daxs") {
			if i == 0 {
				isEncrypted = true
			} else {
				return handle(awserr.New(request.ErrCodeRequestError, "Inconsistency between the schemes of provided endpoints.", nil))
			}
		}
		if scheme == "daxs" && i > 0 {
			return handle(awserr.New(request.InvalidParameterErrCode, "Only one cluster discovery endpoint may be provided for encrypted cluster", nil))
		}
		out[i] = hostPort{host, port}
		hostname = host
	}
	return out, hostname, isEncrypted, nil
}

func parseHostPort(hostPort string) (host string, port int, scheme string, err error) {
	uriString := hostPort
	colon := strings.Index(hostPort, "://")

	handle := func(e error) (host string, port int, scheme string, err error) {
		return "", 0, "", e
	}

	if colon == -1 {
		if strings.Index(hostPort, ":") == -1 {
			return handle(awserr.New(request.ErrCodeRequestError, "Invalid hostport", nil))
		}
		uriString = "dax://" + hostPort
	}
	u, err := url.ParseRequestURI(uriString)
	if err != nil {
		return handle(err)
	}

	host = u.Hostname()
	scheme = u.Scheme
	portStr := u.Port()
	if host == "" {
		return handle(awserr.New(request.ErrCodeRequestError, "Invalid hostport", nil))
	}

	port, err = strconv.Atoi(portStr)
	if err != nil {
		port = defaultPorts[scheme]
	}

	if _, ok := defaultPorts[scheme]; !ok {
		schemes := strings.Join(strings.Fields(fmt.Sprint(reflect.ValueOf(defaultPorts).MapKeys())), ",")
		return handle(awserr.New(request.ErrCodeRequestError, "URL scheme must be one of "+schemes, nil))
	}

	return host, port, scheme, nil
}

func (c *cluster) start() error {
	c.executor.start(c.config.ClusterUpdateInterval, func() error {
		c.safeRefresh(false)
		return nil
	})
	c.executor.start(idleConnectionReapDelay, c.reapIdleConnections)
	c.safeRefresh(false)
	return nil
}

func (c *cluster) Close() error {
	c.executor.stopAll()

	c.lock.Lock()
	defer c.lock.Unlock()
	c.closed = true
	for _, client := range c.routes {
		c.closeClient(client)
	}
	c.routes = nil
	c.active = nil
	return nil
}

func (c *cluster) reapIdleConnections() error {
	c.lock.RLock()
	clients := c.routes
	c.lock.RUnlock()

	for _, c := range clients {
		if d, ok := c.(connectionReaper); ok {
			d.reapIdleConnections()
		}
	}
	return nil
}

func (c *cluster) client(prev DaxAPI) (DaxAPI, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	n := len(c.routes)
	if n == 0 {
		return nil, awserr.New(ErrCodeServiceUnavailable, "No routes found", c.lastRefreshError())
	}
	if n == 1 {
		return c.routes[0], nil
	}
	r := rand.Intn(n)
	if c.routes[r] == prev {
		r++
		if r >= n {
			r = r - n
		}
	}
	return c.routes[r], nil
}

func (c *cluster) safeRefresh(force bool) {
	err := c.refresh(force)
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lastRefreshErr = err
}

func (c *cluster) lastRefreshError() error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lastRefreshErr
}

func (c *cluster) refresh(force bool) error {
	last := atomic.LoadInt64(&c.lastUpdateNs)
	now := time.Now().UnixNano()
	if now-last > c.config.ClusterUpdateThreshold.Nanoseconds() || force {
		if atomic.CompareAndSwapInt64(&c.lastUpdateNs, last, now) {
			return c.refreshNow()
		}
	}
	return nil
}

func (c *cluster) refreshNow() error {
	cfg, err := c.pullEndpoints()
	if err != nil {
		return err
	}
	if !c.hasChanged(cfg) {
		return nil
	}
	return c.update(cfg)
}

func (c *cluster) update(config []serviceEndpoint) error {
	newEndpoints := make(map[hostPort]struct{}, len(config))
	for _, cfg := range config {
		newEndpoints[cfg.hostPort()] = struct{}{}
	}

	newActive := make(map[hostPort]DaxAPI, len(config))
	newRoutes := make([]DaxAPI, len(config))

	c.lock.RLock()
	cls := c.closed
	oldActive := c.active
	c.lock.RUnlock()
	if cls {
		return nil
	}

	var toClose []DaxAPI
	for ep, cli := range oldActive {
		_, ok := newEndpoints[ep]
		if !ok {
			toClose = append(toClose, cli)
		}
	}
	for i, ep := range config {
		cli, ok := oldActive[ep.hostPort()]
		var err error
		if !ok {
			cli, err = c.newSingleClient(ep)
			if err != nil {
				return nil
			}
		}
		newActive[ep.hostPort()] = cli
		newRoutes[i] = cli
	}
	c.lock.Lock()
	c.active = newActive
	c.routes = newRoutes
	c.lock.Unlock()

	go func() {
		for _, client := range toClose {
			c.closeClient(client)
		}
	}()
	return nil
}

func (c *cluster) hasChanged(cfg []serviceEndpoint) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	for _, se := range cfg {
		_, ok := c.active[se.hostPort()]
		if !ok {
			return true
		}
	}
	return len(cfg) != len(c.active)
}

func (c *cluster) pullEndpoints() ([]serviceEndpoint, error) {
	var lastErr error // TODO chain errors?
	for _, s := range c.seeds {
		ips, err := net.LookupIP(s.host)
		if err != nil {
			lastErr = err
			continue
		}

		if len(ips) > 1 {
			// randomize multiple addresses; in-place fischer-yates shuffle.
			for j := len(ips) - 1; j > 0; j-- {
				k := rand.Intn(j + 1)
				ips[k], ips[j] = ips[j], ips[k]
			}
		}

		for _, ip := range ips {
			endpoints, err := c.pullEndpointsFrom(ip, s.port)
			if err != nil {
				lastErr = err
				continue
			}
			if len(endpoints) > 0 {
				return endpoints, nil
			}
		}
	}
	return nil, lastErr
}

func (c *cluster) pullEndpointsFrom(ip net.IP, port int) ([]serviceEndpoint, error) {
	client, err := c.clientBuilder.newClient(ip, port, c.config.connConfig, c.config.Region, c.config.Credentials, c.config.MaxPendingConnectionsPerHost, c.config.DialContext)
	if err != nil {
		return nil, err
	}
	defer c.closeClient(client)
	ctx, cfn := context.WithTimeout(aws.BackgroundContext(), 5*time.Second)
	defer cfn()
	return client.endpoints(RequestOptions{MaxRetries: 2, Context: ctx})
}

func (c *cluster) closeClient(client DaxAPI) {
	if d, ok := client.(io.Closer); ok {
		d.Close()
	}
}

func (c *cluster) newSingleClient(cfg serviceEndpoint) (DaxAPI, error) {
	return c.clientBuilder.newClient(net.IP(cfg.address), cfg.port, c.config.connConfig, c.config.Region, c.config.Credentials, c.config.MaxPendingConnectionsPerHost, c.config.DialContext)
}

type clientBuilder interface {
	newClient(net.IP, int, connConfig, string, *credentials.Credentials, int, dialContext) (DaxAPI, error)
}

type singleClientBuilder struct{}

func (*singleClientBuilder) newClient(ip net.IP, port int, connConfigData connConfig, region string, credentials *credentials.Credentials, maxPendingConnects int, dialContextFn dialContext) (DaxAPI, error) {
	endpoint := fmt.Sprintf("%s:%d", ip, port)
	return newSingleClientWithOptions(endpoint, connConfigData, region, credentials, maxPendingConnects, dialContextFn)
}

type taskExecutor struct {
	tasks int32
	close chan struct{}
}

func newExecutor() *taskExecutor {
	return &taskExecutor{
		close: make(chan struct{}),
	}
}

func (e *taskExecutor) start(d time.Duration, action func() error) {
	ticker := time.NewTicker(d)
	atomic.AddInt32(&e.tasks, 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				action() // TODO recover from panic()?
			case <-e.close:
				ticker.Stop()
				atomic.AddInt32(&e.tasks, -1)
				return
			}
		}
	}()
}

func (e *taskExecutor) numTasks() int32 {
	return atomic.LoadInt32(&e.tasks)
}

func (e *taskExecutor) stopAll() {
	close(e.close)
}
