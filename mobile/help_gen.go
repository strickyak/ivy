// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/build"
	"go/doc"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	pkg, err := build.Import("github.com/strickyak/ivy", "", build.ImportComment)
	if err != nil {
		log.Fatal(err)
	}
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, pkg.Dir, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	astPkg := pkgs[pkg.Name]
	if astPkg == nil {
		log.Fatalf("failed to locate %s package", pkg.Name)
	}

	docPkg := doc.New(astPkg, pkg.ImportPath, doc.AllDecls)

	htmlBuf := new(bytes.Buffer)
	fmt.Fprintln(htmlBuf, `<!-- auto-generated from github.com/strickyak/ivy package doc -->`)
	fmt.Fprintln(htmlBuf, head)
	fmt.Fprintln(htmlBuf, `<body>`)
	doc.ToHTML(htmlBuf, docPkg.Doc, nil)
	fmt.Fprintln(htmlBuf, `</body></html>`)

	goBuf := new(bytes.Buffer)
	fmt.Fprintf(goBuf, "package mobile\n\n")
	fmt.Fprintf(goBuf, "// GENERATED; DO NOT EDIT\n")
	fmt.Fprintf(goBuf, "const help = `%s`\n", sanitize(htmlBuf.Bytes()))

	buf, err := format.Source(goBuf.Bytes())
	if err != nil {
		log.Fatalf("failed to gofmt: %v", err)
	}
	os.Stdout.Write(buf)
}

func sanitize(b []byte) []byte {
	// Replace ` with `+"`"+`
	return bytes.Replace(b, []byte("`"), []byte("`+\"`\"+`"), -1)
}

const head = `
<head>
    <style>
        body {
                font-family: Arial, sans-serif;
	        font-size: 10pt;
                line-height: 1.3em;
                max-width: 950px;
                word-break: normal;
                word-wrap: normal;
        }

        pre {
                border-radius: 10px;
                border: 2px solid #8AC007;
		font-family: monospace;
		font-size: 10pt;
                overflow: auto;
                padding: 10px;
                white-space: pre;
        }
    </style>
</head>`
