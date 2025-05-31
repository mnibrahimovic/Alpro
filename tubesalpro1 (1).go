package main

import "fmt"

const pakaian int = 100

var daftarPakaian tabPakaian
var jumlahPakaian int
var daftarOutfit tabOutfit
var jumlahOutfit int

type data struct {
	jenis    string
	kategori string
	warna    string
	cuaca    string
}

type outfit struct {
	atasan  int
	bawahan int
	alas    int
}

type tabPakaian [pakaian]data
type tabOutfit [pakaian]outfit

func main() {
	var pilihan int
	for pilihan != 10 {
		fmt.Println("==============================")
		fmt.Println("         OOTD PLANNER         ")
		fmt.Println("==============================")
		fmt.Println("1. Menambahkan Jenis Pakaian")
		fmt.Println("2. Tampilkan Daftar Pakaian")
		fmt.Println("3. Buat Kombinasi Outfit")
		fmt.Println("4. Tampilkan Kombinasi Outfit")
		fmt.Println("5. Mengubah Detail Pakaian")
		fmt.Println("6. Cari Pakaian (Sequential Search)")
		fmt.Println("7. Urutkan Pakaian (Selection Sort)")
		fmt.Println("8. Cari Pakaian Berdasarkan Cuaca")
		fmt.Println("9. Hapus Pakaian")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			tambahPakaian()
		case 2:
			lihatPakaian()
		case 3:
			kombinasiOutfit()
		case 4:
			lihatKombinasi()
		case 5:
			ubahPakaian()
		case 6: 
			sequentialSearch()
		case 7: 
			selectionSort()
		case 8: 
			cariPakaianByCuaca()
		case 9:
			hapusPakaian()
		case 10:
			fmt.Println("Terima kasih telah menggunakan OOTD Planner!")
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
		fmt.Println()
	}
}

func tambahPakaian() {
	if jumlahPakaian >= pakaian {
		fmt.Println("Kapasitas pakaian penuh!")
		return
	}
	fmt.Print("Jenis Pakaian (satu kata): ")
	fmt.Scan(&daftarPakaian[jumlahPakaian].jenis)
	fmt.Print("Kategori Pakaian (Formal/Casual/Santai/Lainnya, satu kata): ")
	fmt.Scan(&daftarPakaian[jumlahPakaian].kategori)
	fmt.Print("Warna Pakaian (satu kata): ")
	fmt.Scan(&daftarPakaian[jumlahPakaian].warna)
	fmt.Print("Cuaca Cocok (Panas/Dingin/Hujan/Lainnya, satu kata): ")
	fmt.Scan(&daftarPakaian[jumlahPakaian].cuaca)

	jumlahPakaian++
	fmt.Println("Daftar Pakaian Berhasil Ditambahkan")
}

func lihatPakaian() {
	if jumlahPakaian == 0 {
		fmt.Println("Belum ada pakaian yang terdaftar.")
		return
	}
	fmt.Println("===========================================================")
	fmt.Println("                         DAFTAR PAKAIAN                    ")
	fmt.Println("===========================================================")
	fmt.Printf("%-4s %-15s %-15s %-10s %-10s\n", "No", "Jenis Pakaian", "Kategori", "Warna", "Cuaca")
	fmt.Println("----------------------------------------------------------")
	for i := 0; i < jumlahPakaian; i++ {
		fmt.Printf("%-4d %-15s %-15s %-10s %-10s\n",
			i+1,
			daftarPakaian[i].jenis,
			daftarPakaian[i].kategori,
			daftarPakaian[i].warna,
			daftarPakaian[i].cuaca,
		)
	}
	fmt.Println("----------------------------------------------------------")
	fmt.Println()
}

