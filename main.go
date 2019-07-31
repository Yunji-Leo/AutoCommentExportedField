package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		fi, err := os.Stat(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			err = filepath.Walk(fi.Name(), func(pathX string, infoX os.FileInfo, errX error) error {
				return fixComment(pathX)
			})
		case mode.IsRegular():
			err = fixComment(fi.Name())
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

func fixComment(filename string) error {
	if !strings.HasSuffix(filename, ".go") {
		//fmt.Println("Skip " + filename)
		return nil
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	comments := []*ast.CommentGroup{}
	ast.Inspect(node, func(n ast.Node) bool {

		c, ok := n.(*ast.CommentGroup)
		if ok {
			comments = append(comments, c)
		}

		switch n.(type) {
		case *ast.FuncDecl:
			fn := n.(*ast.FuncDecl)
			if cg := createCommentGroup(fn.Name, fn.Doc, fn.Pos(), filename, fset.Position(fn.Pos()).Line, "function"); cg != nil {
				fn.Doc = cg
			}
		case *ast.GenDecl:
			gd := n.(*ast.GenDecl)
			for i := range gd.Specs {
				switch gd.Specs[i].(type) {
				case *ast.TypeSpec:
					ts := gd.Specs[i].(*ast.TypeSpec)
					if cg := createCommentGroup(ts.Name, gd.Doc, gd.Pos(), filename, fset.Position(ts.Pos()).Line, "type"); cg != nil {
						ts.Doc = cg
					}
				case *ast.ValueSpec:
					vs := gd.Specs[i].(*ast.ValueSpec)
					var pos token.Pos
					var doc *ast.CommentGroup
					if len(gd.Specs) > 1 {
						pos = vs.Pos()
						doc = vs.Doc
					} else {
						pos = gd.Pos()
						doc = gd.Doc
					}
					for j := range vs.Names {
						if cg := createCommentGroup(vs.Names[j], doc, pos, filename, fset.Position(vs.Pos()).Line, "value"); cg != nil {
							vs.Doc = cg
						}
					}
				}
			}
		}
		return true
	})

	node.Comments = comments

	f, err := os.Create(filename)
	defer f.Close()
	err = printer.Fprint(f, fset, node)
	return err
}

func createCommentGroup(ident *ast.Ident, doc *ast.CommentGroup, pos token.Pos, filename string, line int, declType string) *ast.CommentGroup {
	if ident.IsExported() && doc.Text() == "" {
		fmt.Printf("%s: exported "+declType+" declaration without documentation found on line %d: \n\t%s\n", filename, line, ident.Name)
		comment := &ast.Comment{
			Text:  "//" + ident.Name + " TODO: document exported " + declType,
			Slash: pos - 1,
		}

		cg := &ast.CommentGroup{
			List: []*ast.Comment{comment},
		}
		return cg
	}
	return nil
}
