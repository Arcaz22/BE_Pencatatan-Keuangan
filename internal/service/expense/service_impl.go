package expense

import (
	"pencatatan_keuangan/internal/repository"
	"pencatatan_keuangan/internal/service/expense/dto"
	"pencatatan_keuangan/internal/service/expense/impl"
	"pencatatan_keuangan/internal/service/expense/mapper"
	"pencatatan_keuangan/pkg/response"
)

type ExpenseService interface {
    Create(req dto.CreateExpenseRequest) (*dto.ExpenseResponse, error)
	GetAll(params response.PaginationParams, userID string) ([]dto.ExpenseResponse, int, error)
	Update(id string, userID string, req dto.UpdateExpenseRequest) (*dto.ExpenseResponse, error)
    Delete(id string, userID string) error
}

type expenseService struct {
    repository   repository.ExpenseRepository
    categoryRepo repository.CategoryRepository
    expenseMapper *mapper.ExpenseMapper
}

func NewExpenseService(repository repository.ExpenseRepository, categoryRepo repository.CategoryRepository) ExpenseService {
    return &expenseService{
        repository:   repository,
        categoryRepo: categoryRepo,
        expenseMapper: mapper.NewExpenseMapper(),
    }
}

func (s *expenseService) Create(req dto.CreateExpenseRequest) (*dto.ExpenseResponse, error) {
    return impl.CreateExpense(s.repository, s.categoryRepo, s.expenseMapper, req)
}

func (s *expenseService) GetAll(params response.PaginationParams, userID string) ([]dto.ExpenseResponse, int, error) {
    return impl.GetAllExpenses(s.repository, s.categoryRepo, s.expenseMapper, params, userID)
}

func (s *expenseService) Update(id string, userID string, req dto.UpdateExpenseRequest) (*dto.ExpenseResponse, error) {
    return impl.UpdateExpense(s.repository, s.categoryRepo, s.expenseMapper, id, userID, req)
}

func (s *expenseService) Delete(id string, userID string) error {
    return impl.DeleteExpense(s.repository, id, userID)
}
