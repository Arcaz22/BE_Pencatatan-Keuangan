package mapper

import (
	"pencatatan_keuangan/internal/domain"
	"pencatatan_keuangan/internal/service/budget/dto"
	"time"

	"github.com/google/uuid"
)

type BudgetMapper struct{}

func NewBudgetMapper() *BudgetMapper {
	return &BudgetMapper{}
}

const dateFormat = "2006-01-02"

func (m *BudgetMapper) ToBudgetEntity(req dto.CreateBudgetRequest) *domain.Budget {
    categoryID, err := uuid.Parse(req.CategoryID)
    if err != nil {
        return nil
    }

    effectiveFrom, err := time.Parse(dateFormat, req.EffectiveFrom)
    if err != nil {
        return nil
    }
    effectiveTo, err := time.Parse(dateFormat, req.EffectiveTo)
    if err != nil {
        return nil
    }

    return &domain.Budget{
        CategoryID:     categoryID,
        Amount:        req.Amount,
        EffectiveFrom: effectiveFrom,
        EffectiveTo:   effectiveTo,
        IsActive:      true,
    }
}

func (m *BudgetMapper) ToBudgetResponse(budget *domain.Budget, category *domain.Category) dto.BudgetResponse {
    return dto.BudgetResponse{
        ID:            budget.ID.String(),
        CategoryID:    budget.CategoryID.String(),
        Category:      category.Name,
        Amount:        budget.Amount,
        EffectiveFrom: budget.EffectiveFrom.Format(dateFormat),
        EffectiveTo:   budget.EffectiveTo.Format(dateFormat),
        IsActive:      budget.IsActive,
    }
}
