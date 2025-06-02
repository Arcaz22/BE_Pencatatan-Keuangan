package handler

import (
    "pencatatan_keuangan/internal/service/expense"
    "pencatatan_keuangan/internal/service/expense/dto"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
    "pencatatan_keuangan/pkg/utils"

    "github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
    service expense.ExpenseService
}

func NewExpenseHandler(service expense.ExpenseService) *ExpenseHandler {
    return &ExpenseHandler{service: service}
}

// @Summary      Create a new expense
// @Description  Create a new expense record
// @Tags         expenses
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateExpenseRequest true "Expense Data"
// @Success      201 {object} dto.ExpenseResponse
// @Security     BearerAuth
// @Router       /expenses/create [post]
func (h *ExpenseHandler) Create(c *gin.Context) {
    var request dto.CreateExpenseRequest
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

    response.Success(c, constant.MsgCreateExpenseSuccess, result)
}

// @Summary      Get all expenses
// @Description  Get paginated list of expenses with filters
// @Tags         expenses
// @Produce      json
// @Param        page query int false "Page number" minimum(1)
// @Param        limit query int false "Items per page" minimum(1) maximum(100)
// @Param        search query string false "Search in date and amount"
// @Param        sort_by query string false "Sort field"
// @Param        sort_dir query string false "Sort direction" Enums(asc, desc)
// @Success      200 {object} response.PaginatedResponse{data=[]dto.ExpenseResponse}
// @Security     BearerAuth
// @Router       /expenses/all [get]
func (h *ExpenseHandler) GetAll(c *gin.Context) {
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

    response.SuccessPaginated(c, constant.MsgGetExpensesSuccess, result, params, totalRecords)
}

// @Summary      Update expense
// @Description  Update an existing expense record
// @Tags         expenses
// @Accept       json
// @Produce      json
// @Param        id path string true "Expense ID"
// @Param        request body dto.UpdateExpenseRequest true "Expense Data"
// @Success      200 {object} dto.ExpenseResponse
// @Security     BearerAuth
// @Router       /expenses/{id} [put]
func (h *ExpenseHandler) Update(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.Error(utils.NewValidationError(constant.MsgInvalidInput, nil))
        return
    }

    var request dto.UpdateExpenseRequest
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

    response.Success(c, constant.MsgUpdateExpenseSuccess, result)
}

// @Summary      Delete expense
// @Description  Delete an existing expense record
// @Tags         expenses
// @Produce      json
// @Param        id path string true "Expense ID"
// @Success      200 {object} response.Response
// @Security     BearerAuth
// @Router       /expenses/{id} [delete]
func (h *ExpenseHandler) Delete(c *gin.Context) {
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

    response.Success(c, constant.MsgDeleteExpenseSuccess, nil)
}
