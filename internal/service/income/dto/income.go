package dto

type CreateIncomeRequest struct {
    UserID      string  `json:"-"`
    CategoryID  string  `json:"category_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
    Amount      float64 `json:"amount" validate:"required,gt=0" example:"1000.00"`
    Description string  `json:"description" validate:"omitempty,max=255" example:"Salary for September"`
    Date        string  `json:"date" validate:"required,datetime=2006-01-02" example:"2023-09-30"`
}

type UpdateIncomeRequest struct {
	UserID      string  `json:"-"`
	CategoryID string  `json:"category_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Amount      float64 `json:"amount" validate:"required,gt=0" example:"1000.00"`
	Description string  `json:"description" validate:"omitempty,max=255" example:"Salary for September"`
	Date        string  `json:"date" validate:"required,datetime=2006-01-02" example:"2023-09-30"`
}

type IncomeResponse struct {
	ID          string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CategoryID  string  `json:"category_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Category    string `json:"category" example:"Salary"`
	Amount      float64 `json:"amount" example:"1000.00"`
	Description string `json:"description" example:"Salary for September"`
	Date        string `json:"date" example:"2023-09-30"`
}
