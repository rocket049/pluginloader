package main

import (
	"testing"
	"time"

	"github.com/rocket049/pluginloader"
)

func TestPlugin(t *testing.T) {
	p, err := pluginloader.NewPluginLoader("foo.so")
	if err != nil {
		t.Fatal(err)
	}

	f := GetfooFuncs(p)

	t.Log("call func Help:")
	res, err := f.Help("help friend")
	t.Log(res, err)

	t.Log("call func GetTime:")
	tm := f.GetTime()
	t.Log(tm)

	t.Log("call func Bar:")
	f.Bar()

	iface, err := p.Call("NewMan")
	if err != nil {
		t.Fatal(err)
	}
	man := iface.(IMan)
	t.Log("call man.Hello:")
	man.Hello()

	iface, err = p.Call("NewFoo")
	if err != nil {
		t.Fatal(err)
	}
	foo := iface.(IFoo)
	t.Log("call foo.Say:")
	foo.Say("hello")
	foo.Set(nil)

	iface, err = p.Call("NewLittle")
	if err != nil {
		t.Fatal(err)
	}
	lit := iface.(Ilittle)
	t.Log("call little.Hello:")
	lit.Hello()

	if f.GetRenterer(nil) != nil {
		t.Fatal("err GetRenterer")
	}

	tm1 := f.GetTime()
	if tm1.Unix() < time.Now().Unix() {
		t.Fatal("err GetTime")
	}

	vi := f.GetArray()
	for i := 0; i < 3; i++ {
		if vi[i] != i+1 {
			t.Fatal("err GetArray")
		}
	}
}

func BenchmarkWrap(b *testing.B) {
	p, err := pluginloader.NewPluginLoader("foo.so")
	if err != nil {
		b.Fatal(err)
	}

	f := GetfooFuncs(p)
	for i := 0; i < b.N; i++ {
		r1, r2 := f.SwapInt(i, i+10)
		b.Log(r1, r2)
	}
}

func BenchmarkCall(b *testing.B) {
	p, err := pluginloader.NewPluginLoader("foo.so")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		res, err := p.CallValue("SwapInt", i, i+10)
		b.Log(res[0].Int(), res[1].Int(), err)
	}
}
