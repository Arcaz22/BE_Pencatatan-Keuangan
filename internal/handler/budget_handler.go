package handler

import (
    "pencatatan_keuangan/internal/service/budget"
    "pencatatan_keuangan/internal/service/budget/dto"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
    "pencatatan_keuangan/pkg/utils"

    "github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	service budget.BudgetService
}

func NewBudgetHandler(service budget.BudgetService) *BudgetHandler {
	return &BudgetHandler{service: service}
}

// @Summary      Create a new budget
// @Description  Create a new budget record
// @Tags         budgets
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateBudgetRequest true "Budget Data"
// @Success      201 {object} dto.BudgetResponse
// @Security     BearerAuth
// @Router       /budgets/create [post]
func (h *BudgetHandler) Create(c *gin.Context) {
	var request dto.CreateBudgetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
		return
	}

	userID, exists := utils.GetUserID(c)
	if !exists {
		c.Error(utils.NewValidationError(constant.ErrCodeUnauthorized, nil))
		return
	}
	request.UserID = userID

	result, err := h.service.Create(request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, constant.MsgCreateBudgetSuccess, result)
}

// @Summary      Get all budgets
// @Description  Get paginated list of budgets with filters
// @Tags         budgets
// @Produce      json
// @Param        page query int false "Page number" minimum(1)
// @Param        limit query int false "Items per page" minimum(1) maximum(100)
// @Param        search query string false "Search in EffectiveFrom, EffectiveTo and amount"
// @Param        sort_by query string false "Sort field"
// @Param        sort_dir query string false "Sort direction" Enums(asc, desc)
// @Success      200 {object} response.PaginatedResponse{data=[]dto.BudgetResponse}
// @Security     BearerAuth
// @Router       /budgets/all [get]
func (h *BudgetHandler) GetAll(c *gin.Context) {
	params := response.NewPaginationParams(c)
	if err := c.ShouldBindQuery(&params); err != nil {
        c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
        return
    }

	userID, exists := utils.GetUserID(c)
	if !exists {
		c.Error(utils.NewValidationError(constant.ErrCodeUnauthorized, nil))
		return
	}

	result, totalRecords, err := h.service.GetAll(params, userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessPaginated(c, constant.MsgGetBudgetsSuccess, result, params, totalRecords)
}

// @Summary      Update Budget
// @Description  Update an existing budget by ID
// @Tags         budgets
// @Accept       json
// @Produce      json
// @Param        id path string true "Budget ID"
// @Param        request body dto.UpdateBudgetRequest true "Updated Budget Data"
// @Success      200 {object} dto.BudgetResponse
// @Security     BearerAuth
// @Router       /budgets/{id} [put]
func (h *BudgetHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(utils.NewValidationError(constant.MsgInvalidInput, nil))
		return
	}

	var request dto.UpdateBudgetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
		return
	}

	userID, exists := utils.GetUserID(c)
	if !exists {
		c.Error(utils.NewValidationError(constant.ErrCodeUnauthorized, nil))
		return
	}

	result, err := h.service.Update(id, userID, request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, constant.MsgUpdateBudgetSuccess, result)
}

// @Summary      Delete Budget
// @Description  Delete a existing budget record
// @Tags         budgets
// @Param        id path string true "Budget ID"
// @Success      200 {object} response.Response
// @Security     BearerAuth
// @Router       /budgets/{id} [delete]
func (h *BudgetHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(utils.NewValidationError(constant.MsgInvalidInput, nil))
		return
	}

	userID, exists := utils.GetUserID(c)
	if !exists {
		c.Error(utils.NewValidationError(constant.ErrCodeUnauthorized, nil))
		return
	}

	err := h.service.Delete(id, userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, constant.MsgDeleteBudgetSuccess, nil)
}
