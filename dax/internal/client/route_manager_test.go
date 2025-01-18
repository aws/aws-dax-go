package client

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type mockDaxAPI struct {
	id int
}

func (m mockDaxAPI) PutItemWithOptions(_ context.Context, _ *dynamodb.PutItemInput, _ RequestOptions) (*dynamodb.PutItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) DeleteItemWithOptions(_ context.Context, _ *dynamodb.DeleteItemInput, _ RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) UpdateItemWithOptions(_ context.Context, _ *dynamodb.UpdateItemInput, _ RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) GetItemWithOptions(_ context.Context, _ *dynamodb.GetItemInput, _ RequestOptions) (*dynamodb.GetItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) ScanWithOptions(_ context.Context, _ *dynamodb.ScanInput, _ RequestOptions) (*dynamodb.ScanOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) QueryWithOptions(_ context.Context, _ *dynamodb.QueryInput, _ RequestOptions) (*dynamodb.QueryOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) BatchWriteItemWithOptions(_ context.Context, _ *dynamodb.BatchWriteItemInput, _ RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) BatchGetItemWithOptions(_ context.Context, _ *dynamodb.BatchGetItemInput, _ RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) TransactWriteItemsWithOptions(_ context.Context, _ *dynamodb.TransactWriteItemsInput, _ RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) TransactGetItemsWithOptions(_ context.Context, _ *dynamodb.TransactGetItemsInput, _ RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) endpoints(_ context.Context, _ RequestOptions) ([]serviceEndpoint, error) {
	panic("implement me")
}

func Test_disabledRouteManager(t *testing.T) {
	rm := newRouteManager(false, time.Second, nil)
	defer rm.close()
	if rm.isEnabled {
		t.Errorf("Expected route manager to be disabled")
	}

	rm.addRoute("dummy", mockDaxAPI{})
	if len(rm.routes) != 0 {
		t.Errorf("addRoute getting called even with routeManager disabled")
	}

	rm.removeRoute("dummy", mockDaxAPI{}, map[hostPort]clientAndConfig{hostPort{"dummy", 9111}: {client: mockDaxAPI{}}})
	if len(rm.routes) != 0 {
		t.Errorf("addRoute getting called even with routeManager disabled")
	}
}

func Test_setRoutes(t *testing.T) {
	rm := newRouteManager(true, time.Second, nil)
	defer rm.close()
	if len(rm.routes) != 0 {
		t.Errorf("Expected empty routes list, got %v", rm.routes)
	}
	rm.setRoutes(append([]DaxAPI{}, mockDaxAPI{}))
	if len(rm.routes) != 1 {
		t.Errorf("Expected one route but got %v", rm.routes)
	}
}
func Test_getRoute(t *testing.T) {
	rm := newRouteManager(true, time.Second, nil)
	defer rm.close()
	if rm.getRoute(nil) != nil {
		t.Errorf("Expected nil route, got other")
	}

	daxAPI1 := mockDaxAPI{1}
	daxAPI2 := mockDaxAPI{2}
	rm.setRoutes(append([]DaxAPI{}, daxAPI1, daxAPI2))

	if rm.getRoute(daxAPI1) != daxAPI2 {
		t.Errorf("Expected route to be daxAPI2, got other")
	}

	if rm.getRoute(daxAPI1) == daxAPI1 {
		t.Errorf("Expected route to be daxAPI2, got daxAPI1")
	}
}

func Test_addRoute(t *testing.T) {
	rm := newRouteManager(true, time.Second, nil)
	defer rm.close()
	if len(rm.routes) != 0 {
		t.Errorf("Expected empty routes list, got %v", rm.routes)
	}
	rm.addRoute("dummy", mockDaxAPI{})
	if len(rm.routes) != 1 {
		t.Errorf("Expected one route but got %v", rm.routes)
	}
}

func Test_removeRoute(t *testing.T) {
	rm := newRouteManager(true, time.Second, nil)
	defer rm.close()
	if len(rm.routes) != 0 {
		t.Errorf("Expected empty routes list, got %v", rm.routes)
	}
	daxAPI1 := mockDaxAPI{}
	daxAPI2 := mockDaxAPI{}
	daxAPI3 := mockDaxAPI{}
	dummyHostClientMap := map[hostPort]clientAndConfig{
		hostPort{"dummy.1", 9111}: {client: daxAPI1},
		hostPort{"dummy.2", 9111}: {client: daxAPI2},
		hostPort{"dummy.3", 9111}: {client: daxAPI3},
	}
	rm.setRoutes(append([]DaxAPI{}, daxAPI1, daxAPI2, daxAPI3))
	if len(rm.routes) != 3 {
		t.Errorf("Expected three routes but got %v", rm.routes)
	}

	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	if len(rm.routes) != 2 {
		t.Errorf("Expected two routes but got %v", rm.routes)
	}

	// removing same route again should do nothing
	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	if len(rm.routes) != len(dummyHostClientMap) {
		t.Errorf("Expected two routes but got %v", rm.routes)
	}
}

