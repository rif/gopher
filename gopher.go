package main

import (
	//"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	base := "http://localhost:8080/api/add"
	v := url.Values{}
	v.Set("name", "cgrates")
	v.Add("repo", "github.com/cgrates/cgrates")
	v.Add("description", "A rating system")
	v.Add("friend", "Zoe")
	resp, err := http.Get(base + "?" + v.Encode())
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(string(body))
}
