package web_frame

import "net/http"

type Handler struct {
	handlers map[string]func(ctx *Context)
}

func (h *Handler) JoinKey(req *http.Request) string {
	return req.Method + "#" + req.URL.Path
}

func (h *Handler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	key := h.JoinKey(request)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(write, request))
	} else {
		write.WriteHeader(http.StatusNotFound)
		write.Write([]byte("page not found"))
	}
}