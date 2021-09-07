package sale

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model `json:"-"`
	ID         uint64     `json:"id"`
	SaleLines  []SaleLine `json:"sale_lines"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
