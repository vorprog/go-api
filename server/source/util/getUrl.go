package util

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func GetURL(URL string) (result string, err error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	response, httpGetError := http.Get(URL)
	if httpGetError != nil {
		return "", httpGetError
	}

	defer response.Body.Close()
	body, responseBodyError := ioutil.ReadAll(response.Body)
	if responseBodyError != nil {
		return "", responseBodyError
	}

	return string(body), nil
}
