package web

import (
	"net/http"
	"testing"
)

func TestRouter_AddRoute(t *testing.T) {
	var mockFunc HandleFunc = func(ctx Context) {}
	// 构造路由树
	testRoutes := []struct{
		method     string
		path       string
		handleFunc HandleFunc
	}{
		{method: http.MethodGet, path: "/ping", handleFunc: mockFunc},
		{method: http.MethodGet, path: "/ping", handleFunc: mockFunc},
		{method: http.MethodGet, path: "/ping", handleFunc: mockFunc},
		{method: http.MethodGet, path: "/ping", handleFunc: mockFunc},
		{method: http.MethodGet, path: "/ping", handleFunc: mockFunc},
	}
	r := newRouter()
	for _, route := range testRoutes {
		r.AddRoute(route.method, route.path, mockFunc)
	}
	// 验证路由树
	// 断言路由树符合预期
	wantRouter := &router{
		trees: map[string]*node{},
	}
	if err := wantRouter.isEqual(*r); err != nil {
		t.Error(err)
	}
}
