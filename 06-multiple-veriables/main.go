package main

import (
	"fmt"
)

// Çoklu değer döndüren bir fonksiyon
func bolme(x, y int) (int, int) {
	return x / y, x % y
}

// Paket düzeyinde çoklu değişken tanımlama
var (
	name     string = "Jeremy"
	age      int    = 23
	isActive bool   = true
)

func main() {
	// 1. var ile çoklu aynı türde değişken tanımlama
	var x, y, z int
	x = 1
	y = 2
	z = 3
	fmt.Println("x, y, z:", x, y, z)

	// 2. := ile çoklu farklı türde tanımlama
	a, b, c := 42, 3.14, "Hello, World!"
	fmt.Println("a (int):", a)
	fmt.Println("b (float64):", b)
	fmt.Println("c (string):", c)

	// 3. Fonksiyondan çoklu dönüş değeri alma
	division, remainder := bolme(10, 3)
	fmt.Println("Bölme işlemi → Bölüm:", division, ", Kalan:", remainder)

	// 4. Blank identifier (_) ile kullanılmayan dönüş değerini atlama
	_, onlyRemainder := bolme(25, 4)
	fmt.Println("Sadece kalan değeri:", onlyRemainder)

	// 5. Paket düzeyindeki değişkenlerin kullanımı
	fmt.Println("İsim:", name)
	fmt.Println("Yaş:", age)
	fmt.Println("Aktif mi?:", isActive)
}
