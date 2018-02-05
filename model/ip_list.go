package model

import (
	"database/sql"
	"fmt"

	"github.com/CX1ng/proxypool/common"
)

func InsertIP(db *sql.DB, info *ProxyIP) error {
	_SQL := `INSERT INTO ip_list(ip, port, type, origin, raw_time, region, capture_time, last_verify_time, create_time) VALUES (?,?,?,?,?,?,now(),now(),now())`
	_, err := db.Exec(_SQL, info.IP, info.Port, info.Type, info.Origin, info.RawTime, info.Region)
	if err != nil {
		return err
	}
	return nil
}

func GetLimitProxyIP(db *sql.DB, limit int) ([]string, error) {
	var ip, _SQL string
	var port int
	var rows *sql.Rows
	var err error
	if limit < 0 || limit > common.GetLimit {
		return nil, common.ModelLimitInvalid
	} else if limit == 0 {
		_SQL = `select ip,port from ip_list `
		rows, err = db.Query(_SQL)
	} else {
		_SQL = `select ip,port from ip_list limit ?`
		rows, err = db.Query(_SQL, limit)
	}
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
