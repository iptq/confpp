package app

import (
	"confpp/app/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Tracker struct {
	db *gorm.DB
}

func (tracker Tracker) Run() {
	// first, get a list of users
	var users []models.User
	tracker.db.Find(&users)

	fmt.Println("users", users)
}
