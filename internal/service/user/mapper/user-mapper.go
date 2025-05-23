package mapper

import (
    "pencatatan_keuangan/internal/domain"
    "pencatatan_keuangan/internal/service/user/dto"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
    return &UserMapper{}
}

func (m *UserMapper) ToUserEntity(req dto.RegisterRequest, hashedPassword string) *domain.User {
    return &domain.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashedPassword,
    }
}

func (m *UserMapper) ToRegisterResponse(user *domain.User) dto.RegisterResponse {
    return dto.RegisterResponse{
        ID:    user.ID.String(),
        Name:  user.Name,
        Email: user.Email,
    }
}

func (m *UserMapper) ToSignInResponse(user *domain.User, token string) dto.SignInResponse {
    return dto.SignInResponse{
        Token: token,
        User: dto.UserInfo{
            ID:    user.ID.String(),
            Name:  user.Name,
            Email: user.Email,
        },
    }
}

func (m *UserMapper) ToProfileResponse(user *domain.User) dto.UserInfo {
    return dto.UserInfo{
        ID:    user.ID.String(),
        Name:  user.Name,
        Email: user.Email,
    }
}
