package main

import (
	"fmt"
	"go/ast"
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
	structs       map[string][]funcDefine //structs map[structName]methodArray
	funcs         map[string]funcDefine   //funcs map[funcName]funcType
	typs          map[string]string       //typs map[typeName]name,user define tpes
	imports       map[string]string       //imports map[name]importString
	importsPicked map[string]string       //importsPicked map[name]importString
)

func init() {
	structs = make(map[string][]funcDefine)
	funcs = make(map[string]funcDefine)
	typs = make(map[string]string)
	imports = make(map[string]string)
	importsPicked = make(map[string]string)
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

func addTypeFromLine(line string) {
	sp := strings.SplitN(line, " ", 3)
	tv := strings.Split(sp[1], ".")
	if len(tv) == 2 {
		typs[sp[1]] = tv[0]
	} else {
		typs[sp[1]] = ""
	}
	//fmt.Println(sp[1])
}

func addImport(pkg *ast.ImportSpec) {
	var pkgName, importString string
	pkgPath := strings.Trim(pkg.Path.Value, `"`)
	if pkg.Name != nil {
		pkgName = pkg.Name.Name
		importString = pkgName + ` "` + pkgPath + `"`
	} else {
		sp := strings.Split(pkgPath, "/")
		pkgName = sp[len(sp)-1]
		importString = `"` + pkgPath + `"`
	}
	imports[pkgName] = importString
	//fmt.Println(importString)
}

func createFunc(f *funcDefine) string {
	res1 := strings.Join(f.Res, ",")
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

	afunc := []resFunc{}
	for fn, f := range funcs {
		typ := createFunc(&f)
		afunc = append(afunc, resFunc{fn, typ})
	}
	vimports := []string{}
	for _, v := range importsPicked {
		vimports = append(vimports, v)
	}
	data := make(map[string]interface{})
	data["imports"] = vimports
	data["pkg"] = pkgName
	data["funcs"] = afunc
	t := template.New("")
	t.Parse(fnTpl)
	t.Execute(fp, data)

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
}

type resFunc struct {
	Name string
	Typ  string
}

const fnTpl = `
import (
	{{range .imports}}{{.}}
	{{end}}
)
{{range .funcs}}var {{.Name}} {{.Typ}}
{{end}}

func Init{{.pkg}}Funcs(p *pluginloader.PluginLoader) {
{{range .funcs}}	p.MakeFunc(&{{.Name}}, "{{.Name}}")
{{end}}
}
`
