Release v1.2.0 (2020-04-17)
===
* Expose CancellationReasons through dynamodb.TransactionCanceledException
* Add public method to create a DAX Config by merging DAX relevant configurations from an AWS Session

Release v1.1.6 (2020-01-31)
===
* Add logic for QueryPages, QueryPagesWithContext, ScanPages, ScanPagesWithContext

Release v1.1.5
===
* Add retry logic for throttling exceptions and support for newer aws-sdk-go v1.25.48

Release v1.1.4 (2019-11-13)
===
* Add a notion of session to the socket pool and associated sockets. It allows to close idle sockets on an error.

Release v1.1.3 (2019-11-08)
===
* Fix potential for goroutine starvation where a goroutine can get stuck waiting for an available socket to do a request.

Release v1.1.2 (2019-10-08)
===
* Fix error thrown from executing an Updateitem() request with multiple Delete actions

Release v1.1.1 (2019-03-20)
===
* Fix retry logic to make sure errors are always propagated
* Improves error handling

Release v1.1.0 (2018-12-20)
===
* Add TransactGetItems and TransactWriteItems Operations

Release v1.0.2 (2018-11-30)
===
* Add DescribeEndpoints functions

Release v1.0.1 (2018-07-12)
===
* Fixes [#1](https://github.com/aws/aws-dax-go/issues/1)
* Improves error handling

Release v1.0.0 (2018-06-26)
===
* Initial version
