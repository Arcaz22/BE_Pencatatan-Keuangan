package handler

import (
    "pencatatan_keuangan/internal/service/income"
    "pencatatan_keuangan/internal/service/income/dto"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
    "pencatatan_keuangan/pkg/utils"

    "github.com/gin-gonic/gin"
)

type IncomeHandler struct {
    service income.IncomeService
}

func NewIncomeHandler(service income.IncomeService) *IncomeHandler {
    return &IncomeHandler{service: service}
}

// @Summary      Create a new income
// @Description  Create a new income record
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateIncomeRequest true "Income Data"
// @Success      201 {object} dto.IncomeResponse
// @Security     BearerAuth
// @Router       /incomes/create [post]
func (h *IncomeHandler) Create(c *gin.Context) {
    var request dto.CreateIncomeRequest
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

    response.Success(c, constant.MsgCreateIncomeSuccess, result)
}

// @Summary      Get all incomes
// @Description  Get paginated list of incomes with filters
// @Tags         incomes
// @Produce      json
// @Param        page query int false "Page number" minimum(1)
// @Param        limit query int false "Items per page" minimum(1) maximum(100)
// @Param        search query string false "Search in date and amount"
// @Param        sort_by query string false "Sort field"
// @Param        sort_dir query string false "Sort direction" Enums(asc, desc)
// @Success      200 {object} response.PaginatedResponse{data=[]dto.IncomeResponse}
// @Security     BearerAuth
// @Router       /incomes/all [get]
func (h *IncomeHandler) GetAll(c *gin.Context) {
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

    response.SuccessPaginated(c, constant.MsgGetIncomesSuccess, result, params, totalRecords)
}

// @Summary      Update income
// @Description  Update an existing income record
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Param        id path string true "Income ID"
// @Param        request body dto.UpdateIncomeRequest true "Income Data"
// @Success      200 {object} dto.IncomeResponse
// @Security     BearerAuth
// @Router       /incomes/{id} [put]
func (h *IncomeHandler) Update(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.Error(utils.NewValidationError(constant.MsgInvalidInput, nil))
        return
    }

    var request dto.UpdateIncomeRequest
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

    response.Success(c, constant.MsgUpdateIncomeSuccess, result)
}

// @Summary      Delete income
// @Description  Delete an existing income record
// @Tags         incomes
// @Produce      json
// @Param        id path string true "Income ID"
// @Success      200 {object} response.Response
// @Security     BearerAuth
// @Router       /incomes/{id} [delete]
func (h *IncomeHandler) Delete(c *gin.Context) {
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

    response.Success(c, constant.MsgDeleteIncomeSuccess, nil)
}
