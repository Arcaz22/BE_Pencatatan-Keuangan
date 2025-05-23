package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/user/dto"
    "pencatatan_keuangan/internal/service/user/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
    "golang.org/x/crypto/bcrypt"
)

func RegisterUser(repo repository.UserRepository, userMapper *mapper.UserMapper, req dto.RegisterRequest) (*dto.RegisterResponse, error) {
    existingUser, err := repo.FindByEmail(req.Email)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    if existingUser != nil {
        return nil, utils.NewBusinessError(constant.MsgDuplicateEntry)
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, utils.NewSystemError(constant.MsgInternalError, err)
	}

    user := userMapper.ToUserEntity(req, string(hashedPassword))

    err = repo.Create(user)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    response := userMapper.ToRegisterResponse(user)
    return &response, nil
}
