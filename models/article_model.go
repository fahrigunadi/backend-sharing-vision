package models

import "time"

type Article struct {
	Id          uint64    `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(200);not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	CreatedDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100);not null" json:"status"`
}
