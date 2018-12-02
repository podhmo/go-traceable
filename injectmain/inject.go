package injectmain

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

// Process :
func Process(targetFile, template string, modifiers ...func(*ast.File) error) ([]byte, error) {
	fset := token.NewFileSet()
	source, err := ioutil.ReadFile(targetFile)
	if err != nil {
		return nil, err
	}

	log.Println("parse", targetFile)
	f, err := parser.ParseFile(fset, "main", source, parser.AllErrors)
	if err != nil {
		return nil, errors.Wrap(err, "parse main")
	}
	f2, err := parser.ParseFile(fset, "extra", template, parser.AllErrors)
	if err != nil {
		return nil, errors.Wrap(err, "parse template")
	}

	log.Println("transform AST")
	if err := InjectMain(f, f2); err != nil {
		return nil, errors.Wrap(err, "inject")
	}

	for _, m := range modifiers {
		// for example, calling astutil.AddImport()
		if err := m(f); err != nil {
			return nil, errors.Wrap(err, "modify")
		}
	}

	log.Println("format")
	buf := new(bytes.Buffer)
	if err := printer.Fprint(buf, fset, f); err != nil {
		return nil, err
	}

	return imports.Process(targetFile, buf.Bytes(), &imports.Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	})
}

// InjectMain :
func InjectMain(targetT *ast.File, templateT *ast.File) error {
	return Inject(targetT, templateT, "main", "mainInner")
}

// Inject :
func Inject(targetT *ast.File, templateT *ast.File, from, to string) error {
	f := targetT
	f2 := templateT

	if f.Scope.Lookup(to) == nil {
		main := f.Scope.Lookup(from)
		if main.Kind != ast.Fun {
			return fmt.Errorf("unexpected type %s", main.Kind)
		}
		main.Decl.(*ast.FuncDecl).Name.Name = to
	} else {
		main := f.Scope.Lookup(from)
		for i, decl := range f.Decls {
			if main.Decl == decl {
				f.Decls = append(f.Decls[:i], f.Decls[i+1:]...)
				break
			}
		}
	}

	fn := f2.Scope.Lookup("main").Decl.(*ast.FuncDecl)
	f.Decls = append(f.Decls, fn)
	return nil
}
