package models

import (
	"errors"
)

var proxyIPFields = []string{"ip", "port", "type", "origin", "raw_time", "region", "capture_time", "last_verify_time", "create_time"}

type ProxyIP map[string]interface{}

func NewProxyIP() ProxyIP {
	proxyIPMap := make(map[string]interface{})
	for _, name := range proxyIPFields {
		proxyIPMap[name] = nil
	}

	return proxyIPMap
}

func (p ProxyIP) GetFields() []string {
	keys := make([]string, 0, len(p))
	for key, _ := range p {
		keys = append(keys, key)
	}
	return keys
}

func (p ProxyIP) FieldsNum() int {
	return len(p)
}

func (p ProxyIP) Set(name string, value interface{}) {
	p[name] = value
}

func (p ProxyIP) IP() (string, error) {
	ip, ok := p["ip"].(string)
	if !ok {
		return "", errors.New("Get IP Failed")
	}
	return ip, nil
}

func (p ProxyIP) Port() (int64, error) {
	port, ok := p["ip"].(float64)
	if !ok {
		return 0, errors.New("Get Port Failed")
	}
	return int64(port), nil
}

func (p ProxyIP) Type() (string, error) {
	typeValue, ok := p["type"].(string)
	if !ok {
		return "", errors.New("Get Type Failed")
	}
	return typeValue, nil
}

func (p ProxyIP) Region() (string, error) {
	region, ok := p["region"].(string)
	if !ok {
		return "", errors.New("Get Region Failed")
	}
	return region, nil
}

func (p ProxyIP) Origion() (string, error) {
	origin, ok := p["origin"].(string)
	if !ok {
		return "", errors.New("Get Origin Failed")
	}
	return origin, nil
}

func (p ProxyIP) RawTime() (string, error) {
	rawTime, ok := p["raw_time"].(string)
	if !ok {
		return "", errors.New("Get RawTime Failed")
	}
	return rawTime, nil
}

func (p ProxyIP) CaptureTime() (string, error) {
	captureTime, ok := p["capture_time"].(string)
	if !ok {
		return "", errors.New("Get CaptureTime Failed")
	}
	return captureTime, nil
}

func (p ProxyIP) CreateTime() (string, error) {
	createTime, ok := p["create_time"].(string)
	if !ok {
		return "", errors.New("Get CreateTime Failed")
	}
	return createTime, nil
}

func (p ProxyIP) LastVerifyTime() (string, error) {
	lastVerifyTime, ok := p["last_verify_time"].(string)
	if !ok {
		return "", errors.New("Get LastVerifyTime Failed")
	}
	return lastVerifyTime, nil
}
