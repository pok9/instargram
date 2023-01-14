package migrations

import (
	"instargram/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1673725714CreateFollowersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1673725714",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Follower{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("Followers")
		},
	}
}
