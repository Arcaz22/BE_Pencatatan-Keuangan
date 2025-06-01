package mapper

import (
	"pencatatan_keuangan/internal/domain"
	"pencatatan_keuangan/internal/service/income/dto"
	"time"

	"github.com/google/uuid"
)

type IncomeMapper struct{}

func NewIncomeMapper() *IncomeMapper {
	return &IncomeMapper{}
}

func (m *IncomeMapper) ToIncomeEntity(req dto.CreateIncomeRequest) *domain.Income {
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return nil
	}
    date, err := time.Parse("2006-01-02", req.Date)
    if err != nil {
        return nil
    }
	return &domain.Income{
		CategoryID:  categoryID,
		Amount:      req.Amount,
		Description: req.Description,
		Date:        date,
	}
}

func (m *IncomeMapper) ToIncomeResponse(income *domain.Income, category *domain.Category) dto.IncomeResponse {
	return dto.IncomeResponse{
		ID:          income.ID.String(),
		CategoryID:  income.CategoryID.String(),
        Category:    category.Name,
		Amount:      income.Amount,
		Description: income.Description,
		Date:        income.Date.Format("2006-01-02"),
	}
}
