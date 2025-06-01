package response

import (
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

const (
    EnvDevelopment = "development"
    EnvProduction  = "production"
)
type ErrorResponse struct {
    Message      string `json:"message"`
    ErrorCode    string `json:"error_code,omitempty"`
    ErrorRef     string `json:"error_ref,omitempty"`
    ErrorDetails string `json:"error_details,omitempty"`
}

func Error(c *gin.Context, statusCode int, message string, errorCode string, err error) {
    errorRef := uuid.New().String()

    env := os.Getenv("APP_ENV")
    if env == "" {
        env = EnvDevelopment
    }

    res := ErrorResponse{
        Message:   message,
        ErrorCode: chooseErrorCode(errorCode),
        ErrorRef:  errorRef,
    }

    if env == EnvDevelopment && err != nil {
        res.ErrorDetails = err.Error()
    }

    if err != nil {
        logError(errorRef, err.Error())
    }

    c.JSON(statusCode, res)
}

func logError(errorRef string, errorMessage string) {
    timeStr := time.Now().Format(time.RFC3339)
    logMessage := "ERROR [" + timeStr + "] REF: " + errorRef + " - " + errorMessage
    print(logMessage + "\n")
}

func chooseErrorCode(code string) string {
    return code
}

func BadRequest(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusBadRequest, message, errorCode, err)
}

func Unauthorized(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusUnauthorized, message, errorCode, err)
}

func Forbidden(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusForbidden, message, errorCode, err)
}

func NotFound(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusNotFound, message, errorCode, err)
}

func InternalError(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusInternalServerError, message, errorCode, err)
}

func ValidationError(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusUnprocessableEntity, message, errorCode, err)
}

func Conflict(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusConflict, message, errorCode, err)
}
