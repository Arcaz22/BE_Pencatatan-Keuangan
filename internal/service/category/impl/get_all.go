package impl

import (
	"pencatatan_keuangan/internal/domain"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/category/dto"
    "pencatatan_keuangan/internal/service/category/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
    "pencatatan_keuangan/pkg/utils"
)

var searchableFields = []string{"name", "description"}

func GetAllCategories(
    repo repository.CategoryRepository,
    categoryMapper *mapper.CategoryMapper,
    params response.PaginationParams,
) ([]dto.CategoryResponse, int, error) {
    query := repo.GetQueryBuilder()

    // Apply filters and search
    query = response.ApplyFilters(query, params, searchableFields)

    // Count total before pagination
    total, err := response.CountTotalRecords(query, &domain.Category{})
    if err != nil {
        return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
    }

    // Apply pagination
    query = response.ApplyPaginationToQuery(query, params)

    // Execute query
    var categories []domain.Category
    if err := query.Find(&categories).Error; err != nil {
        return nil, 0, utils.NewSystemError(constant.MsgInternalError, err)
    }

    // Map to DTOs
    var categoriesResponse []dto.CategoryResponse
    for _, category := range categories {
        categoriesResponse = append(categoriesResponse, categoryMapper.ToCategoryResponse(&category))
    }

    return categoriesResponse, total, nil
}
