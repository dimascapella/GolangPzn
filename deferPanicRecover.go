package main

import "fmt"

func main() {
	deferApps()
	// panicApps(true)
	recoverApps(true)
}

func endApps() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	fmt.Println("Terjadi Error ", message)
}

func deferApps() {
	defer endApps()
	fmt.Println("App Berjalan")
	nums := 0 / 2
	fmt.Println(nums)
}

// func panicApps(error bool) {
// 	defer endApps()
// 	if error {
// 		panic("Error Apps")
// 	}
// 	fmt.Println("Hella")
// }

func recoverApps(error bool) {
	defer endApps()
	if error {
		panic("Error Apps")
	}
}
