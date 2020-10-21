package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"time"
)

func main() {
	client, err := minio.NewV2("localhost:9000", "V0XJLY639O9PH2FS9LOW", "/fXAJXJqyADI2AATVcZVA+zL6f5yyIqXXqST31nr", false)
	if err != nil {
		panic(err)
	}

	url, err := client.PresignedGetObject("file-contents", "test123", time.Duration(60) * time.Second, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(url.String())
}