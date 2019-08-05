package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)
	fileWrite, err := bodyWrite.CreateFormFile("uploadfile", filename)
	if err != nil {
		return nil
	}
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()
	_, err = io.Copy(fileWrite, fh)
	contentType := bodyWrite.FormDataContentType()
	if err != nil {
		return err
	}
	bodyWrite.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil

}

func main() {
	target_url := "http://localhost:9090/upload"
	filename := "./src/day6/ti.go"
	postFile(filename, target_url)
}
