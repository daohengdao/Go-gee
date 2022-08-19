package gee

import (
	"net/http"
)

// HandlerFunc 定义路由映射的处理方法
type HandlerFunc func(c *Context)

// Engine 实现ServeHTTP的接口
type Engine struct {
	router *router
}

// New Engine的构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

// GET 添加GET请求
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST 添加POST请求
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// Run 启动服务
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

//ServeHTTP 解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	e.router.handle(c)
}
