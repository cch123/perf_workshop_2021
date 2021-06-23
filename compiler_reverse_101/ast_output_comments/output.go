package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./main.go", nil, parser.Mode(4))

	for _, d := range f.Comments {
		ast.Print(fset, d)
	}
}

