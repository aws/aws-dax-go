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

package dax

import (
	"context"
	"crypto/tls"
	"net"
	"net/url"

	"github.com/aws/aws-dax-go/dax/internal/client"
	"github.com/aws/aws-dax-go/dax/internal/proxy"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go/logging"
)

// Dax makes requests to the Amazon DAX API, which conforms to the DynamoDB API.
//
// Dax methods are safe to use concurrently
type Dax struct {
	client client.DaxAPI
	config Config
}

var _ DynamoDBAPI = (*Dax)(nil)

const ServiceName = "dax"

type Config struct {
	client.Config

	// Default request options
	WriteRetries int
	ReadRetries  int

	Logger logging.Logger
}

// DefaultConfig returns the default DAX configuration.
//
// Config.Region and Config.HostPorts still need to be configured properly
// to start up a DAX client.
func DefaultConfig() Config {
	return Config{
		Config:       client.DefaultConfig(),
		WriteRetries: 2,
		ReadRetries:  2,
		Logger:       &logging.Nop{},
	}
}

// NewConfigWithSDKConfig creates a new instance of the DAX config with an aws.Config.
func NewConfigWithSDKConfig(config aws.Config) Config {
	dc := DefaultConfig()
	dc.mergeFrom(config)
	return dc
}

// New creates a new instance of the DAX client with a DAX configuration.
func New(ctx context.Context, cfg Config) (*Dax, error) {
	cfg.Config.SetLogger(cfg.Logger)
	c, err := client.New(ctx, cfg.Config)
	if err != nil {
		if cfg.Logger != nil {
			cfg.Logger.Logf(client.ClassificationError, "Exception in initialisation of DAX Client : %s", err)
		}
		return nil, err
	}
	return &Dax{client: c, config: cfg}, nil
}

// SecureDialContext creates a secure DialContext for connecting to encrypted cluster
func SecureDialContext(endpoint string, skipHostnameVerification bool) (func(ctx context.Context, network string, address string) (net.Conn, error), error) {
	dialer := &proxy.Dialer{}
	var cfg tls.Config
	if skipHostnameVerification {
		cfg = tls.Config{InsecureSkipVerify: true}
	} else {
		u, err := url.ParseRequestURI(endpoint)
		if err != nil {
			return nil, err
		}
		cfg = tls.Config{ServerName: u.Hostname()}
	}
	dialer.Config = &cfg
	return dialer.DialContext, nil
}

// NewWithSDKConfig creates a new instance of the DAX client with an aws.Config.
//
// Example:
//		config := aws.Config{
//			Region: "us-east-1",
//			EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...any) (aws.Endpoint, error) {
//				return aws.Endpoint{
//					URL: "dax://mycluster.frfx8h.clustercfg.dax.usw2.amazonaws.com:8111",
//				}, nil
//			}),
//		}
//
// 		// Create a DAX client from just a session.
// 		svc := dax.NewWithSDKConfig(ctx, config)
func NewWithSDKConfig(ctx context.Context, config aws.Config) (*Dax, error) {
	dc := DefaultConfig()
	dc.mergeFrom(config)
	return New(ctx, dc)
}

func (c *Config) mergeFrom(ac aws.Config) {
	if r := ac.RetryMaxAttempts; r > 0 {
		c.WriteRetries = r
		c.ReadRetries = r
	}

	if ac.Logger != nil {
		c.Logger = ac.Logger
	}

	if ac.Credentials != nil {
		c.Credentials = ac.Credentials
	}
	if ac.EndpointResolverWithOptions != nil {
		c.EndpointResolver = ac.EndpointResolverWithOptions
	}
	if ac.Region != "" {
		c.Region = ac.Region
	}
}

func (c *Config) requestOptions(read bool, opts ...func(*dynamodb.Options)) client.RequestOptions {
	r := c.WriteRetries
	if read {
		r = c.ReadRetries
	}

	opt := client.RequestOptions{}
	opt.Logger = c.Logger
	opt.RetryMaxAttempts = r

	// merge from request options
	for _, o := range opts {
		o(&opt.Options)
	}

	if opt.Retryer != nil {
		opt.Retryer = retry.NewStandard(
			func(options *retry.StandardOptions) {
				options.MaxAttempts = r
				options.Retryables = append(options.Retryables, retry.IsErrorRetryableFunc(client.IsErrorRetryable))
			},
		)
	}
	return opt
}
