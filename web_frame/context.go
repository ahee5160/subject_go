package web_frame

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(obj interface{}) error {
	body, err := ioutil.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, obj)
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

func (c Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

func (c Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func NewContext(write http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: write,
		R: request,
	}
}
