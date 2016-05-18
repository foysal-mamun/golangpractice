package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"time"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}
func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	h := md5.New()
	io.WriteString(h, "Foysal")
	io.WriteString(h, strconv.FormatInt(time.Now().Unix(), 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(token)

	sh := sha256.New()
	io.WriteString(sh, "His money")
	fmt.Printf("% x\n", sh.Sum(nil))

	sh = sha1.New()
	io.WriteString(sh, "His money")
	fmt.Printf("% x\n", sh.Sum(nil))

	hello := "Foysalo"
	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)

	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))

	plainText := []byte("My name is Foysal")
	keyText := "key3456789012345"

	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyText), err)
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	fmt.Printf("%s=>%x\n", plainText, cipherText)

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plainTextCopy := make([]byte, len(plainText))
	cfbdec.XORKeyStream(plainTextCopy, cipherText)
	fmt.Printf("%x => %s\n", cipherText, plainTextCopy)
}
