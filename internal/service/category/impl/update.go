package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/category/dto"
    "pencatatan_keuangan/internal/service/category/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
    "github.com/google/uuid"
)

func UpdateCategory(
    repo repository.CategoryRepository,
    categoryMapper *mapper.CategoryMapper,
    id string,
    req dto.UpdateCategoryRequest,
) (*dto.CategoryResponse, error) {
    // Convert string ID to UUID
    uuid, err := uuid.Parse(id)
    if err != nil {
        return nil, utils.NewValidationError(constant.MsgInvalidInput, err)
    }

    category, err := repo.FindByID(uuid)
    if err != nil {
        return nil, utils.NewNotFoundError(constant.MsgCategoryNotFound)
    }

    category.Name = req.Name
    category.Description = req.Description
    category.Type = req.Type

    if err := repo.Update(category); err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    response := categoryMapper.ToCategoryResponse(category)
    return &response, nil
}
