package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
    "github.com/google/uuid"
)

func DeleteCategory(repo repository.CategoryRepository, id string) error {
    uid, err := uuid.Parse(id)
    if err != nil {
        return utils.NewValidationError(constant.MsgInvalidInput, err)
    }

    category, err := repo.FindByID(uid)
    if err != nil {
        return utils.NewNotFoundError(constant.MsgCategoryNotFound)
    }

    if err := repo.Delete(category); err != nil {
        return utils.NewSystemError(constant.MsgInternalError, err)
    }

    return nil
}
