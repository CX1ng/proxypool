package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/CX1ng/proxypool/common"
	. "github.com/CX1ng/proxypool/dao"
)

type IPList struct {
	IPs   []string `json:"iplist"`
	Count int      `json:"count"`
}

func getProxyIPWithLimit(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		w.Write([]byte("Error:" + err.Error()))
		return
	}
	fmt.Printf("limit:%d\n", limit)
	initializer, ok := StorageInitializer[strings.ToLower(common.GetConfigHandler().Storage)]
	if !ok {
		w.Write([]byte("Error: storage handler not init"))
		return
	}
	db := initializer()
	resp, err := db.GetLimitProxyIP(limit)
	if err != nil {
		w.Write([]byte("Error:" + err.Error()))
		return
	}
	ipList := IPList{
		IPs:   resp,
		Count: len(resp),
	}
	respJson, err := json.Marshal(ipList)
	if err != nil {
		w.Write([]byte("Error:" + err.Error()))
		return
	}
	w.Write(respJson)
}
