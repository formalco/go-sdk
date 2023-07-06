package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const packageName = "buf.build/gen/go/formal/admin/bufbuild/connect-go"
const path = "/admin/v1/adminv1connect"
const url = "https://adminv2.api.formalcloud.net"

var directoryPath = os.Getenv("DIRECTORY_PATH")

func main() {
	goModPath := ""
	if directoryPath != "" {
		goModPath = filepath.Join(directoryPath, "go.mod")
	} else {
		goModPath = "go.mod"
	}
	oldContent, err := os.ReadFile(goModPath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	oldVersion, err := getVersion(string(oldContent), packageName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	cmd := exec.Command("go", "get", packageName+"@latest")
	p, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	cmd.Dir = filepath.Join(p, "sdk")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	content, err := os.ReadFile(goModPath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	version, err := getVersion(string(content), packageName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("New version: %s\r\n", version)
	fmt.Printf("Old version: %s\r\n", oldVersion)
	//if version == oldVersion {
	//	os.Exit(0)
	//}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(os.Getenv("HOME"), "go")
	}
	packageDir := filepath.Join(gopath, "pkg", "mod", packageName+"@"+version)
	dirPath := filepath.Join(packageDir, path)

	var fields []*jen.Statement
	var fieldstwo []string

	f := jen.NewFile("sdk")

	f.Const().Defs(jen.Id("FORMAL_HOST_URL").String().Op("=").Lit(url))

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") {
			continue
		}

		fset := token.NewFileSet()
		source, err := parser.ParseFile(fset, dirPath+"/"+file.Name(), nil, 0)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		for _, decl := range source.Decls {
			if fun, ok := decl.(*ast.FuncDecl); ok {
				if strings.HasPrefix(fun.Name.Name, "New") && strings.HasSuffix(fun.Name.Name, "Client") {
					fields = append(fields, jen.Id(fun.Type.Results.List[0].Type.(*ast.Ident).Name).Qual(packageName+path, fun.Type.Results.List[0].Type.(*ast.Ident).Name))
					fieldstwo = append(fieldstwo, fun.Type.Results.List[0].Type.(*ast.Ident).Name)
				}

			}
		}
	}

	var s []jen.Code
	for _, field := range fields {
		s = append(s, field)
	}
	f.Type().Id("FormalSDK").Struct(s...)

	sdkBody := []jen.Code{
		jen.Id("httpClient").Op(":=").Op("&").Qual("net/http", "Client").Values(jen.Dict{
			//jen.Id("Timeout"): jen.Lit(100),
			jen.Id("Transport"): jen.Op("&").Id("transport").Values(jen.Dict{
				jen.Id("underlyingTransport"): jen.Qual("net/http", "DefaultTransport"),
				jen.Id("apiKey"):              jen.Id("apiKey"),
			}),
		}),
		jen.Return(jen.Id("&FormalSDK").Values(jen.DictFunc(func(d jen.Dict) {
			for _, field := range fieldstwo {
				d[jen.Id(field)] = jen.Qual(packageName+path, "New"+field).Params(
					jen.Id("httpClient"),
					jen.Id("FORMAL_HOST_URL"),
				)
			}
		}))),
	}

	f.Func().Id("New").Params(jen.Id("apiKey").String()).Op("*").Id("FormalSDK").Block(sdkBody...).Line()

	sdkWithUrlBody := []jen.Code{
		jen.Id("httpClient").Op(":=").Op("&").Qual("net/http", "Client").Values(jen.Dict{
			jen.Id("Transport"): jen.Op("&").Id("transport").Values(jen.Dict{
				jen.Id("underlyingTransport"): jen.Qual("net/http", "DefaultTransport"),
				jen.Id("apiKey"):              jen.Id("apiKey"),
			}),
		}),
		jen.Return(jen.Id("&FormalSDK").Values(jen.DictFunc(func(d jen.Dict) {
			for _, field := range fieldstwo {
				d[jen.Id(field)] = jen.Qual(packageName+path, "New"+field).Params(
					jen.Id("httpClient"),
					jen.Id("url"),
				)
			}
		}))),
	}

	f.Func().Id("NewWithUrl").Params(jen.Id("apiKey").String(), jen.Id("url").String()).Op("*").Id("FormalSDK").Block(sdkWithUrlBody...).Line()

	generateRoundTrip(f)
	fmt.Println("Saving file!")
	savePath := ""
	if directoryPath != "" {
		savePath = filepath.Join(directoryPath, "sdk.go")
	} else {
		savePath = "sdk.go"
	}
	err = f.Save(savePath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}

func getVersion(input string, pkgName string) (string, error) {
	re := regexp.MustCompile(pkgName + ` (v[^ \n]+)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 2 {
		return "", fmt.Errorf("no version for package name %s found in string", pkgName)
	}
	return matches[1], nil
}

func generateRoundTrip(f *jen.File) {
	transportBody := []jen.Code{
		jen.Id("underlyingTransport").Qual("net/http", "RoundTripper"),
		jen.Id("apiKey").String(),
	}
	f.Type().Id("transport").Struct(transportBody...)

	f.Func().Params(jen.Id("t").Op("*").Id("transport")).Id("RoundTrip").Params(jen.Id("req").Op("*").Qual("net/http", "Request")).Parens(jen.Op("*").Qual("net/http", "Response").Op(",").Error()).Block(
		jen.Id("req").Dot("Header").Dot("Add").Params(jen.Lit("X-Api-Key"), jen.Id("t").Dot("apiKey")),
		jen.Return(jen.Id("t").Dot("underlyingTransport").Dot("RoundTrip").Call(jen.Id("req"))),
	)
}
