package main

import (
	"fmt"
	"flag"
	"os"
	"net"
	"net/http"
	"log"
	)

func localIP() string {
	var iplist []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				iplist = append(iplist, ipnet.IP.String())
			}
		}
	}
	return iplist[0]
}

func docServer(docpath string, port string) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(docpath))))
	fmt.Printf("Web Server Start : http://%s%s\n", localIP(), port)
	http.ListenAndServe(port, nil)
}

func main() {
	pathPtr := flag.String("path", "", "docpath")
	portPtr := flag.String("http", "", "service port ex):8080")
	flag.Parse()
	if *pathPtr == "" || *portPtr == "" {
		fmt.Println("Docs is simple doc server.")
		fmt.Println("Copyright (C) 2015  kimhanwoong")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if _, err := os.Stat(*pathPtr); err != nil {
		log.Fatal("Target path do not exist")
	}
	docServer(*pathPtr, *portPtr)
}
