package impl

import (
	"pencatatan_keuangan/internal/repository"
	"pencatatan_keuangan/internal/service/category/dto"
	"pencatatan_keuangan/internal/service/category/mapper"
	"pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
)

func CreateCategory(repo repository.CategoryRepository, categoryMapper *mapper.CategoryMapper, req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	exsitingCategory, err := repo.FindByName(req.Name)
	if err != nil {
		return nil, utils.NewSystemError(constant.MsgInternalError, err)
	}

	if exsitingCategory != nil {
		return nil, utils.NewBusinessError(constant.MsgDuplicateEntry)
	}

	category := categoryMapper.ToCategoryEntity(req)

	err = repo.Create(category)
	if err != nil {
		return nil, utils.NewSystemError(constant.MsgInternalError, err)
	}

	response := categoryMapper.ToCategoryResponse(category)
	return &response, nil
}
