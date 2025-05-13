-- Hapus tabel jika sudah ada (untuk testing)
DROP TABLE IF EXISTS financial_records;
DROP TABLE IF EXISTS users;

-- Tabel Users
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

-- Tabel Financial Records
CREATE TABLE IF NOT EXISTS financial_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    barang VARCHAR(255) NOT NULL,
    harga DECIMAL(10,2) NOT NULL,
    tipe ENUM('pemasukan', 'pengeluaran') NOT NULL,
    penjualan INT NOT NULL,
    margin DECIMAL(5,2) NOT NULL,
    supply INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;