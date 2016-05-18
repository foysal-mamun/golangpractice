package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers []Server
}

func ParseJson() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func ProduceJson() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(b))
}

func main() {
	ParseJson()
	ProduceJson()
}