func Test_removeRouteFailOpen(t *testing.T) {
	daxAPI1 := mockDaxAPI{}
	daxAPI2 := mockDaxAPI{}
	daxAPI3 := mockDaxAPI{}
	dummyHostClientMap := map[hostPort]clientAndConfig{
		hostPort{"dummy.1", 9111}: {client: daxAPI1},
		hostPort{"dummy.2", 9111}: {client: daxAPI2},
		hostPort{"dummy.3", 9111}: {client: daxAPI3},
	}
	rm := newRouteManager(true, time.Second, nil)
	defer rm.close()
	if len(rm.routes) != 0 {
		t.Errorf("Expected empty routes list, got %v", rm.routes)
	}
	rm.setRoutes(append([]DaxAPI{}, daxAPI1, daxAPI2, daxAPI3))
	if len(rm.routes) != 3 {
		t.Errorf("Expected three routes but got %v", rm.routes)
	}

	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	rm.removeRoute("dummy.2:9111", daxAPI2, dummyHostClientMap)
	if len(rm.routes) != len(dummyHostClientMap) {
		t.Errorf("Fail Open didn't work as expected")
	}

	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	rm.removeRoute("dummy.2:9111", daxAPI2, dummyHostClientMap)
	if len(rm.routes) != len(dummyHostClientMap) {
		t.Errorf("Fail Open didn't work as expected")
	}

	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	rm.removeRoute("dummy.2:9111", daxAPI2, dummyHostClientMap)
	if rm.isEnabled {
		t.Errorf("Fail Open didn't work as expected")
	}
}

func Test_verifyAndDisable(t *testing.T) {
	rm := newRouteManager(true, time.Second, nil)
	defer rm.close()
	rm.disableDuration = 100 * time.Millisecond
	rm.failOpenTimeList = []time.Time{time.Now(), time.Now(), time.Now()}
	rm.verifyAndDisable(time.Now())
	if rm.isEnabled {
		t.Errorf("Expected isRouteManagerEnabled false but got true")
	}

	// this part tests the timer function
	time.Sleep(105 * time.Millisecond)
	if !rm.isEnabled {
		t.Errorf("Fail Open Callback didn't re-open the routeManager")
	}

	rm.failOpenTimeList = []time.Time{time.Now(), time.Now().Add(-5 * time.Second), time.Now().Add(-5 * time.Second)}
	curTime := time.Now()
	rm.verifyAndDisable(curTime)
	if !rm.isEnabled {
		t.Errorf("Fail Open are not continuous so, it shouldn't disable routeManager")
	}
}

func Test_rebuildRoutes(t *testing.T) {
	rm := newRouteManager(true, time.Second, nil)
	defer rm.close()
	daxAPI1 := mockDaxAPI{}
	daxAPI2 := mockDaxAPI{}
	daxAPI3 := mockDaxAPI{}
	dummyHostClientMap := map[hostPort]clientAndConfig{
		hostPort{"dummy.1", 9111}: {client: daxAPI1},
		hostPort{"dummy.2", 9111}: {client: daxAPI2},
		hostPort{"dummy.3", 9111}: {client: daxAPI3},
	}
	if len(rm.routes) != 0 {
		t.Errorf("Expected zero routes but got %v", rm.routes)
	}
	rm.rebuildRoutes(dummyHostClientMap)
	if len(rm.routes) != len(dummyHostClientMap) {
		t.Errorf("Expected %v routes but got %v", len(dummyHostClientMap), len(rm.routes))
	}
}

func Test_stopTimer(t *testing.T) {
	rm := newRouteManager(true, time.Second, nil)
	defer rm.close()
	timer := time.AfterFunc(rm.disableDuration, func() { rm.isEnabled = true })
	rm.timer = timer
	rm.stopTimer()
	if rm.timer != nil {
		t.Errorf("stopTimer didn't set timer to nil")
	}
	if timer.Stop() {
		t.Errorf("stopTimer didn't stop the timer")
	}
}
