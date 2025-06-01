package domain

import (
    "github.com/google/uuid"
	    "pencatatan_keuangan/pkg/utils"
)

type Budget struct {
    utils.BaseDomain
    UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    User       User      `json:"user" gorm:"foreignKey:UserID"`
    CategoryID uuid.UUID `json:"category_id" gorm:"type:uuid;not null"`
    Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
    Amount     float64   `json:"amount" gorm:"not null"`
    Month      int       `json:"month" gorm:"not null"`
    Year       int       `json:"year" gorm:"not null"`
}
