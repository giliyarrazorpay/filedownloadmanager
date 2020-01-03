package main

import (
	"fmt"
	"filedownloadmanager/server"
	"filedownloadmanager/service"
	"log"
	"net/http"
	)

func main() {
	theserver := &server.Server{}
	twirpHandler:= filedownloadmanager.NewFiledownloadmanagerServer(theserver,nil)
	fmt.Println("test")
	log.Fatal(http.ListenAndServe(":8080", twirpHandler))
}
