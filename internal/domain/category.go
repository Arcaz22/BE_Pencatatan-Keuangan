package domain

import (
    "pencatatan_keuangan/pkg/utils"
    "pencatatan_keuangan/pkg/constant"
)

type Category struct {
    utils.BaseDomain
    Name        string `json:"name" gorm:"not null"`
    Description string `json:"description"`
    Type        string `json:"type" gorm:"type:varchar(10);not null;check:type IN ('income', 'expense')"`
}

func (c *Category) IsValidType() bool {
    return c.Type == constant.CategoryTypeIncome || c.Type == constant.CategoryTypeExpense
}
