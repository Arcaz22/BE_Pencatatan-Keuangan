package repository

import (
    "errors"
    "pencatatan_keuangan/internal/domain"
    // "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type BudgetRepository interface {
    FindByID(id uuid.UUID) (*domain.Budget, error)
    Create(budget *domain.Budget) error
    GetQueryBuilder() *gorm.DB
    Update(budget *domain.Budget) error
	Delete(budget *domain.Budget) error
    // Delete(budget *domain.Budget) error
    // FindByUserAndCategory(userID uuid.UUID, categoryID uuid.UUID, date time.Time) (*domain.Budget, error)
}

type budgetRepository struct {
    db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) BudgetRepository {
    return &budgetRepository{db: db}
}

func (r *budgetRepository) FindByID(id uuid.UUID) (*domain.Budget, error) {
    var budget domain.Budget
    err := r.db.Where("id = ?", id).First(&budget).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &budget, nil
}

func (r *budgetRepository) Create(budget *domain.Budget) error {
    return r.db.Create(budget).Error
}

func (r *budgetRepository) GetQueryBuilder() *gorm.DB {
    return r.db.Model(&domain.Budget{})
}

func (r *budgetRepository) Update(budget *domain.Budget) error {
    err := r.db.Save(budget).Error
    if err != nil {
        return err
    }
    return nil
}

func (r *budgetRepository) Delete(budget *domain.Budget) error {
    err := r.db.Delete(budget).Error
    if err != nil {
        return err
    }
    return nil
}

// func (r *budgetRepository) FindByUserAndCategory(userID uuid.UUID, categoryID uuid.UUID, date time.Time) (*domain.Budget, error) {
//     var budget domain.Budget
//     err := r.db.Where("user_id = ? AND category_id = ? AND effective_from <= ? AND effective_to >= ? AND is_active = ?",
//         userID, categoryID, date, date, true).
//         First(&budget).Error
//     if err != nil {
//         if errors.Is(err, gorm.ErrRecordNotFound) {
//             return nil, nil
//         }
//         return nil, err
//     }
//     return &budget, nil
// }
