package test

import (
	"io"
	"log"
	"net/http"
)

func PatchRequest(url string, data io.Reader) (*http.Response, error)  {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url, data)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp , err
}
func DeleteRequest(url string) (*http.Response, error)  {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp , err
}
