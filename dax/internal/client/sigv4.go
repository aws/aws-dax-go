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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	headerHost          = "host"
	headerDate          = "x-amz-date"
	headerSecurityToken = "x-amz-security-token"
	method              = "POST"
	service             = "dax"
	signerTerminator    = "aws4_request"
	signVersion         = "AWS4"
	signMethod          = "AWS4-HMAC-SHA256"

	dateTimeFormat = "20060102T150405Z"
	dateFormat     = "20060102"
)

// headers in the canonical request should be sorted by lowercase character code
var signedHeaders = []string{headerHost, headerDate}
var signedHeadersBytes = []byte(strings.Join(signedHeaders, ";"))

func generateSigV4(credentials aws.Credentials, hostname, region, payload string) (string, string) {
	return generateSigV4WithTime(credentials, hostname, region, payload, time.Now().UTC())
}

func generateSigV4WithTime(credentials aws.Credentials, hostname, region, payload string, time time.Time) (string, string) {
	headers := sigv4Headers(hostname, time, credentials.SessionToken)

	canonicalRequest := make([]byte, 0, 256) // 4 + 1 + 1 + 1 + 1 + (74+1+28+1) + 1 + 16 + 1 + 64 + 25%
	canonicalRequest = appendCanonicalRequest(canonicalRequest, method, headers, signedHeadersBytes, payload)

	// credentialScope = 8 + 1 + (15) + 1 + 3 + 1 + 13 + (50%)
	// stringToSign = 20 + 1 + 20 + 1 + 64 + 1 + 64
	stringToSign := make([]byte, 0, 180)
	stringToSign = appendStringToSign(stringToSign, signMethod, time, region, service, signerTerminator, canonicalRequest)

	signingKey := sigv4SigningKey(credentials.SecretAccessKey, time.Format(dateFormat), region, service)
	signature := hex.EncodeToString(hmacSha256(signingKey, stringToSign, nil))

	return string(stringToSign), signature
}

func sigv4SigningKey(key, datestamp, region, service string) []byte {
	ksecret := []byte(signVersion + key)
	kdate := hmacSha256(ksecret, []byte(datestamp), nil)
	kregion := hmacSha256(kdate, []byte(region), nil)
	kservice := hmacSha256(kregion, []byte(service), nil)
	ksigning := hmacSha256(kservice, []byte(signerTerminator), nil)
	return ksigning
}

func hmacSha256(key, msg, out []byte) []byte {
	m := hmac.New(sha256.New, key)
	m.Write(msg)
	return m.Sum(out)
}

func appendSha256Hex(in, data []byte) []byte {
	sum := sha256.Sum256(data)
	hexlen := hex.EncodedLen(len(sum))

	out := ensureCapacity(in, hexlen)
	slen := len(out)
	out = out[0 : slen+hexlen]

	b := out[slen:]
	hex.Encode(b, sum[:])
	return out
}

func ensureCapacity(in []byte, n int) []byte {
	l := len(in)
	c := cap(in)
	if c-l >= n {
		return in
	}
	// do not extend more than required, as this is used to ensure capacity before last append
	out := make([]byte, l, l+n)
	copy(out, in)
	return out
}

func appendStringToSign(in []byte, signMethod string, time time.Time, region string, service, terminator string, canonicalRequest []byte) []byte {
	out := append(in, signMethod...)
	out = append(out, '\n')
	out = time.AppendFormat(out, dateTimeFormat)
	out = append(out, '\n')
	out = appendCredentialScope(out, time, region, service, terminator)
	out = append(out, '\n')
	out = appendSha256Hex(out, canonicalRequest)
	return out
}

func appendCredentialScope(in []byte, time time.Time, region, service, terminator string) []byte {
	out := time.AppendFormat(in, dateFormat)
	out = append(out, '/')
	out = append(out, region...)
	out = append(out, '/')
	out = append(out, service...)
	out = append(out, '/')
	out = append(out, terminator...)
	return out
}

func appendCanonicalRequest(in []byte, method string, headers map[string]string, signedHeaders []byte, payload string) []byte {
	out := append(in, method...)
	out = append(out, '\n')
	out = append(out, '/')
	out = append(out, '\n')
	// uri = ""
	out = append(out, '\n')
	out = appendCanonicalHeaders(out, headers)
	out = append(out, '\n')
	out = append(out, signedHeaders...)
	out = append(out, '\n')
	out = appendSha256Hex(out, []byte(payload))
	return out
}

func appendCanonicalHeaders(in []byte, headers map[string]string) []byte {
	out := in
	for _, h := range signedHeaders {
		out = append(out, []byte(h)...)
		out = append(out, ':')
		out = append(out, []byte(headers[h])...)
		out = append(out, '\n')
	}
	return out
}

func sigv4Headers(hostname string, time time.Time, token string) map[string]string {
	headers := make(map[string]string)
	headers[headerHost] = hostname
	headers[headerDate] = time.Format(dateTimeFormat)
	if token != "" {
		headers[headerSecurityToken] = token
	}
	return headers
}
