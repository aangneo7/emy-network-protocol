# Security Policy - Canopy Network

Keamanan adalah fondasi dari ekosistem **Canopy Network**. Sebagai protokol Layer 1 yang menyediakan layanan *Shared Security* dan *Bridge-less Interoperability*, kami berkomitmen untuk melindungi integritas jaringan dan aset pengguna dari segala bentuk ancaman.

Kami sangat menghargai kontribusi dari peneliti keamanan dan pengembang dalam mengidentifikasi kerentanan secara bertanggung jawab.

## Versi yang Didukung (Supported Versions)

Kami secara aktif melakukan patch keamanan pada versi-versi berikut:

| Versi | Status | Catatan |
| :--- | :--- | :--- |
| **v1.0.x (Mainnet)** | ✅ Supported | Target Rilis 2026 |
| **v0.5.x (Betanet)** | ✅ Supported | Fokus Utama Saat Ini |
| < v0.5.x | ❌ Unsupported | Mohon segera upgrade node Anda |

## Melaporkan Celah Keamanan (Reporting a Vulnerability)

**PERINGATAN: Jangan pernah melaporkan celah keamanan melalui GitHub Issues publik.**

Untuk menjaga keamanan seluruh validator dan pengguna, kami meminta Anda untuk melaporkan temuan secara privat melalui jalur berikut:

* **Email:** `security@canopy.network`
* **Subject:** `[VULNERABILITY REPORT] - <Deskripsi Singkat>`
* **PGP Key:** (Opsional) Silakan gunakan Public Key kami yang tersedia di [Link/ID Key] untuk komunikasi terenkripsi.

### Informasi yang Harus Disertakan:
Agar tim kami dapat merespons dengan cepat, mohon sertakan:
1. **Jenis Kerentanan:** (Contoh: Double Spending, Remote Code Execution, Denial of Service).
2. **Lokasi:** File kode, modul (x/), atau endpoint API yang terdampak.
3. **Proof of Concept (PoC):** Langkah-langkah detail atau skrip untuk mereproduksi celah tersebut.
4. **Dampak:** Analisis mengenai risiko terhadap dana pengguna atau stabilitas jaringan.

## Kebijakan "Responsible Disclosure"

Kami berkomitmen untuk bekerja sama secara transparan. Kami meminta Anda untuk:
* Memberikan waktu yang cukup bagi kami untuk memperbaiki celah sebelum dipublikasikan secara umum.
* Tidak mengeksploitasi kerentanan untuk merugikan pengguna atau mencuri dana selama proses investigasi.
* Menghindari serangan yang bersifat merusak infrastruktur fisik atau kampanye phishing.

## Program Bug Bounty (Apresiasi)

Canopy Network memberikan imbalan (reward) bagi laporan yang valid dan berdampak signifikan. Besaran hadiah ditentukan berdasarkan tingkat keparahan (**Severity**) menggunakan standar CVSS:

* **Critical (Kritikal):** Kehilangan dana secara langsung, penghentian rantai (Chain halt), atau eksekusi kode jarak jauh.
* **High (Tinggi):** Manipulasi tata kelola (Governance), sensor transaksi, atau pencurian data sensitif.
* **Medium/Low (Sedang/Rendah):** Bug fungsionalitas yang tidak mengancam dana tetapi merusak pengalaman pengguna.

*Reward dapat diberikan dalam bentuk token **CNPY** atau Stablecoin (USDT/USDC).*

## Komitmen Kami

Setelah menerima laporan, tim keamanan Canopy Network akan:
1. Mengirim konfirmasi penerimaan laporan dalam waktu **24 jam**.
2. Memberikan penilaian awal mengenai dampak dalam **3 hari kerja**.
3. Memberikan pembaruan berkala selama proses perbaikan berlangsung.
4. Memberikan atribusi (penghargaan nama) dalam catatan rilis kami (kecuali Anda ingin tetap anonim).

---
*Terima kasih telah membantu menjaga ekosistem Canopy Network tetap aman.*
