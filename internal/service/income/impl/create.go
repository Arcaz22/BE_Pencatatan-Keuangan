package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/income/dto"
    "pencatatan_keuangan/internal/service/income/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
    "github.com/google/uuid"
)

func CreateIncome(repo repository.IncomeRepository, categoryRepo repository.CategoryRepository, incomeMapper *mapper.IncomeMapper, req dto.CreateIncomeRequest) (*dto.IncomeResponse, error) {
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

    income := incomeMapper.ToIncomeEntity(req)
    if income == nil {
        return nil, utils.NewValidationError(constant.MsgInvalidInput, nil)
    }
    income.UserID = userID

    err = repo.Create(income)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    response := incomeMapper.ToIncomeResponse(income, category)
    return &response, nil
}
