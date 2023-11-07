package jee

import "net/http"

// 框架引擎需要实现的接口
type jee interface {
	// 这个是必须要实现的
	http.Handler
	// Listen() 这个方法用户启动服务
	Listen(addr string) error
	// Shutdown() 这个方法用于关闭服务
	Shutdown() error
	// addRoute() 此方法用户添加路由
	addRoute(method string, pattern string, handlefunc HandleFunc)
}
