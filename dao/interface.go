package dao

import (
	. "github.com/CX1ng/proxypool/models"
)

type Import interface {
	BulkInsertProxyIPs(ips []ProxyIP) error
}

type Export interface {
	GetLimitProxyIP(limit int) ([]string, error)
}
