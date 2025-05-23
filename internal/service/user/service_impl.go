package user

import (
    "github.com/google/uuid"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/user/dto"
    "pencatatan_keuangan/internal/service/user/impl"
    "pencatatan_keuangan/internal/service/user/mapper"
)

type UserService interface {
    Register(req dto.RegisterRequest) (*dto.RegisterResponse, error)
    SignIn(req dto.SignInRequest) (*dto.SignInResponse, error)
    GetUserByID(userID uuid.UUID) (*dto.UserInfo, error)
}

type userService struct {
    repository repository.UserRepository
    userMapper *mapper.UserMapper
}

func NewUserService(repository repository.UserRepository) UserService {
    return &userService{
        repository: repository,
        userMapper: mapper.NewUserMapper(),
    }
}

func (s *userService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {
    return impl.RegisterUser(s.repository, s.userMapper, req)
}

func (s *userService) SignIn(req dto.SignInRequest) (*dto.SignInResponse, error) {
    return impl.SignInUser(s.repository, s.userMapper, req)
}

func (s *userService) GetUserByID(userID uuid.UUID) (*dto.UserInfo, error) {
    return impl.GetUserProfile(s.repository, s.userMapper, userID)
}
