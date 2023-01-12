package migrations

import (
	"instargram/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1673448560AddUserIDToPostMains() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1673448560",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.PostMain{})
		},
		Rollback: func(tx *gorm.DB) error {
			//drop column
			return tx.Migrator().DropColumn(&models.PostMain{}, "user_id")
		},
	}
}
