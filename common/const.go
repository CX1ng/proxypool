package common

const (
	StorageChannelTimeout  = 60
	StorageChannelCapacity = 100
)

// common
const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"
)

// model
const (
	GetLimit = 20
)

var WebUrl = make(map[string]string)

// Validator
const (
	VerifyUrl = "http://www.baidu.com"
	TIME_OUT  = 5
	TITLE     = "百度一下，你就知道"
)

var ProxyIPFields []string

func init() {
	WebUrl["kuaidaili"] = "https://www.kuaidaili.com/free/inha/"
	WebUrl["xici"] = "http://www.xicidaili.com/nn/"

	ProxyIPFields = []string{"ip", "port", "type", "origin", "raw_time", "region", "capture_time", "last_verify_time", "create_time"}
}
