package processord

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
)

func BulkVerifyProxyIPs(insertChan chan models.ProxyIP, infos []models.ProxyIP) {
	var verifyGroup sync.WaitGroup
	verifyGroup.Add(len(infos))
	for _, info := range infos {
		go VerifyProxy(insertChan, &verifyGroup, info)
	}
	verifyGroup.Wait()
}

func VerifyProxy(insertChan chan models.ProxyIP, verifyGroup *sync.WaitGroup, info models.ProxyIP) {
	defer verifyGroup.Done()
	schema, _ := info.Type()
	ip, _ := info.IP()
	port, _ := info.Port()

	request, err := http.NewRequest("Get", common.VerifyUrl, nil)
	if err != nil {
		return
	}
	request.Header.Add("User-Agent", common.UserAgent)

	proxyUrl, err := url.Parse(strings.ToLower(schema) + "://" + net.JoinHostPort(ip, port))
	if err != nil {
		return
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
		return
	}
	defer resp.Body.Close()

	if verifyHttpResponse(resp) {
		info.Set("last_verify_time", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Printf("insert info:%+v\n", info)
		insertChan <- info
	}
}

func verifyHttpResponse(resp *http.Response) bool {
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
