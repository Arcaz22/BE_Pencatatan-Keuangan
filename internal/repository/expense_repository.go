package repository

import (
	"errors"
	"github.com/google/uuid"
	"pencatatan_keuangan/internal/domain"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	FindByID(id uuid.UUID) (*domain.Expense, error)
	Create(expense *domain.Expense) error
	GetQueryBuilder() *gorm.DB
	Update(expense *domain.Expense) error
	Delete(expense *domain.Expense) error
}

type expenseRepository struct {
	db * gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) FindByID(id uuid.UUID) (*domain.Expense, error) {
	var expense domain.Expense
	err := r.db.Where("id = ?", id).First(&expense).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &expense, nil
}

func (r *expenseRepository) Create(expense *domain.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) GetQueryBuilder() *gorm.DB {
	return r.db.Model(&domain.Expense{})
}

func (r *expenseRepository) Update(expense *domain.Expense) error {
	err := r.db.Save(expense).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *expenseRepository) Delete(expense *domain.Expense) error {
	err := r.db.Delete(expense).Error
	if err != nil {
		return err
	}
	return nil
}
