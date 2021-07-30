package server

func OK() (httpStatusCode int, responseBodyMessage string) {
	return 200, "HTTP 200 - OK"
}

func NotFound() (httpStatusCode int, responseBodyMessage string) {
	return 404, "HTTP 404 - Not Found"
}

func InternalServerError() (httpStatusCode int, responseBodyMessage string) {
	return 500, "HTTP 500 - Internal Server Error"
}
