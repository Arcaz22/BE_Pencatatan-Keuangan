package impl

import (
    "time"
    "github.com/google/uuid"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/budget/dto"
    "pencatatan_keuangan/internal/service/budget/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
)

func UpdateBudget(
    repo repository.BudgetRepository,
    categoryRepo repository.CategoryRepository,
    budgetMapper *mapper.BudgetMapper,
    id string,
    userID string,
    req dto.UpdateBudgetRequest,
) (*dto.BudgetResponse, error) {
    budgetID, err := uuid.Parse(id)
    if err != nil {
        return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
    }

    budget, err := repo.FindByID(budgetID)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }
    if budget == nil {
        return nil, utils.NewNotFoundError(constant.MsgBudgetNotFound)
    }

    if budget.UserID.String() != userID {
        return nil, utils.NewAuthError(constant.MsgUnauthorized)
    }

    if req.Amount > 0 {
        budget.Amount = req.Amount
    }

    if req.EffectiveFrom != "" {
        effectiveFrom, err := time.Parse("2006-01-02", req.EffectiveFrom)
        if err != nil {
            return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
        }
        budget.EffectiveFrom = effectiveFrom
    }

    if req.EffectiveTo != "" {
        effectiveTo, err := time.Parse("2006-01-02", req.EffectiveTo)
        if err != nil {
            return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
        }
        budget.EffectiveTo = effectiveTo
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
        budget.CategoryID = categoryID
    }

    if err := repo.Update(budget); err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    category, _ := categoryRepo.FindByID(budget.CategoryID)
    response := budgetMapper.ToBudgetResponse(budget, category)
    return &response, nil
}
