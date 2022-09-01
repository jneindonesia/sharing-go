package main

import (
	"fmt"
	"time"
)

func main() {

	defer fmt.Println("PULANG YUK")

	time.AfterFunc(time.Second*2, func() {
		fmt.Println("await")
	})

	time.Sleep(time.Second * 5)

	// var namaVariabel string
	// namaVariabel = "alan"
	// namaVariabel2 := 1000

	// var (
	// 	X_1 string = "string"
	// 	X_2 string = "string2"
	// 	X_3 string = "string3"
	// 	Y_1 int    = 1
	// )

	// a, b := 1, 2

	// const konstanta = 1000

	// fmt.Println(namaVariabel)
	// fmt.Println(namaVariabel2)
	// fmt.Println(konstanta)
	// fmt.Println(X_1)
	// fmt.Println(X_2)
	// fmt.Println(X_3)
	// fmt.Println(Y_1)

	// // OPERASI MATEMATIKA
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a / b)
	// fmt.Println(a * b)
	// fmt.Println(a % b)

	// alan := "pria"
	// if alan == "pria" || alan == "laki-laki" {
	// 	fmt.Println("benar")
	// } else if alan == "non binary" {
	// 	fmt.Println("mungkin")
	// } else {
	// 	fmt.Println("salah")
	// }

	// switch alan {
	// case "alan":
	// 	fmt.Println("benar")
	// case "non binary":
	// 	fmt.Println("mungkin")
	// default:
	// 	fmt.Println("salah")
	// }

	// arrayOfInt := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(arrayOfInt)

	// mapData := map[string]int{
	// 	"index_0": 0,
	// 	"index_1": 1,
	// }

	// mapData2 := map[string]interface{}{
	// 	"index_0": 2,
	// 	"index_1": "1",
	// }

	// fmt.Println("mapData", mapData)
	// fmt.Println("mapData", mapData["index_0"])
	// fmt.Println("mapData2", mapData2["index_0"])
	// fmt.Println("pertambahan", mapData["index_0"]+mapData2["index_0"].(int))

	// rizki := "rizky"
	// var syauqi *string
	// syauqi = &rizki

	// rizki = "rizki duplicated"
	// rizki = "dayat"

	// fmt.Println(&rizki, rizki)
	// fmt.Println(syauqi, *syauqi)

	// var mhs1 mahasiswa.Mahasiswa
	// mhs1.NIM = "64920"
	// mhs1.Umur = 22
	// mhs1.Nama = "dayat"

	// fmt.Println(mhs1)
	// fmt.Println(mhs1.Umur)
	// mhs1.Absen()

	// for i := 0; i < len(arrayOfInt); i++ {
	// 	fmt.Println(mhs1.Umur, arrayOfInt[i])
	// }

	// // twoDimension := [2][2]int{
	// // 	[2]int{0, 1},
	// // 	[2]int{2, 3},
	// // }

	// for key, value := range mapData {
	// 	fmt.Println(key, value)
	// }

	// fmt.Println(mhs1.Daftar(20))
	// mhs1.ChangeUmur(30)
	// fmt.Println("Umur", mhs1.Umur)
	// city, kabupaten, no := mhs1.Alamat()
	// fmt.Println(city, kabupaten, no)

	// mhs1.TambahMatkul("matematika", 0)
	// mhs1.TambahMatkul("ipa", 100)
	// fmt.Println(mhs1.MataKuliah)

	// slice1 := arrayOfInt[0:2]

	// arrayOfInt[0] = 500
	// fmt.Println(slice1)
}
