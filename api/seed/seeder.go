package seed

import (
	"log"
	"time"

	"github.com/aulkiller/Deall_BE_Test/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Username:  "admin",
		Name:      "Administrator",
		Password:  "passadmin",
		Role:      "Admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	models.User{
		Username:  "auliaihza",
		Name:      "Aulia Ihza",
		Password:  "pass",
		Role:      "User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		HashedPass, err := models.HashPassword(users[i].Password)
		if err != nil {
			log.Fatalf("failed to encrypt seeded pass: %v", err)
		}
		users[i].Password = HashedPass
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
