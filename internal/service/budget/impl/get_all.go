package impl

import (
    "fmt"
    "pencatatan_keuangan/internal/domain"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/budget/dto"
    "pencatatan_keuangan/internal/service/budget/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
    "pencatatan_keuangan/pkg/utils"
)

func GetAllBudgets(
	repo repository.BudgetRepository,
	categoryRepo repository.CategoryRepository,
	budgetMapper *mapper.BudgetMapper,
	params response.PaginationParams,
	userID string,
) ([]dto.BudgetResponse, int, error) {
    query := repo.GetQueryBuilder()

    query = query.Where("user_id = ?", userID)

    if params.Search != "" {
        query = query.Where(`(
            CAST(amount AS TEXT) ILIKE ? OR
            CAST(effective_from AS TEXT) ILIKE ? OR
            CAST(effective_to AS TEXT) ILIKE ?
        )`,
            fmt.Sprintf("%%%s%%", params.Search),
            fmt.Sprintf("%%%s%%", params.Search),
            fmt.Sprintf("%%%s%%", params.Search),
        )
    }

	if params.SortBy != "" {
		direction := "ASC"
		if params.SortDir == "desc" {
			direction = "DESC"
		}
		query = query.Order(fmt.Sprintf("%s %s", params.SortBy, direction))
	}

	var total int64
	if err := query.Model(&domain.Budget{}).Count(&total).Error; err != nil {
		return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
	}

	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		query = query.Offset(offset).Limit(params.Limit)
	}

	var budgets []domain.Budget
	if err := query.Preload("Category").Find(&budgets).Error; err != nil {
		return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
	}

	var budgetsResponse []dto.BudgetResponse
    for _, budget := range budgets {
        category, err := categoryRepo.FindByID(budget.CategoryID)
        if err != nil {
            return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
        }
        budgetsResponse = append(budgetsResponse, budgetMapper.ToBudgetResponse(&budget, category))
    }

    return budgetsResponse, int(total), nil
}
