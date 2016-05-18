package main

import (
	_ "fmt"
	"html/template"
	"log"
	"net/http"
	_ "os"
	"strings"
)

func sayi18n(rw http.ResponseWriter, req *http.Request) {
	t := template.New("i18n")
	t, err := t.Parse(`
		Host: {{.Host}}
		Port: {{.Port}}
		Al: {{.Al}}
		Pea: {{.Pea}}
		Bean: {{.Bean}}
	`)
	if err != nil {
		log.Fatal("Parse error: ", err)
	}

	domain := strings.Split(req.Host, ":")

	t.Execute(rw, struct {
		Host string
		Port string
		Al   string
		Pea  string
		Bean string
	}{
		req.Host,
		domain[1],
		req.Header.Get("Accept-Language"),
		msg("bd", "pea"),
		msg("bd", "bean"),
	})
}

var locales map[string]map[string]string

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}

	return ""
}
func initLang() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	locales["en"] = en

	bd := make(map[string]string, 10)
	bd["pea"] = "bdpea"
	bd["bean"] = "bdbean"
	locales["bd"] = bd
}

func main() {

	initLang()

	http.HandleFunc("/i18n", sayi18n)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listen and serve: ", err)
	}
}
