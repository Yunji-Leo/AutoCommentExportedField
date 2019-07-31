package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	filepath := os.Args[1]

	// parse file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	comments := []*ast.CommentGroup{}
	ast.Inspect(node, func(n ast.Node) bool {
		// collect comments
		c, ok := n.(*ast.CommentGroup)
		if ok {
			comments = append(comments, c)
		}

		switch n.(type) {
		case *ast.FuncDecl:
			fn := n.(*ast.FuncDecl)
			if cg := createCommentGroup(fn.Name, fn.Doc, fn.Pos(), fset.Position(fn.Pos()).Line, "function"); cg != nil {
				fn.Doc = cg
			}
		case *ast.GenDecl:
			gd := n.(*ast.GenDecl)
			for i := range gd.Specs {
				switch gd.Specs[i].(type) {
				case *ast.TypeSpec:
					ts := gd.Specs[i].(*ast.TypeSpec)
					if cg := createCommentGroup(ts.Name, gd.Doc, gd.Pos(), fset.Position(ts.Pos()).Line, "type"); cg != nil {
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
						if cg := createCommentGroup(vs.Names[j], doc, pos, fset.Position(vs.Pos()).Line, "value"); cg != nil {
							vs.Doc = cg
						}
					}
				}
			}
		}
		return true
	})
	// set ast's comments to the collected comments
	node.Comments = comments
	// write new ast to file
	f, err := os.Create(filepath)
	defer f.Close()
	if err := printer.Fprint(f, fset, node); err != nil {
		log.Fatal(err)
	}
}

func createCommentGroup(ident *ast.Ident, doc *ast.CommentGroup, pos token.Pos, line int, declType string) *ast.CommentGroup {
	if ident.IsExported() && doc.Text() == "" {
		fmt.Printf("exported "+declType+" declaration without documentation found on line %d: \n\t%s\n", line, ident.Name)
		comment := &ast.Comment{
			Text:  "//" + ident.Name + " TODO: document exported " + declType,
			Slash: pos - 1,
		}
		// create CommentGroup and set it to the function's documentation comment
		cg := &ast.CommentGroup{
			List: []*ast.Comment{comment},
		}
		return cg
	}
	return nil
}
