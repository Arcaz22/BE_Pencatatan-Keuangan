package mapper

import (
	"pencatatan_keuangan/internal/domain"
	"pencatatan_keuangan/internal/service/expense/dto"
	"time"

	"github.com/google/uuid"
)

type ExpenseMapper struct{}

func NewExpenseMapper() *ExpenseMapper {
	return &ExpenseMapper{}
}

func (m *ExpenseMapper) ToExpenseEntity(req dto.CreateExpenseRequest) *domain.Expense {
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return nil
	}
    date, err := time.Parse("2006-01-02", req.Date)
    if err != nil {
        return nil
    }
	return &domain.Expense{
		CategoryID:  categoryID,
		Amount:      req.Amount,
		Description: req.Description,
		Date:        date,
	}
}

func (m *ExpenseMapper) ToExpenseResponse(expense *domain.Expense, category *domain.Category) dto.ExpenseResponse {
	return dto.ExpenseResponse{
		ID:          expense.ID.String(),
		CategoryID:  expense.CategoryID.String(),
        Category:    category.Name,
		Amount:      expense.Amount,
		Description: expense.Description,
		Date:        expense.Date.Format("2006-01-02"),
	}
}
