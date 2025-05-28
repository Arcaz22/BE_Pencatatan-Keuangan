package category

import (
	"pencatatan_keuangan/internal/repository"
	"pencatatan_keuangan/internal/service/category/dto"
	"pencatatan_keuangan/internal/service/category/impl"
	"pencatatan_keuangan/internal/service/category/mapper"
	"pencatatan_keuangan/pkg/response"
)

type CategoryService interface {
	Create(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	GetAll(params response.PaginationParams) ([]dto.CategoryResponse, int, error)
	Update(id string, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error)
    Delete(id string) error
}

type categoryService struct {
	repository repository.CategoryRepository
	categoryMapper *mapper.CategoryMapper
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &categoryService{
		repository: repository,
		categoryMapper: mapper.NewCategoryMapper(),
	}
}

func (s *categoryService) Create(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	return impl.CreateCategory(s.repository, s.categoryMapper, req)
}

func (s *categoryService) GetAll(params response.PaginationParams) ([]dto.CategoryResponse, int, error) {
    return impl.GetAllCategories(s.repository, s.categoryMapper, params)
}

func (s *categoryService) Update(id string, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
    return impl.UpdateCategory(s.repository, s.categoryMapper, id, req)
}

func (s *categoryService) Delete(id string) error {
    return impl.DeleteCategory(s.repository, id)
}
