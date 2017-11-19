package momentum

import (
	//iogos-replace
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cheikhshift/db"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/fatih/color"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"html"
	"html/template"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

var store = sessions.NewCookieStore([]byte("a very very very very secret key"))

type NoStruct struct {
	/* emptystruct */
}

func NetsessionGet(key string, s *sessions.Session) string {
	return s.Values[key].(string)
}

func UrlAtZ(url, base string) (isURL bool) {
	isURL = strings.Index(url, base) == 0
	return
}

func NetsessionDelete(s *sessions.Session) string {
	//keys := make([]string, len(s.Values))

	//i := 0
	for k := range s.Values {
		// keys[i] = k.(string)
		NetsessionRemove(k.(string), s)
		//i++
	}

	return ""
}

func NetsessionRemove(key string, s *sessions.Session) string {
	delete(s.Values, key)
	return ""
}
func NetsessionKey(key string, s *sessions.Session) bool {
	if _, ok := s.Values[key]; ok {
		//do something here
		return true
	}

	return false
}

func Netadd(x, v float64) float64 {
	return v + x
}

func Netsubs(x, v float64) float64 {
	return v - x
}

func Netmultiply(x, v float64) float64 {
	return v * x
}

func Netdivided(x, v float64) float64 {
	return v / x
}

func NetsessionGetInt(key string, s *sessions.Session) interface{} {
	return s.Values[key]
}

func NetsessionSet(key string, value string, s *sessions.Session) string {
	s.Values[key] = value
	return ""
}
func NetsessionSetInt(key string, value interface{}, s *sessions.Session) string {
	s.Values[key] = value
	return ""
}

func dbDummy() {
	smap := db.O{}
	smap["key"] = "set"
	log.Println(smap)
}

func Netimportcss(s string) string {
	return fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\" /> ", s)
}

func Netimportjs(s string) string {
	return fmt.Sprintf("<script type=\"text/javascript\" src=\"%s\" ></script> ", s)
}

func formval(s string, r *http.Request) string {
	return r.FormValue(s)
}

func renderTemplate(w http.ResponseWriter, p *Page) bool {
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path : web%s.tmpl reason : %s", p.R.URL.Path, n))

			DebugTemplate(w, p.R, fmt.Sprintf("web%s", p.R.URL.Path))
			w.WriteHeader(http.StatusInternalServerError)

			pag, err := loadPage("")

			if err != nil {
				log.Println(err.Error())
				return
			}

			if pag.isResource {
				w.Write(pag.Body)
			} else {
				pag.R = p.R
				pag.Session = p.Session
				renderTemplate(w, pag) //"

			}
		}
	}()

	t := template.New("PageWrapper")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
	t, _ = t.Parse(ReadyTemplate(p.Body))
	outp := new(bytes.Buffer)
	err := t.Execute(outp, p)
	if err != nil {
		log.Println(err.Error())
		DebugTemplate(w, p.R, fmt.Sprintf("web%s", p.R.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		pag, err := loadPage("")

		if err != nil {
			log.Println(err.Error())
			return false
		}
		pag.R = p.R
		pag.Session = p.Session
		p = nil
		if pag.isResource {
			w.Write(pag.Body)
		} else {
			renderTemplate(w, pag) // ""

		}
		return false
	}

	p.Session.Save(p.R, w)

	fmt.Fprintf(w, html.UnescapeString(outp.String()))
	p.Session = nil
	p.Body = nil
	p.R = nil
	p = nil
	return true

}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string, *sessions.Session)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var session *sessions.Session
		var er error
		if session, er = store.Get(r, "session-"); er != nil {
			session, _ = store.New(r, "session-")
		}
		if attmpt := apiAttempt(w, r, session); !attmpt {
			fn(w, r, "", session)
		} else {
			context.Clear(r)
		}

	}
}

