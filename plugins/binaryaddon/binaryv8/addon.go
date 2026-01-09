package binaryv8

import (
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/plugins/addons/binaryaddon"
	v8local "github.com/herb-go/v8local"
)

type Addon struct {
	Addon *binaryaddon.Addon
}

func (a *Addon) Base64Encode(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	data := call.GetArg(0)
	if data != nil {
		return call.Local().NewString(a.Addon.Base64Encode(data.ArrayBufferContent()))
	}
	return nil
}
func (a *Addon) Base64Decode(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	data := call.GetArg(0)
	if data != nil {
		return call.Local().NewArrayBuffer(a.Addon.Base64Decode(data.String()))
	}
	return nil
}
func (a *Addon) Md5Sum(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	data := call.GetArg(0)
	if data != nil {
		return call.Local().NewString(a.Addon.Md5Sum(data.ArrayBufferContent()))
	}
	return nil
}
func (a *Addon) Sha1Sum(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	data := call.GetArg(0)
	if data != nil {
		return call.Local().NewString(a.Addon.Sha1Sum(data.ArrayBufferContent()))
	}
	return nil
}
func (a *Addon) Sha256Sum(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	data := call.GetArg(0)
	if data != nil {
		return call.Local().NewString(a.Addon.Sha256Sum(data.ArrayBufferContent()))
	}
	return nil
}
func (a *Addon) Sha512Sum(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	data := call.GetArg(0)
	if data != nil {
		return call.Local().NewString(a.Addon.Sha512Sum(data.ArrayBufferContent()))
	}
	return nil
}
func (a *Addon) Convert(r *v8local.Local) *v8local.JsValue {
	obj := r.NewObject()
	obj.SetObjectMethod("Base64Encode", a.Base64Encode)
	obj.SetObjectMethod("Base64Decode", a.Base64Decode)
	obj.SetObjectMethod("Md5Sum", a.Md5Sum)
	obj.SetObjectMethod("Sha1Sum", a.Sha1Sum)
	obj.SetObjectMethod("Sha256Sum", a.Sha256Sum)
	obj.SetObjectMethod("Sha512Sum", a.Sha512Sum)
	return obj
}

func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		Addon: binaryaddon.Create(p),
	}
}
