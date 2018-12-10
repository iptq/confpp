package models

import "github.com/jinzhu/gorm"

type Beatmap struct {
	gorm.Model
	BeatmapID    int
	BeatmapSetID int
}
