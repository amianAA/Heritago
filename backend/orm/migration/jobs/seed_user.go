package jobs

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"heritago/backend/orm/models"
	"strings"
)

var (
	fname    = "FirstName"
	lname    = "LastName"
	country  = "Spain"
	rawPass  = "0987654321"
	password = models.HashAndSaltPwd(&rawPass)
	userID   = strings.ToLower(fname[:1] + lname)

	firstUser = &models.User{
		Email:     "test@gmail.com",
		UserID:    &userID,
		FirstName: &fname,
		LastName:  &lname,
		Country:   &country,
		Password:  &password,
	}
)

// SeedUsers inserts the first users
var SeedUsers = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}
