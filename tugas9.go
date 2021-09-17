package main

import (
	"errors"
	"fmt"
	"strings"
)

func validasi(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Menu tidak boleh kosong")
	}
	return true, nil
}

var menu, next string

func main() {
	menus()
	input()
	pesan()
	fmt.Println("Terimakasih atas pesanannya")

}

func menus() {
	fmt.Println("Toko Makanan Indonesia")
	fmt.Println("==========================")
	fmt.Println("Tahu")
	fmt.Println("Tempe")
	fmt.Println("Sambal")
	fmt.Println("Nasi")
	fmt.Println("Lele")
	fmt.Println("Ayam")
	fmt.Println("==========================")
}

func input() {
	fmt.Print("Masukkan menu pesanan anda dalam huruf ( eg: Tahu ) : ")
	fmt.Scanln(&menu)

	if valid, err := validasi(menu); valid {
		defer fmt.Println("Pesanan anda : ", menu)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Print("Lanjutkan Memesan ? (Y/T) ")
	fmt.Scanln(&next)

	pesan()
}

func pesan() {
	if next == "y" || next == "Y" {
		input()

	} else {
		fmt.Print()
	}
}
