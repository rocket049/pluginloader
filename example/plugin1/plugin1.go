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

func NewObjType(name string) *ObjType {
	return &ObjType{name}
}
