package repository

import (
	"errors"
	"github.com/google/uuid"
	"pencatatan_keuangan/internal/domain"
	"gorm.io/gorm"
)

type IncomeRepository interface {
	FindByID(id uuid.UUID) (*domain.Income, error)
	Create(income *domain.Income) error
	GetQueryBuilder() *gorm.DB
	Update(income *domain.Income) error
	Delete(income *domain.Income) error
}

type incomeRepository struct {
	db * gorm.DB
}

func NewIncomeRepository(db *gorm.DB) IncomeRepository {
	return &incomeRepository{db: db}
}

func (r *incomeRepository) FindByID(id uuid.UUID) (*domain.Income, error) {
	var income domain.Income
	err := r.db.Where("id = ?", id).First(&income).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &income, nil
}

func (r *incomeRepository) Create(income *domain.Income) error {
	return r.db.Create(income).Error
}

func (r *incomeRepository) GetQueryBuilder() *gorm.DB {
	return r.db.Model(&domain.Income{})
}

func (r *incomeRepository) Update(income *domain.Income) error {
	err := r.db.Save(income).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *incomeRepository) Delete(income *domain.Income) error {
	err := r.db.Delete(income).Error
	if err != nil {
		return err
	}
	return nil
}
