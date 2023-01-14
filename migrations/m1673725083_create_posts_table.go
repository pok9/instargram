package migrations

import (
	"instargram/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1673725083CreatePostsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1673725083",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Post{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("posts")
		},
	}
}
