package models

import (
	"gorm.io/gorm"
)

type Deposit struct {
	gorm.Model
	Amount float64 `json:"amount"`
	UserID int
	User   User
}
