package util

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
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

	if response.StatusCode >= 400 {
		return "", errors.New("Unexpected response (HTTP STATUS CODE " + strconv.Itoa(response.StatusCode) + ") " + string(body))
	}

	return string(body), nil
}
