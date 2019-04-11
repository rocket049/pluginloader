package main

import (
	"fmt"
	"regexp"
	"strings"
)

//parseLine parse line to decide it is a struct method or a func
func parseLine(line string, res []string, doc string) {
	rf := `^func\s+\w+\(.+$`
	rs := `^func\s+\(\w+\s+\**\w+\)\s*\w+\(.+$`
	rfp, err := regexp.Compile(rf)
	if err != nil {
		panic(err)
	}
	rsp, err := regexp.Compile(rs)
	if err != nil {
		panic(err)
	}
	if rfp.MatchString(line) {
		//pickFunc(line, res, doc)
	} else if rsp.MatchString(line) {
		pickMethod(line, res, doc)
	} else {
		fmt.Println("not match:", line)
	}
}

//pickMethod get struct name and method from a line,then append to structs
func pickMethod(line string, res []string, doc string) {
	rs := `func\s+\(\w+\s+\**(\w+)\)`
	rsp, err := regexp.Compile(rs)
	if err != nil {
		panic(err)
	}
	ps := rsp.FindStringSubmatch(line)
	name := ps[1]

	rs = `func\s+\(\w+\s+\**\w+\)\s*(\w+\(.*\))`
	rsp, err = regexp.Compile(rs)
	if err != nil {
		panic(err)
	}
	ps = rsp.FindStringSubmatch(line)
	method := strings.TrimSpace(ps[1])
	//fmt.Printf("%s.%s\n", name, method)
	for i := 0; i < len(res); i++ {
		res[i] = convertTypeName(res[i])
	}
	addStructMethod(name, method, res, doc)
}

//pickFunc get func from a line,then append to funcs
func pickFunc(line string, res []string, doc string) {
	rf := `func[^\)]*\)`
	rfp, err := regexp.Compile(rf)
	if err != nil {
		panic(err)
	}
	ps := rfp.FindString(line)

	n1 := strings.Index(ps, " ")
	n2 := strings.Index(ps, "(")
	name := ps[n1+1 : n2]
	name = strings.TrimSpace(name)
	typ := "func" + ps[n2:]
	typ = strings.TrimSpace(typ)
	//fmt.Printf("type T%s %s\n", name, typ)
	for i := 0; i < len(res); i++ {
		res[i] = convertTypeName(res[i])
	}
	addFunc(name, typ, res, doc)
}

func convertTypeName(typ string) string {
	t := strings.TrimSpace(typ)
	n := strings.Index(t, "*")
	if n > -1 {
		t = t[n+1:]
	}
	n = strings.Index(t, " ")
	if n > -1 {
		t = t[n+1:]
	}
	t = strings.TrimSpace(t)
	if isBuiltin(t) {
		return typ
	} else {
		return "interface{}"
	}
}

func isBuiltin(typ string) bool {
	builtintyps := []string{"ComplexType",
		"FloatType",
		"IntegerType",
		"Type",
		"Type1",
		"bool",
		"byte",
		"complex128",
		"complex64",
		"error",
		"float32",
		"float64",
		"int",
		"int16",
		"int32",
		"int64",
		"int8",
		"rune",
		"string",
		"uint",
		"uint16",
		"uint32",
		"uint64",
		"uint8",
		"uintptr"}
	length := len(builtintyps)
	for i := 0; i < length; i++ {
		if typ == builtintyps[i] {
			return true
		}
	}
	return false
}
