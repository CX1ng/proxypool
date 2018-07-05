package common

const (
	StorageChannelCapacity = 100
	StorageConcurrencyRoutineCount = 20
)

// common
const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"
)

// model
const (
	GetLimit = 20
)

var StorageMaps = make(map[string]bool)

// Validator
const (
	VerifyUrl = "http://www.baidu.com"
	TIME_OUT  = 5
	TITLE     = "百度一下，你就知道"
)

func init() {
	StorageMaps["mysql"] = true

}
