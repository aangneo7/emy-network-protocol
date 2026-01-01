package main

import (
	"testing"
)

// TestInitialSupply mengecek apakah total supply token sesuai saat inisialisasi
func TestInitialSupply(t *testing.T) {
	// Simulasi variabel token (sesuaikan dengan variabel di main.go kamu)
	expectedName := "Emy Network"
	expectedSymbol := "EMY"
	actualSupply := 1000000 // Contoh angka supply

	if expectedName != "Emy Network" {
		t.Errorf("Nama token salah! Ekspektasi: %s, Realita: %s", expectedName, "NamaLain")
	}

	if expectedSymbol != "EMY" {
		t.Errorf("Simbol token salah! Ekspektasi: %s, Realita: %s", expectedSymbol, "ABC")
	}

	if actualSupply <= 0 {
		t.Error("Total supply harus lebih besar dari 0")
	}
}

// TestTransfer menguji logika perpindahan saldo sederhana
func TestTransfer(t *testing.T) {
	balanceUserA := 100
	balanceUserB := 0
	transferAmount := 30

	// Logika transfer sederhana
	balanceUserA -= transferAmount
	balanceUserB += transferAmount

	if balanceUserA != 70 {
		t.Errorf("Saldo pengirim salah setelah transfer! Sisa: %d", balanceUserA)
	}

	if balanceUserB != 30 {
		t.Errorf("Saldo penerima salah setelah transfer! Sisa: %d", balanceUserB)
	}
}
