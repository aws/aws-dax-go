package client

import (
	"sync"
)

const timeoutErrorThreshold = 5 // remove the client from route list if it has seen 5 consecutive timeout errors

type HealthStatus interface {
	onErrorInReadRequest(err error, route DaxAPI)
	onSuccessInReadRequest()
	onHealthCheckSuccess(route DaxAPI)
}

type enabledHealthStatus struct {
	routeListener       RouteListener
	endpoint            string
	lock                sync.RWMutex
	isHealthy           bool // is the client healthy?
	curReadTimeoutCount int  // total timeout in read requests
}

func newHealthStatus(endpoint string, routeListener RouteListener) HealthStatus {
	if routeListener != nil && routeListener.isRouteManagerEnabled() {
		return &enabledHealthStatus{routeListener: routeListener, endpoint: endpoint, lock: sync.RWMutex{}, isHealthy: true}
	}
	return &disabledHealthStatus{}
}

func (hs *enabledHealthStatus) onErrorInReadRequest(err error, route DaxAPI) {
	if !isIOError(err) {
		return
	}
	hs.lock.RLock()
	if !hs.isHealthy {
		hs.lock.RUnlock()
		return
	}
	hs.lock.RUnlock()

	hs.lock.Lock()
	defer hs.lock.Unlock()
	hs.curReadTimeoutCount += 1
	if hs.curReadTimeoutCount >= timeoutErrorThreshold {
		hs.isHealthy = false
		hs.routeListener.removeRoute(hs.endpoint, route)
	}
}

func (hs *enabledHealthStatus) onSuccessInReadRequest() {
	// Acquire exclusive lock only if there are some read timeouts and route is healthy.
	// Otherwise if either route is unhealthy or curReadTimeoutCount is already zero, we don't need exclusive lock.
	hs.lock.RLock()
	if hs.curReadTimeoutCount == 0 || !hs.isHealthy {
		hs.lock.RUnlock()
		return
	}
	hs.lock.RUnlock()

	hs.lock.Lock()
	defer hs.lock.Unlock()
	hs.curReadTimeoutCount = 0
}

func (hs *enabledHealthStatus) onHealthCheckSuccess(route DaxAPI) {
	hs.lock.RLock()
	if hs.curReadTimeoutCount == 0 && hs.isHealthy {
		hs.lock.RUnlock()
		return
	}
	hs.lock.RUnlock()

	hs.lock.Lock()
	defer hs.lock.Unlock()
	hs.curReadTimeoutCount = 0
	if !hs.isHealthy {
		hs.isHealthy = true
		hs.routeListener.addRoute(hs.endpoint, route)
	}
}

type disabledHealthStatus struct{}

func (hs *disabledHealthStatus) onErrorInReadRequest(err error, route DaxAPI) {}

func (hs *disabledHealthStatus) onSuccessInReadRequest() {}

func (hs *disabledHealthStatus) onHealthCheckSuccess(route DaxAPI) {}
