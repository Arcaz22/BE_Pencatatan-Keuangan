package dto

type SignInRequest struct {
    Email    string `json:"email" binding:"required,email" example:"john.doe@example.com"`
    Password string `json:"password" binding:"required,min=6" example:"strongpass123"`
}

type SignInResponse struct {
    Token string   `json:"token"`
    User  UserInfo `json:"user"`
}
