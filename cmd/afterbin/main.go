package main

import (
	"github.com/cheikhshift/gos/core"
	"fmt"
	"strings"
)

func  main() {
	fmt.Println("Welcome to Momentum Afterbin.")
	fmt.Println("Remember to load momentum on HTML your pages : {{ server }} or <script src=\"/funcfactory.js\"></script>")
	fmt.Println("Converting Templates and Func tags into accessible JS functions.")
	fmt.Println("Any web service service URI can be accessed with Ape. ")
	cfg,err := core.Config()
	if err != nil {
		panic(err)
	}
	//add template paths
	var strfuncs, strtemplate,jsstr,jstrbits string

	for _, v := range cfg.Templates.Templates {
		strfuncs +=  fmt.Sprintf(`} else if r.FormValue("name") == "%s" {
			w.Header().Set("Content-Type", "text/html")
			tmplRendered := Net%s(r.FormValue("payload"))
			w.Write([]byte(tmplRendered))
		`, v.Name, v.Name)

		jstrbits += fmt.Sprintf(`function %s(dataOfInterface, cb){ jsrequestmomentum("/momentum/templates", {name: "%s", payload: JSON.stringify(dataOfInterface)},"POST",  cb) }
`, v.Name,v.Name)
		
	}

	strtemplate = fmt.Sprintf(`http.HandleFunc("/momentum/templates", func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("name") == "reset" {
			return
		%s
		}
	})`, strfuncs)


	
	cfg.AddToMainFunc(strtemplate)
	
	strfuncs = ""
	for _,v := range cfg.Methods.Methods {


			if v.Man == "exp" {
			varss := strings.Split(v.Variables, ",")
			responseformat := ``
			fnFormat := ``
			funcfields := []string{}
			binderString := ``
			jssetters := ``
			varsslen := len(varss) - 1
			for ind, variabl := range varss {
				varname := strings.Split(variabl, " ")
				fnFormat += fmt.Sprintf("tmvv.%s",strings.Title(varname[0]) )
				jssetters += fmt.Sprintf(`
	t.%s = %s`,strings.Title(varname[0]),strings.Title(varname[0]) )
				if ind < varsslen {
					fnFormat += ","
				}
				funcfields = append(funcfields, fmt.Sprintf("%s %s", strings.Title(varname[0]), varname[1] ))
			}
			jstrbits += fmt.Sprintf(`function %s(%s, cb){
	var t = {}
	%s
	jsrequestmomentum("/momentum/funcs?name=%s", t, "POSTJSON", cb)
}
`, v.Name, strings.Replace(fnFormat, "tmvv.", "",  -1  ),jssetters, v.Name)
			if !strings.Contains( v.Returntype ,"(" ){
				responseformat  = fmt.Sprintf("resp[\"%s\"]", v.Returntype )
			} else {
				parsable := strings.Split(strings.Replace(strings.Replace(strings.Replace(v.Returntype, "(","", -1), ")", "" , -1), ", ",",", -1 ), ",")
				parsablelen := len(parsable) - 1
				for ind, variabl := range parsable {
				varname := strings.Split(variabl, " ")

				responseformat += fmt.Sprintf("resp%s%v",varname[0],ind)
				if strings.Contains(variabl, "error"){
				binderString += fmt.Sprintf(`
				if resp%s%v != nil {
					resp["%s"] = resp%s%v.Error()
				} else {
					resp["%s"] = nil
				}`,varname[0], ind, varname[0], varname[0], ind, varname[0] )
				} else {
				binderString += fmt.Sprintf(`
				resp["%s"] = resp%s%v`, varname[0], varname[0], ind )
				}
				if ind < parsablelen {
					responseformat += ","
				}
				}
			} 
			

			strfuncs +=  fmt.Sprintf(`} else if r.FormValue("name") == "%s" {
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
		`, v.Name, v.Name, strings.Join(funcfields, "\n" ) ,  v.Name, responseformat  ,v.Name, fnFormat,binderString)
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
	})`, fmt.Sprintf( "w.Write([]byte(`%s`) )", jstrbits)	)

	cfg.AddToMainFunc(jsstr)
	cfg.AddToMainFunc(strtemplate)

}