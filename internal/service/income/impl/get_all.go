package impl

import (
    "fmt"
    "pencatatan_keuangan/internal/domain"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/income/dto"
    "pencatatan_keuangan/internal/service/income/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
    "pencatatan_keuangan/pkg/utils"
)

func GetAllIncomes(
    repo repository.IncomeRepository,
    categoryRepo repository.CategoryRepository,
    incomeMapper *mapper.IncomeMapper,
    params response.PaginationParams,
    userID string,
) ([]dto.IncomeResponse, int, error) {
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
    if err := query.Model(&domain.Income{}).Count(&total).Error; err != nil {
        return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
    }

    if params.Page > 0 && params.Limit > 0 {
        offset := (params.Page - 1) * params.Limit
        query = query.Offset(offset).Limit(params.Limit)
    }

    var incomes []domain.Income
    if err := query.Preload("Category").Find(&incomes).Error; err != nil {
        return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
    }

    var incomesResponse []dto.IncomeResponse
    for _, income := range incomes {
        category := income.Category
        incomesResponse = append(incomesResponse, incomeMapper.ToIncomeResponse(&income, &category))
    }

    return incomesResponse, int(total), nil
}
