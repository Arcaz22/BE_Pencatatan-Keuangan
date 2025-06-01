package domain

import (
    "time"
    "pencatatan_keuangan/pkg/utils"
    "github.com/google/uuid"
)

type Expense struct {
    utils.BaseDomain
    UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    CategoryID  uuid.UUID `json:"category_id" gorm:"type:uuid"`
    Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
    Amount      float64   `json:"amount" gorm:"not null"`
    Description string    `json:"description"`
    Date        time.Time `json:"date" gorm:"not null"`
}
