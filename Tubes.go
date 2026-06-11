package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const NMAX int = 100

type poLu struct {
	kota     string
	provinsi string
	aqi      float64
	sumber   string
	tanggal  string
}

type tabpol [NMAX]poLu

// ============ FUNGSI BANTUAN ============

// Fungsi untuk membaca input teks dengan spasi
func inputTeks() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		teks, _ := reader.ReadString('\n')
		teks = strings.TrimSpace(teks)
		if teks != "" {
			return teks
		}
	}
}

func kategoriPolusi(aqi float64) string {
	if aqi <= 50 {
		return "Baik"
	} else if aqi <= 100 {
		return "Sedang"
	} else if aqi <= 150 {
		return "Tidak sehat untuk kelompok sensitif"
	} else if aqi <= 200 {
		return "Tidak sehat"
	} else if aqi <= 300 {
		return "Sangat tidak sehat"
	} else if aqi > 300{
		return "Berbahaya"
	} else {
		return "Data Tidak Valid"
	}
}

func round2(f float64) float64 {
	return math.Round(f*100) / 100
}

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

// ============ FUNGSI VALIDASI TANGGAL ============
func inputTanggal() string {
	for {
		fmt.Print("Tanggal (YYYY/MM/DD): ")
		raw := inputTeks()

		parts := strings.Split(raw, "/")
		if len(parts) != 3 {
			fmt.Println("   [!] Format salah. Gunakan YYYY/MM/DD (contoh: 2024/6/1)")
			continue
		}

		tahunStr := strings.TrimSpace(parts[0])
		bulanStr := strings.TrimSpace(parts[1])
		hariStr := strings.TrimSpace(parts[2])

		tahun, errT := strconv.Atoi(tahunStr)
		bulan, errB := strconv.Atoi(bulanStr)
		hari, errH := strconv.Atoi(hariStr)

		if errT != nil || errB != nil || errH != nil {
			fmt.Println("   [!] Tahun, bulan, dan hari harus berupa angka.")
			continue
		}

		if tahun < 1000 || tahun > 9999 {
			fmt.Println("   [!] Tahun harus 4 digit (contoh: 2024).")
			continue
		}

		if bulan < 1 || bulan > 12 {
			fmt.Println("   [!] Bulan harus antara 01 sampai 12.")
			continue
		}

		tgl := time.Date(tahun, time.Month(bulan), hari, 0, 0, 0, 0, time.UTC)
		if tgl.Year() != tahun || int(tgl.Month()) != bulan || tgl.Day() != hari {
			maxHari := time.Date(tahun, time.Month(bulan+1), 0, 0, 0, 0, 0, time.UTC).Day()
			fmt.Printf("   [!] Tanggal tidak valid. Bulan %02d tahun %d maksimal %d hari.\n",
				bulan, tahun, maxHari)
			continue
		}

		hasil := fmt.Sprintf("%04d/%02d/%02d", tahun, bulan, hari)
		fmt.Printf("   >> Tanggal tersimpan: %s\n", hasil)
		return hasil
	}
}

// ============ FUNGSI PILIH SUMBER POLUSI ============
func scanInt() int {
	var hasil int
	fmt.Sscan(inputTeks(), &hasil)
	return hasil
}

func pilihSumber() string {
	for {
		fmt.Println("\n--- Pilih Faktor Sumber Polusi ---")
		fmt.Println("1. Aktivitas Manusia")
		fmt.Println("2. Faktor Alamiah")
		fmt.Print("Pilihan faktor (1-2): ")
		pilihanFaktor := scanInt()

		if pilihanFaktor == 1 {
			for {
				fmt.Println("\n  -- Aktivitas Manusia --")
				fmt.Println("  1. Sektor Transportasi")
				fmt.Println("  2. Industri Dan Manufaktur")
				fmt.Println("  3. Pembangkit Listrik")
				fmt.Println("  4. Pembakaran Sampah Terbuka")
				fmt.Println("  5. Sektor Konstruksi")
				fmt.Println("  6. Aktivitas Rumah Tangga Dan Pertanian")
				fmt.Println("  7. Lainnya (masukkan sendiri)")
				fmt.Print("  Pilihan opsi (1-7): ")
				pilihanOpsi := scanInt()

				switch pilihanOpsi {
				case 1:
					return "Aktivitas Manusia - Sektor Transportasi"
				case 2:
					return "Aktivitas Manusia - Industri Dan Manufaktur"
				case 3:
					return "Aktivitas Manusia - Pembangkit Listrik"
				case 4:
					return "Aktivitas Manusia - Pembakaran Sampah Terbuka"
				case 5:
					return "Aktivitas Manusia - Sektor Konstruksi"
				case 6:
					return "Aktivitas Manusia - Aktivitas Rumah Tangga Dan Pertanian"
				case 7:
					fmt.Print("  Masukkan faktor lain: ")
					lainnya := titleCase(inputTeks())
					return "Aktivitas Manusia - " + lainnya
				default:
					fmt.Println("  Pilihan tidak valid, coba lagi!")
				}
			}

		} else if pilihanFaktor == 2 {
			for {
				fmt.Println("\n  -- Faktor Alamiah --")
				fmt.Println("  1. Aktivitas Vulkanik")
				fmt.Println("  2. Kebakaran Hutan Alami")
				fmt.Println("  3. Kondisi Cuaca")
				fmt.Println("  4. Lainnya (masukkan sendiri)")
				fmt.Print("  Pilihan opsi (1-4): ")
				pilihanOpsi := scanInt()

				switch pilihanOpsi {
				case 1:
					return "Faktor Alamiah - Aktivitas Vulkanik"
				case 2:
					return "Faktor Alamiah - Kebakaran Hutan Alami"
				case 3:
					return "Faktor Alamiah - Kondisi Cuaca"
				case 4:
					fmt.Print("  Masukkan faktor lain: ")
					lainnya := titleCase(inputTeks())
					return "Faktor Alamiah - " + lainnya
				default:
					fmt.Println("  Pilihan tidak valid, coba lagi!")
				}
			}

		} else {
			fmt.Println("Pilihan faktor tidak valid, coba lagi!")
		}
	}
}

