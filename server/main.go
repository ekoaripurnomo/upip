package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"

	"github.com/FlowerWrong/go-hostsfile"
	"github.com/txn2/txeh"
)

type Ipname struct {
	Ip       string
	Hostname string
}

type API int

// name func must be Capital first
func (a *API) Writeiptotext(iphostname Ipname, reply *Ipname) error {

	subhostname := strings.ToLower(iphostname.Hostname)

	resiphost, err := hostsfile.Lookup(subhostname)
	if err != nil {
		panic(err)
	}

	fmt.Println("hostname sebelum :", subhostname)
	fmt.Println("ip sebelum :", resiphost)
	fmt.Println("ip terbaru:", iphostname.Ip)

	if resiphost == "" {
		fmt.Println("IP and Hostname not in list hosts file")
		fmt.Println("Proccessing add new line IP and Hostname")
		addhost(iphostname.Ip, iphostname.Hostname)
		fmt.Println("IP computer:", iphostname.Ip)
		fmt.Println("Hostname:", subhostname)
		fmt.Println("----------------")
	} else if resiphost != iphostname.Ip {
		fmt.Println("IP beda")
		fmt.Println("Proccessing remove IP and Hostname")
		removehost(resiphost)
		fmt.Println("Proccessing add new line IP and Hostname")
		addhost(iphostname.Ip, iphostname.Hostname)
		fmt.Println("IP computer:", iphostname.Ip)
		fmt.Println("Hostname:", subhostname)
		fmt.Println("----------------")
	} else {
		fmt.Println("IP is same, nothing to changes")
		fmt.Println("IP computer:", iphostname.Ip)
		fmt.Println("Hostname:", subhostname)
		fmt.Println("----------------")
	}

	return nil
}

func removehost(resip string) {
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ip will be remove: %v\n", resip)

	hosts.RemoveAddress(resip)
	hosts.Save()

	fmt.Printf("Success remove IP Hostname for hosts file\n")

}

func addhost(resip string, reshost string) {
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ip will be add: %v\n", resip)

	hosts.AddHost(resip, reshost)
	hosts.Save()

	fmt.Printf("Success add IP Hostname for hosts file\n")
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
