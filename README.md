# AWS DAX SDK for Go

aws-dax-go is the official AWS DAX SDK for the Go programming language. https://aws.amazon.com/dynamodb/dax

Checkout our [release notes](https://github.com/aws/aws-dax-go/releases) for
information about the latest bug fixes, updates, and features added to the SDK.

## Getting started
The best way to get started working with the SDK is to use go get to add the SDK
to your Go Workspace manually.

    go get github.com/aws/aws-dax-go

You could also use [Dep](https://github.com/golang/dep) to add the SDK to your
application's dependencies. Using Dep will simplify your update story and help
your application keep pinned to a specific version of the SDK.

    dep ensure -add github.com/aws/aws-dax-go

## Making API requests
This example shows how you can use the AWS DAX SDK to make an API request.

```go
package main

import (
	"fmt"
	"github.com/aws/aws-dax-go/dax"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	
	cfg := dax.DefaultConfig()
	cfg.HostPorts = []string{"dax://mycluster.frfx8h.clustercfg.dax.usw2.amazonaws.com:8111"}
	cfg.Region = "us-west-2"
	client, err := dax.New(cfg)

	if err != nil {
		panic(fmt.Errorf("unable to initialize client %v", err))
	}
	
	
	//Connecion to a secure cluster
	secureEndpoint := "daxs://mycluster.frfx8h.clustercfg.dax.usw2.amazonaws.com"
	secureCfg := dax.DefaultConfig()
	secureCfg.HostPorts = []string{secureEndpoint}
	secureCfg.Region = "us-west-2"
	
	//WARN: Skip hostname verification of TLS connections. 
	//The default is to perform hostname verification, setting this to True will skip verification. 
	//Be sure you understand the implication of doing so, which is the inability to authenticate
	//the cluster that you are connecting to.
	secureCfg.SkipHostnameVerification = false
	
	// DialContext is an optional field in Config.
	// If DialContext is being set in Config for a secure/ encrypted cluster, then use dax.SecureDialContext to 
	// return DialContext. An example of how DailContext can be set using dax.SecureDialContext is shown below.
	secureCfg.DialContext = func(ctx context.Context, network string, address string) (net.Conn, error) {
		//    fmt.Println("Write your custom logic here")
		dialCon, err := dax.SecureDialContext(secureEndpoint, secureCfg.SkipHostnameVerification)
		if err != nil {
			panic(fmt.Errorf("secure dialcontext creation failed %v", err))
		}
		return dialCon(ctx, network, address)
	}
	secureClient, err := dax.New(secureCfg)
	if err != nil {
		panic(fmt.Errorf("unable to initialize client %v", err))
	}
	fmt.Println("secure client created", secureClient)
	
	

	input := &dynamodb.PutItemInput{
		TableName: aws.String("TryDaxGoTable"),
		Item: map[string]*dynamodb.AttributeValue{
			"pk":    {S: aws.String("mykey")},
			"sk":    {N: aws.String("0")},
			"value": {S: aws.String("myvalue")},
		},
	}

	output, err := client.PutItem(input)
	if err != nil {
		panic(fmt.Errorf("unable to make request %v", err))
	}

	fmt.Println("Output: ", output)
}
```

## Feedback and contributing
**GitHub issues:** To provide feedback or report bugs, file GitHub
[Issues](https://github.com/aws/aws-dax-go/issues) on the SDK.
This is the preferred mechanism to give feedback so that other users can engage in
the conversation, +1 issues, etc. Issues you open will be evaluated, and included
in our roadmap.

**Contributing:** You can open pull requests for fixes or additions to the
AWS DAX SDK for Go. All pull requests must be submitted under the Apache 2.0
license and will be reviewed by an SDK team member before being merged in.
Accompanying unit tests, where possible, are appreciated.

## License

This library is licensed under the Apache 2.0 License. 
