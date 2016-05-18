package main

import (
	"fmt"
	"html/template"
	_ "net/http"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func StringTemplate() {

	t := template.New("fieldname example")
	t, _ = t.Parse("Hello, {{.UserName}}!")
	p := Person{UserName: "Foysal"}
	t.Execute(os.Stdout, p)
	// t.Execute(w, p)
}

func NestedFileds() {
	f1 := Friend{Fname: "Sony"}
	f2 := Friend{Fname: "Numa"}
	t := template.New("Nested Fields Example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith, "toUpper": ToUpper})
	t, _ = t.Parse(`
		Hello {{.UserName | toUpper}}!
		{{range .Emails}}
			an email {{. | emailDeal}}
		{{end}}

		{{with .Friends}}
		{{range .}}
			My Friend name is {{.Fname}}
		{{end}}
		{{end}}
	`)

	p := Person{UserName: "Foysal",
		Emails:  []string{"f@y.c", "f@y.c"},
		Friends: []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}

func IFElse() {
	t := template.New("If")
	t = template.Must(t.Parse("Empty pipeline if demo: {{if ``}} will not {{end}}\n"))
	t.Execute(os.Stdout, nil)

	t1 := template.New("If Else")
	t1 = template.Must(t1.Parse("if else demo: {{if `a`}}if{{else}}else{{end}}"))
	t1.Execute(os.Stdout, nil)
}

func ToUpper(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	return strings.ToUpper(s)
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	return (substrs[0] + " at " + substrs[1])
}

func NestedTemplate() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}

func main() {

	// StringTemplate()
	// NestedFileds()
	// IFElse()
	NestedTemplate()

	// http.HandleFunc("/st", StringTemplate)

	// err := http.ListenAndServe(":9090", nil)
	// if err != nil {
	// 	fmt.Println("Server linster error: ", err)
	// }
}
