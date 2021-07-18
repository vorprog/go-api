package util

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

func GetURL(URL string) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	response, httpGetError := http.Get(URL)
	if httpGetError != nil {
		log.Fatal(httpGetError)
	}

	defer response.Body.Close()
	body, responseBodyError := ioutil.ReadAll(response.Body)
	if responseBodyError != nil {
		log.Fatal(responseBodyError)
	}

	return string(body)
}
