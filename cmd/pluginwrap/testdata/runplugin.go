package main

import (
	"fmt"

	"github.com/rocket049/pluginloader"
)

func main() {
	test1()
}

func test1() {
	p, err := pluginloader.NewPluginLoader("foo.so")
	if err != nil {
		panic(err)
	}

	fmt.Println("call func Help:")
	Help := FnHelp(p.P)
	res, err := Help("help friend")
	fmt.Println(res, err)

	fmt.Println("call func Bar:")
	Bar := FnBar(p.P)
	Bar()

	iface, err := p.Call("NewMan")
	if err != nil {
		panic(err)
	}
	man := iface.(IMan)
	fmt.Println("call man.Hello:")
	man.Hello()

	iface, err = p.Call("NewFoo")
	if err != nil {
		panic(err)
	}
	foo := iface.(IFoo)
	fmt.Println("call foo.Say:")
	foo.Say("aaa")

	iface, err = p.Call("NewLittle")
	if err != nil {
		panic(err)
	}
	lit := iface.(Ilittle)
	fmt.Println("call little.Hello:")
	lit.Hello()
}
