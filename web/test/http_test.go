package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetPic(t *testing.T) {

	for i := 0; i < 10; i++ {
		doHttp(i)
	}
}

func doHttp(i int) {
	url := fmt.Sprintf("http://127.0.0.1:9989/api/v1/getpic?page=%d", i)
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
