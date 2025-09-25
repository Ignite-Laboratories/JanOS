package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

var input string
var output string
var source string
var target string

var b *strings.Builder = new(strings.Builder)

//go:generate go run ./ -source MyStruct -target MyOtherStruct[T] -input ../input.go -output ../output.go

func main() {
	flag.StringVar(&input, "input", "", "path to the entry .go source file")
	flag.StringVar(&output, "output", "", "path to the output .go file")
	flag.StringVar(&source, "source", "", "path to the source structure")
	flag.StringVar(&target, "target", "", "path to the target structure")
	flag.Parse()

	var errs []error
	errs = appendIfNil(errs, sanityCheck(input, "input"))
	errs = appendIfNil(errs, sanityCheck(output, "output"))
	errs = appendIfNil(errs, sanityCheck(source, "source"))
	errs = appendIfNil(errs, sanityCheck(target, "target"))

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		os.Exit(2)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, input, nil, parser.ParseComments)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	/**
	Step 0 - Get the methods hung off of the source structure
	*/

	var methods []*ast.FuncDecl
	var receiverNames []string
	var receiverTypes [][]ast.Expr
	var isPointer []bool
	var args [][]ast.Expr
	requiredImports := map[string]struct{}{}

	ast.Inspect(file, func(n ast.Node) bool {
		fNode, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		if fNode.Recv == nil || len(fNode.Recv.List) == 0 {
			return false // no receiver; not a method
		}

		match := false
		receiverName := fNode.Recv.List[0].Names[0].Name

		ast.Inspect(fNode.Recv.List[0].Type, func(rn ast.Node) bool {
			if id, ok := rn.(*ast.Ident); ok && id.Name == source {
				match = true

				/**
				Step 1 - Update the receiver name
				*/
				var indexListExpr *ast.IndexListExpr
				switch t := fNode.Recv.List[0].Type.(type) {
				case *ast.StarExpr:
					// If it's a pointer, get the underlying type expression.
					if expr, ok := t.X.(*ast.IndexListExpr); ok {
						indexListExpr = expr
					}
				case *ast.IndexListExpr:
					indexListExpr = t
				}
				receiverTypes = append(receiverTypes, indexListExpr.Indices)

				id.Name = target
				return false
			}
			return true
		})

		if match {
			methods = append(methods, fNode)
			receiverNames = append(receiverNames, receiverName)

			if _, ptr := fNode.Recv.List[0].Type.(*ast.StarExpr); ptr {
				isPointer = append(isPointer, true)
			} else {
				isPointer = append(isPointer, false)
			}

			a := make([]ast.Expr, len(fNode.Type.Params.List))
			for j, param := range fNode.Type.Params.List {
				a[j] = &ast.Ident{Name: param.Names[0].Name}
				ast.Inspect(param.Type, func(n ast.Node) bool {
					if sel, ok := n.(*ast.SelectorExpr); ok {
						if id, ok := sel.X.(*ast.Ident); ok {
							requiredImports[id.Name] = struct{}{}
						}
					}
					return true
				})
			}
			args = append(args, a)
		}

		return false
	})

	/**
	Step 2 - Rewrite the bodies
	*/

	var specs []*ast.ImportSpec
	for k, _ := range requiredImports {
		for _, i := range file.Imports {
			if strings.HasSuffix(i.Path.Value, fmt.Sprintf(`%s"`, k)) {
				specs = append(specs, &ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: i.Path.Value,
					},
				})
			}
		}
	}

	outMethods := make([]*ast.FuncDecl, len(methods))

	for i, method := range methods {
		var cast ast.Expr
		cast = &ast.CallExpr{
			Fun:  &ast.Ident{Name: source},
			Args: []ast.Expr{&ast.Ident{Name: receiverNames[i]}},
		}

		if isPointer[i] {
			cast = &ast.StarExpr{
				X: cast,
			}
		}

		call := &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   cast,
				Sel: &ast.Ident{Name: method.Name.Name},
			},
			Args: args[i],
		}

		body := []ast.Stmt{
			&ast.ExprStmt{
				call,
			},
		}

		if method.Type.Results != nil {
			body = []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						call,
					},
				},
			}
		}

		method.Body = &ast.BlockStmt{
			List: body,
		}

		outMethods[i] = method
	}

	var decls []ast.Decl

	if len(specs) > 0 {
		impSpecs := make([]ast.Spec, len(specs))
		for i, s := range specs {
			impSpecs[i] = s
		}
		decls = append(decls, &ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: impSpecs,
		})
	}

	for _, m := range outMethods {
		decls = append(decls, m)
	}

	_ = printer.Fprint(b, fset, &ast.File{
		Name:  file.Name,
		Decls: decls,
	})

	fmt.Println(b.String())

	/**
	Step 3 - Write the output file
	*/

	err = os.WriteFile(output, []byte(b.String()), 0o644)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	os.Exit(0)
}

func sanityCheck(input string, name string) error {
	if len(input) == 0 {
		return fmt.Errorf("Missing input parameter '%s'", name)
	}
	return nil
}

func appendIfNil(data []error, element error) []error {
	if element == nil {
		return data
	}
	return append(data, element)
}

func fprintf(format string, a ...any) {
	_, _ = fmt.Fprintf(b, format, a...)
}
