package impl

import (
    "fmt"
    "pencatatan_keuangan/internal/domain"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/expense/dto"
    "pencatatan_keuangan/internal/service/expense/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
    "pencatatan_keuangan/pkg/utils"
)

func GetAllExpenses(
    repo repository.ExpenseRepository,
    categoryRepo repository.CategoryRepository,
    expenseMapper *mapper.ExpenseMapper,
    params response.PaginationParams,
    userID string,
) ([]dto.ExpenseResponse, int, error) {
    query := repo.GetQueryBuilder()

    query = query.Where("user_id = ?", userID)

    if params.Search != "" {
        query = query.Where(`(
            CAST(amount AS TEXT) ILIKE ? OR
            CAST(date AS TEXT) ILIKE ? OR
            description ILIKE ?
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
    if err := query.Model(&domain.Expense{}).Count(&total).Error; err != nil {
        return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
    }

    if params.Page > 0 && params.Limit > 0 {
        offset := (params.Page - 1) * params.Limit
        query = query.Offset(offset).Limit(params.Limit)
    }

    var expenses []domain.Expense
    if err := query.Preload("Category").Find(&expenses).Error; err != nil {
        return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
    }

    var expensesResponse []dto.ExpenseResponse
    for _, expense := range expenses {
        category := expense.Category
        expensesResponse = append(expensesResponse, expenseMapper.ToExpenseResponse(&expense, &category))
    }

    return expensesResponse, int(total), nil
}
