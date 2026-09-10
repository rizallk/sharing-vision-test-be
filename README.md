# Sharing Vision - Backend Test

Dibuat oleh : Mohalim Rizal Kadamong

Backend ini dibuat menggunakan bahasa **Go** dengan framework **Fiber**.

## Spesifikasi

- **Framework:** Go Fiber v2
- **ORM:** GORM
- **Database:** MySQL

---

## Struktur Folder Project

```text
sharing-vision-test-be/
├── config/         # Konfigurasi Database & Environment
├── controllers/    # Handler HTTP Request & Response
├── models/         # Struct Database & Request DTO
├── repositories/   # Komunikasi langsung dengan Database (GORM)
├── routes/         # Definisi Routing Endpoint API
├── services/       # Business Logic & Validasi
├── main.go         # Entry point aplikasi
└── go.mod          # Go modules dependencies
```

## Cara Menjalankan Project

1. **Prasyarat:**
   - Pastikan Golang dan MySQL telah terinstal.
2. **Setup Database:**
   - Buat database baru dengan nama `article` dan tabel `posts` sesuai dengan struktur yang ada di file `models/post.go`.
3. **Install Dependencies:**
   ```bash
   go mod tidy
   ```
4. **Jalankan Aplikasi:**
   ```bash
   go run main.go
   ```
