package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/user/dto"
    "pencatatan_keuangan/internal/service/user/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
    "github.com/google/uuid"
)

func GetUserProfile(repo repository.UserRepository, userMapper *mapper.UserMapper, userID uuid.UUID) (*dto.UserInfo, error) {
    user, err := repo.FindByID(userID)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    if user == nil {
        return nil, utils.NewNotFoundError(constant.MsgNotFound)
    }

    response := userMapper.ToProfileResponse(user)
    return &response, nil
}
