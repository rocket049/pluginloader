package pluginloader

import (
	"plugin"
	"reflect"
)

type PluginLoader struct {
	p *plugin.Plugin
}

//Call return type must be: (res,error)
func (p *PluginLoader) Call(funcName string, p0 ...interface{}) (interface{}, error) {
	f0, err := p.p.Lookup(funcName)
	if err != nil {
		return nil, err
	}
	f1 := reflect.ValueOf(f0)
	pnum := len(p0)
	param := make([]reflect.Value, pnum)
	for i := 0; i < pnum; i++ {
		param[i] = reflect.ValueOf(p0[i])
	}
	res := f1.Call(param)
	if res == nil {
		return nil, nil
	}
	if res[1].Interface() != nil {
		return res[0].Interface(), res[1].Interface().(error)
	} else {
		return res[0].Interface(), nil
	}
}

func NewPluginLoader(pathName string) (*PluginLoader, error) {
	plug, err := plugin.Open(pathName)
	if err != nil {
		return nil, err
	}
	return &PluginLoader{p: plug}, nil
}
