package jee

import "net/http"

type Engine struct {
	srv    *http.Server
	stop   func() error
	router *router
}

// 实现这个方法
// Engine才可以称为一个Handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// HTTP的请求方法
func (e *Engine) Get(pattern string, handlefunc HandleFunc)     {}
func (e *Engine) Post(pattern string, handlefunc HandleFunc)    {}
func (e *Engine) Delete(pattern string, handlefunc HandleFunc)  {}
func (e *Engine) Put(pattern string, handlefunc HandleFunc)     {}
func (e *Engine) Patch(pattern string, handlefunc HandleFunc)   {}
func (e *Engine) Trace(pattern string, handlefunc HandleFunc)   {}
func (e *Engine) Options(pattern string, handlefunc HandleFunc) {}
func (e *Engine) Connect(pattern string, handlefunc HandleFunc) {}
func (e *Engine) Head(pattern string, handlefunc HandleFunc)    {}

// 实现服务器启动的方法
func (e *Engine) Listen(addr string) error {
	return nil
}

// 实现服务器平滑关闭的方法
func (e *Engine) Shutdown() error {
	return e.stop()
}

func New() *Engine {
	return &Engine{
		srv:    &http.Server{},
		router: newRouter(),
	}
}
