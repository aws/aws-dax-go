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
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestSigV4(t *testing.T) {
	creds := aws.Credentials{AccessKeyID: "ak", SecretAccessKey: "sk"}
	endpoint := "dynamodb.us-east-1.amazonaws.com"
	region := "us-east-1"
	payload := "payload"
	time := time.Unix(1519755552, 0).UTC()

	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s", "AWS4-HMAC-SHA256", "20180227T181912Z", "20180227/us-east-1/dax/aws4_request", "059996a9cb1161028c24513d3ea105ae4b06076943a02d9a0f72e2b5d7e6fef0")
	signature := "a7ee7b960e8c11f149d55288149b834cd5f18583401c04b3dac24f4579c3e647"
	actualStringToSign, actualSignature := generateSigV4WithTime(creds, endpoint, region, payload, time)
	if actualStringToSign != stringToSign {
		t.Errorf("expected %v, got %v", stringToSign, actualStringToSign)
	}
	if actualSignature != signature {
		t.Errorf("expected %v, got %v", signature, actualSignature)
	}

	// repeat with session token
	creds = aws.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
		SessionToken:    "st",
	}
	actualStringToSign, actualSignature = generateSigV4WithTime(creds, endpoint, region, payload, time)
	if actualStringToSign != stringToSign {
		t.Errorf("expected %v, got %v", stringToSign, actualStringToSign)
	}
	if actualSignature != signature {
		t.Errorf("expected %v, got %v", signature, actualSignature)
	}
}

func BenchmarkSigV4(b *testing.B) {
	creds := aws.Credentials{AccessKeyID: "ak", SecretAccessKey: "sk"}
	endpoint := "dynamodb.us-east-1.amazonaws.com"
	region := "us-east-1"
	payload := "payload"
	time := time.Unix(1519755552, 0).UTC()

	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s", "AWS4-HMAC-SHA256", "20180227T181912Z", "20180227/us-east-1/dax/aws4_request", "059996a9cb1161028c24513d3ea105ae4b06076943a02d9a0f72e2b5d7e6fef0")
	signature := "a7ee7b960e8c11f149d55288149b834cd5f18583401c04b3dac24f4579c3e647"

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		actualStringToSign, actualSignature := generateSigV4WithTime(creds, endpoint, region, payload, time)
		if actualStringToSign != stringToSign {
			b.Errorf("expected %v, got %v", stringToSign, actualStringToSign)
		}
		if actualSignature != signature {
			b.Errorf("expected %v, got %v", signature, actualSignature)
		}
	}
}
