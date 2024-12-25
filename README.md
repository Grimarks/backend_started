# backend_started

Berikut adalah contoh isi file README.txt untuk masing-masing folder:

Nama Proyek: Backend Started  
Deskripsi:  
Proyek ini adalah backend sederhana yang dibangun menggunakan Golang dengan framework Gin dan GORM untuk mengelola data produk, pengguna, transaksi, dan pembayaran.  

http://localhost:8080/users
http://localhost:8080/products
http://localhost:8080/transactions
http://localhost:8080/payments


Struktur Direktori:  
- config: Berisi konfigurasi database.  
- controllers: Berisi fungsi untuk menangani request endpoint.  
- models: Berisi struktur data dan model database.  
- routes: Berisi konfigurasi rute untuk endpoint API.  

Cara Penggunaan:  
1. Clone repository ini ke komputer Anda.  
2. Pastikan Anda telah menginstal Golang, MySQL, dan Postman.  
3. Buat database baru di MySQL dengan nama sesuai yang didefinisikan di file .env.  
4. Isi file .env dengan informasi berikut:  
   - DB_USER: Nama pengguna database MySQL Anda.  
   - DB_PASSWORD: Password database MySQL Anda.  
   - DB_HOST: Host database (default: localhost).  
   - DB_PORT: Port database (default: 3306).  
   - DB_NAME: Nama database yang telah Anda buat.  
5. Jalankan perintah berikut untuk menginstal dependensi:  
   ```
   go mod tidy
   ```  
6. Jalankan server backend dengan perintah:  
   ```
   go run main.go
   ```  
7. Gunakan Postman untuk mengakses endpoint API sesuai dokumentasi.  

Fitur:  
- Menambahkan, mengambil, memperbarui, dan menghapus data produk.  
- Otentikasi pengguna menggunakan JWT token.  
- Proses transaksi dan pembayaran.  

Catatan:  
Pastikan file .env berada di root folder proyek ini agar konfigurasi database dapat terdeteksi dengan benar.
