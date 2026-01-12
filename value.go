package v8local

import (
	"math/big"

	"github.com/herb-go/v8go"
)

type JsValue struct {
	raw      *v8go.Value
	local    *Local
	exported bool
}

func (j *JsValue) Local() *Local {
	return j.local
}

// Export marks the JsValue as exported and returns the underlying *v8go.Value.
// You must release the returned *v8go.Value manually when you no longer need it.
func (j *JsValue) Export() *v8go.Value {
	j.exported = true
	return j.raw
}

// Mark the JsValue as exported and prevent it from being released automatically.
func (j *JsValue) AsExported() *JsValue {
	j.exported = true
	return j
}
func (j *JsValue) Release() bool {
	if !j.exported && !j.local.ctx.isNullValue(j) {
		j.raw.Release()
		return true
	}
	return false
}
func mustAsObject(v *v8go.Value) *v8go.Object {
	o, err := v.AsObject()
	if err != nil {
		panic(err)
	}
	return o
}
func (v *JsValue) export() *v8go.Value {
	if v == nil {
		return nil
	}
	return v.raw
}
func (v *JsValue) Call(recvr *JsValue, args ...*JsValue) *JsValue {
	if v.raw == nil {
		return nil
	}
	fn, err := v.export().AsFunction()
	if err != nil {
		panic(err)
	}
	fnargs := make([]v8go.Valuer, len(args))
	for i, val := range args {
		fnargs[i] = val.export()
	}
	val, err := fn.Call(recvr.export(), fnargs...)
	if err != nil {
		panic(err)
	}
	result := v.local.manage(val)
	return result
}

func (v *JsValue) Array() []*JsValue {
	result := []*JsValue{}
	length := v.Get("length")
	if length.IsNullOrUndefined() {
		return result
	}
	ln := int(length.Integer())
	for i := 0; i < ln; i++ {
		item := v.GetIdx(uint32(i))
		if item.IsNullOrUndefined() {
			continue
		}
		result = append(result, item)
	}
	return result
}
func (v *JsValue) StringArrry() []string {
	arr := v.Array()
	result := make([]string, len(arr))
	for i, item := range arr {
		result[i] = item.String()
	}
	return result
}
func (v *JsValue) String() string {
	result := v.export().String()
	return result
}

func (v *JsValue) BigInt() *big.Int {
	result := v.export().BigInt()
	return result
}
func (v *JsValue) Boolean() bool {
	result := v.export().Boolean()
	return result
}
func (v *JsValue) Int32() int32 {
	result := v.export().Int32()
	return result
}
func (v *JsValue) Integer() int64 {
	result := v.export().Integer()
	return result
}
func (v *JsValue) Number() float64 {
	result := v.export().Number()
	return result
}
func (v *JsValue) Uint32() uint32 {
	result := v.export().Uint32()
	return result
}
func (v *JsValue) ArrayBufferContent() []byte {
	result := v8go.ArrayBufferContent(v.export())
	return result
}

func (v *JsValue) SameValue(other *JsValue) bool {
	result := v.export().SameValue(other.raw)
	return result
}
func (v *JsValue) IsUndefined() bool {
	result := v.export().IsUndefined()
	return result
}
func (v *JsValue) IsNull() bool {
	result := v.export().IsNull()
	return result
}
func (v *JsValue) IsNullOrUndefined() bool {
	result := v.export().IsNullOrUndefined()
	return result
}

func (v *JsValue) IsTrue() bool {
	result := v.export().IsTrue()
	return result
}
func (v *JsValue) IsFalse() bool {
	result := v.export().IsFalse()
	return result
}

func (v *JsValue) IsFunction() bool {
	result := v.export().IsFunction()
	return result
}

func (v *JsValue) IsObject() bool {
	result := v.export().IsObject()
	return result
}

func (v *JsValue) IsBigInt() bool {
	result := v.export().IsBigInt()
	return result
}
func (v *JsValue) IsBoolean() bool {
	result := v.export().IsBoolean()
	return result
}

func (v *JsValue) IsNumber() bool {
	result := v.export().IsNumber()
	return result
}
func (v *JsValue) IsInt32() bool {
	result := v.export().IsInt32()
	return result
}
func (v *JsValue) IsUint32() bool {
	result := v.export().IsUint32()
	return result
}
func (v *JsValue) IsDate() bool {
	result := v.export().IsDate()
	return result
}
func (v *JsValue) IsNativeError() bool {
	result := v.export().IsNativeError()
	return result
}
func (v *JsValue) IsRegExp() bool {
	result := v.export().IsRegExp()
	return result
}
func (v *JsValue) IsMap() bool {
	result := v.export().IsMap()
	return result
}
func (v *JsValue) IsSet() bool {
	result := v.export().IsSet()
	return result
}
func (v *JsValue) IsArray() bool {
	result := v.export().IsArray()
	return result
}
func (v *JsValue) IsArrayBuffer() bool {
	result := v.export().IsArrayBuffer()
	return result
}
func (v *JsValue) MustMarshalJSON() []byte {
	data, err := v.export().MarshalJSON()
	if err != nil {
		panic(err)
	}
	return data
}

func (v *JsValue) MethodCall(methodName string, args ...*JsValue) *JsValue {
	fn := v.Get(methodName) // ensure method exists
	result := fn.Call(v, args...)
	return result
}
func (v *JsValue) SetObjectMethod(name string, fn FunctionCallback) {
	f := v.local.ctx.NewFunctionTemplate(fn).GetFunction(v.local.ctx)
	v.Set(name, v.local.manage(f))
}
func (v *JsValue) Get(key string) *JsValue {
	val, err := mustAsObject(v.export()).Get(key)
	if err != nil {
		panic(err)
	}
	result := v.local.manage(val)
	return result
}
func (v *JsValue) GetIdx(idx uint32) *JsValue {
	val, err := mustAsObject(v.export()).GetIdx(idx)
	if err != nil {
		panic(err)
	}
	result := v.local.manage(val)
	return result

}

func (v *JsValue) Set(key string, val *JsValue) {
	err := mustAsObject(v.export()).Set(key, val.export())
	if err != nil {
		panic(err)
	}
}

func (v *JsValue) SetIdx(idx uint32, val *JsValue) {
	err := mustAsObject(v.export()).SetIdx(idx, val.export())
	if err != nil {
		panic(err)
	}
}
func (v *JsValue) Has(key string) bool {
	result := mustAsObject(v.export()).Has(key)
	return result
}
func (v *JsValue) HasIdx(idx uint32) bool {
	result := mustAsObject(v.export()).HasIdx(idx)
	return result
}

func (v *JsValue) Delete(key string) bool {
	result := mustAsObject(v.export()).Delete(key)
	return result
}

func (v *JsValue) DeleteIdx(idx uint32) bool {
	result := mustAsObject(v.export()).DeleteIdx(idx)
	return result
}
