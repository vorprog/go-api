package util

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
)

func GetURL(URL string) (result string) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	response, httpGetError := http.Get(URL)
	if httpGetError != nil {
		os.Exit(1)
	}

	defer response.Body.Close()
	body, responseBodyError := ioutil.ReadAll(response.Body)
	if responseBodyError != nil {
		os.Exit(1)
	}

	return string(body)
}
