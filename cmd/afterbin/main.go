package main

import (
	"fmt"
	"github.com/cheikhshift/gos/core"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"reflect"
	"strings"
)

func getNodeType(node ast.Node) string {
	val := reflect.ValueOf(node).Elem()
	return val.Type().Name()
}

func main() {
	fmt.Println("Welcome to Momentum Afterbin.")
	fmt.Println("Remember to load momentum on HTML your pages : {{ server }} or <script src=\"/funcfactory.js\"></script>")
	fmt.Println("Converting Templates and Func tags into accessible JS functions.")
	cfg, err := core.Config()
	if err != nil {
		panic(err)
	}
	//add template paths

	fnParamMap := make(map[string]string)
	fnReturnMap := make(map[string]string)
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, cfg.Output, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadFile(cfg.Output)
	if err != nil {
		panic(err)
	}

	strbody := string(body)

	for _, d := range f.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if fn.Type.Results != nil && fn.Type.Params != nil {
				if len(fn.Type.Params.List) > 0 && len(fn.Type.Params.List[0].Names) > 0 && len(fn.Type.Results.List) > 0 && len(fn.Type.Results.List[0].Names) > 0 {

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
					fnParamMap[strings.Replace(fn.Name.Name, "Net", "", 1)] = strret
					strret = ""
					limtlen = len(fn.Type.Results.List) - 1
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
					fnReturnMap[strings.Replace(fn.Name.Name, "Net", "", 1)] = strret
				}
			}
		}
	}

	var strfuncs, strtemplate, jsstr, jstrbits string

	for _, v := range cfg.Templates.Templates {
		strfuncs += fmt.Sprintf(`} else if r.FormValue("name") == "%s" {
			w.Header().Set("Content-Type", "text/html")
			tmplRendered := Net%s(r.FormValue("payload"))
			w.Write([]byte(tmplRendered))
		`, v.Name, v.Name)

		jstrbits += fmt.Sprintf(`function %s(dataOfInterface, cb){ jsrequestmomentum("/momentum/templates", {name: "%s", payload: JSON.stringify(dataOfInterface)},"POST",  cb) }
`, v.Name, v.Name)

	}

	strtemplate = fmt.Sprintf(`http.HandleFunc("/momentum/templates", func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("name") == "reset" {
			return
		%s
		}
	})`, strfuncs)

	cfg.AddToMainFunc(strtemplate)

	strfuncs = ""
	for _, v := range cfg.Methods.Methods {

		if v.Man == "exp" {
			if fnParamMap[v.Name] == "" {
				fmt.Println("☣️  Error parsing :", v.Name, "Please make sure you are not using naked return values. Please use Named return values with your <func> tags")
			} else {
				varss := strings.Split(fnParamMap[v.Name], ",")
				responseformat := ``
				fnFormat := ``
				funcfields := []string{}
				binderString := ``
				jssetters := ``
				varsslen := len(varss) - 1
				for ind, variabl := range varss {
					varname := strings.Split(variabl, " ")
					fnFormat += fmt.Sprintf("tmvv.%s", strings.Title(varname[0]))
					jssetters += fmt.Sprintf(`
	t.%s = %s`, strings.Title(varname[0]), strings.Title(varname[0]))
					if ind < varsslen {
						fnFormat += ","
					}
					funcfields = append(funcfields, fmt.Sprintf("%s %s", strings.Title(varname[0]), varname[1]))
				}
				jstrbits += fmt.Sprintf(`function %s(%s, cb){
	var t = {}
	%s
	jsrequestmomentum("/momentum/funcs?name=%s", t, "POSTJSON", cb)
}
`, v.Name, strings.Replace(fnFormat, "tmvv.", "", -1), jssetters, v.Name)
				if !strings.Contains(v.Returntype, "(") {
					responseformat = fmt.Sprintf("resp[\"%s\"]", v.Returntype)
				} else {
					parsable := strings.Split(fnReturnMap[v.Name], ",")
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
				}

				strfuncs += fmt.Sprintf(`} else if r.FormValue("name") == "%s" {
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
			resp := db.O{}
			%s := Net%s(%s)
			%s
			w.Write([]byte(mResponse(resp)))
		`, v.Name, v.Name, strings.Join(funcfields, "\n"), v.Name, responseformat, v.Name, fnFormat, binderString)

			}
		}
	}
	strtemplate = fmt.Sprintf(`http.HandleFunc("/momentum/funcs", func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("name") == "reset" {
			return
		%s
		}
	})`, strfuncs)

	jsstr = fmt.Sprintf(`http.HandleFunc("/funcfactory.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/javascript")
		%s
	})`, fmt.Sprintf("w.Write([]byte(`%s`) )", jstrbits))

	cfg.AddToMainFunc(jsstr)
	cfg.AddToMainFunc(strtemplate)

}
