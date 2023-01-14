package migrations

import (
	"instargram/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1673726461CreateCommentsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1673726461",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Comment{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("Comments")
		},
	}
}
