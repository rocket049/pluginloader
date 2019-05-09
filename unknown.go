package pluginloader

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
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

//Json 把结构体编码为 JSON。 convert the struct to JSON. if error,retur nil.
func (s *UnknownObject) Json() []byte {
	res, err := json.Marshal(s.V.Interface())
	if err != nil {
		return nil
	}
	return res
}

//CopyToStruct 利用 gob 编码技术把结构体的值复制到 v 中，v 必须是相似结构体的指针。 copy the value of a struct to v through gob encoding.
func (s *UnknownObject) CopyToStruct(v interface{}) error {
	buf := bytes.NewBufferString("")
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(s.V.Interface())
	if err != nil {
		return err
	}
	reader := bytes.NewReader(buf.Bytes())
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(v)
	return err
}
