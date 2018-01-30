package model

import (
	"database/sql"
	"fmt"
)

func InsertIP(db *sql.DB, info *ProxyIP) error {
	_SQL := `INSERT INTO ip_list(ip, port, type, origin, raw_time, region, capture_time, last_verify_time, create_time) VALUES (?,?,?,?,?,?,now(),now(),now())`
	_, err := db.Exec(_SQL, info.IP, info.Port, info.Type, info.Origin, info.RawTime, info.Region)
	if err != nil {
		return err
	}
	return nil
}

func ExportAll(db *sql.DB) ([]string, error) {
	var ip string
	var port int
	_SQL := `select ip,port from ip_list`
	rows, err := db.Query(_SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var proxyList []string
	for rows.Next() {
		rows.Scan(&ip, &port)
		proxyList = append(proxyList, fmt.Sprintf("http://%s:%d", ip, port))
	}
	return proxyList, nil
}
