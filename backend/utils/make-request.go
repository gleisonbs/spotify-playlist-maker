package utils

import (
	"io"
	"net/http"
)

func MakeHttpRequest(accessToken, method, url string, body io.Reader) *http.Response {
	request, err := http.NewRequest(method, url, body)
	LogError(err)

	request.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	response, err := client.Do(request)
	LogError(err)

	return response
}
