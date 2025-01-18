package client

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/aws/smithy-go/logging"
)

const failOpenThreshold = 3

type routeManager struct {
	routes                 []DaxAPI
	isEnabled              bool
	failOpenTimeList       []time.Time   // recent times when fail open was enabled
	multipleFailOpenWindow time.Duration // if we see multiple fail open events within this window, we will disable route manager.
	disableDuration        time.Duration // disable route manager for this duration after multiple fail open in a row
	timer                  *time.Timer
	logger                 logging.Logger
}

func newRouteManager(isEnabled bool, healthCheckDuration time.Duration, logger logging.Logger) *routeManager {
	return &routeManager{
		routes:                 make([]DaxAPI, 0),
		isEnabled:              isEnabled,
		failOpenTimeList:       make([]time.Time, 0),
		multipleFailOpenWindow: time.Duration(math.Ceil(failOpenThreshold/2.0)) * healthCheckDuration,
		disableDuration:        10 * time.Minute, // Disable route manager after multiple fail open in a row
		logger:                 logger,
	}
}

func (r *routeManager) debugLog(format string, args ...interface{}) {
	if r.logger != nil {
		r.logger.Logf(ClassificationDebug, format, args)
	}
}

func (r *routeManager) setRoutes(routes []DaxAPI) {
	r.routes = routes
}

func (r *routeManager) getAllRoutes() []DaxAPI {
	return r.routes
}

func (r *routeManager) getRoute(prev DaxAPI) DaxAPI {
	numRoutes := len(r.routes)
	if numRoutes == 0 {
		return nil
	}
	randInt := rand.Intn(numRoutes)
	if r.routes[randInt] == prev {
		randInt++
		randInt = randInt % numRoutes
	}
	return r.routes[randInt]
}

func (r *routeManager) addRoute(endpoint string, route DaxAPI) {
	if !r.isEnabled {
		return
	}
	for _, curRoute := range r.routes {
		if curRoute == route {
			return
		}
	}
	r.routes = append(r.routes, route)
	r.debugLog(fmt.Sprintf("Added route: %s to active routes", endpoint))
}

func (r *routeManager) removeRoute(endpoint string, route DaxAPI, allClients map[hostPort]clientAndConfig) {
	if !r.isEnabled {
		return
	}

	// Never remove more than one third of nodes
	if float32(len(r.routes)-1) < 2*float32(len(allClients))/3 {
		r.debugLog("FailOpen: Added all routes back to active routes")
		curTime := time.Now()
		// Fail Open to all routes.
		r.rebuildRoutes(allClients)
		r.verifyAndDisable(curTime)
		return
	}
	for i, activeRoute := range r.routes {
		if activeRoute == route {
			r.routes = append(r.routes[:i], r.routes[i+1:]...)
			r.debugLog(fmt.Sprintf("Removed route: %s from active routes", endpoint))
			return
		}
	}
}

func (r *routeManager) verifyAndDisable(failOpenTime time.Time) {
	// this method will verify if there are more than failOpenThreshold FailOpens in given window,
	// if yes, then the route manager will be disabled for some time
	newFailOpenList := make([]time.Time, 0)
	for _, t := range r.failOpenTimeList {
		if failOpenTime.Sub(t) < r.multipleFailOpenWindow {
			newFailOpenList = append(newFailOpenList, t)
		}
	}
	newFailOpenList = append(newFailOpenList, failOpenTime)
	r.failOpenTimeList = newFailOpenList
	if len(r.failOpenTimeList) < failOpenThreshold {
		return
	}
	r.stopTimer()
	r.isEnabled = false
	r.timer = time.AfterFunc(r.disableDuration, func() { r.isEnabled = true })
}

func (r *routeManager) rebuildRoutes(allClients map[hostPort]clientAndConfig) {
	newRoutes := make([]DaxAPI, 0, len(allClients))
	for _, cliAndCfg := range allClients {
		newRoutes = append(newRoutes, cliAndCfg.client)
	}
	r.routes = newRoutes
}

func (r *routeManager) stopTimer() {
	if r.timer == nil {
		return
	}
	r.timer.Stop()
	r.timer = nil
}

func (r *routeManager) close() {
	r.stopTimer()
}

type RouteManager interface {
	setRoutes(routes []DaxAPI)
	getAllRoutes() []DaxAPI
	getRoute(prev DaxAPI) DaxAPI
	addRoute(endpoint string, route DaxAPI)
	removeRoute(endpoint string, route DaxAPI, allClients map[hostPort]clientAndConfig)
	close()
}
