package httpv8

import (
	"net/url"
	"strconv"
	"sync"
	"unsafe"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/plugins/addons/httpaddon"
	v8local "github.com/herb-go/v8local"
)

type Builder func(r *v8local.Local, a *Addon, req *Request) *v8local.JsValue

var DefaultBuilder = func(r *v8local.Local, a *Addon, req *Request) *v8local.JsValue {
	obj := r.NewObject()
	obj.Set("GetID", a.Functions["GetID"])
	obj.Set("GetURL", a.Functions["GetURL"])
	obj.Set("SetURL", a.Functions["SetURL"])
	obj.Set("GetProxy", a.Functions["GetProxy"])
	obj.Set("SetProxy", a.Functions["SetProxy"])
	obj.Set("GetMethod", a.Functions["GetMethod"])
	obj.Set("SetMethod", a.Functions["SetMethod"])
	obj.Set("GetBody", a.Functions["GetBody"])
	obj.Set("GetBodyArrayBuffer", a.Functions["GetBodyArrayBuffer"])
	obj.Set("SetBody", a.Functions["SetBody"])
	obj.Set("FinishedAt", a.Functions["FinishedAt"])
	obj.Set("ExecuteStatus", a.Functions["ExecuteStatus"])
	obj.Set("ResetHeader", a.Functions["ResetHeader"])
	obj.Set("SetHeader", a.Functions["SetHeader"])
	obj.Set("AddHeader", a.Functions["AddHeader"])
	obj.Set("DelHeader", a.Functions["DelHeader"])
	obj.Set("GetHeader", a.Functions["GetHeader"])
	obj.Set("HeaderValues", a.Functions["HeaderValues"])
	obj.Set("HeaderFields", a.Functions["HeaderFields"])
	obj.Set("ResponseStatusCode", a.Functions["ResponseStatusCode"])
	obj.Set("ResponseBody", a.Functions["ResponseBody"])
	obj.Set("ResponseBodyArrayBuffer", a.Functions["ResponseBodyArrayBuffer"])
	obj.Set("ResponseHeader", a.Functions["ResponseHeader"])
	obj.Set("ResponseHeaderValues", a.Functions["ResponseHeaderValues"])
	obj.Set("ResponseHeaderFields", a.Functions["ResponseHeaderFields"])
	obj.Set("Execute", a.Functions["Execute"])
	return obj
}

type Request struct {
	RID     string
	Request *httpaddon.Request
}

func RequestGetID(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewString(req.Request.GetID())
	}
}

func RequestGetURL(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewString(req.Request.GetURL())
	}
}
func RequestSetURL(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.SetURL(call.GetArg(0).String())
		return nil
	}
}
func RequestGetProxy(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewString(req.Request.GetProxy())
	}
}
func RequestSetProxy(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.SetProxy(call.GetArg(0).String())
		return nil
	}
}

func RequestGetMethod(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewString(req.Request.GetMethod())
	}
}
func RequestSetMethod(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.SetMethod(call.GetArg(0).String())
		return nil
	}
}
func RequestGetBody(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewString(string(req.Request.GetBody()))
	}
}
func RequestGetBodyArrayBuffer(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewArrayBuffer(req.Request.GetBody())
	}
}

func RequestSetBody(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.SetBody([]byte(call.GetArg(0).String()))
		return nil
	}
}

