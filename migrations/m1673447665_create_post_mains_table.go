package migrations

import (
	"instargram/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1673447665CreatePostMainsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1673447665",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.PostMain{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("post_mains")
		},
	}
}
