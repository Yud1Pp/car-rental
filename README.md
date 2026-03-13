# Car Rental API

Backend API sederhana untuk manajemen rental mobil menggunakan Go, Fiber v3, GORM, PostgreSQL, dan Swagger.

API ini menyediakan:

- CRUD customer
- CRUD car
- CRUD booking

Business rule utama pada booking:

- membuat booking mengurangi stok mobil
- `total_cost` dihitung otomatis dari durasi sewa dan `daily_rent`
- update booking menghitung ulang `total_cost`
- perubahan status `finished` memengaruhi stok mobil
- perpindahan mobil pada booking aktif akan menyesuaikan stok mobil lama dan baru

## Teknologi

- Go
- Fiber v3
- GORM
- PostgreSQL
- Swaggo

## ERD

Ringkasan relasi:

- satu customer dapat memiliki banyak booking
- satu car dapat memiliki banyak booking
- satu booking terkait ke satu customer dan satu car

Gambar ERD:

![ERD Car Rental](./erd_car_rental.png)

## Endpoint

Base URL:

```text
/api/v1
```

Health check:

- `GET /ping`

Customer:

- `GET /api/v1/customers/`
- `GET /api/v1/customers/:id`
- `POST /api/v1/customers/`
- `PUT /api/v1/customers/:id`
- `DELETE /api/v1/customers/:id`

Car:

- `GET /api/v1/cars/`
- `GET /api/v1/cars/:id`
- `POST /api/v1/cars/`
- `PUT /api/v1/cars/:id`
- `DELETE /api/v1/cars/:id`

Booking:

- `GET /api/v1/bookings/`
- `GET /api/v1/bookings/:id`
- `POST /api/v1/bookings/`
- `PUT /api/v1/bookings/:id`
- `DELETE /api/v1/bookings/:id`

## Menjalankan Project

### 1. Siapkan `.env`

```env
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=car_rental
```

### 2. Install dependency dan generate Swagger docs

```bash
go mod tidy
swag init -g cmd/main.go -o docs
```

### 3. Jalankan aplikasi

```bash
go run ./cmd
```

Server:

```text
http://localhost:3000
```

Swagger UI:

```text
http://localhost:3000/swagger/index.html
```

## Menguji API

API bisa diuji menggunakan:

- Swagger UI
- `curl`
- Postman
- REST Client di VS Code

Cara paling mudah untuk project ini adalah lewat Swagger UI.

### Uji dengan Swagger UI

1. Jalankan aplikasi:

```bash
go run ./cmd
```

2. Buka:

```text
http://localhost:3000/swagger/index.html
```

3. Pilih endpoint yang ingin diuji, misalnya `POST /customers`

4. Klik `Try it out`

5. Isi body JSON

Contoh create customer:

```json
{
  "name": "Yudi Pratama",
  "nik": "1234567890",
  "phone_number": "089789101234"
}
```

6. Klik `Execute`

7. Lihat:

- response body
- response code
- curl command yang dihasilkan Swagger

### Uji dengan PowerShell

Contoh `curl.exe`:

```powershell
curl.exe -X POST http://localhost:3000/api/v1/customers/ `
  -H "Content-Type: application/json" `
  -d "{\"name\":\"Yudi Pratama\",\"nik\":\"1234567890\",\"phone_number\":\"089789101234\"}"
```

Contoh `Invoke-RestMethod`:

```powershell
Invoke-RestMethod -Method POST `
  -Uri "http://localhost:3000/api/v1/customers/" `
  -ContentType "application/json" `
  -Body '{"name":"Yudi Pratama","nik":"1234567890","phone_number":"089789101234"}'
