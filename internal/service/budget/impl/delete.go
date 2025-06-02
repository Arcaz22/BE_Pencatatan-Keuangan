package impl

import (
    "github.com/google/uuid"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
)

func DeleteBudget(
    repo repository.BudgetRepository,
    id string,
    userID string,
) error {
    budgetID, err := uuid.Parse(id)
    if err != nil {
        return utils.NewValidationError(constant.MsgInvalidInput, err)
    }

    budget, err := repo.FindByID(budgetID)
    if err != nil {
        return utils.NewSystemError(constant.MsgInternalError, err)
    }
    if budget == nil {
        return utils.NewNotFoundError(constant.MsgBudgetNotFound)
    }

    if budget.UserID.String() != userID {
        return utils.NewAuthError(constant.MsgUnauthorized)
    }

    err = repo.Delete(budget)
    if err != nil {
        return utils.NewSystemError(constant.MsgInternalError, err)
    }

    return nil
}
