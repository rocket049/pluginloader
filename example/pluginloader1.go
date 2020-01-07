package main

import (
	"flag"

	"github.com/rocket049/pluginloader"
)

func main() {
	flag.Parse()
	p, err := pluginloader.NewPluginLoader(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	a, err := p.Call("NewObjType", "Boy")
	if err != nil {
		panic(err)
	}
	ua := pluginloader.NewUnknownObjectFromInterface(a)
	ua.Call("Say", "Hello friends!")
	ua.Call("SaySuccess")
}
