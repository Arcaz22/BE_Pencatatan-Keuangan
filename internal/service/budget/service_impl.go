package budget

import (
	"pencatatan_keuangan/internal/repository"
	"pencatatan_keuangan/internal/service/budget/dto"
	"pencatatan_keuangan/internal/service/budget/impl"
	"pencatatan_keuangan/internal/service/budget/mapper"
	"pencatatan_keuangan/pkg/response"
)

type BudgetService interface {
	Create(req dto.CreateBudgetRequest) (*dto.BudgetResponse, error)
	GetAll(params response.PaginationParams, userID string) ([]dto.BudgetResponse, int, error)
	Update(id string, userID string, req dto.UpdateBudgetRequest) (*dto.BudgetResponse, error)
	Delete(id string, userID string) error
}

type budgetService struct {
	repository   repository.BudgetRepository
	categoryRepo repository.CategoryRepository
	budgetMapper *mapper.BudgetMapper
}

func NewBudgetService(repository repository.BudgetRepository, categoryRepo repository.CategoryRepository) BudgetService {
	return &budgetService{
		repository:   repository,
		categoryRepo: categoryRepo,
		budgetMapper: mapper.NewBudgetMapper(),
	}
}

func (s *budgetService) Create(req dto.CreateBudgetRequest) (*dto.BudgetResponse, error) {
	return impl.CreateBudget(s.repository, s.categoryRepo, s.budgetMapper, req)
}

func (s *budgetService) GetAll(params response.PaginationParams, userID string) ([]dto.BudgetResponse, int, error) {
	return impl.GetAllBudgets(s.repository, s.categoryRepo, s.budgetMapper, params, userID)
}

func (s *budgetService) Update(id string, userID string, req dto.UpdateBudgetRequest) (*dto.BudgetResponse, error) {
	return impl.UpdateBudget(s.repository, s.categoryRepo, s.budgetMapper, id, userID, req)
}

func (s *budgetService) Delete(id string, userID string) error {
	return impl.DeleteBudget(s.repository, id, userID)
}
