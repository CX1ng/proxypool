package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/CX1ng/proxypool/common"
	. "github.com/CX1ng/proxypool/dao"
	. "github.com/CX1ng/proxypool/models"
)

type DBConnector struct {
	DB *sql.DB
}

func init() {
	StorageInitializer["mysql"] = NewMysqlConnector
}

func NewMysqlConnector() ImportExporter {
	return &DBConnector{
		DB: GetDBHandler(),
	}
}

func (d DBConnector) BulkInsertProxyIPs(ips []ProxyIP) error {
	length := len(ips)
	if length == 0 {
		return errors.New("IPs Is Empty")
	}
	keys := make([]string, 0, len(ips[0]))
	values := make([]interface{}, 0, len(ips[0])*len(ips))
	for key, _ := range ips[0] {
		keys = append(keys, key)
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	for _, ip := range ips {
		ip.Set("create_time", now)
		for _, key := range keys {
			values = append(values, ip[key])
		}
	}
	placeholderStr := "(?" + strings.Repeat(",?", len(keys)-1) + "),"

	tpl := "insert into ip_list(%s) values%s"
	_SQL := fmt.Sprintf(tpl, strings.Join(keys, ","), strings.Repeat(placeholderStr, len(ips)))
	_SQL = _SQL[0 : len(_SQL)-1]
	result, err := d.DB.Exec(_SQL, values...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != int64(len(ips)) {
		return errors.New("Affect Rows Not Equal To ip num")
	}
	return nil
}

func (d DBConnector) GetLimitProxyIP(limit int) ([]string, error) {
	var ip, _SQL string
	var port int
	var rows *sql.Rows
	var err error
	if limit < 0 || limit > common.GetLimit {
		return nil, common.ErrModelLimitInvalid
	} else if limit == 0 {
		_SQL = `select ip,port from ip_list `
		rows, err = d.DB.Query(_SQL)
	} else {
		_SQL = `select ip,port from ip_list limit ?`
		rows, err = d.DB.Query(_SQL, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var proxyList []string
	for rows.Next() {
		if err := rows.Scan(&ip, &port); err != nil {
			return nil, err
		}
		proxyList = append(proxyList, joinIPPort(ip, port))
	}
	return proxyList, nil
}

func joinIPPort(ip string, port int) string {
	return fmt.Sprintf("http://%s:%d", ip, port)
}

func Noop() {

}
