package models

import "github.com/jinzhu/gorm"

type RawScore struct {
	gorm.Model
	UserID    int64
	BeatmapID int64

	User    User    `gorm:"foreignkey:UserID"`
	Beatmap Beatmap `gorm:"foreignkey:BeatmapID"`
}
