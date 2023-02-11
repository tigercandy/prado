package system

import "time"

type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"-"`
}
