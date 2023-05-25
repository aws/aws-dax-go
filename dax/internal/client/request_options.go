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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
)

type RequestOptions struct {
	dynamodb.Options
	//
	//Logger logging.Logger
	//
	//RetryDelay time.Duration
	////Retryer implements equal jitter backoff strategy for throttled requests
	//Retryer    DaxRetryer
	//MaxRetries int
	////SleepDelayFn is used for non-throttled retryable requests
	//SleepDelayFn func(time.Duration)
	//Context      aws.Context
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
