package domain

import (
	"pencatatan_keuangan/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type Budget struct {
    utils.BaseDomain
    UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    User          User      `json:"user" gorm:"foreignKey:UserID"`
    CategoryID    uuid.UUID `json:"category_id" gorm:"type:uuid;not null"`
    Category      Category  `json:"category" gorm:"foreignKey:CategoryID"`
    Amount        float64   `json:"amount" gorm:"not null"`
    EffectiveFrom time.Time `json:"effective_from" gorm:"not null"`
    EffectiveTo   time.Time `json:"effective_to" gorm:"not null"`
    IsActive      bool      `json:"is_active" gorm:"default:true"`
}
