package middleware

import (
    "github.com/gin-gonic/gin"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/response"
)

type AppError struct {
    Type      string
    Message   string
    Err       error
    Source    string
    ErrorCode string
}

func (e *AppError) Error() string {
    return e.Message
}

var errorHandlers = map[string]func(*gin.Context, string, string, error){
    constant.ErrTypeValidation: response.ValidationError,
    constant.ErrTypeBusiness:   response.BadRequest,
    constant.ErrTypeSystem:     response.InternalError,
    constant.ErrTypeNotFound:   response.NotFound,
    constant.ErrTypeAuth:       response.Unauthorized,
}

func NewAppErrorWithCode(errType, message, code, source string, err error) *AppError {
    return &AppError{
        Type:      errType,
        Message:   message,
        ErrorCode: code,
        Source:    source,
        Err:       err,
    }
}

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            if appErr, ok := err.(*AppError); ok {
                handleAppError(c, appErr)
                return
            }
            response.InternalError(c, constant.MsgInternalError, constant.ErrCodeInternalError, err)
        }
    }
}

func handleAppError(c *gin.Context, err *AppError) {
    if handler, exists := errorHandlers[err.Type]; exists {
        handler(c, err.Message, err.ErrorCode, err.Err)
        return
    }
    response.InternalError(c, err.Message, err.ErrorCode, err.Err)
}
