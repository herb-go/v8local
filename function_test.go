package v8local

import (
	"fmt"
	"testing"
)

func TestFuncCall(t *testing.T) {
	ctx := NewContext()
	local := ctx.NewLocal()
	defer ctx.Close()
	fn := local.NewFunction(func(call *FunctionCallbackInfo) *JsValue {
		arg0 := call.GetArg(0)
		fmt.Println(arg0)
		if arg0 != nil {
			return call.Local().NewString("Hello " + arg0.String())
		}
		return call.Local().NewString("Hello World")
	})
	fn.Call(fn, local.NewString("V8"))
	local.Global().Set("testfunc", fn)
	script := `
	testfunc("V8")
	`
	local.RunScript(script, "")
	local.Close()
}
