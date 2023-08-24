package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/rdegges/go-ipify"
)

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
	var reply string
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	ip := getippub()
	fmt.Println("Your IP public:", ip)
	writeip(ip)

	// hostnm := fmt.Printf("computer_name %v", ip)

	client.Call("API.Writeiptotext", ip, &reply)
}
