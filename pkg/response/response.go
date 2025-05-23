package response

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Response adalah struktur standar untuk API response
type Response struct {
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

// Send mengirimkan respons dengan status code yang ditentukan
func Send(c *gin.Context, statusCode int, message string, data interface{}) {
    c.JSON(statusCode, Response{
        Message: message,
        Data:    data,
    })
}

// Success mengirimkan respons sukses dengan status 200 OK
func Success(c *gin.Context, message string, data interface{}) {
    Send(c, http.StatusOK, message, data)
}

// Created mengirimkan respons sukses dengan status 201 Created
func Created(c *gin.Context, message string, data interface{}) {
    Send(c, http.StatusCreated, message, data)
}

// NoContent mengirimkan respons tanpa konten (204)
func NoContent(c *gin.Context) {
    c.Status(http.StatusNoContent)
}

// Accepted mengirimkan respons dengan status 202 Accepted
func Accepted(c *gin.Context, message string, data interface{}) {
    Send(c, http.StatusAccepted, message, data)
}
