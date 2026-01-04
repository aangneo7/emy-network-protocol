# Contributing to Canopy Network 🌳

Terima kasih telah tertarik untuk berkontribusi pada Canopy Network! Sebagai proyek Layer 1, kami sangat menghargai kontribusi dari komunitas untuk membangun masa depan *bridge-less interoperability*.

Dengan berkontribusi, Anda setuju bahwa kode Anda akan dilisensikan di bawah [MIT License/Apache 2.0] yang digunakan oleh proyek ini.

## 🚀 Cara Memulai

1. **Cari Issue:** Lihat [Issues tab](link-ke-issues) untuk mencari bug atau fitur yang sedang dikerjakan. Jika ingin mengusulkan sesuatu yang baru, buka issue baru terlebih dahulu.
2. **Fork & Clone:** Fork repositori ini ke akun Anda dan clone ke mesin lokal.
3. **Instalasi:** Ikuti panduan di `README.md` untuk menyiapkan lingkungan pengembangan (Go/Rust/Node.js).

## 🛠️ Aturan Alur Kerja (Workflow)

Kami menggunakan model **Feature Branching**:

1. Buat branch baru dari `develop` (jangan langsung ke `main`):
   `git checkout -b feature/nama-fitur-anda` atau `fix/deskripsi-bug`.
2. Lakukan perubahan kode.
3. Pastikan kode Anda lulus semua pengujian:
   `make test` atau `go test ./...`
4. Commit perubahan Anda dengan pesan yang jelas (lihat standar di bawah).
5. Push ke fork Anda dan buka **Pull Request (PR)** ke branch `develop` kami.

## 📝 Standar Pesan Commit

Kami mengikuti konvensi [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` untuk fitur baru (misal: `feat: add bridge-less module for appchains`)
- `fix:` untuk perbaikan bug (misal: `fix: resolve gas calculation error`)
- `docs:` untuk perubahan dokumentasi.
- `test:` untuk menambah atau memperbaiki testing.
- `refactor:` untuk merapikan kode tanpa mengubah fungsi.

## 💻 Standar Penulisan Kode (Coding Standards)

Agar kualitas kode terjaga, mohon perhatikan hal berikut:

- **Linting:** Jalankan linter sebelum commit (misal: `golangci-lint run`).
- **Dokumentasi:** Setiap fungsi publik atau modul baru wajib memiliki komentar penjelasan.
- **Testing:** Setiap fitur baru harus disertai dengan unit test di folder `/tests`.
- **Keamanan:** Jangan pernah melakukan *hardcode* API key atau private key. Gunakan `.env.example`.

## 🤝 Kode Etik

Harap bersikap profesional dan inklusif dalam setiap diskusi di Issue atau Pull Request. Kami berkomitmen untuk menciptakan lingkungan yang ramah bagi semua pengembang.

## ❓ Butuh Bantuan?

Jika Anda memiliki pertanyaan teknis, jangan ragu untuk:
- Bergabung di [Discord Canopy Network](link-discord).
- Menanyakan langsung di diskusi Pull Request yang Anda buat.

---
*Mari kita bangun ekosistem blockchain yang lebih terhubung bersama-sama!*
