package v8plugin

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/herb-go/v8local"

	"github.com/herb-go/herbplugin"
)

const PluginType = "js"
const DefaultNamespace = "system"

func New() *Plugin {
	return &Plugin{
		Plugin: herbplugin.New(),
	}
}

type Plugin struct {
	sync.RWMutex
	entry   string
	Runtime *v8local.Context
	Top     *v8local.Local //Top Level Local
	herbplugin.Plugin
	DisableBuiltin bool
	startCommand   string
	modules        []*herbplugin.Module
	namespace      string
	Builtin        map[string]*v8local.JsValue
}

func (p *Plugin) PluginType() string {
	return PluginType
}
func (p *Plugin) MustInitPlugin() {
	p.Plugin.MustInitPlugin()
	p.Builtin = map[string]*v8local.JsValue{}
	var processs = make([]herbplugin.Process, 0, len(p.modules))
	for k := range p.modules {
		if p.modules[k].InitProcess != nil {
			processs = append(processs, p.modules[k].InitProcess)
		}
	}
	builtin := p.Top.NewObject()
	for key, fn := range p.Builtin {
		builtin.Set(key, fn)
	}
	herbplugin.Exec(p, processs...)
	if !p.DisableBuiltin {
		global := p.Top.Global()
		global.Set(p.namespace, builtin)

	}
}
func (p *Plugin) MustLoadPlugin() {
	p.Plugin.MustLoadPlugin()
	if p.entry != "" {
		data, err := os.ReadFile(filepath.Join(p.PluginOptions().GetLocation().Path, p.entry))
		if err != nil {
			panic(err)
		}
		p.Top.RunScript(string(data), p.entry)
	}
}

func (p *Plugin) MustBootPlugin() {
	p.Plugin.MustBootPlugin()
	var processs = make([]herbplugin.Process, 0, len(p.modules))
	for k := range p.modules {
		if p.modules[k].BootProcess != nil {
			processs = append(processs, p.modules[k].BootProcess)
		}
	}
	herbplugin.Exec(p, processs...)
	if p.startCommand != "" {
		p.Top.RunScript(p.startCommand, "")
	}
}

func (p *Plugin) MustClosePlugin() {
	var processs = make([]herbplugin.Process, 0, len(p.modules))
	for i := len(p.modules) - 1; i >= 0; i-- {
		if p.modules[i].CloseProcess != nil {
			processs = append(processs, p.modules[i].CloseProcess)
		}
	}
	herbplugin.Exec(p, processs...)
	p.modules = nil
	p.Builtin = nil
	p.Plugin.MustClosePlugin()
	rt := p.Runtime
	top := p.Top
	p.Runtime = nil
	p.Top = nil
	top.Close()
	rt.Close()

}
func (p *Plugin) LoadJsPlugin() *Plugin {
	return p
}

type JsPluginLoader interface {
	LoadJsPlugin() *Plugin
}
