package migrations

import (
	"instargram/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1673725240CreateLikesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1673725240",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Like{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("likes")
		},
	}
}
