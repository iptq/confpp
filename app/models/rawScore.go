package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type RawScore struct {
	gorm.Model
	UserID    int       `gorm:"unique_index:rawscore_id"`
	BeatmapID int       `gorm:"unique_index:rawscore_id"`
	Date      time.Time `gorm:"unique_index:rawscore_id"`

	Score     int64
	Count300  int
	Count100  int
	Count50   int
	CountMiss int
	CountGeki int
	CountKatu int
	MaxCombo  int
	Mods      int
	Rank      string

	User    User    `gorm:"foreignkey:UserID"`
	Beatmap Beatmap `gorm:"foreignkey:BeatmapID"`
}
