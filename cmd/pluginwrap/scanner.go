package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var help = flag.Bool("h", false, "show help")
	flag.Parse()
	if *help == true {
		fmt.Println(`Usage:
		pluginwrap path/to/plugin/src
		`)
		return
	}
	var path1 string = "."
	if len(os.Args) == 2 {
		path1 = os.Args[1]
	}
	cmd := exec.Command("gofmt", "-w", path1)
	cmd.Run()

	var outfile string
	if path1 == "." {
		outfile = "pluginWrap.go"
	} else {
		outfile = fmt.Sprintf("%sWrap.go", filepath.Base(path1))
	}
	fset := token.NewFileSet() // positions are relative to fset

	pkgs, err := parser.ParseDir(fset, path1, nil, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, p := range pkgs {
		fmt.Println("Pkg:", k)
		for fn, f := range p.Files {
			//ast.FileExports(f)
			fh := newFileHandler(fn)
			fmt.Println("File:", fn)
			getTypeDecls(f, fset, fh)
		}
	}
	for k, p := range pkgs {
		fmt.Println("Pkg:", k)
		for fn, f := range p.Files {
			//ast.FileExports(f)
			fh := newFileHandler(fn)
			fmt.Println("File:", fn)
			getFuncDecls(f, fset, fh)
		}
	}
	//printIdents()
	saveIdents(outfile)
	cmd = exec.Command("gofmt", "-w", outfile)
	cmd.Run()
}

func getTypeDecls(f *ast.File, fset *token.FileSet, fh *fileHandler) {
	for _, pkg := range f.Imports {
		addImport(pkg)
	}
	for _, n := range f.Decls {
		switch x := n.(type) {
		case *ast.GenDecl:
			if x.Tok == token.TYPE {
				fmt.Println("TYPE:", fh.GetLineAtPos(fset.Position(x.Pos()).Offset))
				addTypeFromLine(fh.GetLineAtPos(fset.Position(x.Pos()).Offset))
			}
		}
	}
}

func getFuncDecls(f *ast.File, fset *token.FileSet, fh *fileHandler) {
	for _, n := range f.Decls {
		switch x := n.(type) {
		case *ast.FuncDecl:
			res := []string{}
			args := []string{}
			vtyps := []string{}
			recv := ""
			if x.Type != nil {
				if x.Name.IsExported() == false {
					continue
				}
				if x.Recv != nil {
					p := x.Recv.List[0].Type.Pos()
					e := x.Recv.List[0].Type.End()
					recv = fh.GetLinePosEnd(fset.Position(p).Offset, fset.Position(e).Offset)
				}
				if x.Type.Results != nil {
					for _, v := range x.Type.Results.List {
						res = append(res, fh.GetLinePosEnd(fset.Position(v.Pos()).Offset, fset.Position(v.End()).Offset))
						vtyps = appendTypeFromExpr(vtyps, v.Type, fset, fh)
					}
				}
				if x.Type.Params != nil {
					for _, v := range x.Type.Params.List {
						args = append(args, fh.GetLinePosEnd(fset.Position(v.Pos()).Offset, fset.Position(v.End()).Offset))
						vtyps = appendTypeFromExpr(vtyps, v.Type, fset, fh)
					}
				}
			}

			//fmt.Println(x.Name.Name, "args:", args)
			parseLine(recv, x.Name.Name, args, res, vtyps, x.Doc.Text())
		}
	}
}

func appendTypeFromExpr(vtyps []string, expr ast.Expr, fset *token.FileSet, fh *fileHandler) []string {
	switch node := expr.(type) {
	case *ast.MapType:
		res1 := appendTypeFromExpr(vtyps, node.Key, fset, fh)
		return appendTypeFromExpr(res1, node.Value, fset, fh)
	default:
		return append(vtyps, fh.GetLinePosEnd(fset.Position(expr.Pos()).Offset, fset.Position(expr.End()).Offset))
	}
}

type fileHandler struct {
	buffer []byte
}

func newFileHandler(name string) *fileHandler {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return nil
	}
	return &fileHandler{buffer: buf}
}

func (s *fileHandler) GetLineAtPos(pos int) string {
	n := bytes.Index(s.buffer[pos:], []byte{'\n'})
	if n == -1 {
		return string(s.buffer[pos:])
	}
	return string(s.buffer[pos : pos+n])
}

func (s *fileHandler) GetLinePosEnd(pos, end int) string {
	return string(s.buffer[pos:end])
}
