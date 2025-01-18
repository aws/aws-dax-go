package client

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

// Generate all unit tests for onReadRequest method in healthStatus file

type mockRouteListener struct {
	mock.Mock
}

func (mrl *mockRouteListener) addRoute(endpoint string, route DaxAPI) {
	mrl.Called()
}

func (mrl *mockRouteListener) removeRoute(endpoint string, route DaxAPI) {
	mrl.Called()
}

func (mrl *mockRouteListener) isRouteManagerEnabled() bool {
	return true
}

func Test_nilRouteListener(t *testing.T) {
	hs := newHealthStatus("dummy", nil)
	_, ok := hs.(*disabledHealthStatus)
	if !ok {
		t.Errorf("disabledHealthStatus not initialized with empty routeListener")
	}
}

func Test_onErrorInReadRequest_differentError(t *testing.T) {
	mrl := &mockRouteListener{}
	hs := newHealthStatus("dummy", mrl)
	ehs, ok := hs.(*enabledHealthStatus)
	if !ok {
		t.Errorf("enabledHealthStatus not initialized with empty routeListener")
	}

	if ehs.curReadTimeoutCount != 0 {
		t.Errorf("curReadTimeoutCount should be initially 0")
	}
	prevReadTimeoutCount := ehs.curReadTimeoutCount
	hs.onErrorInReadRequest(context.DeadlineExceeded, nil)
	if ehs.curReadTimeoutCount != prevReadTimeoutCount+1 {
		t.Errorf("onErrorInReadRequest failed to increment curReadTimeoutCount on timeout error")
	}

	prevReadTimeoutCount = ehs.curReadTimeoutCount
	hs.onErrorInReadRequest(errors.New("non timeout error"), nil)
	if ehs.curReadTimeoutCount != prevReadTimeoutCount {
		t.Errorf("onErrorInReadRequest incremented curReadTimeoutCount on non timeout error")
	}
}

func Test_onErrorInReadRequest_removeRouteCall(t *testing.T) {
	mrl := &mockRouteListener{}
	mrl.On("removeRoute").Return(nil).Times(1)
	hs := newHealthStatus("dummy", mrl)
	ehs, _ := hs.(*enabledHealthStatus)
	for i := 1; i <= timeoutErrorThreshold; i++ {
		hs.onErrorInReadRequest(context.DeadlineExceeded, nil)
		if i < timeoutErrorThreshold {
			mrl.AssertNotCalled(t, "removeRoute")
			if !ehs.isHealthy {
				t.Errorf("isHealthy should be true")
			}
		} else {
			mrl.AssertCalled(t, "removeRoute")
			if ehs.isHealthy {
				t.Errorf("isHealthy should be false")
			}
		}

	}
}

func Test_onSuccessInReadRequest(t *testing.T) {
	mrl := &mockRouteListener{}
	hs := newHealthStatus("dummy", mrl)
	ehs, _ := hs.(*enabledHealthStatus)
	ehs.curReadTimeoutCount = 5
	hs.onSuccessInReadRequest()
	if ehs.curReadTimeoutCount != 0 {
		t.Errorf("onSuccessInReadRequest failed to set curReadTimeoutCount to 0")
	}

	ehs.isHealthy = false
	ehs.curReadTimeoutCount = 5
	hs.onSuccessInReadRequest()
	if ehs.curReadTimeoutCount != 5 {
		t.Errorf("onSuccessInReadRequest reset the curReadTimeoutCount on unhealthy client")
	}
}

func Test_onHealthCheckSuccess(t *testing.T) {
	mrl := &mockRouteListener{}
	mrl.On("addRoute").Return(nil).Times(1)
	hs := newHealthStatus("dummy", mrl)
	ehs, _ := hs.(*enabledHealthStatus)
	ehs.isHealthy = false
	ehs.curReadTimeoutCount = 5

	ehs.onHealthCheckSuccess(nil)
	if !ehs.isHealthy {
		t.Errorf("onHealthCheckSuccess failed to set isHealthy to true")
	}
	if ehs.curReadTimeoutCount != 0 {
		t.Errorf("onHealthCheckSuccess failed to set curReadTimeoutCount to 0")
	}
	mrl.AssertCalled(t, "addRoute")
}
