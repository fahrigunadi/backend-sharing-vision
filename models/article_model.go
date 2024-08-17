package models

import "time"

type Article struct {
	Id          uint64    `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(200)" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	CreatedDate time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"type:timestamp;autoCreateTime;autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100)" json:"status"`
}
