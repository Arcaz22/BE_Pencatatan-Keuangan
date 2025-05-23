package dto

type RegisterRequest struct {
    Name            string `json:"name" binding:"required,min=3,max=100" example:"John Doe"`
    Email           string `json:"email" binding:"required,email" example:"john.doe@example.com"`
    Password        string `json:"password" binding:"required,min=8" example:"strongpass123"`
    ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password" example:"strongpass123"`
}

type RegisterResponse struct {
    ID    string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
    Name  string `json:"name" example:"John Doe"`
    Email string `json:"email" example:"john.doe@example.com"`
}
