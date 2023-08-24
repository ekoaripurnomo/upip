package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type API int

// name func must be Capital first
func (a *API) Writeiptotext(ip string, reply *string) error {
	file, err := os.Create("clientip.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes := []byte(ip)
	file.Write(bytes)
	*reply = ip
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
