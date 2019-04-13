package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
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

func saveIdents(name string) {
	fp, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	pkgName := name[:len(name)-7]
	w := fmt.Sprintf("package main\n\n")
	fp.WriteString(w)
	if len(funcs) > 0 {
		w = fmt.Sprintf("import \"github.com/rocket049/pluginloader\"\n\n")
		fp.WriteString(w)
	}
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

	afunc := []resFunc{}
	for fn, f := range funcs {
		typ := createFunc(&f)
		afunc = append(afunc, resFunc{fn, typ})
	}
	data := make(map[string]interface{})
	data["pkg"] = pkgName
	data["funcs"] = afunc
	t := template.New("")
	t.Parse(fnTpl)
	t.Execute(fp, data)
}

type resFunc struct {
	Name string
	Typ  string
}

const fnTpl = `
{{range .funcs}}var {{.Name}} {{.Typ}}
{{end}}

func Init{{.pkg}}Funcs(p *pluginloader.PluginLoader) {
{{range .funcs}}	p.MakeFunc(&{{.Name}}, "{{.Name}}")
{{end}}
}
`
