package main

import (
	"fmt"
	"time"
)

type Foo struct{}

//Say say something
func (s *Foo) Say(s1 string) (string, error) {
	fmt.Println(time.Now())
	fmt.Println(s1)
	return s1, nil
}

func (s *Foo) Hi(s1 string) *Foo {
	fmt.Println(time.Now())
	fmt.Println(s1)
	return s
}

func (s *Foo) Set(p *Foo) {
	fmt.Println("Foo.Set", time.Now())
}

func (s *Foo) Swap(p *Foo) {
	fmt.Println("Foo.Set", time.Now())
}

func NewFoo() *Foo {
	return new(Foo)
}

func GoFoo(p *Foo) error {
	fmt.Println("GoFoo", time.Now())
	return nil
}

func hello() {
	fmt.Println(time.Now())
}
