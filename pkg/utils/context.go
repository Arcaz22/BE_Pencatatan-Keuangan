package utils

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func GetUserID(c *gin.Context) (string, bool) {
    userID, exists := c.Get("userID")
    if !exists {
        return "", false
    }

    switch v := userID.(type) {
    case string:
        return v, true
    case uuid.UUID:
        return v.String(), true
    default:
        return "", false
    }
}
