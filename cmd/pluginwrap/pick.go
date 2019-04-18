package main

import (
	"fmt"
	"strings"
)

//parseLine parse line to decide it is a struct method or a func
func parseLine(recv, name string, args, res, vtypes []string, doc string) {
	frecv := strings.TrimLeft(recv, "*")
	if len(recv) == 0 {
		pickFunc(name, args, res, vtypes, doc)
	} else {
		pickMethod(frecv, name, args, res, vtypes, doc)
	}
}

//pickMethod get struct name and method from a line,then append to structs
func pickMethod(recv, name string, args, res, vtypes []string, doc string) {
	fargs := strings.Join(args, ",")
	method := fmt.Sprintf("%s(%s)", name, fargs)

	var pick bool = true
	for i := 0; i < len(vtypes); i++ {
		if isUserDefType(vtypes[i]) {
			pick = false
			break
		}
	}

	if pick {
		addStructMethod(recv, method, res, doc)
	}
}

//pickFunc get func from a line,then append to funcs
func pickFunc(name string, args, res, vtypes []string, doc string) {
	fargs := strings.Join(args, ",")
	body := fmt.Sprintf("func(%s)", fargs)

	var pick bool = true
	for i := 0; i < len(vtypes); i++ {
		if isUserDefType(vtypes[i]) {
			pick = false
			break
		}
	}

	if pick {
		addFunc(name, body, res, doc)
		fmt.Println("add func:", name)
	}
}

func convertTypeName(typ string) string {
	if isUserDefType(typ) == false {
		return typ
	} else {
		return "interface{}"
	}
}

func clearType(typ string) string {
	t := strings.TrimLeft(typ, " .[]*")
	fmt.Println("trim:", typ, "->", t)
	return t
}

func isUserDefType(typ string) bool {
	t := clearType(typ)
	_, ok := typs[t]
	if !ok {
		pickImport(t)
	}
	return ok

}

func pickImport(typ string) {
	sp := strings.Split(typ, ".")

	if len(sp) != 2 {
		return
	}

	importsPicked[sp[0]] = imports[sp[0]]
}
