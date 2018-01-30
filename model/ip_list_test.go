package model

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/db"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_IP      = "1.1.1.1"
	TEST_TYPE    = "HTTP"
	TEST_REGION  = "北京"
	TEST_ORIGIN  = "test"
	TEST_RAWTIME = "18-01-29 18:27"
)

const (
	TEST_CONFIGPATH = "../config/config.dev.toml"
)

func init() {
	common.InitConfig(TEST_CONFIGPATH)
	db.InitMysql(common.Config.Mysql)
}

func testServer_InsertIP() (int, error) {
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	port := rd.Intn(65535)
	err := InsertIP(db.Mysql, &ProxyIP{
		IP:      TEST_IP,
		Port:    port,
		Type:    TEST_TYPE,
		Origin:  TEST_ORIGIN,
		RawTime: TEST_RAWTIME,
		Region:  TEST_REGION,
	})
	return port, err
}

func TestServer_InsertIP(t *testing.T) {
	_assert := assert.New(t)
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	port := rd.Intn(65535)
	_, err := testServer_InsertIP()
	if err != nil {
		t.Errorf("Error:%v\n", err)
	}
	_SQL := `select ip,port,type,region,raw_time,origin from ip_list where ip=? and port=?`
	row := db.Mysql.QueryRow(_SQL, TEST_IP, port)
	var resIP, resType, resRegion, resRawTime, resOrigin string
	var resPort int
	err = row.Scan(&resIP, &resPort, &resType, &resRegion, &resRawTime, &resOrigin)
	_assert.Equal(err, nil)
	_assert.Equal(resIP, TEST_IP)
	_assert.Equal(resPort, port)
	_assert.Equal(resType, TEST_TYPE)
	_assert.Equal(resRegion, TEST_REGION)
	_assert.Equal(resRawTime, TEST_RAWTIME)
	_assert.Equal(resOrigin, TEST_ORIGIN)
}

func TestServer_ExportAll(t *testing.T) {
	_assert := assert.New(t)
	_SQL := `delete from ip_list`
	_, err := db.Mysql.Exec(_SQL)
	if err != nil {
		t.Fatalf("Error:%v\n", err)
	}
	port1, err := testServer_InsertIP()
	if err != nil {
		t.Fatalf("Error:%v\n", err)
	}
	time.Sleep(1 * time.Second)
	port2, err := testServer_InsertIP()
	if err != nil {
		t.Fatalf("Error:%v\n", err)
	}
	time.Sleep(1 * time.Second)
	port3, err := testServer_InsertIP()
	if err != nil {
		t.Fatalf("Error:%v\n", err)
	}

	resp, err := ExportAll(db.Mysql)
	_assert.Equal(err, nil)
	_assert.Contains(resp, fmt.Sprintf("http://%s:%d", TEST_IP, port1))
	_assert.Contains(resp, fmt.Sprintf("http://%s:%d", TEST_IP, port2))
	_assert.Contains(resp, fmt.Sprintf("http://%s:%d", TEST_IP, port3))
}
