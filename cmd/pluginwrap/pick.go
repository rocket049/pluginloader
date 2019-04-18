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
		pickFunc(line, res, doc)
	} else if rsp.MatchString(line) {
		pickMethod(line, res, doc)
	} else {
		fmt.Println("not match:", line)
	}
}

//pickMethod get struct name and method from a line,then append to structs
func pickMethod(line string, res []string, doc string) {
	rs := `^func\s+\(\w+\s+\**(\w+)\)`
	rsp, err := regexp.Compile(rs)
	if err != nil {
		panic(err)
	}
	ps := rsp.FindStringSubmatch(line)
	name := ps[1]

	rs = `^func\s+\(\w+\s+\**\w+\)\s*(\w+\([^\)]*\))`
	rsp, err = regexp.Compile(rs)
	if err != nil {
		panic(err)
	}
	ps = rsp.FindStringSubmatch(line)
	method := strings.TrimSpace(ps[1])

	rs = `^func\s+\(\w+\s+\**\w+\)\s*\w+\(([^\)]*)\)`
	rsp, err = regexp.Compile(rs)
	if err != nil {
		panic(err)
	}
	ps = rsp.FindStringSubmatch(line)
	args := strings.Split(ps[1], ",")
	//fmt.Printf("%s.%s\n", name, method)

	var pick bool = true
	for i := 0; i < len(res); i++ {
		if isUserDefType(res[i]) {
			pick = false
			break
		}
	}
	for i := 0; i < len(args); i++ {
		if isUserDefType(args[i]) {
			pick = false
			break
		}
	}
	if pick {
		addStructMethod(name, method, res, doc)
	}
}

//pickFunc get func from a line,then append to funcs
func pickFunc(line string, res []string, doc string) {
	rf := `^func[^\)]*\)`
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

	rs := `^func\s+\w+\(([^\)]*)\)`
	rsp, err := regexp.Compile(rs)
	if err != nil {
		panic(err)
	}
	psv := rsp.FindStringSubmatch(line)
	args := strings.Split(psv[1], ",")
	var pick bool = true
	for i := 0; i < len(res); i++ {
		if isUserDefType(res[i]) {
			pick = false
			break
		}
	}
	for i := 0; i < len(args); i++ {
		if isUserDefType(args[i]) {
			pick = false
			break
		}
	}
	if pick {
		addFunc(name, typ, res, doc)
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
	t := strings.TrimSpace(typ)
	n := strings.Index(t, "*")
	if n > -1 {
		t = t[n+1:]
	}
	n = strings.Index(t, "[]")
	if n > -1 {
		t = t[n+2:]
		fmt.Println("[]", t)
	}
	n = strings.Index(t, " ")
	if n > -1 {
		t = t[n+1:]
	}
	t = strings.TrimSpace(t)
	return t
}

func isUserDefType(typ string) bool {
	t := clearType(typ)
	if strings.HasPrefix(t, "map[") == false {
		_, ok := typs[t]
		if !ok {
			pickImport(t)
		}
		return ok
	} else {
		//map[type1]type2
		typs := strings.SplitN(t[4:], "]", 2)
		for _, t1 := range typs {
			if isUserDefType(t1) == true {
				return true
			}
		}
		return false
	}
}

func pickImport(typ string) {
	sp := strings.Split(typ, ".")

	if len(sp) != 2 {
		return
	}

	importsPicked[sp[0]] = imports[sp[0]]
}
