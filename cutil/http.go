package cutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = http.Client{Timeout: 5 * time.Second}

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setAuth(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("request filter err")
		return nil, err
	}
	body := res.Body
	data, err := ioutil.ReadAll(body)
	defer func() {
		if body != nil {
			body.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Post(url string, params map[string]interface{}) ([]byte, error) {
	marshal, _ := json.Marshal(params)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(marshal))
	setAuth(req)
	res, _ := client.Do(req)
	body := res.Body
	data, err := ioutil.ReadAll(body)
	defer func() {
		if body != nil {
			body.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setAuth(req *http.Request) {
}
