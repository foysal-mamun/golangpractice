package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " "))
	}

	fmt.Println(w, "Hello Foysal!")
}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {

		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		fmt.Println("login GET")
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()

		username := r.Form.Get("username")
		password := r.Form.Get("password")
		age := r.Form.Get("age")
		email := r.Form.Get("email")
		gender := r.Form.Get("gender")

		if len(username) == 0 {
			fmt.Println("UserName required")
		}
		_, err := strconv.Atoi(age)
		if err != nil {
			fmt.Println("Age not a number")
		}

		if m, _ := regexp.MatchString(`^([\w\.\_]{2, 10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
			fmt.Println("Invalid Email")
		}

		genders := []string{"male", "female"}
		for _, v := range genders {
			if v == gender {
				fmt.Println("gender found")
			}
		}

		fmt.Println("UserName: ", template.HTMLEscapeString(username))
		fmt.Println("Password: ", template.HTMLEscapeString(password))
		template.HTMLEscape(w, []byte(username))
	}

}

func upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("fileupload.gtpl")
		t.Execute(w, token)

	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(""+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
