package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

const (
	listen  = ":7090"
	ProxyIP = "proxy_ip"
)

const (
	OK               = 0
	NotFoundProxyIP  = 100000
	SplitHostPortErr = 100001
)

type Resp struct {
	Code int    `json:"code"`
	Body string `json:body`
}

func RenderJson(w http.ResponseWriter, code int, body string) {
	resp := Resp{
		Code: code,
		Body: body,
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Json Marshal Err" + err.Error()))
	}
	w.Write(respBytes)
}

func JudgeIPAnonymous(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	if ip, ok := params[ProxyIP]; ok {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			RenderJson(w, SplitHostPortErr, err.Error())
		}
		if host == "::1" {
			host = "127.0.0.1"
		}
		fmt.Printf("host:%v ip:%v\n", host, ip[0])
		if host == ip[0] {
			RenderJson(w, OK, "true")
		} else {
			RenderJson(w, OK, "false")
		}
		return
	}
	RenderJson(w, NotFoundProxyIP, "Not Found Param With "+ProxyIP)
}

func main() {
	fmt.Println("listen in " + listen)
	http.HandleFunc("/anonymous", JudgeIPAnonymous)
	if err := http.ListenAndServe(listen, nil); err != nil {
		panic(err)
	}
}