func mResponse(v interface{}) string {
	data, _ := json.Marshal(&v)
	return string(data)
}
func apiAttempt(w http.ResponseWriter, r *http.Request, session *sessions.Session) (callmet bool) {
	var response string
	response = ""

	if callmet {
		session.Save(r, w)
		if response != "" {
			//Unmarshal json
			//w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
		}
		return
	}
	return
}
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}
func DebugTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
	lastline := 0
	linestring := ""
	defer func() {
		if n := recover(); n != nil {
			log.Println()
			// log.Println(n)
			log.Println("Error on line :", lastline+1, ":"+strings.TrimSpace(linestring))
			//http.Redirect(w,r,"",307)
		}
	}()

	p, err := loadPage(r.URL.Path)
	filename := tmpl + ".tmpl"
	body, err := Asset(filename)
	session, er := store.Get(r, "session-")

	if er != nil {
		session, er = store.New(r, "session-")
	}
	p.Session = session
	p.R = r
	if err != nil {
		log.Print(err)

	} else {

		lines := strings.Split(string(body), "\n")
		// log.Println( lines )
		linebuffer := ""
		waitend := false
		open := 0
		for i, line := range lines {

			processd := false

			if strings.Contains(line, "{{with") || strings.Contains(line, "{{ with") || strings.Contains(line, "with}}") || strings.Contains(line, "with }}") || strings.Contains(line, "{{range") || strings.Contains(line, "{{ range") || strings.Contains(line, "range }}") || strings.Contains(line, "range}}") || strings.Contains(line, "{{if") || strings.Contains(line, "{{ if") || strings.Contains(line, "if }}") || strings.Contains(line, "if}}") || strings.Contains(line, "{{block") || strings.Contains(line, "{{ block") || strings.Contains(line, "block }}") || strings.Contains(line, "block}}") {
				linebuffer += line
				waitend = true

				endstr := ""
				processd = true
				if !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}")) {

					open++

				}
				for i := 0; i < open; i++ {
					endstr += "\n{{end}}"
				}
				//exec
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
				t, _ = t.Parse(ReadyTemplate(body))
				lastline = i
				linestring = line
				erro := t.Execute(outp, p)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}
			}

			if waitend && !processd && !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end")) {
				linebuffer += line

				endstr := ""
				for i := 0; i < open; i++ {
					endstr += "\n{{end}}"
				}
				//exec
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
				t, _ = t.Parse(ReadyTemplate(body))
				lastline = i
				linestring = line
				erro := t.Execute(outp, p)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}

			}

			if !waitend && !processd {
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
				t, _ = t.Parse(ReadyTemplate(body))
				lastline = i
				linestring = line
				erro := t.Execute(outp, p)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}
			}

			if !processd && (strings.Contains(line, "{{end") || strings.Contains(line, "{{ end")) {
				open--

				if open == 0 {
					waitend = false

				}
			}
		}

	}

}

