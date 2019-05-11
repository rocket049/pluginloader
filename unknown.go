package pluginloader

import (
	"encoding/json"
	"errors"
	"reflect"
)

//UnknownObject field 'V' MUST be valueof object type *struct{...}
//成员'V' 必须是结构体指针的 Value: *struct{...}
type UnknownObject struct {
	V reflect.Value
}

//NewUnknownObject parameter 'v' MUST be valueof object type *struct{...}, or it will return nil
//参数'v' 必须是结构体指针的 Value: *struct{...}， 否则返回 nil
func NewUnknownObject(v reflect.Value) *UnknownObject {
	if v.Type().Kind() != reflect.Ptr {
		return nil
	}
	if v.Type().Elem().Kind() != reflect.Struct {
		return nil
	}
	return &UnknownObject{v}
}

//Get 得到结构体成员的 Value
//get the value of a field
func (s *UnknownObject) Get(name string) reflect.Value {
	return s.V.Elem().FieldByName(name)
}

//Call 运行结构体的 method
//call the method of the struct
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

//Json 把结构体编码为 JSON。 convert the struct to JSON. if error,return nil.
func (s *UnknownObject) Json() []byte {
	res, err := json.Marshal(s.V.Interface())
	if err != nil {
		return nil
	}
	return res
}

//CopyToStruct 利用 reflect 技术把结构体的可 export 值复制到 v 中，v 必须是相似结构体的指针。 copy the exported value of a struct to v through gob encoding.
func (s *UnknownObject) CopyToStruct(v interface{}) error {
	return structCopy(s.V.Interface(), v)
}

func structCopy(src, dst interface{}) error {
	srcV, err := srcFilter(src)
	if err != nil {
		return err
	}
	dstV, err := dstFilter(dst)
	if err != nil {
		return err
	}
	srcKeys := make(map[string]bool)
	for i := 0; i < srcV.NumField(); i++ {
		srcKeys[srcV.Type().Field(i).Name] = true
	}
	for i := 0; i < dstV.Elem().NumField(); i++ {
		fName := dstV.Elem().Type().Field(i).Name
		if _, ok := srcKeys[fName]; ok {
			v := srcV.FieldByName(dstV.Elem().Type().Field(i).Name)
			if v.CanInterface() {
				dstV.Elem().Field(i).Set(v)
			}
		}
	}

	return nil
}

func srcFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not a struct or a pointer to struct")
	}
	return v, nil
}

func dstFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() != reflect.Ptr {
		return reflect.Zero(v.Type()), errors.New("src type error: not a pointer to struct")
	}
	if v.Elem().Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not point to struct")
	}
	return v, nil
}
