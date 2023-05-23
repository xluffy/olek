package main

import (
	"context"
	"fmt"
	"net/http"
	"log"
	"flag"
)

var (
	server         *http.Server
	addrPort       string
	killSwitch     chan bool
)

func getClientIpAddr(req *http.Request) string {
	clientIp := req.Header.Get("X-FORWARDED-FOR")
	if clientIp != "" {
		return clientIp
	}
	return req.RemoteAddr
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := "ok"
	fmt.Fprintf(w, data+"\n\n")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	data := "pong"
	fmt.Fprintf(w, data+"\n\n")
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	clientIp := getClientIpAddr(r)

	fmt.Fprintf(w, clientIp+"\n\n")
}

func main() {
	killSwitch = make(chan bool, 1)

	flag.StringVar(&addrPort, "s", ":6969", "Server addr:port")
	flag.Parse()

	mainCtx := context.Background()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/ip", ipHandler)

	server = &http.Server{Addr: addrPort, Handler: nil}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	<-killSwitch
	_ = server.Shutdown(mainCtx)
}
