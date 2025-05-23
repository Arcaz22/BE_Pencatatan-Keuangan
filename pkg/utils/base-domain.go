package utils

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseDomain struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (base *BaseDomain) BeforeCreate(tx *gorm.DB) error {
    if base.ID == uuid.Nil {
        base.ID = uuid.New()
    }
    return nil
}
