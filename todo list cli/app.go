package main

import (
	"fmt"
)

func menu() int {
	var i int

	fmt.Println("Silahkan Pilih ... :")
	fmt.Println("1. Tambah task")
	fmt.Println("2. Hapus task")
	fmt.Println("3. Lihat daftar task")
	fmt.Println("4. Selesai")
	fmt.Println("5. Exit")

	fmt.Print("Masukan: ")
	fmt.Scan(&i)
	return i
}

func addtask() {

}

func deletetask() {

}

func showtask() {

}

func donetask() {

}

func main() {
	selected := menu()
	fmt.Println("Selected: ", selected)

	switch selected {
	case 1:
		addtask()
	case 2:
		deletetask()
	case 3:
		showtask()
	case 4:
		donetask()
	case 5:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
