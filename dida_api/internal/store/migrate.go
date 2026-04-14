package store

import (
	"dida_api/internal/focus"
	"dida_api/internal/list"
	"dida_api/internal/tag"
	"dida_api/internal/task"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&task.Task{},
		&list.List{},
		&tag.Tag{},
		&focus.Session{},
	)
}
