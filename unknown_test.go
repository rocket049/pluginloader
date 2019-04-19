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

type testStruct struct{}
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
