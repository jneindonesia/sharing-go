package mahasiswa

import (
	"fmt"
	"log"
	"time"
)

type namaOrang string

type behaviourMahasiswa interface {
	Absen()
	Daftar(nilai1 int) int
	Variadic(n ...string)
}

type mataKuliah struct {
	Nama  string
	Nilai int
}

type alamat struct {
	City      string
	Kabupatan string
}

type mahasiswa struct {
	NIM        string
	Umur       int
	Nama       namaOrang
	Address    alamat
	MataKuliah []mataKuliah
}

func NewMahasiswa() behaviourMahasiswa {
	return &mahasiswa{}
}

func (m *mahasiswa) Absen() {
	fmt.Println(time.Now())
}

func (m *mahasiswa) Daftar(nilai1 int) int {
	return m.Umur + nilai1
}

func (m *mahasiswa) ChangeUmur(umur int) {
	m.Umur = umur
}

func (m *mahasiswa) Alamat() (string, string, int) {
	return "Jakarta Barat", "Tomang", 6
}

func (m *mahasiswa) TambahMatkul(nama string, nilai int) {
	m.MataKuliah = append(m.MataKuliah, mataKuliah{
		Nama:  nama,
		Nilai: nilai,
	})
}

func (m *mahasiswa) Variadic(n ...string) {
	log.Println(n)
}
