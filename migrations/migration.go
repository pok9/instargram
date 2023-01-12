package migrations

import (
	"instargram/config"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Migrate() {
	db := config.GetDB()
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		m1673267911CreateUsersTable(),
		m1673447665CreatePostMainsTable(),
		m1673448560AddUserIDToPostMains(),
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