func kombinasiOutfit() {
	if jumlahPakaian == 0 {
		fmt.Println("Tidak ada pakaian yang terdaftar untuk dikombinasikan.")
		return
	}
	if jumlahOutfit >= pakaian {
		fmt.Println("Kapasitas outfit penuh!")
		return
	}

	var atasanNum, bawahanNum, alasNum int
	fmt.Println("Buat Kombinasi Outfit: ")
	lihatPakaian()
	fmt.Print("Masukkan Nomor Atasan Yang Ingin Dikombinasi (1-", jumlahPakaian, "): ")
	fmt.Scan(&atasanNum)
	fmt.Print("Masukkan Nomor Bawahan Yang Ingin Dikombinasi (1-", jumlahPakaian, "): ")
	fmt.Scan(&bawahanNum)
	fmt.Print("Masukkan Nomor Alas Kaki Yang Ingin Dikombinasi (1-", jumlahPakaian, "): ")
	fmt.Scan(&alasNum)

	if atasanNum < 1 || atasanNum > jumlahPakaian ||
		bawahanNum < 1 || bawahanNum > jumlahPakaian ||
		alasNum < 1 || alasNum > jumlahPakaian {
		fmt.Println("Nomor pakaian tidak valid. Kombinasi dibatalkan.")
		return
	}

	daftarOutfit[jumlahOutfit].atasan = atasanNum - 1
	daftarOutfit[jumlahOutfit].bawahan = bawahanNum - 1
	daftarOutfit[jumlahOutfit].alas = alasNum - 1

	jumlahOutfit++
	fmt.Println("Kombinasi Outfit Berhasil Ditambahkan.")
	fmt.Println()
}

func lihatKombinasi() {
	if jumlahOutfit == 0 {
		fmt.Println("Belum ada kombinasi outfit yang terdaftar.")
		return
	}

	fmt.Println("===========================================================")
	fmt.Println("                   DAFTAR KOMBINASI OUTFIT                 ")
	fmt.Println("===========================================================")
	fmt.Printf("%-4s %-20s %-20s %-20s %-15s %-15s %-15s\n",
		"No", "Atasan", "Bawahan", "Alas Kaki", "Warna Atasan", "Warna Bawahan", "Warna Alas Kaki")
	fmt.Println("---------------------------------------------------------------------------------------------------")

	for i := 0; i < jumlahOutfit; i++ {
		if daftarOutfit[i].atasan >= 0 && daftarOutfit[i].atasan < jumlahPakaian &&
			daftarOutfit[i].bawahan >= 0 && daftarOutfit[i].bawahan < jumlahPakaian &&
			daftarOutfit[i].alas >= 0 && daftarOutfit[i].alas < jumlahPakaian {

			atasanPakaian := daftarPakaian[daftarOutfit[i].atasan]
			bawahanPakaian := daftarPakaian[daftarOutfit[i].bawahan]
			alasPakaian := daftarPakaian[daftarOutfit[i].alas]

			fmt.Printf("%-4d %-20s %-20s %-20s %-15s %-15s %-15s\n",
				i+1,
				atasanPakaian.jenis,
				bawahanPakaian.jenis,
				alasPakaian.jenis,
				atasanPakaian.warna,
				bawahanPakaian.warna,
				alasPakaian.warna,
			)
		} else {
			fmt.Printf("%-4d [Kombinasi Tidak Valid - Pakaian Mungkin Dihapus]\n", i+1)
		}
	}
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Println()
}

func ubahPakaian() {
	if jumlahPakaian == 0 {
		fmt.Println("Belum ada pakaian yang terdaftar untuk diubah.")
		return
	}
	lihatPakaian()

	var nomorPakaian int
	fmt.Print("Masukkan Nomor Pakaian yang ingin diubah: ")
	fmt.Scan(&nomorPakaian)

	idx := nomorPakaian - 1

	if idx < 0 || idx >= jumlahPakaian {
		fmt.Println("Nomor Pakaian Tidak Valid. Pembaharuan dibatalkan.")
		return
	}

	pakaianUpdate := &daftarPakaian[idx]

	fmt.Println("\n--- Masukkan Detail Pakaian Baru ---")
	fmt.Printf("Jenis Pakaian (Saat ini: %s, satu kata): ", pakaianUpdate.jenis)
	fmt.Scan(&pakaianUpdate.jenis)
	fmt.Printf("Kategori Pakaian (Saat ini: %s, satu kata): ", pakaianUpdate.kategori)
	fmt.Scan(&pakaianUpdate.kategori)
	fmt.Printf("Warna Pakaian (Saat ini: %s, satu kata): ", pakaianUpdate.warna)
	fmt.Scan(&pakaianUpdate.warna)
	fmt.Printf("Cuaca Cocok (Saat ini: %s, satu kata): ", pakaianUpdate.cuaca)
	fmt.Scan(&pakaianUpdate.cuaca)

	fmt.Println("Pakaian Berhasil Diperbarui!")
}


