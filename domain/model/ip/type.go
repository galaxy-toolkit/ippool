package ip

// Status IP 状态
type Status string

const (
	Normal  Status = "normal"  // 正常
	Invalid Status = "invalid" // 已失效
)