func DebugTemplatePath(tmpl string, intrf interface{}) {
	lastline := 0
	linestring := ""
	defer func() {
		if n := recover(); n != nil {

			log.Println("Error on line :", lastline+1, ":"+strings.TrimSpace(linestring))
			log.Println(n)
			//http.Redirect(w,r,"",307)
		}
	}()

	filename := tmpl
	body, err := Asset(filename)

	if err != nil {
		log.Print(err)

	} else {

		lines := strings.Split(string(body), "\n")
		// log.Println( lines )
		linebuffer := ""
		waitend := false
		open := 0
		for i, line := range lines {

			processd := false

			if strings.Contains(line, "{{with") || strings.Contains(line, "{{ with") || strings.Contains(line, "with}}") || strings.Contains(line, "with }}") || strings.Contains(line, "{{range") || strings.Contains(line, "{{ range") || strings.Contains(line, "range }}") || strings.Contains(line, "range}}") || strings.Contains(line, "{{if") || strings.Contains(line, "{{ if") || strings.Contains(line, "if }}") || strings.Contains(line, "if}}") || strings.Contains(line, "{{block") || strings.Contains(line, "{{ block") || strings.Contains(line, "block }}") || strings.Contains(line, "block}}") {
				linebuffer += line
				waitend = true

				endstr := ""
				if !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}")) {

					open++

				}

				for i := 0; i < open; i++ {
					endstr += "\n{{end}}"
				}
				//exec

				processd = true
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
				t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s", linebuffer, endstr))))
				lastline = i
				linestring = line
				erro := t.Execute(outp, intrf)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}
			}

			if waitend && !processd && !(strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}")) {
				linebuffer += line

				endstr := ""
				for i := 0; i < open; i++ {
					endstr += "\n{{end}}"
				}
				//exec
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
				t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s", linebuffer, endstr))))
				lastline = i
				linestring = line
				erro := t.Execute(outp, intrf)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}

			}

			if !waitend && !processd {
				outp := new(bytes.Buffer)
				t := template.New("PageWrapper")
				t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
				t, _ = t.Parse(ReadyTemplate([]byte(fmt.Sprintf("%s%s", linebuffer))))
				lastline = i
				linestring = line
				erro := t.Execute(outp, intrf)
				if erro != nil {
					log.Println("Error on line :", i+1, line, erro.Error())
				}
			}

			if !processd && (strings.Contains(line, "{{end") || strings.Contains(line, "{{ end") || strings.Contains(line, "end}}") || strings.Contains(line, "end }}")) {
				open--

				if open == 0 {
					waitend = false

				}
			}
		}

	}

}
func Handler(w http.ResponseWriter, r *http.Request, contxt string, session *sessions.Session) {
	var p *Page
	p, err := loadPage(r.URL.Path)

	if err != nil {
		log.Println(err.Error())

		w.WriteHeader(http.StatusNotFound)

		pag, err := loadPage("")

		if err != nil {
			log.Println(err.Error())
			//context.Clear(r)
			return
		}
		pag.R = r
		pag.Session = session
		if p != nil {
			p.Session = nil
			p.Body = nil
			p.R = nil
			p = nil
		}
		if pag.isResource {
			w.Write(pag.Body)
		} else {
			renderTemplate(w, pag) //""
		}
		context.Clear(r)
		return
	}

	if !p.isResource {
		w.Header().Set("Content-Type", "text/html")
		p.Session = session
		p.R = r
		renderTemplate(w, p) //fmt.Sprintf("web%s", r.URL.Path)

		// log.Println(w)
	} else {
		w.Header().Set("Cache-Control", "public")
		if strings.Contains(r.URL.Path, ".css") {
			w.Header().Add("Content-Type", "text/css")
		} else if strings.Contains(r.URL.Path, ".js") {
			w.Header().Add("Content-Type", "application/javascript")
		} else {
			w.Header().Add("Content-Type", http.DetectContentType(p.Body))
		}

		w.Write(p.Body)
	}

	p.Session = nil
	p.Body = nil
	p.R = nil
	p = nil
	context.Clear(r)
	return
}

func loadPage(title string) (*Page, error) {

	if roottitle := (title == "/"); roottitle {
		webbase := "web/"
		fname := fmt.Sprintf("%s%s", webbase, "index.html")
		body, err := Asset(fname)
		if err != nil {
			fname = fmt.Sprintf("%s%s", webbase, "index.tmpl")
			body, err = Asset(fname)
			if err != nil {
				return nil, err
			}
			return &Page{Body: body, isResource: false}, nil
		}

		return &Page{Body: body, isResource: true}, nil

	}

	filename := fmt.Sprintf("web%s.tmpl", title)

	if body, err := Asset(filename); err != nil {
		filename = fmt.Sprintf("web%s.html", title)

		if body, err = Asset(filename); err != nil {
			filename = fmt.Sprintf("web%s", title)

			if body, err = Asset(filename); err != nil {
				return nil, err
			} else {
				if strings.Contains(title, ".tmpl") {
					return nil, nil
				}
				return &Page{Body: body, isResource: true}, nil
			}
		} else {
			return &Page{Body: body, isResource: true}, nil
		}
	} else {
		return &Page{Body: body, isResource: false}, nil
	}

}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
func equalz(args ...interface{}) bool {
	if args[0] == args[1] {
		return true
	}
	return false
}
func nequalz(args ...interface{}) bool {
	if args[0] != args[1] {
		return true
	}
	return false
}

func netlt(x, v float64) bool {
	if x < v {
		return true
	}
	return false
}
func netgt(x, v float64) bool {
	if x > v {
		return true
	}
	return false
}
func netlte(x, v float64) bool {
	if x <= v {
		return true
	}
	return false
}

