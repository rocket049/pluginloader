package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type funcDefine struct {
	Body    string
	Res     []string
	Comment string
}

var (
	structs map[string][]funcDefine //structs map[structName]methodArray
	funcs   map[string]funcDefine   //funcs map[funcName]funcType
)

func init() {
	structs = make(map[string][]funcDefine)
	funcs = make(map[string]funcDefine)
}

func addStructMethod(name, body string, res []string, doc string) {
	ms, ok := structs[name]
	var p []funcDefine
	if !ok {
		p = []funcDefine{}
	} else {
		p = ms
	}
	structs[name] = append(p, funcDefine{body, res, doc})
}

func addFunc(name, body string, res []string, doc string) {
	funcs[name] = funcDefine{body, res, doc}
}

func createFunc(f *funcDefine) string {
	res := bytes.NewBufferString("")
	for i, _ := range f.Res {
		if i > 0 {
			res.WriteString(",")
		}
		res.WriteString(f.Res[i])
	}
	res1 := res.String()
	if strings.Contains(res1, ",") || strings.Contains(res1, " ") {
		res1 = fmt.Sprintf("(%s)", res1)
	}
	return fmt.Sprintf("%s %s", f.Body, res1)
}

// func printIdents() {
// 	fmt.Printf("package main\n\n")
// 	for name, ms := range structs {
// 		fmt.Printf("type I%s interface {\n", name)
// 		for _, v := range ms {
// 			fmt.Printf("\t%s\n", createFunc(&v))
// 		}
// 		fmt.Printf("}\n\n")
// 	}

// 	for name, f := range funcs {
// 		fmt.Printf("type Fn%s %s\n\n", name, createFunc(&f))
// 	}
// }

func saveIdents(name string) {
	fp, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	w := fmt.Sprintf("package main\n\n")
	fp.WriteString(w)
	// w = fmt.Sprintf("import \"plugin\"\n\n")
	// fp.WriteString(w)
	for name, ms := range structs {
		w = fmt.Sprintf("type I%s interface {\n", name)
		fp.WriteString(w)
		for _, v := range ms {
			w = fmt.Sprintf("\t%s\n", createFunc(&v))
			fp.WriteString(w)
		}
		w = fmt.Sprintf("}\n\n")
		fp.WriteString(w)
	}

	// for name, f := range funcs {
	// 	typ := createFunc(&f)
	// 	data := make(map[string]string)
	// 	data["name"] = name
	// 	data["typ"] = typ
	// 	t := template.New("")
	// 	t.Parse(fnTpl)
	// 	t.Execute(fp, data)
	// }
}

const fnTpl = `func Fn{{.name}}(p *plugin.Plugin) {{.typ}} {
	f,err := p.Lookup("{{.name}}")
	if err != nil {
		panic(err)
	}
	return f.({{.typ}})
}

`