```

## Contoh Request Body

Create / update customer:

```json
{
  "name": "Yudi Pratama",
  "nik": "1234567890",
  "phone_number": "089789101234"
}
```

Create / update car:

```json
{
  "name": "Toyota Camry",
  "stock": 1,
  "daily_rent": 1000000
}
```

Create booking:

```json
{
  "customer_id": 1,
  "car_id": 1,
  "start_rent": "2026-03-13T09:00:00Z",
  "end_rent": "2026-03-15T09:00:00Z"
}
```

Update booking:

```json
{
  "customer_id": 1,
  "car_id": 1,
  "start_rent": "2026-03-13T09:00:00Z",
  "end_rent": "2026-03-16T09:00:00Z",
  "finished": false
}
```

## Skenario Pengujian End-to-End

Cerita pengujian:

Sebuah rental mobil menerima customer baru bernama Yudi Pratama. Admin menambahkan dua mobil ke katalog. Setelah itu Yudi membuat booking. Saat booking dibuat, stok mobil harus turun dan total biaya harus terhitung otomatis. Lalu durasi booking diperpanjang, mobil booking dipindah ke mobil lain, dan akhirnya booking diselesaikan sehingga stok mobil kembali naik. Setelah semua valid, admin menghapus data booking, mobil, dan customer.

### Menjalankan Skenario Ini di Swagger UI

Skenario end-to-end di bawah ini juga bisa dijalankan sepenuhnya lewat Swagger UI, bukan hanya dengan `curl`.

Langkahnya:

1. Jalankan aplikasi:

```bash
go run ./cmd
```

2. Buka Swagger UI:

```text
http://localhost:3000/swagger/index.html
```

3. Jalankan endpoint sesuai urutan skenario di bawah, mulai dari:

- `POST /customers`
- `GET /customers`
- `PUT /customers/{id}`
- `POST /cars`
- `GET /cars`
- `POST /bookings`
- `PUT /bookings/{id}`
- `GET /bookings/{id}`
- `DELETE /bookings/{id}`

4. Untuk setiap endpoint:

- klik endpoint-nya
- klik `Try it out`
- isi parameter path jika ada, misalnya `id`
- isi request body JSON jika endpoint menerima body
- klik `Execute`

5. Simpan `id` dari hasil create:

- `customer id` dari `POST /customers`
- `car id` dari `POST /cars`
- `booking id` dari `POST /bookings`

6. Gunakan `id` tersebut pada langkah berikutnya, misalnya:

- `GET /customers/{id}`
- `PUT /cars/{id}`
- `PUT /bookings/{id}`
- `DELETE /bookings/{id}`

7. Setelah tiap langkah penting, cek response body untuk memastikan business rule berjalan:

- setelah create booking, stok mobil berkurang
- setelah pindah mobil, stok mobil lama dan baru berubah sesuai aturan
- setelah `finished=true`, stok mobil kembali bertambah
- setelah update tanggal atau mobil, `total_cost` berubah sesuai kondisi terbaru

8. Jika ingin memeriksa data akhir dengan cepat, jalankan endpoint list/detail:

- `GET /customers`
- `GET /cars`
- `GET /bookings`

### Langkah Uji

1. Cek health endpoint

```bash
curl http://localhost:3000/ping
```

2. Buat customer

```bash
curl -X POST http://localhost:3000/api/v1/customers/ \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Yudi Pratama\",\"nik\":\"1234567890\",\"phone_number\":\"089789101234\"}"
```

3. Lihat semua customer

```bash
curl http://localhost:3000/api/v1/customers/
```

4. Lihat detail customer

```bash
curl http://localhost:3000/api/v1/customers/1
```

5. Update customer

```bash
curl -X PUT http://localhost:3000/api/v1/customers/1 \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Yudi Pratama Update\",\"nik\":\"1234567890\",\"phone_number\":\"081234567890\"}"
```

6. Tambah mobil pertama

```bash
curl -X POST http://localhost:3000/api/v1/cars/ \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Toyota Camry\",\"stock\":1,\"daily_rent\":1000000}"
```

7. Tambah mobil kedua

```bash
curl -X POST http://localhost:3000/api/v1/cars/ \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Honda Civic\",\"stock\":2,\"daily_rent\":800000}"
```

8. Lihat semua mobil

```bash
curl http://localhost:3000/api/v1/cars/
```

9. Lihat detail mobil

```bash
curl http://localhost:3000/api/v1/cars/1
curl http://localhost:3000/api/v1/cars/2
```

10. Update mobil

```bash
curl -X PUT http://localhost:3000/api/v1/cars/2 \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Honda Civic RS\",\"stock\":2,\"daily_rent\":850000}"
```

11. Buat booking

```bash
curl -X POST http://localhost:3000/api/v1/bookings/ \
  -H "Content-Type: application/json" \
  -d "{\"customer_id\":1,\"car_id\":1,\"start_rent\":\"2026-03-13T09:00:00Z\",\"end_rent\":\"2026-03-15T09:00:00Z\"}"
```

12. Lihat semua booking

```bash
curl http://localhost:3000/api/v1/bookings/
```

13. Lihat detail booking

```bash
curl http://localhost:3000/api/v1/bookings/1
```

14. Perpanjang booking

```bash
curl -X PUT http://localhost:3000/api/v1/bookings/1 \
  -H "Content-Type: application/json" \
  -d "{\"customer_id\":1,\"car_id\":1,\"start_rent\":\"2026-03-13T09:00:00Z\",\"end_rent\":\"2026-03-16T09:00:00Z\",\"finished\":false}"
```

15. Pindahkan booking ke mobil lain

```bash
curl -X PUT http://localhost:3000/api/v1/bookings/1 \
  -H "Content-Type: application/json" \
  -d "{\"customer_id\":1,\"car_id\":2,\"start_rent\":\"2026-03-13T09:00:00Z\",\"end_rent\":\"2026-03-16T09:00:00Z\",\"finished\":false}"
```

16. Selesaikan booking

```bash
curl -X PUT http://localhost:3000/api/v1/bookings/1 \
  -H "Content-Type: application/json" \
  -d "{\"customer_id\":1,\"car_id\":2,\"start_rent\":\"2026-03-13T09:00:00Z\",\"end_rent\":\"2026-03-16T09:00:00Z\",\"finished\":true}"
```

17. Verifikasi data akhir

```bash
curl http://localhost:3000/api/v1/bookings/1
curl http://localhost:3000/api/v1/cars/1
curl http://localhost:3000/api/v1/cars/2
```

18. Hapus booking

```bash
curl -X DELETE http://localhost:3000/api/v1/bookings/1
```

19. Hapus mobil

```bash
curl -X DELETE http://localhost:3000/api/v1/cars/1
curl -X DELETE http://localhost:3000/api/v1/cars/2
```

20. Hapus customer

```bash
curl -X DELETE http://localhost:3000/api/v1/customers/1
```

## Checklist Hasil yang Diharapkan

- customer berhasil dibuat, dilihat, diupdate, dan dihapus
- car berhasil dibuat, dilihat, diupdate, dan dihapus
- booking berhasil dibuat, dilihat, diupdate, dan dihapus
- stok mobil turun saat booking dibuat
- stok mobil menyesuaikan saat booking dipindah ke mobil lain
- stok mobil naik kembali saat booking selesai
- `total_cost` selalu dihitung otomatis oleh service
- response booking memuat relasi `customer` dan `car`

## Catatan

- project memakai `AutoMigrate` saat aplikasi start
- setelah mengubah anotasi Swagger, generate ulang docs:

```bash
swag init -g cmd/main.go -o docs
```
