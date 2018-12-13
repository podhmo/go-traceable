package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	_ "github.com/podhmo/go-traceable/httptrace" // for go get
	"github.com/podhmo/go-traceable/injectmain"
	"golang.org/x/tools/go/ast/astutil"
)

const template = `
package p


func main() {
    fmt.Println("running via github.com/podhmo/go-traceable/cmd/go-run-httptrace/")
	teardown := httptrace.Patch()
	defer teardown()
	mainInner()
}
`

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("go-run-httptrace <main.go> [extra arguments]")
		os.Exit(1)
	}

	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	mainFile := os.Args[1]
	source, err := injectmain.Process(mainFile, template, func(fset *token.FileSet, f *ast.File) error {
		astutil.AddNamedImport(fset, f, "httptrace", "github.com/podhmo/go-traceable/httptrace")
		return nil
	})
	if err != nil {
		return err
	}

	originalSource, err := ioutil.ReadFile(mainFile)
	if err != nil {
		return err
	}

	var once sync.Once
	rollback := func() {
		once.Do(func() {
			log.Printf("rollback %s\n", mainFile)
			if err := ioutil.WriteFile(mainFile, originalSource, 0744); err != nil {
				log.Fatal(err)
			}
		})
	}
	defer rollback()
	setAtExitRollback(rollback)

	if err := ioutil.WriteFile(mainFile, source, 0744); err != nil {
		return err
	}

	args := []string{"run"}
	args = append(args, os.Args[1:]...)
	return RunTransformed("go", args)
}

// RunTransformed :
func RunTransformed(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Env = os.Environ()

	found := false
	for _, x := range c.Env {
		if strings.Contains(x, "TRACE=") {
			found = true
			break
		}
	}
	if !found {
		c.Env = append(c.Env, "TRACE=1")
	}
	return c.Run()
}

func setAtExitRollback(rollback func()) {
	c := make(chan os.Signal)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	go func() {
		<-c
		rollback()
		signal.Stop(c)
	}()
}

// TODO: handing permission, correctly
// TODO: handling TRACE varenv, correctly
