package ingredient

import (
	"time"

	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model `json:"-"`
	ID         uint64         `json:"id"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
