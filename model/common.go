package model

import (
	"gorm.io/gorm"
	"time"
)

// ID 自增id主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamps 创建更新时间
type Timestamps struct {
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	DeleteAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
