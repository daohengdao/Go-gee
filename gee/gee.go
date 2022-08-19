package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义路由映射的处理方法
type HandlerFunc func(c *Context)

// Engine 实现ServeHTTP的接口
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup //存储所有的Group
}

type RouterGroup struct {
	prefix      string        //前缀
	middlewares []HandlerFunc //中间件
	parent      *RouterGroup  //父节点
	engine      *Engine
}

// New Engine的构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 创建新的RouterGroup
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 添加GET请求
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 添加POST请求
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
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
