package mapper

import (
	"pencatatan_keuangan/internal/domain"
    "pencatatan_keuangan/internal/service/category/dto"
)

type CategoryMapper struct{}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (m *CategoryMapper) ToCategoryEntity(req dto.CreateCategoryRequest) *domain.Category {
	return &domain.Category{
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
	}
}

func (m *CategoryMapper) ToCategoryResponse(category *domain.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:          category.ID.String(),
		Name:        category.Name,
		Description: category.Description,
		Type:        category.Type,
	}
}
