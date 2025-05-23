package utils

import (
    "pencatatan_keuangan/internal/middleware"
    "pencatatan_keuangan/pkg/constant"
)

func NewValidationError(message string, err error) error {
    return middleware.NewAppErrorWithCode(
        constant.ErrTypeValidation,
        message,
        constant.ErrCodeValidation,
        "",
        err,
    )
}

func NewBusinessError(message string) error {
    return middleware.NewAppErrorWithCode(
        constant.ErrTypeBusiness,
        message,
        constant.ErrCodeBadRequest,
        "",
        nil,
    )
}

func NewSystemError(message string, err error) error {
    return middleware.NewAppErrorWithCode(
        constant.ErrTypeSystem,
        message,
        constant.ErrCodeInternalError,
        "",
        err,
    )
}

func NewNotFoundError(message string) error {
    return middleware.NewAppErrorWithCode(
        constant.ErrTypeNotFound,
        message,
        constant.ErrCodeNotFound,
        "",
        nil,
    )
}

func NewAuthError(message string) error {
    return middleware.NewAppErrorWithCode(
        constant.ErrTypeAuth,
        message,
        constant.ErrCodeUnauthorized,
        "",
        nil,
    )
}
