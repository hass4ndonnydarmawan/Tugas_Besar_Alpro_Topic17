package main

import "fmt"

const MAX int = 100

type Kandidat struct {
	NoUrut   int
	Nama     string
	VisiMisi string
	Suara    int
}

type DaftarKandidat [MAX]Kandidat

func SequentialSearch(A DaftarKandidat, n int, x int) int {
	for i := 0; i < n; i++ {
		if A[i].NoUrut == x {
			return i
		}
	}
	return -1
}

func BinarySearch(A DaftarKandidat, n int, x int) int {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if A[tengah].NoUrut == x {
			return tengah
		} else if A[tengah].NoUrut < x {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func SelectionSortNoUrut(A *DaftarKandidat, n int) {
	var i, j, min int
	var temp Kandidat

	for i = 0; i < n-1; i++ {
		min = i

		for j = i + 1; j < n; j++ {
			if A[j].NoUrut < A[min].NoUrut {
				min = j
			}
		}

		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

func InsertionSortSuara(A *DaftarKandidat, n int) {
	var i, j int
	var temp Kandidat

	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1

		for j >= 0 && A[j].Suara < temp.Suara {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}

func CetakStatistik(A DaftarKandidat, n int) {
	var total int

	for i := 0; i < n; i++ {
		total += A[i].Suara
	}

	fmt.Println("\n+++ HASIL STATISTIK E-VOTING +++")

	if total == 0 {
		fmt.Println("Belum ada suara masuk.")
		return
	}

	for i := 0; i < n; i++ {
		persen := (float64(A[i].Suara) / float64(total)) * 100

		fmt.Printf("No Urut : %d\n", A[i].NoUrut)
		fmt.Printf("Nama    : %s\n", A[i].Nama)
		fmt.Printf("Persen  : %.2f%%\n", persen)
		fmt.Printf("Suara   : %d\n\n", A[i].Suara)
	}

	fmt.Printf("Total Suara Masuk : %d\n", total)
}

func TampilkanKandidat(A DaftarKandidat, n int) {
	if n == 0 {
		fmt.Println("Belum ada kandidat.")
		return
	}

	fmt.Println("\n=== DATA KANDIDAT ===")

	for i := 0; i < n; i++ {
		fmt.Printf("No Urut   : %d\n", A[i].NoUrut)
		fmt.Printf("Nama      : %s\n", A[i].Nama)
		fmt.Printf("VisiMisi  : %s\n", A[i].VisiMisi)
		fmt.Printf("Suara     : %d\n\n", A[i].Suara)
	}
}

func main() {
	var tabel DaftarKandidat
	var jumlah int
	var menu int

	for {
		fmt.Println("\n    APLIKASI E-VOTING    ")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Edit Kandidat")
		fmt.Println("3. Hapus Kandidat")
		fmt.Println("4. Cari Kandidat Sequential Search")
		fmt.Println("5. Cari Kandidat Binary Search")
		fmt.Println("6. Voting Kandidat")
		fmt.Println("7. Tampilkan Ranking Suara")
		fmt.Println("8. Tampilkan Statistik")
		fmt.Println("9. Tampilkan Semua Kandidat")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		if menu == 1 {

			if jumlah >= MAX {
				fmt.Println("Kapasitas penuh.")
			} else {

				var k Kandidat
				var cek int

				fmt.Print("Masukkan Nomor Urut: ")
				fmt.Scan(&k.NoUrut)

				cek = SequentialSearch(tabel, jumlah, k.NoUrut)

				if cek != -1 {
					fmt.Println("Nomor urut sudah digunakan.")
				} else {

					fmt.Print("Masukkan Nama Kandidat: ")
					fmt.Scan(&k.Nama)

					fmt.Print("Masukkan Visi Misi: ")
					fmt.Scan(&k.VisiMisi)

					k.Suara = 0

					tabel[jumlah] = k
					jumlah++

					fmt.Println("Kandidat berhasil ditambahkan.")
				}
			}

		} else if menu == 2 {

			var target int

			fmt.Print("Masukkan Nomor Urut Kandidat: ")
			fmt.Scan(&target)

			pos := SequentialSearch(tabel, jumlah, target)

			if pos != -1 {

				fmt.Print("Masukkan Nama Baru: ")
				fmt.Scan(&tabel[pos].Nama)

				fmt.Print("Masukkan Visi Misi Baru: ")
				fmt.Scan(&tabel[pos].VisiMisi)

				fmt.Println("Data berhasil diubah.")

			} else {
				fmt.Println("Kandidat tidak ditemukan.")
			}

		} else if menu == 3 {

			var target int

			fmt.Print("Masukkan Nomor Urut Kandidat: ")
			fmt.Scan(&target)

			pos := SequentialSearch(tabel, jumlah, target)

			if pos != -1 {

				for i := pos; i < jumlah-1; i++ {
					tabel[i] = tabel[i+1]
				}

				jumlah--

				fmt.Println("Data berhasil dihapus.")

			} else {
				fmt.Println("Kandidat tidak ditemukan.")
			}

		} else if menu == 4 {

			var target int

			fmt.Print("Masukkan Nomor Urut Kandidat: ")
			fmt.Scan(&target)

			pos := SequentialSearch(tabel, jumlah, target)

			if pos != -1 {

				fmt.Println("\n=== DATA DITEMUKAN ===")
				fmt.Printf("No Urut  : %d\n", tabel[pos].NoUrut)
				fmt.Printf("Nama     : %s\n", tabel[pos].Nama)
				fmt.Printf("VisiMisi : %s\n", tabel[pos].VisiMisi)
				fmt.Printf("Suara    : %d\n", tabel[pos].Suara)

			} else {
				fmt.Println("Data tidak ditemukan.")
			}

		} else if menu == 5 {

			var target int
			var tabelCari DaftarKandidat

			tabelCari = tabel

			SelectionSortNoUrut(&tabelCari, jumlah)

			fmt.Print("Masukkan Nomor Urut Kandidat: ")
			fmt.Scan(&target)

			pos := BinarySearch(tabelCari, jumlah, target)

			if pos != -1 {

				fmt.Println("\n=== DATA DITEMUKAN ===")
				fmt.Printf("No Urut  : %d\n", tabelCari[pos].NoUrut)
				fmt.Printf("Nama     : %s\n", tabelCari[pos].Nama)
				fmt.Printf("VisiMisi : %s\n", tabelCari[pos].VisiMisi)
				fmt.Printf("Suara    : %d\n", tabelCari[pos].Suara)

			} else {
				fmt.Println("Data tidak ditemukan.")
			}

		} else if menu == 6 {

			var target int

			if jumlah == 0 {

				fmt.Println("Belum ada kandidat.")

			} else {

				TampilkanKandidat(tabel, jumlah)

				fmt.Print("Masukkan Nomor Urut Pilihan: ")
				fmt.Scan(&target)

				pos := SequentialSearch(tabel, jumlah, target)

				if pos != -1 {

					tabel[pos].Suara++
					fmt.Println("Voting berhasil.")

				} else {
					fmt.Println("Pilihan tidak ditemukan.")
				}
			}

		} else if menu == 7 {

			if jumlah == 0 {

				fmt.Println("Belum ada kandidat.")

			} else {

				var tabelUrut DaftarKandidat

				tabelUrut = tabel

				InsertionSortSuara(&tabelUrut, jumlah)

				fmt.Println("\n=== RANKING SUARA ===")

				for i := 0; i < jumlah; i++ {

					fmt.Printf("%d. %s - %d suara\n",
						i+1,
						tabelUrut[i].Nama,
						tabelUrut[i].Suara)
				}
			}

		} else if menu == 8 {

			CetakStatistik(tabel, jumlah)

		} else if menu == 9 {

			TampilkanKandidat(tabel, jumlah)

		} else if menu == 10 {

			fmt.Println("Program selesai.")
			break

		} else {

			fmt.Println("Menu tidak valid.")
		}
	}
}