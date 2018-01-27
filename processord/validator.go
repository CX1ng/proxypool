package processord

import (
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/CX1ng/proxypool/common"

	"github.com/PuerkitoBio/goquery"
)

const (
	TIME_OUT = 10
)

//现在只是单条验证，以后是不是可以搞成批量验证
func VerifyProxy(ip, port string) bool {
	request, err := http.NewRequest("Get", common.VerifyUrl, nil)
	if err != nil {
		return false
	}
	request.Header.Add("User-Agent", common.UserAgent)

	proxyUrl, err := url.Parse("http://" + net.JoinHostPort(ip, port))
	if err != nil {
		return false
	}

	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	client := http.Client{
		Transport: tr,
		Timeout:   time.Second * common.TIME_OUT,
	}

	resp, err := client.Do(request)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	dot, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return false
	}
	if dot.Find("title").Text() == common.TITLE {
		return true
	}
	return false
}
