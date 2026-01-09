package v8local

import "github.com/herb-go/v8go"

type callback struct {
	cb  FunctionCallback
	ctx *Context
}

func (c *callback) call(info *v8go.FunctionCallbackInfo) (output *v8go.Value) {
	defer func() {
		if r := recover(); r != nil {
			errmsg, _ := v8go.NewValue(info.Context().Isolate(), r.(error).Error())
			output = info.Context().Isolate().ThrowException(errmsg)
		}
	}()
	local := c.ctx.NewLocal()
	defer local.Close()
	rawargs := info.Args()
	args := make([]*JsValue, len(rawargs))
	for k, v := range rawargs {
		args[k] = local.manage(v, true)
	}
	this := local.manage(info.This().Value, true)
	fi := NewFunctionCallbackInfo(local, this, args...)
	result := c.cb(fi)
	if result == nil {
		return nil
	}
	return result.Export()
}

type FunctionCallback func(info *FunctionCallbackInfo) *JsValue

func (f FunctionCallback) newCallback(ctx *Context) *callback {
	return &callback{cb: f, ctx: ctx}
}

func NewFunctionCallbackInfo(local *Local, this *JsValue, args ...*JsValue) *FunctionCallbackInfo {
	return &FunctionCallbackInfo{
		local: local,
		this:  this,
		args:  args,
	}
}

type FunctionCallbackInfo struct {
	local *Local
	args  []*JsValue
	this  *JsValue
}

func (i *FunctionCallbackInfo) Local() *Local {
	return i.local
}

// This returns the receiver object "this".
func (i *FunctionCallbackInfo) This() *JsValue {
	return i.this
}

// Args returns a slice of the value arguments that are passed to the JS function.
func (i *FunctionCallbackInfo) Args() []*JsValue {
	return i.args
}
func (i *FunctionCallbackInfo) GetArg(idx int) *JsValue {
	if idx < 0 || idx >= len(i.args) {
		return i.local.NullValue()
	}
	return i.args[idx]
}

type FunctionTemplate struct {
	tmpl *v8go.FunctionTemplate
}

func (t *FunctionTemplate) GetFunction(ctx *Context) *v8go.Value {
	fn := t.tmpl.GetFunction(ctx.Raw)
	return fn.Value
}
func (t *FunctionTemplate) GetLocalFunction(local *Local) *JsValue {
	fn := t.tmpl.GetFunction(local.ctx.Raw)
	return local.manage(fn.Value, true)

}
func newFunctionTemplate(ctx *Context, callback FunctionCallback) *FunctionTemplate {
	tmpl := v8go.NewFunctionTemplate(ctx.Raw.Isolate(), callback.newCallback(ctx).call)
	return &FunctionTemplate{
		tmpl: tmpl,
	}
}
