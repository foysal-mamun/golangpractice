package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type RecurlyServers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server1 struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

type RecurlyServers1 struct {
	XMLName     xml.Name  `xml:"servers"`
	Version     string    `xml:"version,attr"`
	Svs         []server1 `xml:"server"`
	Description string    `xml:",innerxml"`
}

func ReadXmlFile() {

	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	v := RecurlyServers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

func ProduceXML() {
	v := &RecurlyServers1{Version: "1", Description: "Foysal"}
	v.Svs = append(v.Svs, server1{"shanghai_vpn", "127.1"})
	v.Svs = append(v.Svs, server1{"Dhaka_vpn", "127.2"})

	output, err := xml.MarshalIndent(v, " ", "	")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func main() {
	ReadXmlFile()
	ProduceXML()
}
