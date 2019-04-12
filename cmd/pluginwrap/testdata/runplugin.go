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
	//p.Call("Help", "good boy")
	res, _ := p.CallValue("Help", "good boy")
	fmt.Println(res[0].String(), res[1].IsNil())

	fmt.Println("call func Bar:")
	p.Call("Bar")

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
	fmt.Println("call foo.Hi:")
	foo.Hi("bbbb")

	iface, err = p.Call("NewLittle")
	if err != nil {
		panic(err)
	}
	lit := iface.(Ilittle)
	fmt.Println("call little.Hello:")
	lit.Hello()
}