func RequestFinishedAt(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewInt64(req.Request.FinishedAt())
	}
}
func RequestExecuteStatus(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewInt32(int32(req.Request.ExecuteStauts()))
	}
}
func RequestResetHeader(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.ResetHeader()
		return nil
	}
}
func RequestSetHeader(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.SetHeader(call.GetArg(0).String(), call.GetArg(1).String())
		return nil
	}
}
func RequestAddHeader(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.AddHeader(call.GetArg(0).String(), call.GetArg(1).String())
		return nil
	}
}
func RequestDelHeader(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.DelHeader(call.GetArg(0).String())
		return nil
	}

}
func RequestGetHeader(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewString(req.Request.GetHeader(call.GetArg(0).String()))
	}
}
func RequestHeaderValues(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		result := req.Request.HeaderValues(call.GetArg(0).String())
		var output = make([]*v8local.JsValue, len(result))
		for i, v := range result {
			output[i] = call.Local().NewString(v)
		}
		return call.Local().NewArray(output...)
	}
}
func RequestHeaderFields(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		result := req.Request.HeaderFields()
		var output = make([]*v8local.JsValue, len(result))
		for i, v := range result {
			output[i] = call.Local().NewString(v)
		}
		return call.Local().NewArray(output...)
	}
}

func RequestResponseStatusCode(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewInt32(int32(req.Request.ResponseStatusCode()))
	}
}
func RequestResponseBody(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewString(string(req.Request.ResponseBody()))
	}
}
func RequestResponseBodyArrayBuffer(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewArrayBuffer(req.Request.ResponseBody())
	}
}
func RequestResponseHeader(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		return call.Local().NewString(req.Request.ResponseHeader(call.GetArg(0).String()))
	}
}
func RequestResponseHeaderValues(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		result := req.Request.ResponseHeaderValues(call.GetArg(0).String())

		var output = make([]*v8local.JsValue, len(result))
		for i, v := range result {
			output[i] = call.Local().NewString(v)
		}
		return call.Local().NewArray(output...)
	}
}
func RequestResponseHeaderFields(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		result := req.Request.ResponseHeaderFields()

		var output = make([]*v8local.JsValue, len(result))
		for i, v := range result {
			output[i] = call.Local().NewString(v)
		}
		return call.Local().NewArray(output...)
	}
}
func RequestExecute(a *Addon) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		req := a.LoadReq(call.This().Get("id").String())
		req.Request.MustExecute()
		return nil
	}
}

type Addon struct {
	Addon     *httpaddon.Addon
	Builder   Builder
	Functions map[string]*v8local.JsValue
	reqs      sync.Map
}

