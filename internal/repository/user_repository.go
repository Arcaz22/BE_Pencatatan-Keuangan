package repository

import (
    "errors"
    "github.com/google/uuid"
    "pencatatan_keuangan/internal/domain"
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(user *domain.User) error
    FindByEmail(email string) (*domain.User, error)
    FindByID(id uuid.UUID) (*domain.User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
    var user domain.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) FindByID(id uuid.UUID) (*domain.User, error) {
    var user domain.User
    err := r.db.Where("id = ?", id).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}
