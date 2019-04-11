package main

import (
	"fmt"
	"time"
)

type Foo struct{}

//Say say something
func (s *Foo) Say(s1 string) {
	fmt.Println(time.Now())
	fmt.Println(s1)
}

func (s *Foo) Hi(s1 string) {
	fmt.Println(time.Now())
	fmt.Println(s1)
}

func NewFoo() *Foo {
	return new(Foo)
}

func hello() {
	fmt.Println(time.Now())
}