func sequentialSearch() {
	if jumlahPakaian == 0 {
		fmt.Println("Belum ada pakaian untuk dicari.")
		return
	}

	var pilihanKriteria int
	fmt.Println("\n--- Cari Pakaian ---")
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Jenis Pakaian")
	fmt.Println("2. Kategori Pakaian")
	fmt.Println("3. Warna Pakaian")
	fmt.Println("4. Cuaca Cocok")
	fmt.Print("Pilih kriteria (1/2/3/4): ")
	fmt.Scan(&pilihanKriteria)

	var kataKunci string
	fmt.Print("Masukkan kata kunci pencarian (satu kata): ")
	fmt.Scan(&kataKunci)

	fmt.Println("\n-----------------------------------------------------------")
	fmt.Println("Hasil Pencarian untuk Kata Kunci:", kataKunci)
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("%-4s %-15s %-15s %-10s %-10s\n", "No", "Jenis Pakaian", "Kategori", "Warna", "Cuaca")
	fmt.Println("-----------------------------------------------------------")

	found := false
	foundItemsCount := 0

	for i := 0; i < jumlahPakaian; i++ {
		isFound := false
		switch pilihanKriteria {
		case 1:
			if daftarPakaian[i].jenis == kataKunci {
				isFound = true
			}
		case 2:
			if daftarPakaian[i].kategori == kataKunci {
				isFound = true
			}
		case 3:
			if daftarPakaian[i].warna == kataKunci {
				isFound = true
			}
		case 4: 
			if daftarPakaian[i].cuaca == kataKunci {
				isFound = true
			}
		default:
			fmt.Println("Kriteria pencarian tidak valid.")
			return
		}
		if isFound {
			fmt.Printf("%-4d %-15s %-15s %-10s %-10s\n",
				i+1,
				daftarPakaian[i].jenis,
				daftarPakaian[i].kategori,
				daftarPakaian[i].warna,
				daftarPakaian[i].cuaca,
			)
			found = true
			foundItemsCount++
		}
	}
	if !found {
		fmt.Println("Tidak ada pakaian yang cocok dengan kata kunci tersebut.")
	} else {
		fmt.Printf("Ditemukan %d pakaian yang cocok.\n", foundItemsCount)
	}
	fmt.Println("-----------------------------------------------------------")
}

func selectionSort() {
	if jumlahPakaian == 0 { 
		fmt.Println("Tidak ada pakaian untuk diurutkan.")
		return
	}
	for i := 0; i < jumlahPakaian-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlahPakaian; j++ {
			if daftarPakaian[j].jenis < daftarPakaian[minIndex].jenis {
				minIndex = j
			}
		}
		if minIndex != i {
			temp := daftarPakaian[i]
			daftarPakaian[i] = daftarPakaian[minIndex]
			daftarPakaian[minIndex] = temp
		}
	}
	fmt.Println("Daftar pakaian telah diurutkan berdasarkan jenis dengan Selection Sort.")
	lihatPakaian() 
}

