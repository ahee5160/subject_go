package web

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx Context)

type Server interface {
	http.Handler
	// 生命周期管理
	Start() error
	// 路由注册功能
	AddRoute(method string, path string, handleFunc HandleFunc)
}

// 确保 HTTPServer 一定实现了 Server
var _ Server = &HTTPServer{}

type HTTPServer struct {
	addr string
	*router
}

func NewHttpServer(addr string) *HTTPServer {
	return &HTTPServer{
		addr:   addr,
		router: newRouter(),
	}
}

// ServeHTTP 处理请求的入口
func (h *HTTPServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// 构建 Context
	ctx := Context{
		Req:  req,
		Resp: resp,
	}
	h.serve(ctx)
}

// serve 路由匹配并执行命中的业务逻辑
func (h HTTPServer) serve(ctx Context) {
	
}

func (h *HTTPServer) Start() error {
	listener, err := net.Listen("tcp", h.addr)
	if err != nil {
		return err
	}
	// 此处可以让用户注册 after start 回调，或执行一下业务所需的前置条件
	return http.Serve(listener, h)
}

func (h *HTTPServer) Get(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

func (h *HTTPServer) Post(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodPost, path, handleFunc)
}
