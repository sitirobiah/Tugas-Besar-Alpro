package main

import (
    "fmt"
)

type Proyek struct {
    Nama     string
    Klien    string
    Deadline int
    Bayaran  int
    Status   string // "pending", "dikerjakan", "selesai"
}

var daftar []Proyek
var pilihan int

func tambahProyek() {
    var p Proyek
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
    daftar = append(daftar, p)
    fmt.Println("Proyek berhasil ditambahkan.")
}

func ubahStatusProyek() {
    var nama string
    var statusBaru string
    var i int
    fmt.Print("Masukkan nama proyek yang ingin diubah statusnya: ")
    fmt.Scanln(&nama)
    for i = 0; i < len(daftar); i++ {
        if daftar[i].Nama == nama {
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
    var nama string
    var i int
    fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
    fmt.Scanln(&nama)
    for i = 0; i < len(daftar); i++ {
        if daftar[i].Nama == nama {
            daftar = append(daftar[:i], daftar[i+1:]...)
            fmt.Println("Proyek berhasil dihapus.")
            return
        }
    }
    fmt.Println("Proyek tidak ditemukan.")
}

func cariProyekSequential() {
    var kata string
    var p Proyek
    fmt.Print("Masukkan nama proyek atau klien: ")
    fmt.Scanln(&kata)
    for _, p = range daftar {
        if p.Nama == kata || p.Klien == kata {
            fmt.Println("Ditemukan:", p)
            return
        }
    }
    fmt.Println("Proyek tidak ditemukan.")
}

func cariProyekBinary() {
    var i, idx, pass int
    var temp Proyek

    pass = 1
    for pass < len(daftar) {
        idx = pass - 1
        i = pass
        for i < len(daftar) {
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
    var kiri, kanan int
    fmt.Print("Masukkan nama proyek: ")
    fmt.Scanln(&nama)
    kiri = 0
    kanan = len(daftar) - 1
    for kiri <= kanan {
        var tengah int = (kiri + kanan) / 2
        if daftar[tengah].Nama == nama {
            fmt.Println("Ditemukan:", daftar[tengah])
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
    for i = 0; i < len(daftar)-1; i++ {
        min = i
        for j = i + 1; j < len(daftar); j++ {
            if daftar[j].Deadline < daftar[min].Deadline {
                min = j
            }
        }
        daftar[i], daftar[min] = daftar[min], daftar[i]
    }
    fmt.Println("Diurutkan berdasarkan deadline (Selection Sort).")
}

func urutkanBayaranInsertionSort() {
    var i, j int
    var key Proyek
    for i = 1; i < len(daftar); i++ {
        key = daftar[i]
        j = i - 1
        for j >= 0 && daftar[j].Bayaran < key.Bayaran {
            daftar[j+1] = daftar[j]
            j--
        }
        daftar[j+1] = key
    }
    fmt.Println("Diurutkan berdasarkan bayaran (Insertion Sort).")
}

func tampilkanLaporan() {
    var p Proyek
    fmt.Println("\nLaporan Proyek:")
    for _, p = range daftar {
        fmt.Printf("Nama: %s | Klien: %s | Deadline: %d | Bayaran: %d | Status: %s\n",
            p.Nama, p.Klien, p.Deadline, p.Bayaran, p.Status)
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