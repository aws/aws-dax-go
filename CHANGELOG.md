Release v1.2.9 (2021-07-01)
===
* Retry requests on AuthenticationRequiredException

Release v1.2.8 (2021-03-08)
===
* Add TLS support to DAX go client.

Release v1.2.7 (2021-01-16)
===
* Add support for go modules.

Release v1.2.6 (2020-12-23)
===
* Fix problems reading integer values from the wire that overflow an int64.

Release v1.2.5 (2020-12-17)
===
* Incorporate [pull request #26](https://github.com/aws/aws-dax-go/pull/26), which fixes merging an `aws.Config` into a `dax.Config`.

Release v1.2.4 (2020-12-16)
===
* Remove the assignment of a Dax client pointer to the DynamoDBAPI interface, as requested in [issue 30](https://github.com/aws/aws-dax-go/issues/30).

Release v1.2.3 (2020-12-03)
===
* Add new (but unimplemented by DAX) DynamoDB operations.
* Change behavior of unimplemented operations to return an error instead of panicking, per [issue 27](https://github.com/aws/aws-dax-go/issues/27).

Release v1.2.2 (2020-11-12)
===
* Implement `BatchGetItemPage`, `QueryPage`, and `ScanPage` operations

Release v1.2.1 (2020-09-10)
===
* Support clients connecting with a custom dialer

Release v1.2.0 (2020-04-17)
===
* Expose TransactionCancellationReasons through dynamodb.TransactionCanceledException
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
