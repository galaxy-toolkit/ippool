package model

// IPStatus IP 状态
type IPStatus string

const (
	Normal  IPStatus = "normal"  // 正常
	Invalid IPStatus = "invalid" // 已失效
)
