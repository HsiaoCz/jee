package jee

import "net/http"

type Engine struct{}

// 实现这个方法
// Engine才可以称为一个Handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// HTTP的请求方法
func (e *Engine) Get()     {}
func (e *Engine) Post()    {}
func (e *Engine) Delete()  {}
func (e *Engine) Put()     {}
func (e *Engine) Patch()   {}
func (e *Engine) Trace()   {}
func (e *Engine) Options() {}
func (e *Engine) Connect() {}
func (e *Engine) Head()    {}

// 实现服务器启动的方法
func (e *Engine) Listen() error {
	return nil
}

// 实现服务器平滑关闭的方法
func (e *Engine) Shutdown() error {
	return nil
}
