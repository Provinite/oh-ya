package batch

import (
	"time"

	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id" gorm:"primaryKey"`
	RecipeID   uint      `json:"recipe"`
	Quantity   uint      `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
