package main

import "fmt"

func main() {
	var uts, uas, tugas float64

	fmt.Print("Masukkan nilai UTS: ")
	fmt.Scanln(&uts)

	fmt.Print("Masukkan nilai UAS: ")
	fmt.Scanln(&uas)

	fmt.Print("Masukkan nilai Tugas: ")
	fmt.Scanln(&tugas)

	nilaiAkhir := (uts * 0.3) + (uas * 0.4) + (tugas * 0.3)

	fmt.Printf("Nilai Akhir Mahasiswa: %.1f\n", nilaiAkhir)
}
