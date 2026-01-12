package httpv8

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/plugins/addons/httpaddon"
	"github.com/herb-go/v8go"
	"github.com/herb-go/v8local/v8plugin"
)

func TestAddon(t *testing.T) {
	app := &http.ServeMux{}
	s := httptest.NewServer(app)
	defer s.Close()
	u, err := url.Parse(s.URL)
	if err != nil {
		panic(err)
	}
	opt := herbplugin.NewOptions()
	opt.Permissions = append(opt.Permissions, httpaddon.Permission)
	opt.Trusted.Domains = append(opt.Trusted.Domains, u.Host)
	opt.GetLocation().Path = "."
	i := v8plugin.NewInitializer()
	i.Entry = "test.js"
	var addon *Addon
	module := herbplugin.CreateModule(
		"test",
		func(ctx context.Context, p herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			plugin := p.(*v8plugin.Plugin)
			addon = Create(p)
			local := plugin.Runtime.NewLocal()
			defer local.Close()
			local.Global().Set("HTTP", addon.Convert(local))
			next(ctx, p)
		},
		func(ctx context.Context, p herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			next(ctx, p)
		},
		func(ctx context.Context, p herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			next(ctx, p)
		},
	)
	i.Modules = append(i.Modules, module)
	p := v8plugin.MustCreatePlugin(i)
	herbplugin.Lanuch(p, opt)
	local := p.Runtime.NewLocal()
	test := local.Global().Get("test")
	result := test.Call(test, local.NewString(s.URL))
	if result.Local() != local {
		t.Fatal("Invalid return value")
	}
	if addon.getRequestCount() == 0 {
		t.Fatal("Request cleaned up")
	}
	local.Close()
	p.Top.Close()
	v8go.ForceV8GC(p.Runtime.Raw.Isolate())
	p.Runtime.RunIdleTasks(false, 1.0)
	p.Runtime.RunIdleTasks(false, 1.0)
	if addon.getRequestCount() != 0 {
		t.Fatal("Request not cleaned up")
	}
}