func (a *Addon) ParseURL(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	rawurl := call.Args()[0].String()
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil
	}
	result := call.Local().NewObject()
	result.Set("Host", call.Local().NewString(u.Host))
	result.Set("Hostname", call.Local().NewString(u.Host))
	result.Set("Scheme", call.Local().NewString(u.Scheme))
	result.Set("Path", call.Local().NewString(u.Path))
	result.Set("Query", call.Local().NewString(u.RawQuery))
	result.Set("User", call.Local().NewString(u.User.Username()))
	p, _ := u.User.Password()
	result.Set("Password", call.Local().NewString(p))
	result.Set("Port", call.Local().NewString(u.Port()))
	result.Set("Fragment", call.Local().NewString(u.Fragment))
	return result
}
func (a *Addon) NewRequest(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	method := call.GetArg(0).String()
	url := call.GetArg(1).String()
	req := a.Addon.Create(method, url)
	rid := strconv.FormatInt(int64(uintptr(unsafe.Pointer(req))), 16)
	ar := &Request{
		RID:     rid,
		Request: req,
	}
	a.reqs.Store(rid, ar)
	fr := call.This().Get("FinalizationRegistry")
	obj := a.Builder(call.Local(), a, ar)
	obj.Set("id", call.Local().NewString(rid))
	fn := fr.Get("register")
	fn.Call(fr, obj, a.Register(call.Local(), call.This(), req.ID))
	return obj
}
func (a *Addon) Register(r *v8local.Local, addonobj *v8local.JsValue, id string) *v8local.JsValue {
	obj := r.NewObject()
	upload := addonobj.Get("unload")
	obj.Set("unload", upload)
	obj.Set("id", r.NewString(id))
	return obj
}
func (a *Addon) unload(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	a.reqs.Delete(call.GetArg(0).String())
	return nil
}
func (a *Addon) size(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	count := 0
	a.reqs.Range(func(key, value interface{}) bool {
		count++
		return true // continue iteration
	})

	return call.Local().NewInt32(int32(count))
}
func (a *Addon) LoadReq(id string) *Request {
	v, ok := a.reqs.Load(id)
	if !ok {
		panic("v8 http request id " + id + " not found")
	}
	return v.(*Request)
}
func (a *Addon) Convert(r *v8local.Local) *v8local.JsValue {
	obj := r.NewObject()
	obj.SetObjectMethod("New", a.NewRequest)
	obj.SetObjectMethod("ParseURL", a.ParseURL)
	fr := r.RunScript(" new FinalizationRegistry((reg) => {reg.unload(reg.id)})", "FinalizationRegistry")
	obj.Set("FinalizationRegistry", fr)
	obj.Set("unload", r.NewFunction(a.unload))
	obj.SetObjectMethod("Size", a.size)
	a.Functions["GetID"] = r.NewFunction(RequestGetID(a)).AsExported()
	a.Functions["GetURL"] = r.NewFunction(RequestGetURL(a)).AsExported()
	a.Functions["SetURL"] = r.NewFunction(RequestSetURL(a)).AsExported()
	a.Functions["GetProxy"] = r.NewFunction(RequestGetProxy(a)).AsExported()
	a.Functions["SetProxy"] = r.NewFunction(RequestSetProxy(a)).AsExported()
	a.Functions["GetMethod"] = r.NewFunction(RequestGetMethod(a)).AsExported()
	a.Functions["SetMethod"] = r.NewFunction(RequestSetMethod(a)).AsExported()
	a.Functions["GetBody"] = r.NewFunction(RequestGetBody(a)).AsExported()
	a.Functions["GetBodyArrayBuffer"] = r.NewFunction(RequestGetBodyArrayBuffer(a)).AsExported()
	a.Functions["SetBody"] = r.NewFunction(RequestSetBody(a)).AsExported()
	a.Functions["FinishedAt"] = r.NewFunction(RequestFinishedAt(a)).AsExported()
	a.Functions["ExecuteStatus"] = r.NewFunction(RequestExecuteStatus(a)).AsExported()
	a.Functions["ResetHeader"] = r.NewFunction(RequestResetHeader(a)).AsExported()
	a.Functions["SetHeader"] = r.NewFunction(RequestSetHeader(a)).AsExported()
	a.Functions["AddHeader"] = r.NewFunction(RequestAddHeader(a)).AsExported()
	a.Functions["DelHeader"] = r.NewFunction(RequestDelHeader(a)).AsExported()
	a.Functions["GetHeader"] = r.NewFunction(RequestGetHeader(a)).AsExported()
	a.Functions["HeaderValues"] = r.NewFunction(RequestHeaderValues(a)).AsExported()
	a.Functions["HeaderFields"] = r.NewFunction(RequestHeaderFields(a)).AsExported()
	a.Functions["ResponseStatusCode"] = r.NewFunction(RequestResponseStatusCode(a)).AsExported()
	a.Functions["ResponseBody"] = r.NewFunction(RequestResponseBody(a)).AsExported()
	a.Functions["ResponseBodyArrayBuffer"] = r.NewFunction(RequestResponseBodyArrayBuffer(a)).AsExported()
	a.Functions["ResponseHeader"] = r.NewFunction(RequestResponseHeader(a)).AsExported()
	a.Functions["ResponseHeaderValues"] = r.NewFunction(RequestResponseHeaderValues(a)).AsExported()
	a.Functions["ResponseHeaderFields"] = r.NewFunction(RequestResponseHeaderFields(a)).AsExported()
	a.Functions["Execute"] = r.NewFunction(RequestExecute(a)).AsExported()

	return obj
}
func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		Addon:     httpaddon.Create(p),
		Functions: make(map[string]*v8local.JsValue),
		Builder:   DefaultBuilder,
	}
}
