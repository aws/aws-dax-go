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
	"github.com/aws/aws-dax-go/dax/internal/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"time"
)

// Dax makes requests to the Amazon DAX API, which conforms to the DynamoDB API.

// Dax methods are safe to use concurrently
type Dax struct {
	client client.DaxAPI
	config Config
}

const ServiceName = "dax"

type Config struct {
	client.Config

	// Default request options
	RequestTimeout time.Duration
	WriteRetries   int
	ReadRetries    int

	LogLevel aws.LogLevelType
	Logger   aws.Logger
}

var defaultConfig = Config{
	Config: client.DefaultConfig(),

	RequestTimeout: 1 * time.Minute,
	WriteRetries:   2,
	ReadRetries:    2,
	LogLevel:       aws.LogOff,
	Logger:         aws.NewDefaultLogger(),
}

func DefaultConfig() Config {
	return defaultConfig
}

func New(cfg Config) (*Dax, error) {
	c, err := client.New(cfg.Config)
	if err != nil {
		return nil, err
	}
	return &Dax{client: c, config: cfg}, nil
}

func NewWithSession(session session.Session) (*Dax, error) {
	if err := client.ValidateHandlers(session.Handlers, false); err != nil {
		return nil, err
	}
	dc := DefaultConfig()
	if session.Config != nil {
		if err := client.ValidateConfig(*session.Config, false); err != nil {
			return nil, err
		}
		dc.mergeFrom(*session.Config)
	}
	return New(dc)
}

func (c *Config) mergeFrom(ac aws.Config) {
	if r := ac.MaxRetries; r != nil {
		c.WriteRetries = *r
		c.ReadRetries = *r
	}
	if ac.Logger != nil {
		c.Logger = ac.Logger
	}
	if ac.LogLevel != nil {
		c.LogLevel = *ac.LogLevel
	}

	if ac.Credentials != nil {
		c.Credentials = ac.Credentials
	}
	if ac.Endpoint != nil {
		c.HostPorts = []string{*ac.Endpoint}
	}
	if ac.Region != nil {
		c.Region = *ac.Region
	}
}

func (c *Config) requestOptions(read bool, ctx context.Context, opts ...request.Option) (client.RequestOptions, context.CancelFunc, error) {
	r := c.WriteRetries
	if read {
		r = c.ReadRetries
	}
	var cfn context.CancelFunc
	if ctx == nil && c.RequestTimeout > 0 {
		ctx, cfn = context.WithTimeout(aws.BackgroundContext(), c.RequestTimeout)
	}
	opt := client.RequestOptions{
		LogLevel:   c.LogLevel,
		Logger:     c.Logger,
		MaxRetries: r,
	}
	if err := opt.MergeFromRequestOptions(ctx, opts...); err != nil {
		return client.RequestOptions{}, nil, err
	}
	return opt, cfn, nil
}

var _ dynamodbiface.DynamoDBAPI = (*Dax)(nil)
