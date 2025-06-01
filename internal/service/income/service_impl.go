package income

import (
	"pencatatan_keuangan/internal/repository"
	"pencatatan_keuangan/internal/service/income/dto"
	"pencatatan_keuangan/internal/service/income/impl"
	"pencatatan_keuangan/internal/service/income/mapper"
	"pencatatan_keuangan/pkg/response"
)

type IncomeService interface {
    Create(req dto.CreateIncomeRequest) (*dto.IncomeResponse, error)
	GetAll(params response.PaginationParams, userID string) ([]dto.IncomeResponse, int, error)
	Update(id string, userID string, req dto.UpdateIncomeRequest) (*dto.IncomeResponse, error)
    Delete(id string, userID string) error
}

type incomeService struct {
    repository   repository.IncomeRepository
    categoryRepo repository.CategoryRepository
    incomeMapper *mapper.IncomeMapper
}

func NewIncomeService(repository repository.IncomeRepository, categoryRepo repository.CategoryRepository) IncomeService {
    return &incomeService{
        repository:   repository,
        categoryRepo: categoryRepo,
        incomeMapper: mapper.NewIncomeMapper(),
    }
}

func (s *incomeService) Create(req dto.CreateIncomeRequest) (*dto.IncomeResponse, error) {
    return impl.CreateIncome(s.repository, s.categoryRepo, s.incomeMapper, req)
}

func (s *incomeService) GetAll(params response.PaginationParams, userID string) ([]dto.IncomeResponse, int, error) {
    return impl.GetAllIncomes(s.repository, s.categoryRepo, s.incomeMapper, params, userID)
}

func (s *incomeService) Update(id string, userID string, req dto.UpdateIncomeRequest) (*dto.IncomeResponse, error) {
    return impl.UpdateIncome(s.repository, s.categoryRepo, s.incomeMapper, id, userID, req)
}

func (s *incomeService) Delete(id string, userID string) error {
    return impl.DeleteIncome(s.repository, id, userID)
}
