package main

import "fmt"

func main() {
	var nama string
	var a, b float64

	fmt.Print("Masukkan nama karyawan: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan gaji pokok: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan persentase pajak (%): ")
	fmt.Scanln(&b)

	pajak := (b / 100) * a
	bersih := a - pajak

	fmt.Println("\nHasil Perhitungan Gaji:")
	fmt.Printf("Nama Karyawan: %s\n", nama)
	fmt.Printf("Gaji Pokok: %.0f\n", a)
	fmt.Printf("Pajak: %.0f\n", pajak)
	fmt.Printf("Gaji Bersih: %.0f\n", bersih)
}
