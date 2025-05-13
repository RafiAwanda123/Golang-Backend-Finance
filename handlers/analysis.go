package handlers

import (
	"net/http"

	"github.com/RafiAwanda123/Finance-UMKM/database"
	"github.com/RafiAwanda123/Finance-UMKM/ml"
	"github.com/RafiAwanda123/Finance-UMKM/models"

	"github.com/gin-gonic/gin"
)

type PredictionResponse struct {
	Item       string  `json:"item"`
	LinearPred float64 `json:"prediksi_linear"`
	NaivePred  float64 `json:"prediksi_naive"`
}

func AnalysisHandler(c *gin.Context) {
	// Ambil parameter dari URL
	item := c.Param("item")
	userID := c.MustGet("user_id").(int)

	// Ambil data historis dari database
	salesData, err := models.GetSalesByItem(database.DB, userID, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	// Jika tidak ada data
	if len(salesData) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data historis tidak ditemukan"})
		return
	}

	// Ekstrak data penjualan
	var sales []int
	for _, d := range salesData {
		sales = append(sales, d.Penjualan)
	}

	// Lakukan prediksi
	response := PredictionResponse{
		Item:       item,
		LinearPred: ml.PredictLinearRegression(sales),
		NaivePred:  ml.PredictNaive(sales),
	}

	c.JSON(http.StatusOK, response)
}
