package v8local

import (
	"sync"

	"github.com/herb-go/v8go"
)

func NewContext(opt ...v8go.ContextOption) *Context {
	c := &Context{Raw: v8go.NewContext(opt...)}
	c.objectTemplate = v8go.NewObjectTemplate(c.Raw.Isolate())
	c.nullvalue = v8go.Null(c.Raw.Isolate())
	return c
}

type Context struct {
	locker         sync.RWMutex
	objectTemplate *v8go.ObjectTemplate
	Raw            *v8go.Context
	nullvalue      *v8go.Value
}

func (c *Context) Close() {
	c.locker.Lock()
	defer c.locker.Unlock()
	if c.Raw == nil {
		return
	}
	ctx := c.Raw
	c.Raw = nil
	c.nullvalue = nil
	c.objectTemplate = nil
	ctx.Close()
	ctx.Isolate().Dispose()
}
func (c *Context) isNullValue(v *JsValue) bool {
	return v == nil || v.raw == nil || v.raw == c.nullvalue
}
func (c *Context) RunIdleTasks(nowait bool, idleTimeInSeconds float64) {
	iso := c.Raw.Isolate()
	if v8go.PumpMessageLoop(iso, nowait) {
		v8go.RunIdleTasks(c.Raw.Isolate(), idleTimeInSeconds)
	}
}
func (c *Context) NewLocal() *Local {
	return &Local{ctx: c, values: []*JsValue{}}
}

func (c *Context) NewFunctionTemplate(callback FunctionCallback) *FunctionTemplate {
	return newFunctionTemplate(c, callback)
}

func (c *Context) NewFunction(callback FunctionCallback) *v8go.Value {
	tmpl := c.NewFunctionTemplate(callback)
	fn := tmpl.GetFunction(c)
	return fn
}
