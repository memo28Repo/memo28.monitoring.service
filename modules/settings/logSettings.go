package settings

import "gorm.io/gorm"

type LogSettings struct {
	gorm.Model

	// Number log最大数量 超出则删除
	Number int64 `gorm:"column number;DEFAULT 10" json:"number"`

	// ExpirationTime 每个错误log最长的保存时间
	ExpirationTime int `gorm:"column expirationTime;DEFAULT 0" json:"expirationTime"`
}
