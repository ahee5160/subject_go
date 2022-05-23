package main

import (
	"fmt"
	"net/http"
	"test/web_frame"
)

func ping(context *web_frame.Context) {
	text := fmt.Sprintf("%s %s pong", context.R.Host, context.R.Method)
	context.WriteJson(http.StatusOK, text)
}

type signupRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type signupResponse struct {
	UserID int `json:"user_id"`
}

func signup(context *web_frame.Context) {
	req := &signupRequest{}
	err := context.ReadJson(req)
	if err != nil {
		resp := fmt.Sprintf("read body failed, err: %s", err)
		context.WriteJson(http.StatusBadRequest, resp)
		return
	}
	resp := signupResponse{UserID: 123}
	if err = context.WriteJson(http.StatusOK, resp); err != nil {
		fmt.Printf("read body failed, err: %s", err)
	}
}

func main() {
	server := web_frame.NewHttpServer("test-server")
	server.Route("/ping", ping)
	server.Route("/user/signup", signup)
	server.Run(":5160")
}
