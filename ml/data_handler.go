package ml

import (
	"time"

	"github.com/RafiAwanda123/Finance-UMKM/database"
	"github.com/RafiAwanda123/Finance-UMKM/models"
)

// HistoricalData - Struktur data untuk analisis ML
type HistoricalData struct {
	Date   time.Time `json:"date"`
	Sales  int       `json:"sales"`
	Price  float64   `json:"price"`
	Margin float64   `json:"margin"`
}

// FetchSalesData - Ambil data historis untuk analisis
func FetchSalesData(userID int, item string) ([]HistoricalData, error) {
	query := `
        SELECT created_at, penjualan, harga, margin
        FROM financial_records 
        WHERE user_id = ? AND barang = ?
        ORDER BY created_at ASC`

	rows, err := database.DB.Query(query, userID, item)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []HistoricalData
	for rows.Next() {
		var record models.FinancialRecord
		err := rows.Scan(
			&record.CreatedAt,
			&record.Penjualan,
			&record.Harga,
			&record.Margin,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, HistoricalData{
			Date:   record.CreatedAt,
			Sales:  record.Penjualan,
			Price:  record.Harga,
			Margin: record.Margin,
		})
	}
	return data, nil
}

// FetchTimeSeriesData - Ambil data dalam format time series
func FetchTimeSeriesData(userID int, days int) ([]HistoricalData, error) {
	query := `
        SELECT created_at, SUM(penjualan), AVG(harga), AVG(margin)
        FROM financial_records 
        WHERE user_id = ? AND created_at >= ?
        GROUP BY created_at
        ORDER BY created_at ASC`

	startDate := time.Now().AddDate(0, 0, -days)

	rows, err := database.DB.Query(query, userID, startDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timeSeries []HistoricalData
	for rows.Next() {
		var hd HistoricalData
		var date string
		err := rows.Scan(
			&date,
			&hd.Sales,
			&hd.Price,
			&hd.Margin,
		)
		if err != nil {
			return nil, err
		}

		parsedDate, _ := time.Parse("2006-01-02", date)
		hd.Date = parsedDate
		timeSeries = append(timeSeries, hd)
	}
	return timeSeries, nil
}
