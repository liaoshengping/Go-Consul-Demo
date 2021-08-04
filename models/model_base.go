package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	UpdatedAt   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳填纳秒数充更新时间
	CreatedAt   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
}


