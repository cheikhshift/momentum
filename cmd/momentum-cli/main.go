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
	

	flag.Parse()

	err := os.Chdir(*wd)
	if err != nil {
		panic(err)
	}
	
	// Used GoAst example as starter
	fmt.Println("Welcome to Momentum Aftergo.")
	fmt.Println("Converting funcs with `RPC` in comments.")


	
	//add template path
	fnInterfaceMap := make(map[string][]string)
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


									
									fmt.Println(fnInterfaceMap)
									//build rpc method
									varss := strings.Split(fnParamMap[fn.Name.Name], ",")
									responseformat := ``
									fnFormat := ``
									funcfields := []string{}
									binderString := ``
									jssetters := ``
									interfaceMap := fnInterfaceMap
									ObjMap, hasmap := interfaceMap[fn.Name.Name]
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

									if hasmap && !strings.Contains(jstrbits, "function ObjectCallb") {
										jstrbits += `
function ObjectCallb(jsonobj,success) {
	this.Working = false;
	Object.assign(this, jsonobj.Obj)
	this.cb(jsonobj.Result, success)
}
										`
									}

									if !hasmap {
									jstrbits += fmt.Sprintf(`/**
* RPC function : %s
* %s
* @namespace %s
* @method %s
* @params %s
* @return %s
*/
`, fn.Name.Name, strings.Replace(strings.Join(checkforRPC, `
*`), "//","" , -1 ),name,fn.Name.Name, fnParamMap[fn.Name.Name], fnReturnMap[fn.Name.Name] )
								}

									if strings.Contains(name, "main") {

									if hasmap {

										for i := 0; i < len(ObjMap); i++ {
											intr := ObjMap[i]
											intr = strings.Replace(intr, "*", "" , -1 )
											intrint := fmt.Sprintf(`
/**
* %s
* @namespace %s
* @param Obj - Object with Go interface fields.
* @return %s - Object with specified interface.
*/
var %s = function(Obj) {
	Object.assign(this, Obj)
}											
`, intr,name,intr, intr)
											if !strings.Contains(jstrbits,intrint) {
												jstrbits += intrint
											}

											jstrbits += fmt.Sprintf(`%s.prototype.%s = function(%s %s cb){
	var t = {};
	%s
	var payload = {Obj : this, Params : t};
	t.Working = true;
	this.cb = cb;
	jsrequestmomentum("/momentum/funcs?name=%s.%s", payload, "POSTJSON", ObjectCallb.bind(this))
}
`, intr,fn.Name.Name, strings.Replace(fnFormat, "tmvv.", "", -1), comma, jssetters, ObjMap[i], fn.Name.Name)


										}
									} else {
									jstrbits += fmt.Sprintf(`function %s(%s %s cb){
	var t = {}
	%s
	jsrequestmomentum("/momentum/funcs?name=%s", t, "POSTJSON", cb)
}
`, fn.Name.Name, strings.Replace(fnFormat, "tmvv.", "", -1), comma, jssetters, fn.Name.Name)

									}

									} else {

										if hasmap {

										for i := 0; i < len(ObjMap); i++ {
											intr := ObjMap[i]
											intr = strings.Replace(intr, "*", "" , -1 ) 
											intrint := fmt.Sprintf(`
/**
* %s
* @namespace %s
* @param Obj - Object with Go interface fields.
* @return %s - Object with specified interfaces.
*/
var %s["%s"] = function(Obj) {
	Object.assign(this, Obj)
}							
`, intr,name,intr,intr, ObjMap[i])
											if !strings.Contains(jstrbits,intrint) {
												jstrbits += intrint
											}

											jstrbits += fmt.Sprintf(`%s[%s].prototype.%s = function(%s %s cb){
	var t = {}
	%s
	var payload = {Obj : this, Params : t}
	jsrequestmomentum("/momentum/funcs?name=%s.%s", payload, "POSTJSON", cb)
}
`,name , intr,fn.Name.Name, strings.Replace(fnFormat, "tmvv.", "", -1), comma, jssetters, ObjMap[i],fn.Name.Name)


										}
									}	else {
										jstrbits += fmt.Sprintf(`%s["%s"] = function(%s %s cb){
									}
	var t = {}
	%s
	t.Working = true
	jsrequestmomentum("/momentum/funcs?name=%s", t, "POSTJSON", ObjectCallb.bind(this))
}
`, name,fn.Name.Name, strings.Replace(fnFormat, "tmvv.", "", -1), comma, jssetters, name,fn.Name.Name)
										}
									}

									parsable := strings.Split(fnReturnMap[fn.Name.Name], ",")
									var namemod string
									if !strings.Contains(name,"main"){
										namemod = fmt.Sprintf("%s.", name)
									}
									if fnReturnMap[fn.Name.Name] == "" {

										if hasmap {
											var freeifptn string
											if strings.Contains(ObjMap[0], "*"){
												freeifptn = "tmvv.Obj = nil"
											}
												strfuncs += fmt.Sprintf(`} else if r.FormValue("name") == "%s%s.%s" {
											
			w.Header().Set(ContentType, ContentField )
			type Payload%s struct {
				%s
			}
			type Group%s struct {
				Obj %s
				Params Payload%s
			}
			decoder := json.NewDecoder(r.Body)
			 var tmvv Group%s
			 err := decoder.Decode(&tmvv)
			 if err != nil {
			 	w.WriteHeader(http.StatusInternalServerError)
			    w.Write([]byte(fmt.Sprintf("{\"error\":\"%%s\"}",err.Error())))
			    return
			 }
			resp := bson.M{}
			tmvv.Obj.%s(%s)
			resp["Obj"] = tmvv.Obj
			jsonstr := mResponse(resp)
			jsonbytes := []byte(jsonstr)
			w.Write(jsonbytes)
			jsonbytes = nil
			%s
		`, namemod,ObjMap[0],fn.Name.Name, fn.Name.Name, strings.Join(funcfields, "\n"), fn.Name.Name,  ObjMap[0], fn.Name.Name ,fn.Name.Name, fn.Name.Name, strings.Replace(fnFormat, "tmvv.", "tmvv.Params.",1) , freeifptn)
										} else {
														strfuncs += fmt.Sprintf(`} else if r.FormValue("name") == "%s%s" {
			w.Header().Set(ContentType, ContentField )
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
			jsonstr := mResponse(resp)
			jsonbytes := []byte(jsonstr)
			w.Write(jsonbytes)
			jsonbytes = nil
		`, namemod,fn.Name.Name, fn.Name.Name, strings.Join(funcfields, "\n"), fn.Name.Name, fn.Name.Name, fnFormat)
									}

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

								if hasmap {
											var freeifptn string
											if strings.Contains(ObjMap[0], "*"){
												freeifptn = "tmvv.Obj = nil"
											}
												strfuncs += fmt.Sprintf(`} else if r.FormValue("name") == "%s%s.%s" {
			w.Header().Set(ContentType, ContentField )
			type Payload%s struct {
				%s
			}
			type Group%s struct {
				Obj %s
				Params Payload%s
			}
			decoder := json.NewDecoder(r.Body)
			 var tmvv Group%s
			 err := decoder.Decode(&tmvv)
			 if err != nil {
			 	w.WriteHeader(http.StatusInternalServerError)
			    w.Write([]byte(fmt.Sprintf("{\"error\":\"JSON : %%s\"}",err.Error())))
			    return
			 }
			resp := bson.M{}
			%s := tmvv.Obj.%s(%s)
			%s
			respW := bson.M{"Obj" : tmvv.Obj, "Result" : resp}
			jsonstr := mResponse(respW)
			jsonbytes := []byte(jsonstr)
			w.Write(jsonbytes)
			jsonbytes = nil
			%s
		`,namemod, ObjMap[0] , fn.Name.Name, fn.Name.Name, strings.Join(funcfields, "\n"),fn.Name.Name, ObjMap[0], fn.Name.Name, fn.Name.Name, responseformat, fn.Name.Name, strings.Replace(fnFormat, "tmvv.", "tmvv.Params.",1) , binderString,freeifptn)
										}  else {
								

									strfuncs += fmt.Sprintf(`} else if r.FormValue("name") == "%s%s" {
			w.Header().Set(ContentType, ContentField )
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
			jsonstr := mResponse(resp)
			jsonbytes := []byte(jsonstr)
			w.Write(jsonbytes)
			jsonbytes = nil
		`,namemod, fn.Name.Name, fn.Name.Name, strings.Join(funcfields, "\n"), fn.Name.Name, responseformat, fn.Name.Name, fnFormat, binderString)
							

								}


							}

							delete(fnInterfaceMap, fn.Name.Name)

								}

							}
						}
					}

				}

				// Print the modified AST.
			}
		}
		strtemplate = fmt.Sprintf(`
		var notfound =  []byte("{\"error\":\"Function not found!\" } ")
		const ContentType = "Content-Type"
		const ContentField = "application/json"
		const AllowedOrigins = "*"

		func Momentum(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", AllowedOrigins)
		w.Header().Set("Access-Control-Allow-Headers", ContentType)

		
		defer func() {
			if n := recover(); n != nil {
				w.Write([]byte(mResponse(bson.M{"error": fmt.Sprintf("%%s", n)})))
			}
		}()
		if r.FormValue("name") == "reset" || r.Method == "OPTIONS" {
			return
		%s
		} else {
			w.Header().Set(ContentType, ContentField )
			w.WriteHeader(http.StatusNotFound)
			w.Write(notfound)
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
  xhttp.open(type == "POSTJSON" ? "POST" : type, url, true);

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
								   
								   	// http.Handler , Serves javascript library
									func MomentumJS(w http.ResponseWriter, r *http.Request) {
			if !strings.Contains(w.Header().Get("content-type"), "/javascript" ) { 
		w.Header().Set("Content-Type", "text/javascript")
			w.Write( jsinitbytes )
		}
		w.Write(jsfuncs)
	}`, name, fmt.Sprintf("` %s `",jsdepfuncs),  fmt.Sprintf("` %s `",jstrbits ), strtemplate )

		d1 := []byte(jsstr)
		_ = ioutil.WriteFile(fmt.Sprintf("momentum_%s.go", name), d1, 0700)
		core.RunCmd(fmt.Sprintf("gofmt -w momentum_%s.go", name))
	}
}
