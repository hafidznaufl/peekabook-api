# PeekaBook API

Aplikasi "Peekabook" adalah sistem manajemen peminjaman buku yang memungkinkan pengguna untuk meminjam buku, membuat pesan, dan melihat informasi buku yang tersedia. API ini menyediakan layanan yang diperlukan untuk mengakses dan mengelola data dalam sistem ini.

## Daftar Isi
- [Instalasi](#instalasi)
- [Penggunaan API](#penggunaan-api)
- [Endpoints](#endpoints)
- [Desain Database](#desain-database)
- [Kontribusi](#kontribusi)
- [Lisensi](#lisensi)

## Instalasi

Untuk menjalankan API ini, Anda perlu mengikuti langkah-langkah instalasi berikut:

1. Clone repositori ini ke komputer Anda.
2. Instal semua dependensi yang diperlukan dengan menjalankan perintah `go mod tidy`.
3. Konfigurasi koneksi database MySQL di file `.env` atau konfigurasi yang sesuai.
4. Jalankan API dengan perintah `go run app.go`.
5. API akan berjalan di `http://localhost:8000` secara default.

## Penggunaan API

Anda dapat menggunakan API ini untuk berbagai keperluan seperti peminjaman buku, mengirim pesan ke admin, melihat informasi buku, dan lainnya. Pastikan untuk mengamati dokumentasi endpoint di bawah ini.

## Endpoints

### Pengguna (Users):
- GET `/users`: Mendapatkan daftar semua pengguna.
- GET `/users/:id`: Mendapatkan data pengguna berdasarkan ID.
- GET `/users/name/:name`: Mendapatkan data pengguna berdasarkan nama.
- POST `/users`: Membuat pengguna baru.
- PUT `/users/:id`: Mengupdate data pengguna berdasarkan ID.
- DELETE `/users/:id`: Menghapus pengguna berdasarkan ID.

### Admin:
- GET `/admin`: Mendapatkan daftar semua admin.
- GET `/admin/:id`: Mendapatkan data admin berdasarkan ID.
- GET `/admin/name/:name`: Mendapatkan data admin berdasarkan nama.
- POST `/admin`: Membuat admin baru.
- PUT `/admin/:id`: Mengupdate data admin berdasarkan ID.
- DELETE `/admin/:id`: Menghapus admin berdasarkan ID.

### Buku (Books):
- GET `/books`: Mendapatkan daftar semua buku.
- GET `/books/:id`: Mendapatkan data buku berdasarkan ID.
- GET `/books/title/:title`: Mendapatkan data buku berdasarkan judul.
- POST `/books`: Membuat buku baru.
- PUT `/books/:id`: Mengupdate data buku berdasarkan ID.
- DELETE `/books/:id`: Menghapus buku berdasarkan ID.

### Peminjaman (Borrowing):
- GET `/borrow`: Mendapatkan daftar semua peminjaman.
- GET `/borrow/:id`: Mendapatkan data peminjaman berdasarkan ID.
- GET `/borrow/name/:name`: Mendapatkan data peminjaman buku berdasarkan nama pengguna.
- POST `/borrow`: Melakukan peminjaman buku.
- PUT `/borrow/:id`: Mengupdate data peminjaman berdasarkan ID.
- DELETE `/borrow/:id`: Menghapus data peminjaman berdasarkan ID.

### Penulis (Author):
- GET `/authors`: Mendapatkan daftar semua peminjaman.
- GET `/authors/:id`: Mendapatkan data peminjaman berdasarkan ID.
- GET `/users/name/:name`: Mendapatkan data penulis berdasarkan nama.
- POST `/authors`: Melakukan peminjaman buku.
- PUT `/authors/:id`: Mengupdate data peminjaman berdasarkan ID.
- DELETE `/authors/:id`: Menghapus data peminjaman berdasarkan ID.

<!-- 
### Permintaan Buku (Book Requests):
- GET `/chat`: Mendapatkan daftar semua pesan.
- GET `/chat/:id`: Mendapatkan data pesan berdasarkan ID.
- POST `/chat`: Membuat pesan baru.
- PUT `/chat/:id`: Mengupdate data pesan berdasarkan ID.
- DELETE `/chat/:id`: Menghapus data pesan berdasarkan ID. -->

## Dokumentasi

Dokumentasi dapat dilihat pada [tautan ini](https://documenter.getpostman.com/view/23660564/2s9YRFVVZ8#5ad2708b-c887-441f-9c90-e18d05f8c884) untuk detail lebih lanjut


## Desain Database

Aplikasi ini menggunakan desain database yang telah diatur sebelumnya. Berikut adalah skema desain ERD yang digunakan:

[![](https://app.eraser.io/workspace/jIvOglfvfnBAHwnwQba4/preview?elements=BV2s12o7U6oh6doRcDxP2w&type=embed)](https://app.eraser.io/workspace/jIvOglfvfnBAHwnwQba4?elements=BV2s12o7U6oh6doRcDxP2w)

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan kirim pull request atau laporkan masalah (issues) yang Anda temui. Kami sangat menghargai kontribusi Anda!

## Lisensi

Proyek ini dilisensikan di bawah 2023 Peekabook API Team. Silakan merujuk ke file [LICENSE](LICENSE) untuk informasi lebih lanjut.

© 2023 Peekabook API Team
