package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RafiAwanda123/Finance-UMKM/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(cfg *config.AppConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Gagal konek ke database: %v", err))
	}

	// Test koneksi
	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Database tidak merespon: %v", err))
	}

	// Konfigurasi koneksi
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Koneksi database sukses!")

	// Jalankan migrasi
	runMigrations()
}

func runMigrations() {
	migrator, err := os.ReadFile("database/migrations.sql")
	if err != nil {
		panic("Gagal baca file migrasi: " + err.Error())
	}

	// Eksekusi per statement
	statements := strings.Split(string(migrator), ";")

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		_, err = DB.Exec(stmt)
		if err != nil {
			panic("Gagal jalankan migrasi: " + err.Error() + "\nPada statement: " + stmt)
		}
	}

	fmt.Println("Migrasi database berhasil!")
}
