package handler

import (
    "pencatatan_keuangan/internal/service/dashboard"
    "pencatatan_keuangan/internal/service/dashboard/dto"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
    "pencatatan_keuangan/pkg/utils"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type DashboardHandler struct {
    service dashboard.DashboardService
}

func NewDashboardHandler(service dashboard.DashboardService) *DashboardHandler {
    return &DashboardHandler{service: service}
}

// @Summary      Get dashboard data
// @Description  Get financial dashboard data including summaries, charts and monthly breakdowns
// @Tags         dashboard
// @Produce      json
// @Param        year query int false "Year for filtering data (default: current year)"
// @Param        month query int false "Month for filtering data (optional, 1-12)"
// @Success      200 {object} dto.DashboardResponse
// @Security     BearerAuth
// @Router       /dashboard [get]
func (h *DashboardHandler) GetDashboard(c *gin.Context) {
    var filter dto.FilterRequest
    if err := c.ShouldBindQuery(&filter); err != nil {
        c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
        return
    }

    userIDStr, exists := utils.GetUserID(c)
    if !exists {
        c.Error(utils.NewValidationError(constant.ErrCodeUnauthorized, nil))
        return
    }

    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
        return
    }

    dashboardData, err := h.service.GetDashboard(userID, filter)
    if err != nil {
        c.Error(err)
        return
    }

    response.Success(c, constant.MsgGetDashboardSuccess, dashboardData)
}
