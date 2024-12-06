package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x)

	switch {
	case x%2 != 0 && x == 5:
		y = x + 1
		z = x + y
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d ", x, y, z)

	case x%10 == 0:
		z = x / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d ", x, z)

	case x%5 == 0 :
		z = x * x
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil pembagian antara %d ^2 = %d ", x, z)

	case x%2 == 0:
		y = x + 1
		z = x * y
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d ", x, y, z)

	}
}
