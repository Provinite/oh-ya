package purchase

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model    `json:"-"`
	ID            uint           `json:"id" gorm:"primaryKey"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	PurchaseLines []PurchaseLine `json:"purchase_lines"`
	CreatedAt     time.Time      `json:"created_at"`
}
