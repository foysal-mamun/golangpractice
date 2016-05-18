package main

import (
	"fmt"
	"os"
)

func DirAction() {
	os.Mkdir("dirAction", 0777)
	os.MkdirAll("dirAction/test1/test2", 0777)
	// err := os.Remove("dirAction")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err := os.RemoveAll("dirAction")
	if err != nil {
		fmt.Println(err)
	}
}

func CreateWriteFile() {
	userFile := "foysal.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()

	for i := 0; i < 5; i++ {
		fout.WriteString(string(i) + " just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

}

func ReadFile() {

	userFile := "foysal.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func main() {
	// DirAction()
	// CreateWriteFile()
	ReadFile()
}
