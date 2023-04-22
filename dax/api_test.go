package dax

import (
	"testing"

	"github.com/aws/aws-dax-go/dax/internal/client"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// https://github.com/aws/aws-dax-go/issues/27
func TestUnimplementedBehavior(t *testing.T) {
	dax := createClient(t)

	// CreateBackup is not implemented by DAX
	o, err := dax.CreateBackup(nil)

	if o != nil {
		t.Errorf("expect nil from unimplemented method, got %v", o)
	}
	if err == nil || err.Error() != client.ErrCodeNotImplemented {
		t.Errorf("expect not implemented error, got %v", err)
	}
}

func TestUnimplementedRequestBehavior(t *testing.T) {
	dax := createClient(t)

	// CreateGlobalTable is not implemented by DAX
	params := &dynamodb.CreateGlobalTableInput{
		GlobalTableName:  nil,
		ReplicationGroup: []*dynamodb.Replica{},
	}
	req, o := dax.CreateGlobalTableRequest(params)

	// Build() should return an error
	err := req.Build()
	if err == nil || err.Error() != client.ErrCodeNotImplemented {
		t.Errorf("expect not implemented error, got %v", err)
	}
	if o.GlobalTableDescription != nil {
		t.Errorf("expect unfilled response from unimplemented method, got %v", o)
	}

	// Send() should return an error
	err = req.Send()
	if err == nil || err.Error() != client.ErrCodeNotImplemented {
		t.Errorf("expect not implemented error, got %v", err)
	}
	if o.GlobalTableDescription != nil {
		t.Errorf("expect unfilled response from unimplemented method, got %v", o)
	}
}

func createClient(t *testing.T) *Dax {
	cfg := DefaultConfig()
	cfg.HostPorts = []string{"127.0.0.1:8111"}
	cfg.Region = "us-west-2"
	dax, err := New(cfg)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	return dax
}
