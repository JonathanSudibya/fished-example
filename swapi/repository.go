package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var swapiURL = "https://swapi.co/api/"
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Fetch ...
func Fetch(url string) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// FetchJSON ...
func FetchJSON(url string, body interface{}) error {
	b, err := Fetch(url)
	if err != nil {
		return err
	}

	json.Unmarshal(b, &body)
	return nil
}

// FetchPeople ...
func FetchPeople(i int) (*People, error) {
	p := &People{}
	URL := swapiURL + "people/" + strconv.Itoa(i) + "/"
	fmt.Println(URL)

	err := FetchJSON(URL, p)
	return p, err
}
