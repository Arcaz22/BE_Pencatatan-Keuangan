package dto

type CreateBudgetRequest struct {
    UserID      string  `json:"-"`
    CategoryID  string  `json:"category_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
    Amount      float64 `json:"amount" validate:"required,gt=0" example:"1000.00"`
	EffectiveFrom string  `json:"effective_from" validate:"required,datetime=2006-01-02" example:"2023-09-01"`
	EffectiveTo   string  `json:"effective_to" validate:"required,datetime=2006-01-02" example:"2023-09-30"`
	IsActive	  bool    `json:"is_active" validate:"omitempty" example:"true"`
}

type UpdateBudgetRequest struct {
	UserID      string  `json:"-"`
	CategoryID  string  `json:"category_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Amount      float64 `json:"amount" validate:"required,gt=0" example:"1000.00"`
	EffectiveFrom string  `json:"effective_from" validate:"required,datetime=2006-01-02" example:"2023-09-01"`
	EffectiveTo   string  `json:"effective_to" validate:"required,datetime=2006-01-02" example:"2023-09-30"`
	IsActive	  bool    `json:"is_active" validate:"omitempty" example:"true"`
}

type BudgetResponse struct {
	ID          string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CategoryID  string  `json:"category_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Category    string  `json:"category" example:"Salary"`
	Amount      float64 `json:"amount" example:"1000.00"`
	EffectiveFrom string  `json:"effective_from" example:"2023-09-01"`
	EffectiveTo   string  `json:"effective_to" example:"2023-09-30"`
	IsActive	bool    `json:"is_active" example:"true"`
}
