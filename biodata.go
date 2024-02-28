package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var temanKelas = map[int]Teman{
	1: {1, "Ani", "Jakarta", "Web Developer", "Ingin memperdalam pengetahuan tentang Golang"},
	2: {2, "Budi", "Bandung", "Data Scientist", "Menggunakan Golang untuk analisis data"},
	3: {3, "Caca", "Medan", "Mahasiswa", "Memperbesar mendapatkan peluang pekerjaan"},
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Gunakan: go run biodata.go [nomor_absen]")
		return
	}

	absen := args[0]


	var absenInt int
	_, err := fmt.Sscanf(absen, "%v", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}


	teman, ok := temanKelas[absenInt]
	if !ok {
		fmt.Println("Tidak ada data teman dengan nomor absen tersebut")
		return
	}


	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
