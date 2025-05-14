package handlers

import (
	"net/http"

	"github.com/RafiAwanda123/Finance-UMKM/database"
	"github.com/RafiAwanda123/Finance-UMKM/models"
	"github.com/RafiAwanda123/Finance-UMKM/utils"
	"github.com/gin-gonic/gin"

	"errors"

	"gorm.io/gorm"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIError(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	// Revisi: Gunakan blank identifier jika tidak membutuhkan return value
	_, err := models.GetUserByUsername(database.DB, input.Username)
	if err == nil { // Jika tidak ada error = username sudah ada
		utils.APIError(c, http.StatusConflict, "Username sudah terdaftar")
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) { // Jika error selain "not found"
		utils.APIError(c, http.StatusInternalServerError, "Gagal memeriksa username")
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal memproses password")
		return
	}

	newUser := models.User{
		Username:     input.Username,
		PasswordHash: hashedPassword,
	}

	if err := models.CreateUser(database.DB, &newUser); err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal membuat user")
		return
	}

	utils.APISuccess(c, http.StatusCreated, gin.H{
		"message":  "User berhasil dibuat",
		"username": newUser.Username,
	})
}

func Login(c *gin.Context) {
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIError(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	user, err := models.GetUserByUsername(database.DB, input.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.APIError(c, http.StatusUnauthorized, "Username tidak terdaftar")
		} else {
			utils.APIError(c, http.StatusInternalServerError, "Gagal memeriksa user")
		}
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		utils.APIError(c, http.StatusUnauthorized, "Password salah")
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	utils.APISuccess(c, http.StatusOK, gin.H{
		"token":      token,
		"expires_in": 7200,
		"username":   user.Username,
	})
}
