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

	InitfooFuncs(p)

	t.Log("call func Help:")
	res, err := Help("help friend")
	t.Log(res, err)

	t.Log("call func GetTime:")
	tm := GetTime()
	t.Log(tm)

	t.Log("call func Bar:")
	Bar()

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
	foo.Say("aaa")

	iface, err = p.Call("NewLittle")
	if err != nil {
		t.Fatal(err)
	}
	lit := iface.(Ilittle)
	t.Log("call little.Hello:")
	lit.Hello()

	if GetRenterer() != nil {
		t.Fatal("err GetRenterer")
	}

	tm1 := GetTime()
	if tm1.Unix() < time.Now().Unix() {
		t.Fatal("err GetTime")
	}
}

func BenchmarkWrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1, r2 := SwapInt(i, i+10)
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
