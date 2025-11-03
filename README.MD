# 📰 Article API (Golang + Gin + GORM + MySQL)

## 🚀 Cara Menjalankan Project

### 1️⃣ Clone Repository
```bash
git clone https://github.com/username/article-api.git
cd article-api
```

### 2️⃣ Install Dependency
```bash
go mod tidy
```

### 3️⃣ Buat Database di MySQL
Buka terminal MySQL atau phpMyAdmin, lalu jalankan:
```sql
CREATE DATABASE article;
```

### 4️⃣ Cek Koneksi Database
Pastikan di file `database/connection.go` sudah sesuai:
```go
dsn := "root:@tcp(127.0.0.1:3306)/article?charset=utf8mb4&parseTime=True&loc=Local"
```

### 5️⃣ Jalankan Server
```bash
go run main.go
```

Server akan berjalan di:
```
http://localhost:8080
```

---

## 📚 Endpoint API

| Method | Endpoint | Deskripsi |
|---------|-----------|-----------|
| GET | `/article/:limit/:offset` | Menampilkan daftar artikel (paging) |
| GET | `/article/detail/:id` | Menampilkan artikel berdasarkan ID |
| POST | `/article/` | Menambahkan artikel baru |
| PUT | `/article/:id` | Mengubah artikel |
| DELETE | `/article/:id` | Menghapus artikel |

---

## 🧾 Validasi Input

Sebelum membuat atau mengubah artikel, pastikan JSON memenuhi syarat berikut:

| Field | Validasi |
|--------|-----------|
| title | required, minimal 20 karakter |
| content | required, minimal 200 karakter |
| category | required, minimal 3 karakter |
| status | required, hanya boleh `publish`, `draft`, atau `thrash` |

---

## ✅ Contoh Request (POST /article/)

```json
{
  "title": "Judul Artikel Minimal 20 Karakter",
  "content": "Isi artikel minimal 200 karakter...",
  "category": "Teknologi",
  "status": "publish"
}
```

---

## 🧑‍💻 Dibuat oleh
**Mohammad Rafa Adila**  
Fullstack Web Developer
