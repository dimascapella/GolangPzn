package main

import (
	"fmt"
	"math"
)

func main() {
	// Membulatkan Ke angka terdekat
	MathRound := math.Round(8.3)
	fmt.Println(MathRound)

	// Membulatkan angka ke bawah
	MathFloor := math.Floor(4.9)
	fmt.Println(MathFloor)

	// Membulatkan angka ke atas
	MathCeil := math.Ceil(4.3)
	fmt.Println(MathCeil)

	// Membandingkan nilai yang paling besar
	MathMax := math.Max(10, 20)
	fmt.Println(MathMax)

	// Membandingkan nilai yang paling kecil
	MathMin := math.Min(10, 20)
	fmt.Println(MathMin)
}
