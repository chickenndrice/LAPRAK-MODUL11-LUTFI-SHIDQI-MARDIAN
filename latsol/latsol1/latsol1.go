package main

import "fmt"

func main() {
	var ph float64
	var ket string
	fmt.Scan(&ph)
	
	switch {
	case ph >= 0 && ph < 6.5:
		ket = "Air tidak layak minum"
	case ph > 8.6 && ph <= 14:
		ket = "Air tidak layak minum"
	case ph >= 6.5 && ph <= 8.6:
		ket = "Air layak minum"
	default:
		ket = "Nilai pH tidak valid. Nilai pH harus antara 0 dan 14"
	}
	fmt.Println(ket)
}