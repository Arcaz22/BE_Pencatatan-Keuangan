package middleware

import (
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/jwt"
    "pencatatan_keuangan/pkg/response"
    "strings"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader(constant.HeaderAuthorization)
        if authHeader == "" {
            response.Unauthorized(c, "Authorization header is required", "UNAUTHORIZED", nil)
            c.Abort()
            return
        }

        var tokenString string
        if strings.HasPrefix(authHeader, "Bearer ") {
            parts := strings.Split(authHeader, " ")
            if len(parts) != 2 {
                response.Unauthorized(c, "Invalid authorization format", "UNAUTHORIZED", nil)
                c.Abort()
                return
            }
            tokenString = parts[1]
        } else {
            tokenString = authHeader
        }

        userID, err := jwt.ValidateToken(tokenString)
        if err != nil {
            response.Unauthorized(c, "Invalid or expired token", "UNAUTHORIZED", nil)
            c.Abort()
            return
        }

        c.Set("userID", userID)
        c.Next()
    }
}
