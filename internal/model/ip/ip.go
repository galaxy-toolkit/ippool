package ip

// IP 实例
type IP struct {
	ID      int    // ID
	Address string `` // IP 地址
	Status  Status // IP 状态
	Latency int    // 延迟，单位毫秒
	Source  string // 来源
}

// Status IP 状态
type Status string

const (
	Normal  Status = "normal"  // 正常
	Invalid Status = "invalid" // 已失效
)
