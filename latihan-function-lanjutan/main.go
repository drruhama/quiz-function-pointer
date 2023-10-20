package main

import "fmt"

func main() {
	nilaiSiswa := []float32{85.5, 92.0, 78.5, 90.0, 88.5}
	fmt.Println("Nilai Siswa : ", nilaiSiswa)
	var jum int = len(nilaiSiswa)

	rataRata := rataRatakan(jum, nilaiSiswa...)
	fmt.Println("Jumlah nilai siswa dalam kelas :", jum)
	fmt.Println("Nilai rata-rata siswa dalam kelas :", rataRata)
}

func rataRatakan(n int, sliceNilai ...float32) float32 {
	var tot float32 = 0
	var i int = 0
	for i = 0; i < n; i++ {
		tot += (sliceNilai[i])
	}
	var nilaiRatarata float32 = float32(tot) / float32(n)
	return nilaiRatarata
}
