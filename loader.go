package pluginloader

import (
	"errors"
	"plugin"
	"reflect"
)

type PluginLoader struct {
	P *plugin.Plugin
}

//CallValue Allow any number of return values,return type: []reflect.Value,error
func (p *PluginLoader) CallValue(funcName string, p0 ...interface{}) ([]reflect.Value, error) {
	f0, err := p.P.Lookup(funcName)
	if err != nil {
		return nil, err
	}
	f1 := reflect.ValueOf(f0)
	pnum := len(p0)
	param := make([]reflect.Value, pnum)
	for i := 0; i < pnum; i++ {
		param[i] = reflect.ValueOf(p0[i])
	}
	return f1.Call(param), nil
}

//Call return type must be: (res,error)
func (p *PluginLoader) Call(funcName string, p0 ...interface{}) (interface{}, error) {
	f0, err := p.P.Lookup(funcName)
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
	} else if len(res) == 2 {
		if res[1].Interface() != nil {
			return res[0].Interface(), res[1].Interface().(error)
		} else {
			return res[0].Interface(), nil
		}
	} else if len(res) == 1 {
		return res[0].Interface(), nil
	} else {
		return nil, errors.New(funcName + ": Return value format error.")
	}
}

//NewPluginLoader create a loader
func NewPluginLoader(pathName string) (*PluginLoader, error) {
	plug, err := plugin.Open(pathName)
	if err != nil {
		return nil, err
	}
	return &PluginLoader{P: plug}, nil
}

//MakeFunc point a func ptr to plugin
func (s *PluginLoader) MakeFunc(fptr interface{}, name string) error {
	f, err := s.P.Lookup(name)
	if err != nil {
		return err
	}
	vf := reflect.ValueOf(f)
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), vf.Call)
	fn.Set(v)
	return nil
}
