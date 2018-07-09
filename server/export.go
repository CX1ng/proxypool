package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	. "github.com/CX1ng/proxypool/dao/mysql"
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
	db := DBConnector{DB: GetDBHandler()}
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
