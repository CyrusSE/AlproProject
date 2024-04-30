package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/common-nighthawk/go-figure"
)

const NMAX int = 1000

type dataAccount struct {
	nama     string
	password string
	username string
	PIN      int
}

type acc [NMAX]dataAccount

var pilihan int

func main() {
	var account acc
	main_menu(&account)
	fmt.Print(account[0].nama)
}

func main_menu(akun *acc) {
	menu("Selamat datang di aplikasi konsultasi online.", "1. Pasien", "2. Dokter", "3. Liat Forum Konsultasi", "4. Exit", "", "Silahkan pilih keperluan anda.", "MAIN", false, akun)
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		pasien(*akun)
	case 2:
		fmt.Println("SOON")
	case 3:
		fmt.Println("SOON")
	case 4:
		bye()
	default:
		menu("Masukkan angka sesuai option.", "1. Pasien", "2. Dokter", "3. Liat Forum Konsultasi", "4. Exit", "", "Silahkan pilih keperluan anda.", "MAIN", false, akun)
		fmt.Scan(&pilihan)
	}
}

func pasien(akun acc) {
	menu("Selamat datang di menu pasien.", "1. Login Akun Pasien", "2. Buat Akun Baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &akun)
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		menu("", "Login Akun Pasien", "", "", "", "", "Silahkan masukkan username dan password anda.", "PASIEN", true, &akun)
		loginAccount(akun)
	case 2:
		menu("", "Buat Akun Baru Pasien", "", "", "", "", "Silahkan masukkan username dan password anda.", "PASIEN", false, &akun)
	case 3:
		//Recovery
	case 4:
		main_menu(&akun)
	default:
		menu("Masukkan angka sesuai option.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &akun)
		fmt.Scan(&pilihan)
	}
}

func menu(header, kata1, kata2, kata3, kata4, kata5, pesan, menuType string, hide bool, akun *acc) {
	clear()
	color("a")
	title("TUGAS BESAR ALGORITMA PEMPROGRAMAN")
	fmt.Println("    ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("%4s║%73s║\n", "", "")
	fmt.Printf("%4s║%53s%-20s║\n", "", "Aplikasi Konsultasi Kesehatan", "")
	fmt.Printf("%4s║%73s║\n", "", "")
	fmt.Printf("%4s║%55s %17s║\n", "", "Created by : • Faisal Ihsan Santoso", "")
	fmt.Printf("%4s║%59s %13s║\n", "", "• Arie Farchan Fyrzatullah", "")
	fmt.Printf("%4s║%73s║\n", "", "")
	fmt.Println("    ╠═════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("%4s║%73s║\n", "", "")
	if header != "" {
		fmt.Printf("%4s║  %-71s║\n", "", header)
		fmt.Printf("%4s║%73s║\n", "", "")
	}
	if kata1 != "" {
		fmt.Printf("%4s║%4s%-69s║\n", "", "", kata1)
	}
	if kata2 != "" {
		fmt.Printf("%4s║%4s%-69s║\n", "", "", kata2)
	}
	if kata3 != "" {
		fmt.Printf("%4s║%4s%-69s║\n", "", "", kata3)
	}
	if kata4 != "" {
		fmt.Printf("%4s║%4s%-69s║\n", "", "", kata4)
	}
	if kata5 != "" {
		fmt.Printf("%4s║%4s%-69s║\n", "", "", kata5)
	}
	fmt.Printf("%4s║%73s║\n", "", "")
	fmt.Printf("%4s║%73s║\n", "", "")
	fmt.Printf("%4s║  %-71s║\n", "", pesan)
	fmt.Printf("%4s║%73s║\n", "", "")
	if menuType == "PASIEN" {
		fmt.Printf("%4s║%59s╔═════════════╣\n", "", "")
		fmt.Printf("%4s║%73s║\n", "", "║ MENU PASIEN ")
		fmt.Println("    ╚═══════════════════════════════════════════════════════════╩═════════════╝")
	} else if menuType == "DOKTER" {
		fmt.Printf("%4s║%56s╔════════════════╣\n", "", "")
		fmt.Printf("%4s║%73s║\n", "", "║ MENU DOKTER ")
		fmt.Println("    ╚════════════════════════════════════════════════════════╩════════════════╝")
	} else {
		fmt.Printf("    ║%61s╔═══════════╣\n", "")
		fmt.Printf("    ║%72s %s║\n", "║ MAIN MENU", "")
		fmt.Println("    ╚═════════════════════════════════════════════════════════════╩═══════════╝")
	}
	if !hide {
		fmt.Printf("%4sSELECT> ", "")
	}
}

func loginAccount(A acc) {
	var username, password string
	var found bool
	fmt.Printf("%3s %s", "", "Masukkan username Anda : ")
	fmt.Scan(&username)
	fmt.Printf("%3s %s", "", "Masukkan password Anda : ")
	fmt.Scan(&password)
	for i := 0; i < len(A); i++ {
		if A[i].username == username && A[i].password == password {
			//Login Berhasil
			found = true
			break
		} else if A[i].username == username && A[i].password != password {
			menu("Password anda salah, gunakan PIN untuk recovery akun anda.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &A)
			found = true
			break
		}
	}
	if !found {
		menu("Akun anda tidak ditemukan, silahkan buat akun baru.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &A)
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func color(text string) {
	cmd := exec.Command("cmd", "/c", "color ", text)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func title(text string) {
	cmd := exec.Command("cmd", "/c", fmt.Sprintf("title %s", text))
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func bye() {
	clear()
	myFigure := figure.NewFigure("GOOD BYE!", "", true)
	myFigure.Print()
}
