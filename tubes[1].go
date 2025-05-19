package main

import (
	"fmt"
)

type pakaian struct {
	nama     string
	kategori string
}

func main() {
	var a string
	for a != "Exit" {
		fmt.Prinln("Pilih Menu")
		fmt.Println("1. Menu Pakaian")
		fmt.Println("2. Kombinasi Outfit")
		fmt.Println("3. Cari Kombinasi")
		fmt.Println("4. Urutan Outfit")
		fmt.Println("5. Outfit Berdasarkan Cuaca")
		fmt.Println("Exit")
		fmt.Scan(&a)

		switch a {
		case "1":
			pakaian()
		case "2":
			kombinasi()
		case "3":
			search()
		case "4":
			urut()
		case "5":
			cuaca()
		}
	}
}

func pakaian() {
	var x string
	for x != "Exit" {
		fmt.Println("1. Menambah")
		fmt.Println("2. Mengubah")
		fmt.Println("3. Menghapus")
		fmt.Println("4. Exit")
		fmt.Scan(&x)

		switch x {
		case "1":
			menambah()
		case "2":
			mengubah()
		case "3":
			menghapus()
		}
	}
}

func kategori() {
	var kategori string

	fmt.Println("Masukkan kategori (cuaca_hujan, cuaca_panas, meeting_formal, dll):")
	fmt.Scanln(&kategori)

	switch kategori {
	case "cuaca_hujan":
		fmt.Println("Rekomendasi: Jaket anti air, sepatu boots, dan payung.")
	case "cuaca_panas":
		fmt.Println("Rekomendasi: Kaos tipis, celana pendek, dan topi.")
	case "cuaca_dingin":
		fmt.Println("Rekomendasi: Sweater tebal, celana panjang, dan syal.")
	case "meeting_formal":
		fmt.Println("Rekomendasi: Kemeja, celana bahan, jas, dan sepatu pantofel.")
	case "acara_santai":
		fmt.Println("Rekomendasi: Kaos santai, jeans, dan sneakers.")
	default:
		fmt.Println("Kategori tidak dikenali.")
	}
}
