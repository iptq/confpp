package app

import (
	"confpp/app/models"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/revel/modules/jobs/app/jobs"
	osuapi "github.com/thehowl/go-osuapi"
)

type Tracker struct {
	api *osuapi.Client
	db  *gorm.DB
}

func (tracker Tracker) Run() {
	// get a list of users
	var users []models.User
	tracker.db.Find(&users)

	fmt.Println("users", users)

	// for each user, spawn a process looking for new plays
	for _, user := range users {
		jobs.Now(RetrieveNewPlays{api: tracker.api, db: tracker.db, user: user})
	}
}

type RetrieveNewPlays struct {
	api  *osuapi.Client
	db   *gorm.DB
	user models.User
}

func (job RetrieveNewPlays) Run() {
	// get list of recent scores
	results, err := job.api.GetUserRecent(osuapi.GetUserScoresOpts{
		UserID: job.user.OsuID,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
	}

	fmt.Println("recents for", job.user.OsuID, ":", results)

	// try to stick them all into the db
	for _, record := range results {
		score := models.RawScore{
			UserID:    job.user.OsuID,
			BeatmapID: record.BeatmapID,
			Date:      time.Time(record.Date),

			Score:     record.Score.Score,
			Count300:  record.Count300,
			Count100:  record.Count100,
			Count50:   record.Count50,
			CountMiss: record.CountMiss,
			CountGeki: record.CountGeki,
			CountKatu: record.CountKatu,
			MaxCombo:  record.MaxCombo,
			Mods:      int(record.Mods),
			Rank:      record.Rank,
		}
		job.db.Create(&score)
	}
}
