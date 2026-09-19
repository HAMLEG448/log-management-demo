package entity

import "time"

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Username string `gorm:"uniqueIndex;not null" json:"username"`

	PasswordHash string `gorm:"not null" json:"-"`

	Role string `gorm:"not null" json:"role"`

	Tenant string `json:"tenant"`

	CreatedAt time.Time `json:"created_at"`
}
