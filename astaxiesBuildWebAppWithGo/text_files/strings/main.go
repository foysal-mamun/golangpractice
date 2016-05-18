package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Foysal Mamun", "foy"))

	s := []string{"Foysal", "Sony", "Numa"}
	fmt.Println(strings.Join(s, " "))

	fmt.Println(strings.Index("chicken", "ic"))
	fmt.Println(strings.Index("chicken", "as"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))

	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}
