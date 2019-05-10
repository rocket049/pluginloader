package pluginloader

import (
	"os/exec"
	"reflect"
	"testing"
)

func buildFoo() {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "./cmd/pluginwrap/testdata/foo")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

type testStruct struct {
	ID int32
}
type typeNotStruct int

func TestMatchStructPtr(t *testing.T) {
	p1 := new(testStruct)
	p2 := new(typeNotStruct)
	if NewUnknownObject(reflect.ValueOf(p1)) == nil {
		t.Fatal("NewUnknownObject struct ptr")
	}
	if NewUnknownObject(reflect.ValueOf(p2)) != nil {
		t.Fatal("NewUnknownObject Not struct ptr")
	}

	if NewUnknownObject(reflect.ValueOf(*p1)) != nil {
		t.Fatal("NewUnknownObject struct")
	}
}

func TestUnknown(t *testing.T) {
	buildFoo()
	p, err := NewPluginLoader("foo.so")
	if err != nil {
		t.Fatal(err)
	}
	v, err := p.CallValue("NewFoo")
	if err != nil {
		t.Fatal(err)
	}
	obj := NewUnknownObject(v[0])
	if obj.Get("Id").Int() != 100 {
		t.Fatal("get Foo.Id Id != 100")
	}
	ret := obj.Call("Set", nil)
	if ret[0].Bool() != true {
		t.Fatal("call foo.Set")
	}
}

func TestJson(t *testing.T) {
	buildFoo()
	p, err := NewPluginLoader("foo.so")
	if err != nil {
		t.Fatal(err)
	}
	v, err := p.CallValue("NewFoo")
	if err != nil {
		t.Fatal(err)
	}
	obj := NewUnknownObject(v[0])
	if obj == nil {
		t.Fatal("NewUnknownObject struct ptr")
	}
	json1 := obj.Json()
	if json1 == nil {
		t.Fatal("Json error")
	} else {
		t.Log(string(json1))
	}
}

func TestCopy(t *testing.T) {
	buildFoo()
	p, err := NewPluginLoader("foo.so")
	if err != nil {
		t.Fatal(err)
	}
	v, err := p.CallValue("NewFoo")
	if err != nil {
		t.Fatal(err)
	}
	obj := NewUnknownObject(v[0])
	if obj == nil {
		t.Fatal("NewUnknownObject struct ptr")
	}
	pv := &struct {
		Id int
	}{}
	err = obj.CopyToStruct(pv)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%#v\n", *pv)
	}
}
