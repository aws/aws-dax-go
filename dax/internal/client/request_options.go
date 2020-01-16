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
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
)

type RequestOptions struct {
	LogLevel aws.LogLevelType
	Logger   aws.Logger

	RetryDelay time.Duration
	//Retryer implements equal jitter backoff stratergy for throttled requests
	Retryer    DaxRetryer
	MaxRetries int
	//SleepDelayFn is used for non-throttled retryable requests
	SleepDelayFn func(time.Duration)
	Context      aws.Context
}

func (o *RequestOptions) applyTo(r *request.Request) {
	if r != nil {
		r.Config.LogLevel = aws.LogLevel(o.LogLevel)
		r.Config.Logger = o.Logger

		r.RetryDelay = o.RetryDelay
		r.Config.MaxRetries = aws.Int(o.MaxRetries)
		r.Config.SleepDelay = o.SleepDelayFn
		if o.Context != nil {
			r.SetContext(o.Context)
		}
	}
}

func (o *RequestOptions) MergeFromRequestOptions(ctx aws.Context, opts ...request.Option) error {
	if len(opts) == 0 {
		if ctx != nil {
			o.Context = ctx
		}
		return nil
	}

	// New request has to be created to avoid panics when setting fields
	r := request.New(aws.Config{}, metadata.ClientInfo{}, request.Handlers{}, nil, &request.Operation{}, nil, nil)
	r.ApplyOptions(opts...)
	if err := o.mergeFromRequest(r, true); err != nil {
		return err
	}
	if ctx != nil {
		o.Context = ctx
	}
	return nil
}

func (o *RequestOptions) mergeFromRequest(r *request.Request, validate bool) error {
	if r == nil {
		return nil
	}
	if validate {
		if err := ValidateRequest(r); err != nil {
			return err
		}
	}
	if r.Config.LogLevel != nil {
		o.LogLevel = *r.Config.LogLevel
	}
	if r.Config.Logger != nil {
		o.Logger = r.Config.Logger
	}
	if r.RetryDelay >= 0 {
		o.RetryDelay = r.RetryDelay
	}
	if r.Config.MaxRetries != nil {
		o.MaxRetries = *r.Config.MaxRetries
	}
	if r.Config.SleepDelay != nil {
		o.SleepDelayFn = r.Config.SleepDelay
	}
	if r.Context() != nil { // TODO Should the Context() from Request override the one in RequestOptions
		o.Context = r.Context()
	}
	return nil
}

func ValidateRequest(r *request.Request) error {
	if r == nil {
		return nil
	}
	if err := ValidateHandlers(r.Handlers, true); err != nil {
		return err
	}
	if r.Retryable != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: Retryable", nil)
	}
	if len(r.SignedHeaderVals) > 0 {
		return awserr.New(request.InvalidParameterErrCode, "custom signed headers not supported", nil)
	}

	return ValidateConfig(r.Config, true)
}

func ValidateHandlers(h request.Handlers, expectDaxHandlers bool) error {
	if h.Validate.Len() > 0 || h.Sign.Len() > 0 || h.ValidateResponse.Len() > 0 ||
		h.Unmarshal.Len() > 0 || h.UnmarshalMeta.Len() > 0 || h.UnmarshalError.Len() > 0 ||
		h.Retry.Len() > 0 || h.AfterRetry.Len() > 0 || h.Complete.Len() > 0 {
		return awserr.New(request.InvalidParameterErrCode, "custom handlers not supported", nil)
	}
	e := 0
	if expectDaxHandlers {
		e = 1
	}
	if h.Build.Len() > e || h.Send.Len() > e {
		return awserr.New(request.InvalidParameterErrCode, "custom build or send handlers not supported", nil)
	}
	return nil
}

func ValidateConfig(c aws.Config, isRequestConfig bool) error {
	if c.CredentialsChainVerboseErrors != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: CredentialsChainVerboseErrors", nil)
	}
	if c.EndpointResolver != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: EndpointResolver", nil)
	}
	if c.EnforceShouldRetryCheck != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: EnforceShouldRetryCheck", nil)
	}
	if c.DisableSSL != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: DisableSSL", nil)
	}
	if c.HTTPClient != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: HTTPClient", nil)
	}
	if c.Retryer != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: Retryer", nil)
	}
	if c.DisableParamValidation != nil && *c.DisableParamValidation {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: DisableParamValidation", nil)
	}
	if c.DisableComputeChecksums != nil && *c.DisableComputeChecksums {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: DisableComputeChecksums", nil)
	}
	if c.UseDualStack != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: UseDualStack", nil)
	}
	if c.DisableRestProtocolURICleaning != nil {
		return awserr.New(request.InvalidParameterErrCode, "unsupported config: DisableRestProtocolURICleaning", nil)
	}
	// Skip validation of S3* and EC2* options
	if isRequestConfig {
		if c.Credentials != nil {
			return awserr.New(request.InvalidParameterErrCode, "unsupported config: Credentials per request. Set Credentials at client init", nil)
		}
		if c.Endpoint != nil {
			return awserr.New(request.InvalidParameterErrCode, "unsupported config: Endpoint per request. Set Endpoint at client init", nil)
		}
		if c.Region != nil {
			return awserr.New(request.InvalidParameterErrCode, "unsupported config: Region per request. Set Region at client init", nil)
		}
	}
	return nil
}
