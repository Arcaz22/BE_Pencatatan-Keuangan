package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/budget/dto"
    "pencatatan_keuangan/internal/service/budget/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
    "github.com/google/uuid"
)

func CreateBudget(repo repository.BudgetRepository, categoryRepo repository.CategoryRepository, budgetMapper *mapper.BudgetMapper, req dto.CreateBudgetRequest) (*dto.BudgetResponse, error) {
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

	budget := budgetMapper.ToBudgetEntity(req)
	if budget == nil {
		return nil, utils.NewValidationError(constant.MsgInvalidInput, nil)
	}
	budget.UserID = userID

	err = repo.Create(budget)
	if err != nil {
		return nil, utils.NewSystemError(constant.MsgInternalError, err)
	}

	response := budgetMapper.ToBudgetResponse(budget, category)
	return &response, nil
}
