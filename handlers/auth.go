package handlers

import (
	"net/http"

	"github.com/RafiAwanda123/Finance-UMKM/database"
	"github.com/RafiAwanda123/Finance-UMKM/models"
	"github.com/RafiAwanda123/Finance-UMKM/utils"

	"github.com/gin-gonic/gin"
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

	// Cek username sudah ada
	if _, err := models.GetUserByUsername(database.DB, input.Username); err == nil {
		utils.APIError(c, http.StatusConflict, "Username sudah terdaftar")
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal memproses password")
		return
	}

	// Simpan user baru
	newUser := models.User{
		Username:     input.Username,
		PasswordHash: hashedPassword,
	}

	if err := models.CreateUser(database.DB, &newUser); err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal membuat user")
		return
	}

	utils.APISuccess(c, http.StatusCreated, "User berhasil dibuat")
}

func Login(c *gin.Context) {
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIError(c, http.StatusBadRequest, "Format input tidak valid")
		return
	}

	// Cari user
	user, err := models.GetUserByUsername(database.DB, input.Username)
	if err != nil {
		utils.APIError(c, http.StatusUnauthorized, "Username/password salah")
		return
	}

	// Verifikasi password
	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		utils.APIError(c, http.StatusUnauthorized, "Username/password salah")
		return
	}

	// Generate JWT 2 jam
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	utils.APISuccess(c, http.StatusOK, gin.H{
		"token":      token,
		"expires_in": 7200,
	})
}
