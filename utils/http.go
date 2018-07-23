package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/CX1ng/proxypool/common"
)

func HttpRequestWithUserAgent(url string) (string, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("User-Agent", common.UserAgent)
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(html), nil
}
