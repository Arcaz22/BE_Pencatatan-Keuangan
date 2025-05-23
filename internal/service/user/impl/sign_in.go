package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/user/dto"
    "pencatatan_keuangan/internal/service/user/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/jwt"
    "pencatatan_keuangan/pkg/utils"
    "golang.org/x/crypto/bcrypt"
)

func SignInUser(repo repository.UserRepository, userMapper *mapper.UserMapper, req dto.SignInRequest) (*dto.SignInResponse, error) {
    user, err := repo.FindByEmail(req.Email)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    if user == nil {
        return nil, utils.NewBusinessError(constant.MsgInvalidCredential)
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return nil, utils.NewBusinessError(constant.MsgInvalidCredential)
    }

    token, err := jwt.GenerateToken(user.ID)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    response := userMapper.ToSignInResponse(user, token)
    return &response, nil
}
