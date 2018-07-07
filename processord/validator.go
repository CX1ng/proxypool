package processord

import (
	"net"
	"net/http"
	"net/url"
	"time"
	"strconv"
	"github.com/CX1ng/proxypool/common"

	"github.com/PuerkitoBio/goquery"
	"github.com/CX1ng/proxypool/models"
	"strings"
)

func BulkVerifyProxyIPs(infos []models.ProxyIP) []models.ProxyIP{
	ipList := make([]models.ProxyIP,0)
	now := time.Now().Format("2006-01-02 15:04:05")
	for _,info := range infos {
		if ok := VerifyProxy(info["type"].(string), info["ip"].(string), strconv.Itoa(info["port"].(int))); !ok {
			//fmt.Printf("Verify Failed:%+v\n", info)
			continue
		}
		info.Set("last_verify_time", now)
		//fmt.Printf("Proxy_ip:%+v\n", info)
		ipList = append(ipList, info)
	}
	return ipList
}

func VerifyProxy(schema, ip, port string) bool {
	request, err := http.NewRequest("Get", common.VerifyUrl, nil)
	if err != nil {
		return false
	}
	request.Header.Add("User-Agent", common.UserAgent)

	proxyUrl, err := url.Parse(strings.ToLower(schema) +"://" + net.JoinHostPort(ip, port))
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
