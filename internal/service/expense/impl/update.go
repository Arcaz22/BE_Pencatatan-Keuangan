package impl

import (
	"time"
	"github.com/google/uuid"
	"pencatatan_keuangan/internal/repository"
	"pencatatan_keuangan/internal/service/expense/dto"
	"pencatatan_keuangan/internal/service/expense/mapper"
	"pencatatan_keuangan/pkg/constant"
	"pencatatan_keuangan/pkg/utils"
)

func UpdateExpense(
    repo repository.ExpenseRepository,
    categoryRepo repository.CategoryRepository,
    expenseMapper *mapper.ExpenseMapper,
    id string,
    userID string,
    req dto.UpdateExpenseRequest,
) (*dto.ExpenseResponse, error) {
    expenseID, err := uuid.Parse(id)
    if err != nil {
        return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
    }

    expense, err := repo.FindByID(expenseID)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }
    if expense == nil {
        return nil, utils.NewNotFoundError(constant.MsgExpenseNotFound)
    }

    if expense.UserID.String() != userID {
        return nil, utils.NewAuthError(constant.MsgUnauthorized)
    }

    if req.Amount > 0 {
        expense.Amount = req.Amount
    }
    if req.Description != "" {
        expense.Description = req.Description
    }
    if req.Date != "" {
        date, err := time.Parse("2006-01-02", req.Date)
        if err != nil {
            return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
        }
        expense.Date = date
    }

    if req.CategoryID != "" {
        categoryID, err := uuid.Parse(req.CategoryID)
        if err != nil {
            return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
        }
        category, err := categoryRepo.FindByID(categoryID)
        if err != nil {
            return nil, utils.NewSystemError(constant.MsgInternalError, err)
        }
        if category == nil {
            return nil, utils.NewValidationError(constant.MsgCategoryNotFound, nil)
        }
        expense.CategoryID = categoryID
    }

    if err := repo.Update(expense); err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    category, _ := categoryRepo.FindByID(expense.CategoryID)
    response := expenseMapper.ToExpenseResponse(expense, category)
    return &response, nil
}
