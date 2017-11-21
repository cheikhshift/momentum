package main

import (
	"fmt"
	"github.com/cheikhshift/gos/core"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Used GoAst example as starter
	fmt.Println("Welcome to Momentum Aftergo.")
	fmt.Println("Converting funcs with `RPC` in comments.")

	//add template path

	fnParamMap := make(map[string]string)
	fnReturnMap := make(map[string]string)

	// Create the AST by parsing src.
	fsettopl := token.NewFileSet() // positions are relative to fset
	pwd, _ := os.Getwd()
	var strfuncs, strtemplate, jsstr, jstrbits string
	pkgs, err := parser.ParseDir(fsettopl, pwd, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for name, pkg := range pkgs {
		fmt.Println("Processing package :", name)
		for fname, _ := range pkg.Files {
			if strings.Contains(fname, "momentum_") {
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
				if !strings.Contains(name, "main") {
					jstrbits = fmt.Sprintf(`var %s = {};
						`, name)
				} else {
					jstrbits = ""
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
								for i, cmment := range fn.Doc.List {
									checkforRPC[i] = cmment.Text
								}
								if strings.Contains(strings.Join(checkforRPC, ":"), "RPC") {
									if fn.Type.Params != nil {
										strret := ""
										limtlen := len(fn.Type.Params.List) - 1
										for indx, fieldss := range fn.Type.Params.List {

											limtlenv := len(fieldss.Names) - 1
											typeExpr := fieldss.Type
											start := typeExpr.Pos() - 2
											end := typeExpr.End() - 1
											for indxv, fieldnamesubs := range fieldss.Names {
												strret += fmt.Sprintf("%s%s", fieldnamesubs, strbody[start:end])
												if indxv < limtlenv {
													strret += ","
												}
											}

											if indx < limtlen {
												strret += ","
											}
										}
										fnParamMap[fn.Name.Name] = strret
									}
									if fn.Type.Results != nil {
										strret := ""
										limtlen := len(fn.Type.Results.List) - 1
										for indx, fieldss := range fn.Type.Results.List {
											limtlenv := len(fieldss.Names) - 1
											typeExpr := fieldss.Type
											start := typeExpr.Pos() - 2
											end := typeExpr.End() - 1
											for indxv, fieldnamesubs := range fieldss.Names {
												strret += fmt.Sprintf("%s%s", fieldnamesubs, strbody[start:end])
												if indxv < limtlenv {
													strret += ","
												}
											}

											if indx < limtlen {
												strret += ","
											}
										}
										if strret != "" {
											fnReturnMap[fn.Name.Name] = strret
										} else {
											fnReturnMap[fn.Name.Name] = "result "
										fmt.Println("No named returned variable found, assume your result to be in json key `result`.")
										}
									} 

									

									//build rpc method
									varss := strings.Split(fnParamMap[fn.Name.Name], ",")
									responseformat := ``
									fnFormat := ``
									funcfields := []string{}
									binderString := ``
									jssetters := ``
									varsslen := len(varss) - 1
									for ind, variabl := range varss {
										varname := strings.Split(variabl, " ")
										if len(varname) > 1 {
											fnFormat += fmt.Sprintf("tmvv.%s", strings.Title(varname[0]))
											jssetters += fmt.Sprintf(`
	t.%s = %s`, strings.Title(varname[0]), strings.Title(varname[0]))
											if ind < varsslen {
												fnFormat += ","
											}

											funcfields = append(funcfields, fmt.Sprintf("%s %s", fmt.Sprintf("%s", strings.Title(varname[0])), varname[1]))
										}
									}
									comma := ","
									if fnFormat == "" {
										comma = ""
									}
									if strings.Contains(name, "main") {
									jstrbits += fmt.Sprintf(`function %s(%s %s cb){
	var t = {}
	%s
	jsrequestmomentum("/momentum/funcs?name=%s", t, "POSTJSON", cb)
}
`, fn.Name.Name, strings.Replace(fnFormat, "tmvv.", "", -1), comma, jssetters, fn.Name.Name)
									} else {
										jstrbits += fmt.Sprintf(`%s["%s"] = function(%s %s cb){
	var t = {}
	%s
	jsrequestmomentum("/momentum/funcs?name=%s.%s", t, "POSTJSON", cb)
}
`, name,fn.Name.Name, strings.Replace(fnFormat, "tmvv.", "", -1), comma, jssetters, name,fn.Name.Name)
									}

									parsable := strings.Split(fnReturnMap[fn.Name.Name], ",")
									if fnReturnMap[fn.Name.Name] == "" {
															strfuncs += fmt.Sprintf(`} else if r.FormValue("name") == "%s.%s" {
			w.Header().Set("Content-Type", "application/json")
			type Payload%s struct {
				%s
			}
			decoder := json.NewDecoder(r.Body)
			 var tmvv Payload%s
			 err := decoder.Decode(&tmvv)
			 if err != nil {
			 	w.WriteHeader(http.StatusInternalServerError)
			    w.Write([]byte(fmt.Sprintf("{\"error\":\"%%s\"}",err.Error())))
			    return
			 }
			resp := bson.M{}
			%s(%s)
			w.Write([]byte(mResponse(resp)))
		`, name,fn.Name.Name, fn.Name.Name, strings.Join(funcfields, "\n"), fn.Name.Name, fn.Name.Name, fnFormat)
									} else {
									parsablelen := len(parsable) - 1
									for ind, variabl := range parsable {
										varname := strings.Split(variabl, " ")

										responseformat += fmt.Sprintf("resp%s%v", varname[0], ind)
										if strings.Contains(variabl, "error") {
											binderString += fmt.Sprintf(`
				if resp%s%v != nil {
					resp["%s"] = resp%s%v.Error()
				} else {
					resp["%s"] = nil
				}`, varname[0], ind, varname[0], varname[0], ind, varname[0])
										} else {
											binderString += fmt.Sprintf(`
				resp["%s"] = resp%s%v`, varname[0], varname[0], ind)
										}
										if ind < parsablelen {
											responseformat += ","
										}
									}

									strfuncs += fmt.Sprintf(`} else if r.FormValue("name") == "%s.%s" {
			w.Header().Set("Content-Type", "application/json")
			type Payload%s struct {
				%s
			}
			decoder := json.NewDecoder(r.Body)
			 var tmvv Payload%s
			 err := decoder.Decode(&tmvv)
			 if err != nil {
			 	w.WriteHeader(http.StatusInternalServerError)
			    w.Write([]byte(fmt.Sprintf("{\"error\":\"JSON : %%s\"}",err.Error())))
			    return
			 }
			resp := bson.M{}
			%s := %s(%s)
			%s
			w.Write([]byte(mResponse(resp)))
		`,name, fn.Name.Name, fn.Name.Name, strings.Join(funcfields, "\n"), fn.Name.Name, responseformat, fn.Name.Name, fnFormat, binderString)
							}

								}

							}
						}
					}

				}

				// Print the modified AST.
			}
		}
		strtemplate = fmt.Sprintf(`func Momentum(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if n := recover(); n != nil {
				w.Write([]byte(mResponse(bson.M{"error": fmt.Sprintf("%%s", n)})))
			}
		}()
		if r.FormValue("name") == "reset" {
			return
		%s
		}
	}`, strfuncs)

		jsdepfuncs := `  function jsrequestmomentum(url,payload,type,callback){
   var xhttp = new XMLHttpRequest();
  	xhttp.onreadystatechange = function() {
  		if(xhttp.readyState == 4){
   		var success = ( xhttp.status == 200)
    	if (type == "POSTJSON"){
    		try {
    		callback(JSON.parse(xhttp.responseText), success);
    		} catch (e) {
    			console.log("Invalid JSON");
    			callback({error : xhttp.responseText == "" ? "Server wrote no response" : xhttp.responseText}, false )
    		}
    	} else callback(xhttp.responseText, success );
    }
  };

  var serialize = function(obj) {
  var str = [];
  for(var p in obj)
    if (obj.hasOwnProperty(p)) {
      str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
    }
  return str.join("&");
  }
  xhttp.open(type, url, true);

  if(type == "POST"){
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send(serialize(payload));
} else   if(type == "POSTJSON"){
    xhttp.setRequestHeader("Content-type", "application/json");
    xhttp.send(JSON.stringify(payload));
}  else  xhttp.send();
}
`
		jsstr = fmt.Sprintf(`package %s

									import (
			"net/http"
			"strings"
			"encoding/json"
			"fmt"
			"gopkg.in/mgo.v2/bson"
			)
									var (
										FuncFactory = "/funcfactory.js"
										MomentumURI = "/momentum/funcs"
									)
									var jsinitbytes = []byte(%s)
									var jsfuncs = []byte(%s)

									// Momentum http.Handler
									// of your package. Use this with variable MomentumURI.
									// It will give server side access to the generated javascript
									// library.
									%s

									func mResponse(v interface{}) string {
										data, _ := json.Marshal(&v)
										return string(data)
									}

									// Middleware
									// Chain this with other handlers
									// to import other momentum javascript
									// libraries.
								    func MomentumJSChain(f http.HandlerFunc) http.HandlerFunc {
								        return func(w http.ResponseWriter, r *http.Request) {
								        	MomentumJS(w, r)
								        	f(w,r)
								        }
								    }

								    // Middleware
								    // Chain this with other momentum handlers
								    // to use RPC functions from multiple 
								    // libraries.
									func MomentumChain(f http.HandlerFunc) http.HandlerFunc {
								        return func(w http.ResponseWriter, r *http.Request) {
								        	Momentum(w, r)
								        	f(w,r)
								        }
								    }
								   
								   	// http.Handler , Servers javascript library
									func MomentumJS(w http.ResponseWriter, r *http.Request) {
			if !strings.Contains(w.Header().Get("content-type"), "/javascript" ) { 
		w.Header().Set("Content-Type", "text/javascript")
			w.Write( jsinitbytes )
		}
		w.Write(jsfuncs)
	}`, name, fmt.Sprintf("` %s `",jsdepfuncs),  fmt.Sprintf("` %s `",jstrbits), strtemplate )

		d1 := []byte(jsstr)
		_ = ioutil.WriteFile(fmt.Sprintf("momentum_%s.go", name), d1, 0700)
		core.RunCmd(fmt.Sprintf("gofmt -w momentum_%s.go", name))
	}
}
