package dao

import (
	. "github.com/CX1ng/proxypool/models"
)

type Importer interface {
	BulkInsertProxyIPs(ips []ProxyIP) error
}

type Exporter interface {
	GetLimitProxyIP(limit int) ([]string, error)
}

type ImportExporter interface {
	Importer
	Exporter
}

type initializer func() ImportExporter

var StorageInitializer = make(map[string]initializer)
