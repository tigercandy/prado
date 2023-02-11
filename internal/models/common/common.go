package common

import (
	"time"
)

type Operator struct {
	CreatedUser string `gorm:"column:created_user" json:"createdUser"`
	UpdatedUser string `gorm:"column:updated_user" json:"updatedUser"`
}

type Model struct {
	Id        int        `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}