// ============ TAMPILAN MENU ============
func tampilMenuUtama() {
	fmt.Println("\n=========================================")
	fmt.Println("   SISTEM MONITORING POLUSI UDARA")
	fmt.Println("=========================================")
	fmt.Println("1. Kelola Data Polusi")
	fmt.Println("2. Pencarian Data")
	fmt.Println("3. Pengurutan Data")
	fmt.Println("4. Laporan Wilayah")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------------------------")
	fmt.Print("Pilihan (1-5): ")
}

// ============ MENU 1: KELOLA DATA ============
func menuKelolaData(data *tabpol, nData *int) {
	var pilihan int
	for {
		fmt.Println("\n--- Menu Kelola Data ---")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Ubah Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Lihat Semua Data")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		pilihan = scanInt()

		switch pilihan {
		case 1:
			tambahData(data, nData)
		case 2:
			ubahData(data, *nData)
		case 3:
			hapusData(data, nData)
		case 4:
			lihatData(*data, *nData)
		case 0:
			fmt.Println("Kembali ke Menu Utama...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// ============ MENU 2: PENCARIAN ============
func menuPencarian(data *tabpol, nData int) {
	var pilihan int
	for {
		fmt.Println("\n--- Menu Pencarian ---")
		fmt.Println("1. Sequential Search (berdasarkan Kota)")
		fmt.Println("2. Binary Search (berdasarkan Kota)")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		pilihan = scanInt()

		switch pilihan {
		case 1:
			sequentialSearch(*data, nData)
		case 2:
			// EDIT: Data diurutkan otomatis dulu berdasarkan kota (A-Z) sebelum binary search
			sortBerdasarkanKota(data, nData)
			fmt.Println(">> [Info] Data telah diurutkan berdasarkan abjad Kota untuk proses Binary Search.")
			binarySearch(*data, nData)
		case 0:
			fmt.Println("Kembali ke Menu Utama...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// ============ MENU 3: PENGURUTAN ============
func menuPengurutan(data *tabpol, nData int) {
	var pilihan int
	for {
		fmt.Println("\n--- Menu Pengurutan ---")
		fmt.Println("1. Selection Sort (berdasarkan AQI Tertinggi)")
		fmt.Println("2. Insertion Sort (berdasarkan Tanggal Terbaru)")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		pilihan = scanInt()

		switch pilihan {
		case 1:
			selectionSort(data, nData)
			fmt.Println("Data berhasil diurutkan berdasarkan AQI (Tertinggi -> Terendah)!")
			lihatData(*data, nData)
		case 2:
			insertionSort(data, nData)
			fmt.Println("Data berhasil diurutkan berdasarkan Tanggal (Terbaru -> Terlama)!")
			lihatData(*data, nData)
		case 0:
			fmt.Println("Kembali ke Menu Utama...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// ============ MENU 4: LAPORAN ============
func menuLaporan(data tabpol, nData int) {
	var pilihan int
	for {
		fmt.Println("\n--- Menu Laporan ---")
		fmt.Println("1. Wilayah dengan Polusi Tertinggi")
		fmt.Println("2. Peringatan Polusi Berbahaya")
		fmt.Println("3. Kelompok Data per Provinsi")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		pilihan = scanInt()

		switch pilihan {
		case 1:
			laporanTertinggi(data, nData)
		case 2:
			peringatanPolusi(data, nData)
		case 3:
			kelompokPerProvinsi(data, nData)
		case 0:
			fmt.Println("Kembali ke Menu Utama...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// ============ FUNGSI DATA ============
func tambahData(data *tabpol, nData *int) {
	if *nData >= NMAX-1 {
		fmt.Println("Data penuh!")
		return
	}
	fmt.Print("Kota/Kabupaten : ")
	(*data)[*nData].kota = titleCase(inputTeks())

	fmt.Print("Provinsi       : ")
	(*data)[*nData].provinsi = titleCase(inputTeks())

	fmt.Printf("AQI           : ")
	fmt.Sscan(inputTeks(), &(*data)[*nData].aqi)
	(*data)[*nData].aqi = round2((*data)[*nData].aqi)

	(*data)[*nData].sumber = pilihSumber()

	(*data)[*nData].tanggal = inputTanggal()

	*nData++
	fmt.Println(">> Data berhasil ditambahkan!")
	fmt.Printf(">> Status: %s\n", kategoriPolusi((*data)[*nData-1].aqi))
}

func lihatData(data tabpol, nData int) {
	if nData == 0 {
		fmt.Println("Data kosong.")
		return
	}
	fmt.Println("\nNo | Kota/Kabupaten | Provinsi       | AQI    | Sumber  | Tanggal    | Status")
	fmt.Println("-------------------------------------------------------------------------------------")
	for i := 0; i < nData; i++ {
		fmt.Printf("%2d | %-14s | %-14s | %6.2f | %-7s | %-10s | %s\n",
			i+1,
			data[i].kota,
			data[i].provinsi,
			data[i].aqi,
			data[i].sumber,
			data[i].tanggal,
			kategoriPolusi(data[i].aqi))
	}
}

func ubahData(data *tabpol, nData int) {
	lihatData(*data, nData)
	fmt.Print("Pilih nomor data yang diubah: ")
	idx := scanInt()
	idx--
	if idx < 0 || idx >= nData {
		fmt.Println("Nomor tidak valid!")
		return
	}
	fmt.Print("Kota/Kabupaten baru : ")
	(*data)[idx].kota = titleCase(inputTeks())

	fmt.Print("Provinsi baru       : ")
	(*data)[idx].provinsi = titleCase(inputTeks())

	fmt.Print("AQI baru (misal 87.45) : ")
	fmt.Sscan(inputTeks(), &(*data)[idx].aqi)
	(*data)[idx].aqi = round2((*data)[idx].aqi)

	(*data)[idx].sumber = pilihSumber()

	(*data)[idx].tanggal = inputTanggal()
	fmt.Println(">> Data berhasil diubah!")
}

func hapusData(data *tabpol, nData *int) {
	lihatData(*data, *nData)
	fmt.Print("Pilih nomor data yang dihapus: ")
	idx := scanInt()
	idx--
	if idx < 0 || idx >= *nData {
		fmt.Println("Nomor tidak valid!")
		return
	}
	for i := idx; i < *nData-1; i++ {
		(*data)[i] = (*data)[i+1]
	}
	*nData--
	fmt.Println(">> Data berhasil dihapus!")
}

// ============ FUNGSI SORTING (UNTUK BINARY SEARCH) ============
func sortBerdasarkanKota(data *tabpol, nData int) {
	for i := 0; i < nData-1; i++ {
		for j := 0; j < nData-i-1; j++ {
			if (*data)[j].kota > (*data)[j+1].kota {
				temp := (*data)[j]
				(*data)[j] = (*data)[j+1]
				(*data)[j+1] = temp
			}
		}
	}
}

// ============ FUNGSI SEARCHING ============
func sequentialSearch(data tabpol, nData int) {
	fmt.Print("Masukkan nama kota/kabupaten: ")
	cari := titleCase(inputTeks())
	ketemu := false
	for i := 0; i < nData; i++ {
		if data[i].kota == cari {
			fmt.Printf(">> Ditemukan: %s, %s | AQI: %.2f | Tanggal: %s | Status: %s\n",
				data[i].kota, data[i].provinsi, data[i].aqi, data[i].tanggal, kategoriPolusi(data[i].aqi))
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan.")
	}
}

func binarySearch(data tabpol, nData int) {
	fmt.Print("Masukkan nama kota/kabupaten yang dicari: ")
	cari := titleCase(inputTeks())
	kiri, kanan := 0, nData-1
	
	ketemu := false
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah].kota == cari {
			// Jika ketemu, kita harus cek data di sekitar indeks 'tengah' jika ada kota dengan nama sama
			// Binary search standar hanya mengembalikan 1 elemen, agar rapi kita tampilkan datanya
			fmt.Printf(">> Ditemukan: %s, %s | AQI: %.2f | Tanggal: %s | Status: %s\n",
				data[tengah].kota, data[tengah].provinsi, data[tengah].aqi, data[tengah].tanggal, kategoriPolusi(data[tengah].aqi))
			ketemu = true
			break // Keluar dari loop setelah ditemukan
		} else if data[tengah].kota < cari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	
	if !ketemu {
		fmt.Println("Data tidak ditemukan.")
	}
}

// ============ FUNGSI SORTING ============
func selectionSort(data *tabpol, nData int) {
	pass := 1
	for pass < nData {
		idx := pass - 1
		i := pass
		for i < nData {
			// Descending (terbesar di atas)
			if data[idx].aqi < data[i].aqi {
				idx = i
			}
			i++
		}
		temp := data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
		pass++
	}
}

func insertionSort(data *tabpol, nData int) {
	for i := 1; i < nData; i++ {
		key := (*data)[i]
		j := i - 1
		// EDIT: Diubah dari > menjadi < agar Descending (Tanggal Terbaru di atas)
		for j >= 0 && (*data)[j].tanggal < key.tanggal {
			(*data)[j+1] = (*data)[j]
			j--
		}
		(*data)[j+1] = key
	}
}

// ============ FUNGSI LAPORAN ============
func laporanTertinggi(data tabpol, nData int) {
	if nData == 0 {
		fmt.Println("Data kosong.")
		return
	}
	maxIdx := 0
	for i := 1; i < nData; i++ {
		if data[i].aqi > data[maxIdx].aqi {
			maxIdx = i
		}
	}
	fmt.Printf(">> Polusi Tertinggi: %s, %s | AQI: %.2f | Status: %s\n",
		data[maxIdx].kota, data[maxIdx].provinsi, data[maxIdx].aqi, kategoriPolusi(data[maxIdx].aqi))
}

func peringatanPolusi(data tabpol, nData int) {
	fmt.Println("\n>> Peringatan Wilayah Berbahaya:")
	ada := false
	for i := 0; i < nData; i++ {
		if data[i].aqi > 100 {
			fmt.Printf("   [!] %s, %s - AQI: %.2f - %s\n",
				data[i].kota, data[i].provinsi, data[i].aqi, kategoriPolusi(data[i].aqi))
			ada = true
		}
	}
	if !ada {
		fmt.Println("   Tidak ada wilayah dengan polusi berbahaya.")
	}
}

// ============ LAPORAN PER PROVINSI ============
func kelompokPerProvinsi(data tabpol, nData int) {
	if nData == 0 {
		fmt.Println("Data kosong.")
		return
	}

	var provinsiList [NMAX]string
	nProvinsi := 0

	for i := 0; i < nData; i++ {
		sudahAda := false
		for j := 0; j < nProvinsi; j++ {
			if provinsiList[j] == data[i].provinsi {
				sudahAda = true
				break
			}
		}
		if !sudahAda {
			provinsiList[nProvinsi] = data[i].provinsi
			nProvinsi++
		}
	}

	fmt.Println("\n========================================")
	fmt.Println("   LAPORAN DATA POLUSI PER PROVINSI")
	fmt.Println("========================================")

	for p := 0; p < nProvinsi; p++ {
		fmt.Printf("\n>> Provinsi: %s\n", provinsiList[p])
		fmt.Println("   No | Kota/Kabupaten | AQI    | Sumber  | Tanggal    | Status")
		fmt.Println("   ---------------------------------------------------------------")

		count := 0
		totalAQI := 0.0
		maxAQI := -1.0
		minAQI := 99999.0

		for i := 0; i < nData; i++ {
			if data[i].provinsi == provinsiList[p] {
				count++
				totalAQI += data[i].aqi
				if data[i].aqi > maxAQI {
					maxAQI = data[i].aqi
				}
				if data[i].aqi < minAQI {
					minAQI = data[i].aqi
				}
				fmt.Printf("   %2d | %-14s | %6.2f | %-7s | %-10s | %s\n",
					count,
					data[i].kota,
					data[i].aqi,
					data[i].sumber,
					data[i].tanggal,
					kategoriPolusi(data[i].aqi))
			}
		}

		rataRata := round2(totalAQI / float64(count))
		fmt.Printf("   --- Ringkasan: %d data | Rata-rata AQI: %.2f | Tertinggi: %.2f | Terendah: %.2f\n", count, rataRata, maxAQI, minAQI)
	}
}

// ============ MAIN ============
func main() {
	var data tabpol
	var nData, pilihan int

	for {
		tampilMenuUtama()
		pilihan = scanInt()

		switch pilihan {
		case 1:
			menuKelolaData(&data, &nData)
		case 2:
			// Ubah parameter agar menuPencarian bisa mengubah isi data (mem-passing pointer)
			menuPencarian(&data, nData)
		case 3:
			menuPengurutan(&data, nData)
		case 4:
			menuLaporan(data, nData)
		}

		if pilihan == 5 {
			fmt.Println("Terima kasih! Program selesai.")
			break
		}
	}
}