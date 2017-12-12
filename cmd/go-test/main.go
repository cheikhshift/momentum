package main

import (
	"fmt"
	"github.com/cheikhshift/gos/core"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"flag"
	"strings"
)

func Exists(arr []string, lookup string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == lookup {
			return true
		}
	}
	return false
}

func main() {
	wd :=  flag.String("workdir","", "Path of directory with go sources to convert.")
	runtest := flag.Bool("test", false, "Run `$ go test` after source files are created.")

	flag.Parse()

	if *wd != "" {
		err := os.Chdir(*wd)
		if err != nil {
			panic(err)
		}
	}
	// Used GoAst example as starter
	fmt.Println("Welcome to go-test.")
	fmt.Println("Define tests with function comments")


	
	//add template path
	fnInterfaceMap := make(map[string][]string)
	

	// Create the AST by parsing src.
	fsettopl := token.NewFileSet() // positions are relative to fset
	pwd, _ := os.Getwd()
	var strfuncs, jsstr string
	pkgs, err := parser.ParseDir(fsettopl, pwd, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for name, pkg := range pkgs {
		fmt.Println("Processing package :", name)
		for fname, _ := range pkg.Files {
			if strings.Contains(fname, "_test") {
				os.Remove(fname)
			} else {
				fset := token.NewFileSet() // positions are relative to fset
				f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
				if err != nil {
					panic(err)
				}
				
				body, err := ioutil.ReadFile(fname)
				if err != nil {
					panic(err)
				}
			
				strbody := string(body)

				// Create an ast.CommentMap from the ast.File's comments.
				// This helps keeping the association between comments
				// and AST nodes.

				// Remove the first variable declaration from the list of declarations.
				for _, d := range f.Decls {
					if fn, isFn := d.(*ast.FuncDecl); isFn {
						if fn.Doc != nil {
							if len(fn.Doc.List) > 0 {
								var checkforRPC = make([]string, len(fn.Doc.List), len(fn.Doc.List))
								var testcases = []string{}
								for i, cmment := range fn.Doc.List {
									checkforRPC[i] = cmment.Text
									if strings.Contains(cmment.Text, "@case") {
										testcases = append(testcases, strings.Replace(cmment.Text ,"//","", -1 ) )
									}
								}
								if strings.Contains(strings.Join(checkforRPC, ":"), "@test") {

									

									fnString := strbody[(fn.Type.Pos() - 1 ): (fn.Type.End() -1 )]								

									partsstr := strings.Split(fnString, fn.Name.Name)

									if strings.Contains(partsstr[0] , ")" ) {
										partssub := strings.Split(strings.TrimSpace(partsstr[0]), " ")
										if _, exts := fnInterfaceMap[fn.Name.Name] ; !exts {
											fnInterfaceMap[fn.Name.Name] = []string{}
										}
										

										intname :=strings.Replace( strings.Replace(strings.TrimSpace(partssub[len(partssub) - 1]), ")","", -1) , "*","",-1)
										if strings.Contains(fnString, "*"){
											intname = fmt.Sprintf(`*%s`, intname)
										}
										if !Exists(fnInterfaceMap[fn.Name.Name], intname ) {

											fnInterfaceMap[fn.Name.Name] = append(fnInterfaceMap[fn.Name.Name], intname)

										}
									}

							
									
									
									//build rpc method
									
								
									interfaceMap := fnInterfaceMap
									ObjMap, hasmap := interfaceMap[fn.Name.Name]
								
								
									var strtester string

									for i := 0; i < len(testcases); i++ {
										caseset := strings.Split(strings.Replace(testcases[i],"@case", "", 1),"@equal")
										
										expoutput := strings.Split(caseset[1], ",")
										expoutmapped := ""
										for u := 0; u < len(expoutput); u++ {
											

											expoutmapped += fmt.Sprintf("expgen%v%v",i, u)
											if (u + 1) < len(expoutput) {
												expoutmapped += ","
											}
										}
										if hasmap {
											strtester += fmt.Sprintf(`
												%s := obj.%s(%s)`, expoutmapped,fn.Name.Name, caseset[0])
										} else {
										strtester += fmt.Sprintf(`
											%s := %s(%s)`, expoutmapped,fn.Name.Name, caseset[0])
										}
										for u := 0; u < len(expoutput); u++ {
											strtester += fmt.Sprintf(`
											if diff := deep.Equal(expgen%v%v, %s); diff != nil {
												t.Error(diff)
											}
	`,i, u, expoutput[u])
										}

									}

								

								if hasmap {
											
										strfuncs += fmt.Sprintf(`
									func Test%s(t *testing.T) {
										obj := %s{}

										%s
										
									}`,fn.Name.Name,strings.Replace(ObjMap[0], "*","&", 1),strtester)
								

								}  else {
								
									//strfuncs += 
							strfuncs += fmt.Sprintf(`
									func Test%s(t *testing.T) {
										
										%s
										
									}`,fn.Name.Name,strtester)

								}


							

							delete(fnInterfaceMap, fn.Name.Name)

								}

							}
						}
					}

				}

				
			}
		}
	

		jsstr = fmt.Sprintf(`package %s

		import (
			"testing"
			"github.com/go-test/deep"
		)
		

		%s
		`,name,strfuncs)
		d1 := []byte(jsstr)

		_ = ioutil.WriteFile(fmt.Sprintf("%s_test.go", name), d1, 0700)
		core.RunCmd(fmt.Sprintf("gofmt -w %s_test.go", name))
		if *runtest {
			core.RunCmd("go test")
		}
	}
}
