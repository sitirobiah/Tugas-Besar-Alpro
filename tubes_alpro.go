package main

import (
	"fmt"
)

type Proyek struct {
	ID       int
	Nama     string
	Klien    string
	Deadline int
	Bayaran  int
	Status   string
}

var daftar [100]Proyek
var jumlahProyek int = 0
var idTerakhir int = 0
var pilihan int

func tambahProyek() {
	if jumlahProyek >= len(daftar) {
		fmt.Println("Kapasitas proyek penuh.")
		return
	}

	var p Proyek
	idTerakhir++
	p.ID = idTerakhir
	fmt.Print("Nama proyek: ")
	fmt.Scanln(&p.Nama)
	fmt.Print("Klien: ")
	fmt.Scanln(&p.Klien)
	fmt.Print("Deadline (dalam hari): ")
	fmt.Scanln(&p.Deadline)
	fmt.Print("Bayaran: ")
	fmt.Scanln(&p.Bayaran)
	fmt.Print("Status (pending/dikerjakan/selesai): ")
	fmt.Scanln(&p.Status)

	daftar[jumlahProyek] = p
	jumlahProyek++
	fmt.Println("Proyek berhasil ditambahkan.")
}

func ubahStatusProyek() {
	var id int
	var statusBaru string
	var i int

	fmt.Print("Masukkan ID proyek yang ingin diubah statusnya: ")
	fmt.Scanln(&id)
	for i = 0; i < jumlahProyek; i++ {
		if daftar[i].ID == id {
			fmt.Print("Masukkan status baru: ")
			fmt.Scanln(&statusBaru)
			daftar[i].Status = statusBaru
			fmt.Println("Status proyek diperbarui.")
			return
		}
	}
	fmt.Println("Proyek tidak ditemukan.")
}

func hapusProyek() {
	var id, i, j int
	fmt.Print("Masukkan ID proyek yang ingin dihapus: ")
	fmt.Scanln(&id)
	for i = 0; i < jumlahProyek; i++ {
		if daftar[i].ID == id {
			for j = i; j < jumlahProyek-1; j++ {
				daftar[j] = daftar[j+1]
			}
			jumlahProyek--
			fmt.Println("Proyek berhasil dihapus.")
			return
		}
	}
	fmt.Println("Proyek tidak ditemukan.")
}

func cariProyekSequential() {
	var kata string
	var i int

	fmt.Print("Masukkan nama proyek atau klien: ")
	fmt.Scanln(&kata)
	for i = 0; i < jumlahProyek; i++ {
		if daftar[i].Nama == kata || daftar[i].Klien == kata {
			p := daftar[i]
			fmt.Printf("ID: %d, Nama: %s, Klien: %s, Deadline: %d hari, Bayaran: %d, Status: %s\n",
				p.ID, p.Nama, p.Klien, p.Deadline, p.Bayaran, p.Status)
			return
		}
	}
	fmt.Println("Proyek tidak ditemukan.")
}

func cariProyekBinary() {
	var i, idx, pass int
	var temp Proyek

	pass = 1
	for pass < jumlahProyek {
		idx = pass - 1
		i = pass
		for i < jumlahProyek {
			if daftar[i].Nama < daftar[idx].Nama {
				idx = i
			}
			i++
		}
		temp = daftar[pass-1]
		daftar[pass-1] = daftar[idx]
		daftar[idx] = temp
		pass++
	}

	var nama string
	var kiri, kanan, tengah int
	fmt.Print("Masukkan nama proyek: ")
	fmt.Scanln(&nama)
	kiri = 0
	kanan = jumlahProyek - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if daftar[tengah].Nama == nama {
			p := daftar[tengah]
			fmt.Printf("ID: %d, Nama: %s, Klien: %s, Deadline: %d hari, Bayaran: %d, Status: %s\n",
				p.ID, p.Nama, p.Klien, p.Deadline, p.Bayaran, p.Status)
			return
		} else if daftar[tengah].Nama < nama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	fmt.Println("Proyek tidak ditemukan.")
}

func urutkanDeadlineSelectionSort() {
	var i, j, min int
	for i = 0; i < jumlahProyek-1; i++ {
		fmt.Println("Iterasi ke-", i)
		min = i
		for j = i + 1; j < jumlahProyek; j++ {
			if daftar[j].Deadline < daftar[min].Deadline {
				min = j
			}
		}
		daftar[i], daftar[min] = daftar[min], daftar[i]
	}
	fmt.Println("Diurutkan berdasarkan deadline (Selection Sort).")
	tampilkanLaporan()
}

func urutkanBayaranInsertionSort() {
	var i, j int
	var key Proyek
	for i = 1; i < jumlahProyek; i++ {
		key = daftar[i]
		j = i - 1
		for j >= 0 && daftar[j].Bayaran < key.Bayaran {
			daftar[j+1] = daftar[j]
			j--
		}
		daftar[j+1] = key
	}
	fmt.Println("Diurutkan berdasarkan bayaran (Insertion Sort).")
	tampilkanLaporan()
}

func tampilkanLaporan() {
	var i int
	fmt.Println("\nLaporan Proyek:")
	for i = 0; i < jumlahProyek; i++ {
		fmt.Printf("ID: %d, Nama: %s | Klien: %s | Deadline: %d | Bayaran: %d | Status: %s\n",
			daftar[i].ID, daftar[i].Nama, daftar[i].Klien,
			daftar[i].Deadline, daftar[i].Bayaran, daftar[i].Status)
	}
}

func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Ubah Status Proyek")
		fmt.Println("3. Hapus Proyek")
		fmt.Println("4. Cari Proyek (Sequential)")
		fmt.Println("5. Cari Proyek (Binary)")
		fmt.Println("6. Urutkan Proyek (Deadline - Selection Sort)")
		fmt.Println("7. Urutkan Proyek (Bayaran - Insertion Sort)")
		fmt.Println("8. Tampilkan Laporan Proyek")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahProyek()
		case 2:
			ubahStatusProyek()
		case 3:
			hapusProyek()
		case 4:
			cariProyekSequential()
		case 5:
			cariProyekBinary()
		case 6:
			urutkanDeadlineSelectionSort()
		case 7:
			urutkanBayaranInsertionSort()
		case 8:
			tampilkanLaporan()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}
