package framework

type IGroup interface {
	Get(string, Controller)
	Post(string, Controller)
	Put(string, Controller)
	Delete(string, Controller)
	Group(string) IGroup
}

type Group struct {
	prefix string
	core   *Core
	parent *Group
}

func NewGroup(prefix string, core *Core) *Group {
	return &Group{
		prefix: prefix,
		core:   core,
	}
}

func (g *Group) Get(router string, controller Controller) {
	router = g.getAbsolutePrefix() + router
	g.core.Get(router, controller)
}

func (g *Group) Post(router string, controller Controller) {
	router = g.getAbsolutePrefix() + router
	g.core.Post(router, controller)
}

func (g *Group) Put(router string, controller Controller) {
	router = g.getAbsolutePrefix() + router
	g.core.Put(router, controller)
}

func (g *Group) Delete(router string, controller Controller) {
	router = g.getAbsolutePrefix() + router
	g.core.Delete(router, controller)
}

func (g *Group) Group(prefix string) IGroup {
	group := NewGroup(prefix, g.core)
	group.parent = g
	return group
}

// 获取当前group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}
