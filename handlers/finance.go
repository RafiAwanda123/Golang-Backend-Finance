package handlers

import (
	"database/sql"
	"net/http"

	"github.com/RafiAwanda123/Finance-UMKM/database"
	"github.com/RafiAwanda123/Finance-UMKM/models"
	"github.com/RafiAwanda123/Finance-UMKM/utils"

	"github.com/gin-gonic/gin"
)

// GetAllFinance - Get all financial records for current user
func GetAllFinance(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	records, err := models.GetAllFinanceByUser(database.DB, userID)
	if err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal mengambil data")
		return
	}

	utils.APISuccess(c, http.StatusOK, records)
}

// GetFinanceInfo - Get metadata keuangan user
func GetFinanceInfo(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	// Panggil fungsi model untuk ambil metadata
	metadata, err := models.GetFinanceMetadata(database.DB, userID)
	if err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal mengambil metadata")
		return
	}

	utils.APISuccess(c, http.StatusOK, metadata)
}

// GetFinanceByID - Get single record by ID
func GetFinanceByID(c *gin.Context) {
	userID := c.MustGet("user_id").(int)
	recordID := c.Param("id")

	record, err := models.GetFinanceByID(database.DB, recordID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.APIError(c, http.StatusNotFound, "Data tidak ditemukan")
			return
		}
		utils.APIError(c, http.StatusInternalServerError, "Error database")
		return
	}

	utils.APISuccess(c, http.StatusOK, record)
}

// AddFinance - Create new financial record
func AddFinance(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	var record models.FinancialRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		utils.APIError(c, http.StatusBadRequest, "Format JSON tidak valid")
		return
	}

	record.UserID = userID

	if err := models.CreateFinance(database.DB, &record); err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal menyimpan data")
		return
	}

	utils.APISuccess(c, http.StatusCreated, gin.H{
		"message": "Data berhasil ditambahkan",
		"id":      record.ID,
	})
}

// UpdateFinance - Update existing record
func UpdateFinance(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	var record models.FinancialRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		utils.APIError(c, http.StatusBadRequest, "Format JSON tidak valid")
		return
	}

	record.UserID = userID

	if err := models.UpdateFinance(database.DB, &record); err != nil {
		if err == sql.ErrNoRows {
			utils.APIError(c, http.StatusNotFound, "Data tidak ditemukan")
			return
		}
		utils.APIError(c, http.StatusInternalServerError, "Gagal update data")
		return
	}

	utils.APISuccess(c, http.StatusOK, "Data berhasil diperbarui")
}

// DeleteFinance - Delete record
func DeleteFinance(c *gin.Context) {
	userID := c.MustGet("user_id").(int)
	recordID := c.Query("id")

	if err := models.DeleteFinance(database.DB, recordID, userID); err != nil {
		utils.APIError(c, http.StatusInternalServerError, "Gagal menghapus data")
		return
	}

	utils.APISuccess(c, http.StatusOK, "Data berhasil dihapus")
}
