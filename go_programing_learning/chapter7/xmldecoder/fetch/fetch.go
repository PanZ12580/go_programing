package fetch

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	b := resp.Body
	res, err := ioutil.ReadAll(b)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}