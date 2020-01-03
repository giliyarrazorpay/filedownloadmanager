# go-filedownloadmanager

run the mainserver.go using -> go run mainserver.go
  
 at client side u can run using curl replace the file path with name value
  curl --request "POST" --location "http://localhost:8080/twirp/filedownloadmanager.Filedownloadmanager/GetFile" --hta '{"name": "/Users/athreshgiliyar/go/src/filedownloadmanager/text.txt"}' --verbose

or change the filepath in client.go and run -> go run client.go
