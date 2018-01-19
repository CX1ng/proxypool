package model

import (
	"database/sql"
)

func InsertIP(db *sql.DB, info *ProxyIP) error {
	_SQL := `INSERT INTO ip_list(ip, port, type, origin, raw_time, region, capture_time, last_verify_time, create_time) VALUES (?,?,?,?,?,?,now(),now(),now())`
	_, err := db.Exec(_SQL, info.IP, info.Port, info.Type, info.Origin, info.RawTime, info.Region)
	if err != nil {
		return err
	}
	return nil
}
