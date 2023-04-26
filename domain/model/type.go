package model

// IPStatus IP 状态
type IPStatus string

const (
	NotVerify IPStatus = "not-verify" // 还未验证
	Normal    IPStatus = "normal"     // 正常
	Invalid   IPStatus = "invalid"    // 失效
)

// IPProtocol 代理协议
type IPProtocol string

const (
	HTTP  IPProtocol = "HTTP"  // HTTP
	HTTPS IPProtocol = "HTTPS" // HTTPS
	SOCKS IPProtocol = "SOCKS" // SOCKS
)
