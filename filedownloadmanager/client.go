package main

import (
	"context"
	"fmt"
	"filedownloadmanager/service"
	"io/ioutil"
	"net/http"
)

func main(){
	client:= filedownloadmanager.NewFiledownloadmanagerProtobufClient("http://localhost:8080",&http.Client{})
	file,err:= client.GetFile(context.Background(),&filedownloadmanager.Filename{Name : "/Users/athreshgiliyar/go/src/filedownloadmanager/text.txt"})
	if err != nil {
		fmt.Println("error %v",err)
	}
	err=ioutil.WriteFile("/Users/athreshgiliyar/go/src/filedownloadmanager/clientfiles/check.txt",file.Content,0644)
	fmt.Println(err)
}