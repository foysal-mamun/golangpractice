package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "os"
	"regexp"
	"strings"
)

func main() {

	// if len(os.Args) == 1 {
	// 	fmt.Println("Usage: regexp [string]")
	// } else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
	// 	fmt.Println("Number")
	// } else {
	// 	fmt.Println("Not number")
	// }

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	fmt.Println(strings.TrimSpace(src))
}
