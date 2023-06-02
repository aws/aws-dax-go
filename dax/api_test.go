package dax

import (
	"context"
	"testing"

	"github.com/aws/aws-dax-go/dax/internal/client"
)

// https://github.com/aws/aws-dax-go/issues/27
func TestUnimplementedBehavior(t *testing.T) {
	dax := createClient(t)

	// CreateBackup is not implemented by DAX
	o, err := dax.CreateBackup(context.Background(), nil)

	if o != nil {
		t.Errorf("expect nil from unimplemented method, got %v", o)
	}
	if err == nil || err.Error() != client.ErrCodeNotImplemented {
		t.Errorf("expect not implemented error, got %v", err)
	}
}

func createClient(t *testing.T) *Dax {
	cfg := DefaultConfig()
	cfg.HostPorts = []string{"127.0.0.1:8111"}
	cfg.Region = "us-west-2"
	dax, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	return dax
}
