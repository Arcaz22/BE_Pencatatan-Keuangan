package handler

import (
	"pencatatan_keuangan/internal/service/user"
	"pencatatan_keuangan/internal/service/user/dto"
	"pencatatan_keuangan/pkg/constant"
	"pencatatan_keuangan/pkg/response"
	"pencatatan_keuangan/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary      Register a new user
// @Description  Register a new user with name, email and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body dto.RegisterRequest true "User Registration Data"
// @Success      201 {object} dto.RegisterResponse
// @Router       /users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
		return
	}

	result, err := h.service.Register(request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, constant.MsgRegisterSuccess, result)
}

// @Summary      User sign in
// @Description  Authenticates a user and returns a JWT token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body dto.SignInRequest true "User Credentials"
// @Success      200 {object} dto.SignInResponse
// @Router       /users/signin [post]
func (h *UserHandler) SignIn(c *gin.Context) {
	var request dto.SignInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(utils.NewValidationError(constant.MsgInvalidInput, err))
		return
	}

	result, err := h.service.SignIn(request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, constant.MsgLoginSuccess, result)
}

// @Summary      Get user profile
// @Description  Returns the profile information for the authenticated user
// @Tags         users
// @Produce      json
// @Success      200 {object} dto.UserInfo
// @Security     BearerAuth
// @Router       /users/profile [get]
func (h *UserHandler) Profile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.Error(utils.NewAuthError(constant.MsgUnauthorized))
		return
	}

	user, err := h.service.GetUserByID(userID.(uuid.UUID))
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, constant.MsgRetrivedUserSuccess, user)
}

// @Summary      Logout user
// @Description  Logout user (frontend should remove token)
// @Tags         users
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /users/logout [post]
func (h *UserHandler) Logout(c *gin.Context) {
    response.Success(c, constant.MsgLogoutSuccess, nil)
}
