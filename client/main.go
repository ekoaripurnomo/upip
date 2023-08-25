package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/rdegges/go-ipify"
)

type Ipname struct {
	Ip       string
	Hostname string
}

// Get IP public
func getippub() string {
	ip, err := ipify.GetIp()
	if err != nil {
		fmt.Println(err.Error())
	}
	return ip
}

// Write IP Public to text
func writeip(ip string) {
	file, err := os.Create("clientip.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes := []byte(ip)
	file.Write(bytes)
}

func main() {
	var reply Ipname
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	ip := getippub()
	fmt.Println("Your IP public:", ip)
	writeip(ip)

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	iphostname := Ipname{ip, name}

	fmt.Println("IP computer:", iphostname.Ip)
	fmt.Println("Hostname:", iphostname.Hostname)

	client.Call("API.Writeiptotext", iphostname, &reply)
}
