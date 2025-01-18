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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/logging"
)

const (
	schemeDax  = "dax"
	schemeDaxs = "daxs"
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

type Config struct {
	MaxPendingConnectionsPerHost int
	ClusterUpdateThreshold       time.Duration
	ClusterUpdateInterval        time.Duration
	IdleConnectionReapDelay      time.Duration
	ClientHealthCheckInterval    time.Duration

	HostPorts        []string
	Region           string
	EndpointResolver aws.EndpointResolverWithOptions
	Credentials      aws.CredentialsProvider
	DialContext      func(ctx context.Context, network string, address string) (net.Conn, error)
	connConfig       connConfig

	SkipHostnameVerification bool
	logger                   logging.Logger

	RouteManagerEnabled bool // this flag temporarily removes routes facing network errors.
}

type connConfig struct {
	isEncrypted              bool
	hostname                 string
	skipHostnameVerification bool
}

func (cfg *Config) validate(op string) error {
	if len(cfg.HostPorts) == 0 && cfg.EndpointResolver == nil {
		return &smithy.OperationError{
			ServiceID:     service,
			OperationName: op,
			Err:           errors.New("config.HostPorts or config.EndpointResolver is required"),
		}
	}
	if len(cfg.Region) == 0 {
		return &smithy.OperationError{
			ServiceID:     service,
			OperationName: op,
			Err:           errors.New("config.Region is required"),
		}
	}
	if cfg.Credentials == nil {
		return &smithy.OperationError{
			ServiceID:     service,
			OperationName: op,
			Err:           errors.New("config.Credentials is required"),
		}
	}
	if cfg.MaxPendingConnectionsPerHost < 0 {
		return &smithy.OperationError{
			ServiceID:     service,
			OperationName: op,
			Err:           errors.New("config.MaxPendingConnectionsPerHost cannot be negative"),
		}
	}

	return nil
}

func (cfg *Config) validateConnConfig() {
	if cfg.connConfig.isEncrypted && cfg.SkipHostnameVerification {
		cfg.logger.Logf(ClassificationWarn, "Skip hostname verification of TLS connections. The default is to perform hostname verification, setting this to True will skip verification. Be sure you understand the implication of doing so, which is the inability to authenticate the cluster that you are connecting to.")
	}
}

func (cfg *Config) SetLogger(logger logging.Logger) {
	cfg.logger = logger
}

var defaultPorts = map[string]int{
	schemeDax:  8111,
	schemeDaxs: 9111,
}

func DefaultConfig() Config {
	cfg := Config{
		MaxPendingConnectionsPerHost: 10,
		ClusterUpdateInterval:        time.Second * 4,
		ClusterUpdateThreshold:       time.Millisecond * 125,
		ClientHealthCheckInterval:    time.Second * 5,

		connConfig:               connConfig{},
		SkipHostnameVerification: false,
		logger:                   &logging.Nop{},
		IdleConnectionReapDelay:  30 * time.Second,
	}
	if cfg.Credentials == nil {
		conf, err := config.LoadDefaultConfig(context.Background())
		if err != nil {
			panic(fmt.Sprintf("unexpected error: %+v", err))
		}
		cfg.Credentials = conf.Credentials
	}
	return cfg
}

type ClusterDaxClient struct {
	config  Config
	cluster *cluster
}

func New(ctx context.Context, config Config) (*ClusterDaxClient, error) {
	cluster, err := newCluster(config)
	if err != nil {
		return nil, err
	}
	err = cluster.start(ctx)
	if err != nil {
		return nil, err
	}
	client := &ClusterDaxClient{config: config, cluster: cluster}
	return client, nil
}

func (cc *ClusterDaxClient) Close() error {
	return cc.cluster.Close()
}

func (cc *ClusterDaxClient) endpoints(ctx context.Context, opt RequestOptions) ([]serviceEndpoint, error) {
	var out []serviceEndpoint
	var err error
	action := func(client DaxAPI, o RequestOptions) error {
		out, err = client.endpoints(ctx, o)
		return err
	}
	if err = cc.retry(ctx, opEndpoints, action, opt); err != nil {
		return nil, operationError(opEndpoints, err)
	}
	return out, nil
}

func (cc *ClusterDaxClient) PutItemWithOptions(ctx context.Context, input *dynamodb.PutItemInput, opt RequestOptions) (*dynamodb.PutItemOutput, error) {
	var output *dynamodb.PutItemOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.PutItemWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpPutItem, action, opt); err != nil {
		return nil, operationError(OpPutItem, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) DeleteItemWithOptions(ctx context.Context, input *dynamodb.DeleteItemInput, opt RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	var output *dynamodb.DeleteItemOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.DeleteItemWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpDeleteItem, action, opt); err != nil {
		return nil, operationError(OpDeleteItem, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) UpdateItemWithOptions(ctx context.Context, input *dynamodb.UpdateItemInput, opt RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	var output *dynamodb.UpdateItemOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.UpdateItemWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpUpdateItem, action, opt); err != nil {
		return nil, operationError(OpUpdateItem, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) BatchWriteItemWithOptions(ctx context.Context, input *dynamodb.BatchWriteItemInput, opt RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	var output *dynamodb.BatchWriteItemOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.BatchWriteItemWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpBatchWriteItem, action, opt); err != nil {
		return nil, operationError(OpBatchWriteItem, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) TransactWriteItemsWithOptions(ctx context.Context, input *dynamodb.TransactWriteItemsInput, opt RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	var output *dynamodb.TransactWriteItemsOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.TransactWriteItemsWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpTransactWriteItems, action, opt); err != nil {
		return nil, operationError(OpTransactWriteItems, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) TransactGetItemsWithOptions(ctx context.Context, input *dynamodb.TransactGetItemsInput, opt RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	var output *dynamodb.TransactGetItemsOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.TransactGetItemsWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpTransactGetItems, action, opt); err != nil {
		return nil, operationError(OpTransactGetItems, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) GetItemWithOptions(ctx context.Context, input *dynamodb.GetItemInput, opt RequestOptions) (*dynamodb.GetItemOutput, error) {
	var output *dynamodb.GetItemOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.GetItemWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpGetItem, action, opt); err != nil {
		return nil, operationError(OpGetItem, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) QueryWithOptions(ctx context.Context, input *dynamodb.QueryInput, opt RequestOptions) (*dynamodb.QueryOutput, error) {
	var output *dynamodb.QueryOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.QueryWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpQuery, action, opt); err != nil {
		return nil, operationError(OpQuery, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) ScanWithOptions(ctx context.Context, input *dynamodb.ScanInput, opt RequestOptions) (*dynamodb.ScanOutput, error) {
	var output *dynamodb.ScanOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.ScanWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpScan, action, opt); err != nil {
		return nil, operationError(OpScan, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) BatchGetItemWithOptions(ctx context.Context, input *dynamodb.BatchGetItemInput, opt RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	var output *dynamodb.BatchGetItemOutput
	action := func(client DaxAPI, o RequestOptions) error {
		var err error
		output, err = client.BatchGetItemWithOptions(ctx, input, o)
		return err
	}
	if err := cc.retry(ctx, OpBatchGetItem, action, opt); err != nil {
		return nil, operationError(OpBatchGetItem, err)
	}
	return output, nil
}

func (cc *ClusterDaxClient) retry(ctx context.Context, op string, action func(client DaxAPI, o RequestOptions) error, opt RequestOptions) (err error) {
	defer func() {
		if daxErr, ok := err.(daxError); ok {
			err = convertDaxError(daxErr)
		}
	}()

	attempts := opt.RetryMaxAttempts
	opt.RetryMaxAttempts = 0 // disable retries on single node client

	var client DaxAPI
	// Start from 0 to accomodate for the initial request
	for i := 0; i <= attempts; i++ {
		if i > 0 && opt.Logger != nil {
			opt.Logger.Logf(ClassificationDebug, "Retrying Request %s/%s, attempt %d", service, op, i)
		}
		client, err = cc.cluster.client(client, op)
		if err != nil {
			if !isRetryable(opt, i+1, err) {
				return err
			}
			continue
		}

		err = action(client, opt)
		if err == nil {
			// success
			return nil
		}
		if !isRetryable(opt, i+1, err) {
			return err
		}

		d, err := opt.Retryer.RetryDelay(i+1, err)
		if err != nil {
			return &smithy.OperationError{Err: err, OperationName: op}
		}
		if err = Sleep(ctx, op, d); err != nil {
			return err
		}

		if opt.Logger != nil {
			opt.Logger.Logf(ClassificationDebug, "Error in executing request %s/%s. : %s", service, op, err)
		}
	}
	return err
}

type cluster struct {
	lock           sync.RWMutex
	active         map[hostPort]clientAndConfig // protected by lock
	routeManager   RouteManager                 // protected by lock
	closed         bool                         // protected by lock
	lastRefreshErr error                        // protected by lock

	lastUpdateNs int64
	executor     *taskExecutor

	seeds         []hostPort
	config        Config
	clientBuilder clientBuilder
}

type clientAndConfig struct {
	client DaxAPI
	cfg    serviceEndpoint
}

func newCluster(cfg Config) (*cluster, error) {
	const op = "NewClient"
	if err := cfg.validate(op); err != nil {
		return nil, err
	}
	hostPorts := cfg.HostPorts
	if cfg.EndpointResolver != nil {
		endpoint, err := cfg.EndpointResolver.ResolveEndpoint(serviceName, cfg.Region)
		if err != nil {
			return nil, err
		}
		hostPorts = append(hostPorts, endpoint.URL)
	}
	seeds, hostname, isEncrypted, err := getHostPorts(hostPorts, op)
	if err != nil {
		return nil, err
	}
	cfg.connConfig.isEncrypted = isEncrypted
	cfg.connConfig.skipHostnameVerification = cfg.SkipHostnameVerification
	cfg.connConfig.hostname = hostname
	cfg.validateConnConfig()
	routeManager := newRouteManager(cfg.RouteManagerEnabled, cfg.ClientHealthCheckInterval, cfg.logger)
	return &cluster{seeds: seeds, config: cfg, executor: newExecutor(), clientBuilder: &singleClientBuilder{}, routeManager: routeManager}, nil
}

func getHostPorts(hosts []string, op string) (hostPorts []hostPort, hostname string, isEncrypted bool, err error) {
	out := make([]hostPort, len(hosts))

	handle := func(e error) (hostPorts []hostPort, hostname string, isEncrypted bool, err error) {
		return nil, "", false, e
	}

	for i, hp := range hosts {
		host, port, scheme, err := parseHostPort(hp, op)
		if err != nil {
			return handle(err)
		}

		if isEncrypted != (scheme == schemeDaxs) {
			if i == 0 {
				isEncrypted = true
			} else {
				return handle(&smithy.OperationError{
					ServiceID:     service,
					OperationName: op,
					Err:           errors.New("inconsistency between the schemes of provided endpoints"),
				})
			}
		}
		if scheme == schemeDaxs && i > 0 {
			return handle(&smithy.OperationError{
				ServiceID:     service,
				OperationName: op,
				Err:           errors.New("only one cluster discovery endpoint may be provided for encrypted cluster"),
			})
		}
		out[i] = hostPort{host, port}
		hostname = host
	}
	return out, hostname, isEncrypted, nil
}

func parseHostPort(hostPort string, op string) (host string, port int, scheme string, err error) {
	uriString := hostPort
	colon := strings.Index(hostPort, "://")

	handle := func(e error) (host string, port int, scheme string, err error) {
		return "", 0, "", e
	}

	if colon == -1 {
		if strings.Index(hostPort, ":") == -1 {
			return handle(&smithy.OperationError{
				ServiceID:     service,
				OperationName: op,
				Err:           errors.New(hostPort + "is invalid host port."),
			})
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
		return handle(&smithy.OperationError{
			ServiceID:     service,
			OperationName: op,
			Err:           errors.New("invalid host port"),
		})
	}

	port, err = strconv.Atoi(portStr)
	if err != nil {
		port = defaultPorts[scheme]
	}

	if _, ok := defaultPorts[scheme]; !ok {
		schemes := strings.Join(strings.Fields(fmt.Sprint(reflect.ValueOf(defaultPorts).MapKeys())), ",")
		return handle(&smithy.OperationError{
			ServiceID:     service,
			OperationName: op,
			Err:           errors.New("URL scheme must be one of " + schemes),
		})
	}

	return host, port, scheme, nil
}

func (c *cluster) start(ctx context.Context) error {
	c.executor.start(c.config.ClusterUpdateInterval, func() error {
		c.safeRefresh(ctx, false)
		return nil
	})
	c.executor.start(c.config.IdleConnectionReapDelay, c.reapIdleConnections)
	c.safeRefresh(ctx, false)
	return nil
}

func (c *cluster) Close() error {
	c.executor.stopAll()

	c.lock.Lock()
	defer c.lock.Unlock()
	c.closed = true
	for _, config := range c.active {
		c.closeClient(config.client)
	}
	c.active = nil
	c.routeManager.close()
	c.routeManager = nil
	return nil
}

func (c *cluster) reapIdleConnections() error {
	clients := c.getAllRoutes()
	for _, c := range clients {
		if d, ok := c.(connectionReaper); ok {
			d.reapIdleConnections()
		}
	}
	return nil
}

func (c *cluster) client(prev DaxAPI, op string) (DaxAPI, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	route := c.routeManager.getRoute(prev)
	if route == nil {
		return nil, &smithy.OperationError{
			ServiceID:     service,
			OperationName: op,
			Err:           fmt.Errorf("no routes found. lastRefreshError: %v", c.lastRefreshError()),
		}
	}
	return route, nil
}

func (c *cluster) safeRefresh(ctx context.Context, force bool) {
	err := c.refresh(ctx, force)
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lastRefreshErr = err
}

func (c *cluster) lastRefreshError() error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lastRefreshErr
}

func (c *cluster) refresh(ctx context.Context, force bool) error {
	last := atomic.LoadInt64(&c.lastUpdateNs)
	now := time.Now().UnixNano()
	if now-last > c.config.ClusterUpdateThreshold.Nanoseconds() || force {
		if atomic.CompareAndSwapInt64(&c.lastUpdateNs, last, now) {
			return c.refreshNow(ctx)
		}
	}
	return nil
}

func (c *cluster) refreshNow(ctx context.Context) error {
	cfg, err := c.pullEndpoints(ctx)
	if err != nil {
		c.debugLog("ERROR: Failed to refresh endpoint : %s", err)
		return err
	}
	if !c.hasChanged(cfg) {
		return nil
	}
	c.update(ctx, cfg)
	return nil
}

// This method is responsible for updating the set of active routes tracked by
// the clsuter-dax-client in response to updates in the roster.
func (c *cluster) update(ctx context.Context, config []serviceEndpoint) {
	newEndpoints := make(map[hostPort]struct{}, len(config))
	for _, cfg := range config {
		newEndpoints[cfg.hostPort()] = struct{}{}
	}

	newActive := make(map[hostPort]clientAndConfig, len(config))
	newRoutes := make([]DaxAPI, len(config))
	shouldUpdateRoutes := true
	var toClose []clientAndConfig
	// Track the newly created client instances, so that we can clean them up in case of partial failures.
	var newCliCfg []clientAndConfig

	c.lock.Lock()

	cls := c.closed
	oldActive := c.active

	if cls {
		shouldUpdateRoutes = false
	} else {
		// Close the client instances that are no longer part of roster.
		for ep, clicfg := range oldActive {
			_, isPartOfUpdatedEndpointsConfig := newEndpoints[ep]
			if !isPartOfUpdatedEndpointsConfig {
				c.debugLog("Found updated endpoint configs, will close inactive endpoint client : %s", ep.host)
				toClose = append(toClose, clicfg)
			}
		}

		// Create client instances for the new endpoints in roster.
		for i, ep := range config {
			cliAndCfg, alreadyExists := oldActive[ep.hostPort()]
			if !alreadyExists {
				cli, err := c.newSingleClient(ep)
				if err != nil {
					shouldUpdateRoutes = false
					break
				} else {
					cliAndCfg = clientAndConfig{client: cli, cfg: ep}
					newCliCfg = append(newCliCfg, cliAndCfg)
				}

				if sc, ok := cli.(HealthCheckDaxAPI); ok {
					sc.startHealthChecks(ctx, c, ep.hostPort())
				}
			}
			newActive[ep.hostPort()] = cliAndCfg
			newRoutes[i] = cliAndCfg.client
		}
	}

	if shouldUpdateRoutes {
		c.active = newActive
		c.routeManager.setRoutes(newRoutes)
	} else {
		// cleanup newly created clients if they are not going to be tracked further.
		toClose = append(toClose, newCliCfg...)
	}
	c.lock.Unlock()

	go func() {
		for _, client := range toClose {
			c.debugLog("Closing client for : %s", client.cfg.hostname)
			c.closeClient(client.client)
		}
	}()
}

func (c *cluster) onHealthCheckFailed(ctx context.Context, host hostPort) {
	c.lock.Lock()
	c.debugLog("DEBUG: Refreshing cache for host: " + host.host)
	shouldCloseOldClient := true
	var oldClientConfig, ok = c.active[host]
	if ok {
		cli, err := c.newSingleClient(oldClientConfig.cfg)
		if sc, ok := cli.(HealthCheckDaxAPI); ok {
			sc.startHealthChecks(ctx, c, host)
		}

		if err == nil {
			c.active[host] = clientAndConfig{client: cli, cfg: oldClientConfig.cfg}

			newRoutes := make([]DaxAPI, len(c.active))
			i := 0
			for _, cliAndCfg := range c.active {
				newRoutes[i] = cliAndCfg.client
				i++
			}
			c.routeManager.setRoutes(newRoutes)
		} else {
			shouldCloseOldClient = false
			c.debugLog("DEBUG: Failed to refresh cache for host: " + host.host)
		}
	} else {
		c.debugLog("DEBUG: The node is not part of active routes. Ignoring the health check failure for host: " + host.host)
	}
	c.lock.Unlock()

	if shouldCloseOldClient {
		c.debugLog("DEBUG: Closing old instance of a replaced client for endpoint: %s", oldClientConfig.cfg.hostPort().host)
		c.closeClient(oldClientConfig.client)
	}
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

func (c *cluster) pullEndpoints(ctx context.Context) ([]serviceEndpoint, error) {
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
			endpoints, err := c.pullEndpointsFrom(ctx, ip, s.port)
			if err != nil {
				lastErr = err
				continue
			}
			c.debugLog("DEBUG: Pulled endpoints from %s : %v", ip, endpoints)
			if len(endpoints) > 0 {
				return endpoints, nil
			}
		}
	}
	return nil, lastErr
}

func (c *cluster) pullEndpointsFrom(ctx context.Context, ip net.IP, port int) ([]serviceEndpoint, error) {
	client, err := c.clientBuilder.newClient(ip, port, c.config.connConfig, c.config.Region, c.config.Credentials,
		c.config.MaxPendingConnectionsPerHost, c.config.DialContext, c)
	if err != nil {
		return nil, err
	}
	defer c.closeClient(client)
	ctx, cfn := context.WithTimeout(ctx, 5*time.Second)
	defer cfn()
	opts := RequestOptions{}
	opts.RetryMaxAttempts = 2
	return client.endpoints(ctx, opts)
}

func (c *cluster) closeClient(client DaxAPI) {
	if d, ok := client.(io.Closer); ok {
		d.Close()
	}
}

func (c *cluster) debugLog(format string, args ...interface{}) {
	if c.config.logger != nil {
		{
			c.config.logger.Logf(ClassificationDebug, format, args)
		}
	}
}

func (c *cluster) isRouteManagerEnabled() bool {
	return c.config.RouteManagerEnabled
}

func (c *cluster) addRoute(endpoint string, route DaxAPI) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.routeManager.addRoute(endpoint, route)
}

func (c *cluster) removeRoute(endpoint string, route DaxAPI) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.routeManager.removeRoute(endpoint, route, c.active)
}

func (c *cluster) getAllRoutes() []DaxAPI {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.routeManager.getAllRoutes()
}

func (c *cluster) newSingleClient(cfg serviceEndpoint) (DaxAPI, error) {
	return c.clientBuilder.newClient(net.IP(cfg.address), cfg.port, c.config.connConfig, c.config.Region, c.config.Credentials, c.config.MaxPendingConnectionsPerHost, c.config.DialContext, c)
}

type clientBuilder interface {
	newClient(net.IP, int, connConfig, string, aws.CredentialsProvider, int, dialContext, RouteListener) (DaxAPI, error)
}

type singleClientBuilder struct{}

var _ clientBuilder = (*singleClientBuilder)(nil)

func (*singleClientBuilder) newClient(ip net.IP, port int, connConfigData connConfig, region string, credentials aws.CredentialsProvider, maxPendingConnects int, dialContextFn dialContext, routeListener RouteListener) (DaxAPI, error) {
	endpoint := fmt.Sprintf("%s:%d", ip, port)
	return newSingleClientWithOptions(endpoint, connConfigData, region, credentials, maxPendingConnects, dialContextFn, routeListener)
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

type RouteListener interface {
	isRouteManagerEnabled() bool
	addRoute(string, DaxAPI)
	removeRoute(string, DaxAPI)
}
