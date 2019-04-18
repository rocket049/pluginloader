package main

import (
	"bytes"
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
	var path1 string = "."
	if len(os.Args) == 2 {
		path1 = os.Args[1]
	}
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
			showFile(f, fset, fh)
		}
	}
	//printIdents()
	saveIdents(outfile)
	cmd := exec.Command("go", "fmt", outfile)
	cmd.Run()
}

func showFile(f *ast.File, fset *token.FileSet, fh *fileHandler) {
	for _, pkg := range f.Imports {
		addImport(pkg)
	}
	for _, n := range f.Decls {
		switch x := n.(type) {
		case *ast.FuncDecl:
			line := fh.GetLineAtPos(fset.Position(x.Pos()).Offset)
			res := []string{}
			if x.Type != nil {
				if x.Name.IsExported() == false {
					continue
				}
				if x.Type.Results != nil {
					for _, v := range x.Type.Results.List {
						res = append(res, fh.GetLinePosEnd(fset.Position(v.Pos()).Offset, fset.Position(v.End()).Offset))
					}
				}
			}

			//fmt.Println(x.Name.Name, "args:", args)
			parseLine(line, res, x.Doc.Text())
		case *ast.GenDecl:
			if x.Tok == token.TYPE {
				//fmt.Println("TYPE:", fh.GetLineAtPos(fset.Position(x.Pos()).Offset))
				addTypeFromLine(fh.GetLineAtPos(fset.Position(x.Pos()).Offset))
			}
		}
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
