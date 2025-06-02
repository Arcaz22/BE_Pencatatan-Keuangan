package impl

import (
    "github.com/google/uuid"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
)

func DeleteExpense(
    repo repository.ExpenseRepository,
    id string,
    userID string,
) error {
    expenseID, err := uuid.Parse(id)
    if err != nil {
        return utils.NewValidationError(constant.MsgInvalidInput, err)
    }

    expense, err := repo.FindByID(expenseID)
    if err != nil {
        return utils.NewSystemError(constant.MsgInternalError, err)
    }
    if expense == nil {
        return utils.NewNotFoundError(constant.MsgExpenseNotFound)
    }

    if expense.UserID.String() != userID {
        return utils.NewAuthError(constant.MsgUnauthorized)
    }

    err = repo.Delete(expense)
    if err != nil {
        return utils.NewSystemError(constant.MsgInternalError, err)
    }

    return nil
}
