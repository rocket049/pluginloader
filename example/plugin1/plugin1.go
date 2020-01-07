package main

import (
	"fmt"
)

type ObjType struct {
	name string
}

func (p *ObjType) Say(s string) int {
	fmt.Printf("%s say: %s\n", p.name, s)
	return len(s)
}

func (p *ObjType) SaySuccess() {
	fmt.Printf("%s say: Plugin successful loaded!\n", p.name)
}

func NewObjType(name string) *ObjType {
	return &ObjType{name}
}
