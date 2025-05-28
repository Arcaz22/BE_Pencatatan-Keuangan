package dto

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100" example:"Groceries"`
	Description string `json:"description" binding:"max=255" example:"Monthly grocery shopping expenses"`
	Type        string `json:"type" binding:"required,oneof=income expense" example:"expense"`
}

type UpdateCategoryRequest struct {
    Name        string `json:"name" binding:"required,min=3,max=50" example:"Groceries"`
    Description string `json:"description" binding:"omitempty,max=255" example:"Monthly grocery shopping expenses"`
    Type        string `json:"type" binding:"required,oneof=income expense" example:"expense"`
}

type CategoryResponse struct {
	ID          string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string `json:"name" example:"Groceries"`
	Description string `json:"description" example:"Monthly grocery shopping expenses"`
	Type        string `json:"type" example:"EXPENSE"`
}
