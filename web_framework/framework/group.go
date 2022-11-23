package framework

type IGroup interface {
	Get(string, ...Controller)
	Post(string, ...Controller)
	Put(string, ...Controller)
	Delete(string, ...Controller)
	Group(string) IGroup
	Use(...Controller) IGroup
}

type Group struct {
	prefix      string
	core        *Core
	parent      *Group
	middlewares []Controller // 组中间件
}

func NewGroup(prefix string, core *Core) *Group {
	return &Group{
		prefix: prefix,
		core:   core,
	}
}

func (g *Group) Get(router string, controller ...Controller) {
	router = g.getAbsolutePrefix() + router
	controllers := append(g.GetAllMiddlewares(), controller...)
	g.core.Get(router, controllers...)
}

func (g *Group) Post(router string, controller ...Controller) {
	router = g.getAbsolutePrefix() + router
	controllers := append(g.GetAllMiddlewares(), controller...)
	g.core.Post(router, controllers...)
}

func (g *Group) Put(router string, controller ...Controller) {
	router = g.getAbsolutePrefix() + router
	controllers := append(g.GetAllMiddlewares(), controller...)
	g.core.Put(router, controllers...)
}

func (g *Group) Delete(router string, controller ...Controller) {
	router = g.getAbsolutePrefix() + router
	controllers := append(g.GetAllMiddlewares(), controller...)
	g.core.Delete(router, controllers...)
}

func (g *Group) Group(prefix string) IGroup {
	group := NewGroup(prefix, g.core)
	group.parent = g
	return group
}

func (g *Group) Use(middleware ...Controller) IGroup {
	g.middlewares = append(g.middlewares, middleware...)
	return g
}

func (g *Group) GetAllMiddlewares() []Controller {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.GetAllMiddlewares(), g.middlewares...)
}

// 获取当前group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}
