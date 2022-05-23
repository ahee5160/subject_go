package web_frame

import "net/http"

type Server interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
	Run(addr string) error
}

type HttpServer struct {
	Name string
	handler *Handler
}

func (h *HttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		ctx := NewContext(writer, request)
		handleFunc(ctx)
	})
}

func (h *HttpServer) Run(addr string) error {
	handler := Handler{}
	http.Handle("/", &handler)
	return http.ListenAndServe(addr, nil)
}

func NewHttpServer(name string) Server {
	return &HttpServer{
		Name: name,
	}
}
