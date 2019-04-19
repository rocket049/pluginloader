package main

import (
	"fmt"
	"os"
	"time"

	md "github.com/russross/blackfriday"
)

type Foo struct {
	Id int
}

//Say say something
func (s *Foo) Say(s1 string) (string, error) {
	fmt.Println(time.Now())
	return s1, nil
}

func (s *Foo) Hi(s1 string) *Foo {
	fmt.Println(time.Now())
	fmt.Println(s1, md.EXTENSION_FENCED_CODE)
	return s
}

func (s *Foo) Set(p *os.File) bool {
	fmt.Println("Foo.Set", p, time.Now())
	return true
}

func (s *Foo) Swap(p *Foo) {
	fmt.Println("Foo.Swap", time.Now())
}

func NewFoo() *Foo {
	return &Foo{100}
}

func GoFoo(p *Foo) error {
	fmt.Println("GoFoo", time.Now())
	return nil
}

func Hello(m []*Man) {
	fmt.Println(time.Now())
}
func Hello2(m ...string) {
	fmt.Println(time.Now())
}

func SwapInt(a, b int) (x, y int) {
	return b, a
}

func GetTime() time.Time {
	return time.Now()
}

func GetRenterer(m map[string]*os.File) map[string]md.Renderer {
	return nil
}

type IDX int

func GetRentererBad() map[IDX]md.Renderer {
	return nil
}

func GetArray() []int {
	return []int{1, 2, 3}
}
