package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/expense/dto"
    "pencatatan_keuangan/internal/service/expense/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
    "github.com/google/uuid"
)

func CreateExpense(repo repository.ExpenseRepository, categoryRepo repository.CategoryRepository, expenseMapper *mapper.ExpenseMapper, req dto.CreateExpenseRequest) (*dto.ExpenseResponse, error) {
    userID, err := uuid.Parse(req.UserID)
    if err != nil {
        return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
    }

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

    expense := expenseMapper.ToExpenseEntity(req)
    if expense == nil {
        return nil, utils.NewValidationError(constant.MsgInvalidInput, nil)
    }
    expense.UserID = userID

    err = repo.Create(expense)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    response := expenseMapper.ToExpenseResponse(expense, category)
    return &response, nil
}