func binarySearch(kataKunciJenis string) int { 
	if jumlahPakaian == 0 {
		return -1
	}
	low := 0
	high := jumlahPakaian - 1
	for low <= high {
		mid := (low + high) / 2
		if daftarPakaian[mid].jenis == kataKunciJenis {
			return mid
		} else if daftarPakaian[mid].jenis < kataKunciJenis {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func cariPakaianByCuaca() {
	if jumlahPakaian == 0 {
		fmt.Println("Belum ada pakaian untuk dicari.")
		return
	}

	var cuacaCari string
	fmt.Println("\n--- Cari Pakaian Berdasarkan Cuaca ---")
	fmt.Print("Masukkan cuaca yang dicari (Panas/Dingin/Hujan/Lainnya, satu kata): ")
	fmt.Scan(&cuacaCari)

	fmt.Println("\n-----------------------------------------------------------")
	fmt.Println("Hasil Pencarian Pakaian untuk Cuaca:", cuacaCari)
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("%-4s %-15s %-15s %-10s %-10s\n", "No", "Jenis Pakaian", "Kategori", "Warna", "Cuaca")
	fmt.Println("-----------------------------------------------------------")

	found := false
	foundItemsCount := 0

	for i := 0; i < jumlahPakaian; i++ {
		if daftarPakaian[i].cuaca == cuacaCari {
			fmt.Printf("%-4d %-15s %-15s %-10s %-10s\n",
				i+1,
				daftarPakaian[i].jenis,
				daftarPakaian[i].kategori,
				daftarPakaian[i].warna,
				daftarPakaian[i].cuaca,
			)
			found = true
			foundItemsCount++
		}
	}
	if !found {
		fmt.Println("Tidak ada pakaian yang cocok untuk cuaca '", cuacaCari, "'.")
	} else {
		fmt.Printf("Ditemukan %d pakaian yang cocok.\n", foundItemsCount)
	}
	fmt.Println("-----------------------------------------------------------")
}

func hapusPakaian() {
	if jumlahPakaian == 0 {
		fmt.Println("Belum ada pakaian yang dihapus.")
		return
	}
	lihatPakaian()
	var nomorPakaian int
	fmt.Print("Masukkan nomor pakaian yang ingin dihapus: ")
	fmt.Scan(&nomorPakaian)

	pakaianDeleteIdx := nomorPakaian - 1 
	if pakaianDeleteIdx < 0 || pakaianDeleteIdx >= jumlahPakaian {
		fmt.Println("Nomor tidak valid. Penghapusan dibatalkan.")
		return
	}
	fmt.Println("Mengecek apakah pakaian ini digunakan dalam kombinasi outfit...")
	foundInOutfit := false
	for i := 0; i < jumlahOutfit; i++ {
		if daftarOutfit[i].atasan == pakaianDeleteIdx || daftarOutfit[i].bawahan == pakaianDeleteIdx || daftarOutfit[i].alas == pakaianDeleteIdx {
			foundInOutfit = true
		}
	}

	if foundInOutfit {
		fmt.Println("Peringatan: Pakaian ini digunakan dalam satu atau lebih kombinasi outfit.")
		fmt.Println("Menghapus pakaian ini akan membuat kombinasi outfit tersebut tidak valid.")
		var konfirmasi string
		fmt.Print("Lanjutkan penghapusan? (ya/tidak): ")
		fmt.Scan(&konfirmasi)
		if konfirmasi != "ya" {
			fmt.Println("Penghapusan dibatalkan.")
			return
		}
	}
	for i := pakaianDeleteIdx; i < jumlahPakaian-1; i++ {
		daftarPakaian[i] = daftarPakaian[i+1]
	}
	daftarPakaian[jumlahPakaian-1] = data{} 
	jumlahPakaian--
	fmt.Println("Pakaian berhasil dihapus.")

	for i := 0; i < jumlahOutfit; i++ {
		if daftarOutfit[i].atasan > pakaianDeleteIdx {
			daftarOutfit[i].atasan--
		}
		if daftarOutfit[i].bawahan > pakaianDeleteIdx {
			daftarOutfit[i].bawahan--
		}
		if daftarOutfit[i].alas > pakaianDeleteIdx {
			daftarOutfit[i].alas--
		}
	}
	fmt.Println("Outfit sudah selesai diperbarui.")
	fmt.Println()
}