func GetLine(fname string, match string) int {
	intx := 0
	file, err := os.Open(fname)
	if err != nil {
		color.Red("Could not find a source file")
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intx = intx + 1
		if strings.Contains(scanner.Text(), match) {

			return intx
		}

	}

	return -1
}
func netgte(x, v float64) bool {
	if x >= v {
		return true
	}
	return false
}

type Page struct {
	Title      string
	Body       []byte
	isResource bool
	R          *http.Request
	Session    *sessions.Session
}

func ReadyTemplate(body []byte) string {
	return strings.Replace(strings.Replace(strings.Replace(string(body), "/{", "\"{", -1), "}/", "}\"", -1), "`", "\"", -1)
}

func NetLoadWebAsset(args ...interface{}) string {

	data, err := Asset(fmt.Sprintf("web%s", args[0].(string)))
	if err != nil {
		return err.Error()
	}
	return string(data)

}

func Netang(args ...interface{}) string {

	var d NoStruct
	filename := "tmpl/ang.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (ang) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = NoStruct{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("ang")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bang(d NoStruct) string {
	return Netbang(d)
}

func Netbang(d NoStruct) string {

	filename := "tmpl/ang.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("ang")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (ang) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
}
func Netcang(args ...interface{}) (d NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = NoStruct{}
	}
	return
}

func cang(args ...interface{}) (d NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = NoStruct{}
	}
	return
}

func Bang(intstr interface{}) string {
	return Netang(intstr)
}

func Netserver(args ...interface{}) string {

	var d NoStruct
	filename := "tmpl/server.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (server) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = NoStruct{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("server")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bserver(d NoStruct) string {
	return Netbserver(d)
}

func Netbserver(d NoStruct) string {

	filename := "tmpl/server.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("server")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (server) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
}
func Netcserver(args ...interface{}) (d NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = NoStruct{}
	}
	return
}

func cserver(args ...interface{}) (d NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = NoStruct{}
	}
	return
}

func Bserver(intstr interface{}) string {
	return Netserver(intstr)
}

func Netjquery(args ...interface{}) string {

	var d NoStruct
	filename := "tmpl/jquery.tmpl"
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (jquery) : %s", filename))
			// log.Println(n)
			DebugTemplatePath(filename, &d)
			//http.Redirect(w,r,"",307)
		}
	}()
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = NoStruct{}
	}

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("jquery")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
	t, _ = t.Parse(ReadyTemplate(body))

	erro := t.Execute(output, &d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", filename))
		DebugTemplatePath(filename, &d)
	}
	return html.UnescapeString(output.String())

}
func bjquery(d NoStruct) string {
	return Netbjquery(d)
}

func Netbjquery(d NoStruct) string {

	filename := "tmpl/jquery.tmpl"

	body, er := Asset(filename)
	if er != nil {
		return ""
	}
	output := new(bytes.Buffer)
	t := template.New("jquery")
	t = t.Funcs(template.FuncMap{"a": Netadd, "s": Netsubs, "m": Netmultiply, "d": Netdivided, "js": Netimportjs, "css": Netimportcss, "sd": NetsessionDelete, "sr": NetsessionRemove, "sc": NetsessionKey, "ss": NetsessionSet, "sso": NetsessionSetInt, "sgo": NetsessionGetInt, "sg": NetsessionGet, "form": formval, "eq": equalz, "neq": nequalz, "lte": netlt, "LoadWebAsset": NetLoadWebAsset, "ang": Netang, "bang": Netbang, "cang": Netcang, "server": Netserver, "bserver": Netbserver, "cserver": Netcserver, "jquery": Netjquery, "bjquery": Netbjquery, "cjquery": Netcjquery})
	t, _ = t.Parse(ReadyTemplate(body))
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path (jquery) : %s", filename))
			DebugTemplatePath(filename, &d)
		}
	}()
	erro := t.Execute(output, &d)
	if erro != nil {
		log.Println(erro)
	}
	return html.UnescapeString(output.String())
}
func Netcjquery(args ...interface{}) (d NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = NoStruct{}
	}
	return
}

func cjquery(args ...interface{}) (d NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = NoStruct{}
	}
	return
}

func Bjquery(intstr interface{}) string {
	return Netjquery(intstr)
}

func dummy_timer() {
	dg := time.Second * 5
	log.Println(dg)
}
func FileServer() http.Handler {
	return http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "web"})
}
