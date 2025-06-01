package impl

import (
    "github.com/google/uuid"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
)

func DeleteIncome(
    repo repository.IncomeRepository,
    id string,
    userID string,
) error {
    incomeID, err := uuid.Parse(id)
    if err != nil {
        return utils.NewValidationError(constant.MsgInvalidInput, err)
    }

    income, err := repo.FindByID(incomeID)
    if err != nil {
        return utils.NewSystemError(constant.MsgInternalError, err)
    }
    if income == nil {
        return utils.NewNotFoundError(constant.MsgIncomeNotFound)
    }

    if income.UserID.String() != userID {
        return utils.NewAuthError(constant.MsgUnauthorized)
    }

    err = repo.Delete(income)
    if err != nil {
        return utils.NewSystemError(constant.MsgInternalError, err)
    }

    return nil
}
