package main

import (
	"log"
	"sync"
	"net/http"
)

func startupGameServer()  {
	
}

func startupWebServer() {
	port := ":12307"
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	log.Printf("Web service addr %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() { defer wg.Done(); startupWebServer() }()  // 开启web服务器
	wg.Wait()
}