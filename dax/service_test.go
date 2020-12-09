package dax

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestConfigMergeFrom(t *testing.T) {
	testCases := []struct {
		testName             string
		daxConfig            Config
		awsConfig            aws.Config
		expectedWriteRetries int
		expectedReadRetries  int
	}{
		{
			testName:             "DefaultConfig merging with an empty aws config should result in keeping the same default retries",
			daxConfig:            DefaultConfig(),
			awsConfig:            aws.Config{},
			expectedWriteRetries: 2,
			expectedReadRetries:  2,
		},
		{
			testName:             "DefaultConfig merging with an aws config that specifies aws.UseServiceDefaultRetries should result in using default retries",
			daxConfig:            DefaultConfig(),
			awsConfig:            aws.Config{MaxRetries: aws.Int(aws.UseServiceDefaultRetries)},
			expectedWriteRetries: 2,
			expectedReadRetries:  2,
		},
		{
			testName:             "DefaultConfig merging with an aws config that specifies a non-negative MaxRetry should result in using that value as both WriteRetries and ReadRetries",
			daxConfig:            DefaultConfig(),
			awsConfig:            aws.Config{MaxRetries: aws.Int(123)},
			expectedWriteRetries: 123,
			expectedReadRetries:  123,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.testName, func(t *testing.T) {
			testCase.daxConfig.mergeFrom(testCase.awsConfig)
			if testCase.daxConfig.WriteRetries != testCase.expectedWriteRetries {
				t.Errorf("write retries is %d, but expected %d", testCase.daxConfig.WriteRetries, testCase.expectedWriteRetries)
			}

			if testCase.daxConfig.ReadRetries != testCase.expectedReadRetries {
				t.Errorf("read retries is %d, but expected %d", testCase.daxConfig.ReadRetries, testCase.expectedReadRetries)
			}
		})
	}
}
