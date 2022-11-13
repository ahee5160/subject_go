package web

import (
	"fmt"
	"reflect"
)

// 路由森林
type router struct {
	// method -> 路由树根节点
	trees map[string]*node
}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}

// AddRoute 注册路由
func (r *router) AddRoute(method string, path string, handleFunc HandleFunc) {
	fmt.Println("add route")
}

type node struct {
	path string
	// 子 path 到子节点的映射
	children map[string]*node
	// 用户注册的业务逻辑
	handler HandleFunc
}

func (r *router) isEqual(x *router) error {
	if r == nil && x == nil {
		return nil
	}
	if r == nil || x == nil {
		return fmt.Errorf("route tree is nil")
	}
	for method, tree := range r.trees {
		y, ok := x.trees[method]
		if !ok {
			return fmt.Errorf("hasn't %s method", method)
		}
		if !tree.isEqual(y) {
			return fmt.Errorf("%s method route tree are diff", method)
		}
	}
	return nil
}

func (n *node) isEqual(y *node) bool {
	if n == nil && y == nil {
		return true
	}
	if n == nil || y == nil {
		return false
	}
	if n.path != y.path {
		return false
	}
	if reflect.ValueOf(n.handler) != reflect.ValueOf(y.handler) {
		return false
	}
	if len(n.children) != len(y.children) {
		return false
	}
	for subPath, subNode := range n.children {
		z, ok := y.children[subPath]
		if !ok {
			return false
		}
		if !subNode.isEqual(z) {
			return false
		}
	}
	return true
}
