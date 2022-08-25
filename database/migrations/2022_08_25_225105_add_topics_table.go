package migrations

import (
	"database/sql"
	"forum/app/models"
	"forum/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel
	}
	type Category struct {
		models.BaseModel
	}

	type Topic struct {
		models.BaseModel

		Title      string `json:"title,omitempty"`
		Body       string `json:"body,omitempty"`
		UserID     string `json:"user_id,omitempty"`
		CategoryID string `json:"category_id,omitempty"`

		User     User
		Category Category

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Topic{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Topic{})
		if err != nil {
			return
		}
	}

	migrate.Add("2022_08_25_225105_add_topics_table", up, down)
}
