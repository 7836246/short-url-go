package model

import "time"

type ShortUrl struct {
	Id        int32     `gorm:"column:id;primaryKey;not null" json:"id"`
	Lurl      string    `gorm:"column:lurl;not null" json:"lurl"`
	Surl      string    `gorm:"column:surl;not null" json:"surl"`
	GmtCreate time.Time `gorm:"column:gmt_create;default:CURRENT_TIMESTAMP;not null" json:"gmt_create"`
}
