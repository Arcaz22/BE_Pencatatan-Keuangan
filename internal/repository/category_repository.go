package repository

import (
	"errors"
	"github.com/google/uuid"
	"pencatatan_keuangan/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindByName(name string) (*domain.Category, error)
	Create(category *domain.Category) error
	FindByID(id uuid.UUID) (*domain.Category, error)
	GetQueryBuilder() *gorm.DB
	Update(category *domain.Category) error
    Delete(category *domain.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindByName(name string) (*domain.Category, error) {
	var category domain.Category
	err := r.db.Where("name = ?", name).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindByID(id uuid.UUID) (*domain.Category, error) {
	var category domain.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetQueryBuilder() *gorm.DB {
    return r.db.Model(&domain.Category{})
}

func (r *categoryRepository) Update(category *domain.Category) error {
	err := r.db.Save(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) Delete(category *domain.Category) error {
	err := r.db.Delete(category).Error
	if err != nil {
		return err
	}
	return nil
}
