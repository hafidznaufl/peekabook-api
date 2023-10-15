# PeekaBook API

Aplikasi "Peekabook" adalah sistem manajemen peminjaman buku yang memungkinkan pengguna untuk meminjam buku, membuat permintaan buku, dan melihat informasi buku yang tersedia. API ini menyediakan layanan yang diperlukan untuk mengakses dan mengelola data dalam sistem ini.

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
2. Instal semua dependensi yang diperlukan dengan menjalankan perintah `go get`.
3. Konfigurasi koneksi database MySQL di file `.env` atau konfigurasi yang sesuai.
4. Jalankan API dengan perintah `go run main.go`.
5. API akan berjalan di `http://localhost:8000` secara default.

## Penggunaan API

Anda dapat menggunakan API ini untuk berbagai keperluan seperti peminjaman buku, mengirim pesan ke admin, melihat informasi buku, dan lainnya. Pastikan untuk mengamati dokumentasi endpoint di bawah ini.

## Endpoints

### Pengguna (Users):
- GET `/users`: Mendapatkan daftar semua pengguna.
- GET `/users/:id`: Mendapatkan data pengguna berdasarkan ID.
- POST `/users`: Membuat pengguna baru.
- PUT `/users/:id`: Mengupdate data pengguna berdasarkan ID.
- DELETE `/users/:id`: Menghapus pengguna berdasarkan ID.

### Admin:
- GET `/admin`: Mendapatkan daftar semua admin.
- GET `/admin/:id`: Mendapatkan data admin berdasarkan ID.
- POST `/admin`: Membuat admin baru.
- PUT `/admin/:id`: Mengupdate data admin berdasarkan ID.
- DELETE `/admin/:id`: Menghapus admin berdasarkan ID.

### Buku (Books):
- GET `/books`: Mendapatkan daftar semua buku.
- GET `/books/:id`: Mendapatkan data buku berdasarkan ID.
- POST `/books`: Membuat buku baru.
- PUT `/books/:id`: Mengupdate data buku berdasarkan ID.
- DELETE `/books/:id`: Menghapus buku berdasarkan ID.

### Peminjaman (Borrowing):
- GET `/borrow`: Mendapatkan daftar semua peminjaman.
- GET `/borrow/:id`: Mendapatkan data peminjaman berdasarkan ID.
- POST `/borrow`: Melakukan peminjaman buku.
- PUT `/borrow/:id`: Mengupdate data peminjaman berdasarkan ID.
- DELETE `/borrow/:id`: Menghapus data peminjaman berdasarkan ID.

### Permintaan Buku (Book Requests):
- GET `/requests`: Mendapatkan daftar semua permintaan buku.
- GET `/requests/:id`: Mendapatkan data permintaan buku berdasarkan ID.
- POST `/requests`: Membuat permintaan buku baru.
- PUT `/requests/:id`: Mengupdate data permintaan buku berdasarkan ID.
- DELETE `/requests/:id`: Menghapus data permintaan buku berdasarkan ID.

## Desain Database

Aplikasi ini menggunakan desain database yang telah diatur sebelumnya. Berikut adalah skema desain ERD yang digunakan:

[![](https://app.eraser.io/workspace/jIvOglfvfnBAHwnwQba4/preview?elements=5baF7S8UzmCG6QYUc-x8bw&type=embed)](https://app.eraser.io/workspace/jIvOglfvfnBAHwnwQba4?elements=5baF7S8UzmCG6QYUc-x8bw)

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan kirim pull request atau laporkan masalah (issues) yang Anda temui. Kami sangat menghargai kontribusi Anda!

## Lisensi

Proyek ini dilisensikan di bawah 2023 Peekabook API Team. Silakan merujuk ke file [LICENSE](LICENSE) untuk informasi lebih lanjut.

© 2023 Peekabook API Team
