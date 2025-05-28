package handler

import (
	"pencatatan_keuangan/internal/service/category"
	"pencatatan_keuangan/internal/service/category/dto"
	"pencatatan_keuangan/pkg/constant"
	"pencatatan_keuangan/pkg/response"
	"pencatatan_keuangan/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	service category.CategoryService
}

func NewCategoryHandler(service category.CategoryService) *CategoryHandler {
    return &CategoryHandler{service: service}
}

// @Summary      Create a new category
// @Description  Create a new category with name and type
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param request body dto.CreateCategoryRequest true "Category Data"
// @Success      201 {object} dto.CategoryResponse
// @Security     BearerAuth
// @Router       /categories/create [post]
func (h *CategoryHandler) Create(c *gin.Context) {
	var request dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
		return
	}

	result, err := h.service.Create(request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, constant.MsgCreateCategorySuccess, result)
}

// @Summary      Get all categories
// @Description  Get paginated list of categories with filters
// @Tags         categories
// @Produce      json
// @Param        page query int false "Page number" minimum(1)
// @Param        limit query int false "Items per page" minimum(1) maximum(100)
// @Param        search query string false "Search in name and description"
// @Param        sort_by query string false "Sort field"
// @Param        sort_dir query string false "Sort direction" Enums(asc, desc)
// @Param        type query string false "Filter by type" Enums(income, expense)
// @Success      200 {object} response.PaginatedResponse{data=[]dto.CategoryResponse}
// @Security     BearerAuth
// @Router       /categories/all [get]
func (h *CategoryHandler) GetAll(c *gin.Context) {
    params := response.NewPaginationParams(c)

    categories, total, err := h.service.GetAll(params)
    if err != nil {
        c.Error(err)
        return
    }

response.SuccessPaginated(c, constant.MsgRetrievedCategoriesSuccess, categories, params, total)
}

// @Summary      Update a category
// @Description  Update a category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id path string true "Category ID" format(uuid)
// @Param        request body dto.UpdateCategoryRequest true "Category Data"
// @Success      200 {object} dto.CategoryResponse
// @Security     BearerAuth
// @Router       /categories/{id} [put]
func (h *CategoryHandler) Update(c *gin.Context) {
    idStr := c.Param("id")

    id, err := uuid.Parse(idStr)
    if err != nil {
        c.Error(utils.NewValidationError("Invalid category ID format", err))
        return
    }

    var request dto.UpdateCategoryRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
        return
    }

    result, err := h.service.Update(id.String(), request)
    if err != nil {
        c.Error(err)
        return
    }

    response.Success(c, constant.MsgUpdateCategorySuccess, result)
}

// @Summary      Delete a category
// @Description  Delete a category by ID
// @Tags         categories
// @Produce      json
// @Param        id path string true "Category ID" format(uuid)
// @Success      204 {object} nil
// @Security     BearerAuth
// @Router       /categories/{id} [delete]
func (h *CategoryHandler) Delete(c *gin.Context) {
    idStr := c.Param("id")

    _, err := uuid.Parse(idStr)
    if err != nil {
        c.Error(utils.NewValidationError("Invalid category ID format", err))
        return
    }

    err = h.service.Delete(idStr)
    if err != nil {
        c.Error(err)
        return
    }

    response.NoContent(c)
}
