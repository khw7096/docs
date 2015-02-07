package main

import (
	"fmt"
	"flag"
	"os"
	"net"
	"net/http"
	)

func ispath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func LocalIP() string {
	var iplist []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
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

func docserver(docpath string, port string) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(docpath))))
	fmt.Printf("Web Server Start : http://%s%s\n", LocalIP(), port)
	http.ListenAndServe(port, nil)
}

func main() {
	docpathPtr := flag.String("docpath", "", "docpath")
	portPtr := flag.String("server", "", "service port ex):8080")
	flag.Parse()

	if *docpathPtr != "" && *portPtr != "" {
		if ispath(*docpathPtr) == true {
			docserver(*docpathPtr, *portPtr)
		} else {
			fmt.Println("Target path do not exist")
			os.Exit(1)
		}
	} else {
		fmt.Println("Docs is simple doc server.")
		fmt.Println("Copyright (C) 2015  kimhanwoong")
		flag.PrintDefaults()
	}
}
