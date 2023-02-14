package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type MyInterface interface {
	DoSomething() string
	DoSomeElse() string
}

// generateStruct generates a new struct named "MyStruct" that implements all the methods of the provided interface.
func generateStruct(packageName string, iface *ast.InterfaceType, fileName string) string {
	structName := "MyStruct"
	var builder strings.Builder
	// builder.WriteString("// Code generated by \"code-gen \"; DO NOT EDIT.\n")
	builder.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	builder.WriteString(fmt.Sprintf("type %s struct{\n\n}\n", structName))
	receiverName := strings.ToLower(structName[:1])
	for _, method := range iface.Methods.List {
		funcDecl, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}
		builder.WriteString("\n")
		builder.WriteString(fmt.Sprintf("func (%s *%s) %s(", receiverName, structName, method.Names[0]))
		for i, param := range funcDecl.Params.List {
			for _, name := range param.Names {
				builder.WriteString(name.Name)
				if i < len(funcDecl.Params.List)-1 {
					builder.WriteString(", ")
				}
			}
		}
		builder.WriteString(") ")
		if funcDecl.Results != nil && len(funcDecl.Results.List) > 0 {
			for _, result := range funcDecl.Results.List {
				for _, name := range result.Names {
					builder.WriteString(name.Name)
					builder.WriteString(" ")
				}
				builder.WriteString(result.Type.(*ast.Ident).Name)
			}
		}
		builder.WriteString("{\n\tpanic(\"implement me\")\n}\n")
	}
	return builder.String()
}

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", func(fi os.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go") && !strings.HasSuffix(fi.Name(), "_generated.go") && strings.HasSuffix(fi.Name(), ".go")
	}, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			fmt.Println("fileName", fileName)
			fmt.Println("package ", file.Name.Name)
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if iface, ok := typeSpec.Type.(*ast.InterfaceType); ok {
								newFileName := fmt.Sprintf("%s_generated.go", strings.TrimSuffix(fileName, filepath.Ext(fileName)))
								fmt.Println("new file name:", newFileName)
								newFile, err := os.Create(newFileName)
								if err != nil {
									panic(err)
								}
								defer newFile.Close()
								fmt.Fprintf(newFile, generateStruct(file.Name.Name, iface, newFileName))
							}
						}
					}
				}
			}
		}
	}
}
