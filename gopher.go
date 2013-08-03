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
	v.Add("description", "Rating system for telecom providers with deadline. Longer description here!")
	query := base + "?" + v.Encode()
	log.Print(query)
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(string(body))
}
