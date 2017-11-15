package main

import (
	"github.com/cheikhshift/gos/core"
	"fmt"
	"strings"
)

func  main() {
	fmt.Println("Welcome to Momentum Afterbin.")
	fmt.Println("Remember to load momentum on your pages : {{ server }} or <script src=\"/funcfactory.js\"></script>")
	fmt.Println("Converting Templates and Func tags into accessible JS functions.")
	fmt.Println("Any web service service URI can be accessed with Ape. ")
	cfg,err := core.Config()
	if err != nil {
		panic(err)
	}
	//add template paths
	var strfuncs, strtemplate string

	for _, v := range cfg.Templates.Templates {
		strfuncs +=  fmt.Sprintf(`} else if r.FormValue("name") == "%s" {
			w.Header().Set("Content-Type", "text/html")
			tmplRendered := Net%s(r.FormValue("payload"))
			w.Write([]byte(tmplRendered))
		`, v.Name, v.Name)
	}

	strtemplate = fmt.Sprintf(`http.HandleFunc("/momentum/templates", func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("name") == "reset" {
			return
		%s
		}
	})`, strfuncs)
	
	cfg.AddToMainFunc(strtemplate)
	strtemplate = ""
	strfuncs = ""

	/*
			XMLName    xml.Name `xml:"method"`
			Method     string   `xml:",innerxml"`
			Name       string   `xml:"name,attr"`
			Variables  string   `xml:"var,attr"`
			Limit      string   `xml:"limit,attr"`
			Object     string   `xml:"object,attr"`
			Autoface   string   `xml:"autoface,attr"`
			Keeplocal  string   `xml:"keep-local,attr"`
			Testi      string   `xml:"testi,attr"`
			Testo      string   `xml:"testo,attr"`
			Man 	   string 	`xml:"m,attr"`
			Returntype string   `xml:"return,attr"`
	*/
	for _,v := range cfg.Methods.Methods {


			if v.Man == "exp" {
			varss := strings.Split(v.Variables, ",")
			responseformat := ``
			fnFormat := ``
			funcfields := []string{}
			binderString := ``
			varsslen := len(varss) - 1
			for ind, variabl := range varss {
				varname := strings.Split(variabl, " ")
				fnFormat += fmt.Sprintf("t.%s",strings.Title(varname[0]) )
				if ind < varsslen {
					fnFormat += ","
				}
				funcfields = append(funcfields, fmt.Sprintf("%s %s", strings.Title(varname[0]), varname[1] ))
			}
			if !strings.Contains( v.Returntype ,"(" ){
				responseformat  = fmt.Sprintf("resp[\"%s\"]", v.Returntype )
			} else {
				parsable := strings.Split(strings.Replace(strings.Replace(strings.Replace(v.Returntype, "(","", -1), ")", "" , -1), ", ",",", -1 ), ",")
				parsablelen := len(parsable) - 1
				for ind, variabl := range parsable {
				varname := strings.Split(variabl, " ")

				responseformat += fmt.Sprintf("resp%s%v",varname[0],ind)
				binderString += fmt.Sprintf(`
				resp["%s"] = resp%s%v`, varname[0], varname[0], ind )
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
			 var t Payload%s
			 err := decoder.Decode(&t)
			 if err != nil {
			 	w.WriteHeader(http.StatusInternalServerError)
			    w.Write([]byte(fmt.Sprintf("{\"error\":%%s}",err)))
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

	cfg.AddToMainFunc(strtemplate)

}