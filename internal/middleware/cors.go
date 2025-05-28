package middleware

import (
    "pencatatan_keuangan/pkg/constant"
    "github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set(constant.HeaderAccessControlAllowOrigin, constant.AllowOriginAll)
        c.Writer.Header().Set(constant.HeaderAccessControlAllowMethods, constant.AllowMethods)
        c.Writer.Header().Set(constant.HeaderAccessControlAllowHeaders, constant.AllowHeadersBasic)

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}
