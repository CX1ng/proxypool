package processord

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
)

func BulkVerifyProxyIPs(infos []models.ProxyIP) []models.ProxyIP {
	ipList := make([]models.ProxyIP, 0)
	now := time.Now().Format("2006-01-02 15:04:05")
	for _, info := range infos {
		// TODO: 使用并发替换串行
		if ok := VerifyProxy(info["type"].(string), info["ip"].(string), strconv.Itoa(info["port"].(int))); !ok {
			continue
		}
		info.Set("last_verify_time", now)
		fmt.Printf("Proxy_ip:%+v\n", info)
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

	proxyUrl, err := url.Parse(strings.ToLower(schema) + "://" + net.JoinHostPort(ip, port))
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

	return verifyHttpResponse(resp)
}

func verifyHttpResponse(resp *http.Response) bool {
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
