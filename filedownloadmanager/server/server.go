package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/twitchtv/twirp"
	"filedownloadmanager/service"
	"io"
	"os"
)

type Server struct{}

func (s *Server) GetFile(ctx context.Context, filename *filedownloadmanager.Filename) (file *filedownloadmanager.File, err error){
	fileContent, err := os.Open(filename.Name)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
		return &filedownloadmanager.File{
			Content: nil,
		},twirp.NotFoundError("")
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, fileContent)
	return &filedownloadmanager.File{
		Content: buf.Bytes(),
	}, nil
}

