package main

import (
	"fmt"
	"os"

	"github.com/rocket049/pluginloader"
)

func main() {
	if len(os.Args) != 2 {
		panic("main:Must have a parameter.")
	}
	p, err := pluginloader.NewPluginLoader(os.Args[1])
	if err != nil {
		panic(err)
	}

	res, err := p.Call("Hello", "A", "B", "C")
	fmt.Println(res, err)

	res, err = p.Call("Say", "Let's run.")
	fmt.Println(res, err)
}
