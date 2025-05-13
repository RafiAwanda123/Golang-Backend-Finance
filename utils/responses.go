package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// API Response Standard
type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// APISuccess untuk response sukses
func APISuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, APIResponse{
		Status:  status,
		Message: "success",
		Data:    data,
	})
}

// APIError untuk response error
func APIError(c *gin.Context, status int, message string) {
	c.JSON(status, APIResponse{
		Status:  status,
		Message: message,
	})
	c.Abort()
}

// HashPassword mengenkripsi password dengan bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash memverifikasi password dengan hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetQueryParamInt helper untuk mengambil query param integer
func GetQueryParamInt(c *gin.Context, key string, defaultValue int) int {
	param := c.Query(key)
	if param == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(param)
	if err != nil {
		return defaultValue
	}
	return value
}
