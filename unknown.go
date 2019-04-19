package pluginloader

import (
	"reflect"
)

type UnknownObject struct {
	V reflect.Value
}

func NewUnknownObject(v reflect.Value) *UnknownObject {
	return &UnknownObject{v}
}

func (s *UnknownObject) Get(name string) reflect.Value {
	return s.V.Elem().FieldByName(name)
}

func (s *UnknownObject) Call(fn string, args ...interface{}) []reflect.Value {
	f := s.V.MethodByName(fn)
	argn := len(args)
	argv := make([]reflect.Value, argn)
	for i := 0; i < argn; i++ {
		if args[i] == nil {
			argv[i] = reflect.Zero(f.Type().In(i))
		} else {
			argv[i] = reflect.ValueOf(args[i])
		}
	}
	return f.Call(argv)
}
