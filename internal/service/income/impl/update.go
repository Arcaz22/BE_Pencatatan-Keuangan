package impl

import (
	"time"
	"github.com/google/uuid"
	"pencatatan_keuangan/internal/repository"
	"pencatatan_keuangan/internal/service/income/dto"
	"pencatatan_keuangan/internal/service/income/mapper"
	"pencatatan_keuangan/pkg/constant"
	"pencatatan_keuangan/pkg/utils"
)

func UpdateIncome(
    repo repository.IncomeRepository,
    categoryRepo repository.CategoryRepository,
    incomeMapper *mapper.IncomeMapper,
    id string,
    userID string,
    req dto.UpdateIncomeRequest,
) (*dto.IncomeResponse, error) {
    incomeID, err := uuid.Parse(id)
    if err != nil {
        return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
    }

    income, err := repo.FindByID(incomeID)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }
    if income == nil {
        return nil, utils.NewNotFoundError(constant.MsgIncomeNotFound)
    }

    if income.UserID.String() != userID {
        return nil, utils.NewAuthError(constant.MsgUnauthorized)
    }

    if req.Amount > 0 {
        income.Amount = req.Amount
    }
    if req.Description != "" {
        income.Description = req.Description
    }
    if req.Date != "" {
        date, err := time.Parse("2006-01-02", req.Date)
        if err != nil {
            return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
        }
        income.Date = date
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
        income.CategoryID = categoryID
    }

    if err := repo.Update(income); err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    category, _ := categoryRepo.FindByID(income.CategoryID)
    response := incomeMapper.ToIncomeResponse(income, category)
    return &response, nil
}
