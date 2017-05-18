package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func send(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
		newsChan <- string(b)
	}
}
