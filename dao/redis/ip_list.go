package redis

import (
	"encoding/json"
	"net"
	"github.com/fzzy/radix/redis"

	"github.com/CX1ng/proxypool/common"
	. "github.com/CX1ng/proxypool/dao"
	. "github.com/CX1ng/proxypool/models"
)

const (
	ipStorageKey     = "ip_list"
	ipAddrStorageKey = "ip_list_key"
)

func init() {
	StorageInitializer["redis"] = NewRedisConnector
}

type RedisConnector struct {
	conn *redis.Client
}

func NewRedisConnector() ImportExporter {
	return &RedisConnector{
		conn: GetRedisHandler(),
	}
}

func (r RedisConnector) BulkInsertProxyIPs(ips []ProxyIP) error {
	err := r.conn.Cmd("MULTI").Err
	if err != nil {
		return err
	}
	defer r.conn.Cmd("DISCARD")
	if err := r.conn.Cmd("select", common.GetConfigHandler().Redis.Db).Err; err != nil {
		return err
	}
	for _, ip := range ips {
		msg, err := json.Marshal(ip)
		if err != nil {
			return err
		}
		ip, err := ip.IP()
		if err != nil {
			return err
		}
		if err := r.conn.Cmd("HSET", ipStorageKey, ip, msg).Err; err != nil {
			return err
		}
		if err := r.conn.Cmd("SADD", ipAddrStorageKey, ip).Err; err != nil {
			return err
		}
	}
	return r.conn.Cmd("EXEC").Err
}

// TODO： 增加查询IP信息的接口
func (r RedisConnector) GetLimitProxyIP(limit int) ([]string, error) {
	if limit < 0 || limit > common.GetLimit {
		return nil, common.ErrModelLimitInvalid
	}
	reply := r.conn.Cmd("SRANDMEMBER", ipAddrStorageKey, limit)
	if reply.Err != nil {
		return nil, reply.Err
	}
	proxyList := make([]string, 0, limit)
	proxyIP := NewProxyIP()

	ips, err := reply.List()
	if err != nil {
		return nil, err
	}
	ipInterfaces := make([]interface{},len(ips) + 1)
	ipInterfaces[0] = ipStorageKey
	for i,v := range ips {
		ipInterfaces[i+1] = v
	}
	result := r.conn.Cmd("HMGET", ipInterfaces...)
	if result.Err != nil {
		return nil,result.Err
	}
	ipInfos, err := result.List()
	if err != nil {
		return nil, err
	}

	//TODO: 后续可优化，不用多次请求
	for _, ipInfo := range ipInfos {
		if err = json.Unmarshal([]byte(ipInfo), &proxyIP); err != nil {
			return nil, err
		}
		ip, err := proxyIP.IP()
		if err != nil {
			return nil, err
		}
		port, err := proxyIP.Port()
		if err != nil {
			return nil, err
		}
		proxyList = append(proxyList, net.JoinHostPort(ip, port))
	}
	return proxyList, nil
}

func Noop() {

}
