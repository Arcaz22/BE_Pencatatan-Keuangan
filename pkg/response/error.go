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

// ErrorResponse adalah struktur respons untuk error
type ErrorResponse struct {
    Message      string `json:"message"`
    ErrorCode    string `json:"error_code,omitempty"`    // Kode error opsional
    ErrorRef     string `json:"error_ref,omitempty"`     // UUID untuk penelusuran error
    ErrorDetails string `json:"error_details,omitempty"` // Detail error (hanya di development)
}

// Error memformat dan mengembalikan respons error final
func Error(c *gin.Context, statusCode int, message string, errorCode string, err error) {
    errorRef := uuid.New().String()

    env := os.Getenv("APP_ENV")
    if env == "" {
        env = EnvDevelopment
    }

    res := ErrorResponse{
        Message:   message,
        ErrorCode: chooseErrorCode(errorCode), // Bisa kosong
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

// logError mencatat error ke log sistem
func logError(errorRef string, errorMessage string) {
    timeStr := time.Now().Format(time.RFC3339)
    logMessage := "ERROR [" + timeStr + "] REF: " + errorRef + " - " + errorMessage
    print(logMessage + "\n")
}

// chooseErrorCode menentukan kode error yang akan digunakan
func chooseErrorCode(code string) string {
    return code
}

// Fungsi-fungsi helper untuk error

// BadRequest mengirimkan respons error 400
func BadRequest(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusBadRequest, message, errorCode, err)
}

// Unauthorized mengirimkan respons error 401
func Unauthorized(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusUnauthorized, message, errorCode, err)
}

// Forbidden mengirimkan respons error 403
func Forbidden(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusForbidden, message, errorCode, err)
}

// NotFound mengirimkan respons error 404
func NotFound(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusNotFound, message, errorCode, err)
}

// InternalError mengirimkan respons error 500
func InternalError(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusInternalServerError, message, errorCode, err)
}

// ValidationError mengirimkan respons error 422
func ValidationError(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusUnprocessableEntity, message, errorCode, err)
}

// Conflict mengirimkan respons error 409
func Conflict(c *gin.Context, message, errorCode string, err error) {
    Error(c, http.StatusConflict, message, errorCode, err)
}
