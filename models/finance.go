package models

import (
	"database/sql"
	"time"
)

// FinancialRecord - Struktur data keuangan (sesuai database)
type FinancialRecord struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Barang    string    `json:"barang"`
	Harga     float64   `json:"harga"`
	Tipe      string    `json:"tipe"`
	Penjualan int       `json:"penjualan"`
	Margin    float64   `json:"margin"`
	Supply    int       `json:"supply"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAllFinanceByUser - Ambil semua data keuangan user
func GetAllFinanceByUser(db *sql.DB, userID int) ([]FinancialRecord, error) {
	query := `
        SELECT id, barang, harga, tipe, penjualan, margin, supply, created_at
        FROM financial_records 
        WHERE user_id = ? 
        ORDER BY created_at DESC`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []FinancialRecord
	for rows.Next() {
		var r FinancialRecord
		err := rows.Scan(
			&r.ID,
			&r.Barang,
			&r.Harga,
			&r.Tipe,
			&r.Penjualan,
			&r.Margin,
			&r.Supply,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, nil
}

// FinanceMetadata - Struktur untuk metadata keuangan
type FinanceMetadata struct {
	TotalPenjualan   int     `json:"total_penjualan"`
	AverageMargin    float64 `json:"rata_margin"`
	TotalPemasukan   float64 `json:"total_pemasukan"`
	TotalPengeluaran float64 `json:"total_pengeluaran"`
}

// GetFinanceMetadata - Ambil metadata agregat keuangan
func GetFinanceMetadata(db *sql.DB, userID int) (FinanceMetadata, error) {
	query := `
        SELECT 
            SUM(penjualan) AS total_penjualan,
            AVG(margin) AS rata_margin,
            SUM(CASE WHEN tipe = 'pemasukan' THEN harga ELSE 0 END) AS total_pemasukan,
            SUM(CASE WHEN tipe = 'pengeluaran' THEN harga ELSE 0 END) AS total_pengeluaran
        FROM financial_records 
        WHERE user_id = ?`

	var meta FinanceMetadata
	err := db.QueryRow(query, userID).Scan(
		&meta.TotalPenjualan,
		&meta.AverageMargin,
		&meta.TotalPemasukan,
		&meta.TotalPengeluaran,
	)

	return meta, err
}

// GetFinanceByID - Ambil data berdasarkan ID
func GetFinanceByID(db *sql.DB, id string, userID int) (FinancialRecord, error) {
	query := `
        SELECT id, barang, harga, tipe, penjualan, margin, supply, created_at
        FROM financial_records 
        WHERE id = ? AND user_id = ?`

	var record FinancialRecord
	err := db.QueryRow(query, id, userID).Scan(
		&record.ID,
		&record.Barang,
		&record.Harga,
		&record.Tipe,
		&record.Penjualan,
		&record.Margin,
		&record.Supply,
		&record.CreatedAt,
	)

	return record, err
}

// CreateFinance - Tambah data keuangan baru
func CreateFinance(db *sql.DB, record *FinancialRecord) error {
	query := `
        INSERT INTO financial_records 
        (user_id, barang, harga, tipe, penjualan, margin, supply) 
        VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(
		query,
		record.UserID,
		record.Barang,
		record.Harga,
		record.Tipe,
		record.Penjualan,
		record.Margin,
		record.Supply,
	)
	return err
}

// UpdateFinance - Update data keuangan
func UpdateFinance(db *sql.DB, record *FinancialRecord) error {
	query := `
        UPDATE financial_records 
        SET barang = ?, harga = ?, tipe = ?, penjualan = ?, margin = ?, supply = ? 
        WHERE id = ? AND user_id = ?`

	_, err := db.Exec(
		query,
		record.Barang,
		record.Harga,
		record.Tipe,
		record.Penjualan,
		record.Margin,
		record.Supply,
		record.ID,
		record.UserID,
	)
	return err
}

// DeleteFinance - Hapus data keuangan
func DeleteFinance(db *sql.DB, id string, userID int) error {
	query := "DELETE FROM financial_records WHERE id = ? AND user_id = ?"
	_, err := db.Exec(query, id, userID)
	return err
}

// GetSalesByItem - Ambil data penjualan per item (untuk analisis)
func GetSalesByItem(db *sql.DB, userID int, item string) ([]FinancialRecord, error) {
	query := `
        SELECT penjualan, created_at 
        FROM financial_records 
        WHERE user_id = ? AND barang = ?
        ORDER BY created_at ASC`

	rows, err := db.Query(query, userID, item)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []FinancialRecord
	for rows.Next() {
		var r FinancialRecord
		// Langsung scan ke field CreatedAt milik struct FinancialRecord
		err := rows.Scan(&r.Penjualan, &r.CreatedAt) // <-- Perbaikan di sini
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}

	return records, nil
}
