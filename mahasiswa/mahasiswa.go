package mahasiswa

import (
	"fmt"
	"time"
)

type MataKuliah struct {
	Nama  string
	Nilai int
}

type Alamat struct {
	City      string
	Kabupatan string
}

type Mahasiswa struct {
	NIM        string
	Umur       int
	Nama       string
	Address    Alamat
	MataKuliah []MataKuliah
}

func (m Mahasiswa) Absen() {
	fmt.Println(time.Now())
}

func (m Mahasiswa) Daftar(nilai1 int) int {
	return m.Umur + nilai1
}

func (m *Mahasiswa) ChangeUmur(umur int) {
	m.Umur = umur
}

func (m *Mahasiswa) Alamat() (string, string, int) {
	return "Jakarta Barat", "Tomang", 6
}

func (m *Mahasiswa) TambahMatkul(nama string, nilai int) {
	m.MataKuliah = append(m.MataKuliah, MataKuliah{
		Nama:  nama,
		Nilai: nilai,
	})
}
