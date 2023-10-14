package main

import "fmt"

// Fungsi untuk menghitung tip 15% dan 20%
func hitungTip(tagihan float64) float64 {
	if tagihan >= 50 && tagihan <= 300 {
		return tagihan * 0.15 // Tip 15% jika tagihan antara 50 dan 300 (15% -> 15/100 = 0,15)
	} else {
		return tagihan * 0.2 // Tip 20% untuk tagihan lainnya (20% -> 20/100 = 0,2)
	}
}

func main() {
	// Data uji 1 dari case PR
	tagihan1 := 275.0
	tagihan2 := 40.0
	tagihan3 := 430.0

	// Hitung tip untuk masing-masing data uji
	tip1 := hitungTip(tagihan1)
	tip2 := hitungTip(tagihan2)
	tip3 := hitungTip(tagihan3)

	// Hitung total nilai (tagihan + tip) untuk masing-masing data uji
	total1 := tagihan1 + tip1
	total2 := tagihan2 + tip2
	total3 := tagihan3 + tip3

	// Menampilkan hasil dari fungsi ke konsol atau terminal
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total tagihannya %.2f\n", tagihan1, tip1, total1)
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total tagihannya %.2f\n", tagihan2, tip2, total2)
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total tagihannya %.2f\n", tagihan3, tip3, total3)
}
