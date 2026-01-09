package v8local

import (
	"fmt"
	"math/big"
	"runtime"

	"github.com/herb-go/v8go"
)

type Local struct {
	ctx    *Context
	values []*JsValue
}

func (l *Local) Close() {
	for _, v := range l.values {
		v.Release()
	}
}
func (l *Local) Context() *Context {
	return l.ctx
}
func (l *Local) NewLocal() *Local {
	return l.ctx.NewLocal()
}
func (l *Local) NewFunction(callback FunctionCallback) *JsValue {
	tmpl := l.ctx.NewFunctionTemplate(callback)
	fn := tmpl.GetFunction(l.ctx)
	return l.manage(fn, true)
}
func (l *Local) manage(v *v8go.Value, managed bool) *JsValue {
	val := &JsValue{
		raw:      v,
		exported: !managed,
		local:    l,
	}
	l.values = append(l.values, val)
	return val
}

// Import imports an existing *v8go.Value into the Local context.
// The imported value is marked as exported, so it won't be released when the Local context is closed.
func (l *Local) Import(v *v8go.Value) *JsValue {
	val := &JsValue{
		raw:      v,
		local:    l,
		exported: true,
	}
	return val
}

func (l *Local) Global() *JsValue {
	result := l.manage(l.ctx.Raw.Global().Value, false)
	return result
}

func (l *Local) newValue(v interface{}) *JsValue {
	if l == nil || l.ctx == nil || l.ctx.Raw == nil {
		return nil
	}
	val, err := v8go.NewValue(l.ctx.Raw.Isolate(), v)
	runtime.KeepAlive(v)
	if err != nil {
		panic(err)
	}
	return l.manage(val, true)
}
func (l *Local) NewString(val string) *JsValue {
	return l.newValue(val)
}
func (l *Local) NewInt32(val int32) *JsValue {
	return l.newValue(val)
}
func (l *Local) NewInt64(val int64) *JsValue {
	return l.newValue(val)
}
func (l *Local) NewBoolean(val bool) *JsValue {
	return l.newValue(val)
}
func (l *Local) NewBigInt(val *big.Int) *JsValue {
	return l.newValue(val)
}
func (l *Local) NewNumber(val float64) *JsValue {
	return l.newValue(val)
}

func (l *Local) NewStringArray(values ...string) *JsValue {
	args := make([]*JsValue, len(values))
	for i, v := range values {
		args[i] = l.NewString(v)
	}
	return l.NewArray(args...)
}
func (l *Local) NewArray(values ...*JsValue) *JsValue {
	a := l.RunScript("Array", "array")
	result := a.Call(a, values...)
	return result
}
func (l *Local) NewObject() *JsValue {
	obj, err := l.ctx.objectTemplate.NewInstance(l.ctx.Raw)
	if err != nil {
		panic(err)
	}
	result := l.manage(obj.Value, true) //?
	return result
}
func (l *Local) NewArrayBuffer(data []byte) *JsValue {
	v := l.RunScript(fmt.Sprintf("new ArrayBuffer(%d)", len(data)), "arraybuffer.js")
	v8go.WriteToArrayBuffer(v.raw, data)
	return v
}
func (l *Local) RunScript(script string, name string) *JsValue {
	result, err := l.ctx.Raw.RunScript(script, name)
	if err != nil {
		panic(err)
	}
	return l.manage(result, true)
}
func (l *Local) NullValue() *JsValue {
	return l.manage(l.ctx.nullvalue, true)
}
