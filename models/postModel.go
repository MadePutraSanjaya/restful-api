package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title string
	Body  string
	gorm.Model
}
