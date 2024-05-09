# Aplikasi Konsultasi Kesehatan
Tugas Besar Mata Kuliah Algoritma Pemprograman.
## Deskripsi Aplikasi

Not yet.

## Dibuat Oleh

- Faisal Ihsan Santoso
- Arie Farchan Fyrzatullah

## Fitur

Not yet.

## Patch 1.0

- Implementasi menu utama dengan Command Line Interface
  
  ![Base Menu](https://github.com/CyrusSE/AlproProject/assets/80195151/cfb58bf0-8997-4517-ade6-003eefd98c92)
- Implementasi switch case untuk pilihan
- Add fungsi clear console, title console, dan color console
- Add ASCII Good Bye!

  ![GoodBye](https://github.com/CyrusSE/AlproProject/assets/80195151/13e88081-d40a-46b0-9be6-69a2e66172f8)

## Patch 1.1

- **Ganti judul**: Aplikasi Konsultasi Kesehatan
- **Perubahan Terminologi**: Semua referensi dari "mahasiswa" diubah menjadi "Pasien" dan "Admin" menjadi "Dokter".
- **Tamu**: Implementasi fitur untuk pengguna tanpa akun untuk mengakses informasi terbatas.
- **Recovery Akun Pasien**: -
- **Header menu**: Penambahan header pada menu untuk informasi yang lebih jelas.
- **All in one menu**: Semua fitur dalam satu menu utama agar akses yang lebih mudah.
- **Login page Pasien**:
  - Login berhasil: -
  - Login gagal: Pesan error (gunakan PIN untuk recovery) dan opsi untuk mencoba lagi.
  - Akun tidak ditemukan: Pesan error dan opsi untuk registrasi.

## Patch 1.2
- Removed external package.
- Add created total account.
- Fix some bugs.
- Function untuk mencari dan mengecek akun exist atau tidak.
- Function untuk mencari index dari akun yang dicari.
- Better login account algorithm.
- **Add creating account**:
  - Buat akun gagal jika username sudah digunakan.
  - Save nama, nama, username, password, PIN ke array struct.
- **Add recovery account**:
  - Recovery gagal jika input PIN berbeda dengan data pin dari username yang diminta.
  - Recovery gagal jika akun tidak ditemukan.
  - Recovery berhasil dan ganti password baru jika input PIN benar.

## Patch 1.3
- Fix some bugs for invalid option.
- **Login account finally works**:
  - Jika login berhasil maka akan dilempar ke menu utama pasien.
- **Add struct pertanyaan**:
  - tag (Judul pertanyaan)
  - pertanyaan
  - tanggapan
  - date
  - author
- **Logged pasien bisa posting pertanyaan**:
  - Judul pertanyaan adalah Tag.
  - Pertayaan di parse menggunakan library bufio agar pertanyaa bisa lebih dari 1 kata.
  - Total pertanyaan di tambah 1.
  - Save waktu sekarang ke dalam struct pertanyaan dengan library time.
- **Lihat Forum Pertanyaan (half working)**:
  - format : (tag)  (date)  (author)



