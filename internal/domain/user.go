package domain

import (
	"pencatatan_keuangan/pkg/utils"
)

type User struct {
    utils.BaseDomain
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"uniqueIndex"`
    Password string `json:"-"`
}
