package client

import (
	"errors"

	"github.com/aws/smithy-go"
)

func operationError(op string, err error) error {
	// nil error
	if err == nil {
		return nil
	}
	// already smithy.OperationError
	var smithyErr *smithy.OperationError
	if errors.As(err, &smithyErr) {
		return err
	}
	return &smithy.OperationError{
		ServiceID:     service,
		OperationName: op,
		Err:           err,
	}
}
