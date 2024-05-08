package main

import (
	"fmt"
	"os"
	"os/exec"
)

const NMAX int = 1000

type dataAccount struct {
	nama     string
	username string
	password string
	PIN      int
}

type acc [NMAX]dataAccount

var pilihan int

func main() {
	var account acc
	var totalAccount int
	main_menu(&account, &totalAccount)
	fmt.Print(account[0].nama)
}

func main_menu(akun *acc, T *int) {
	menu("Selamat datang di aplikasi konsultasi online.", "1. Pasien", "2. Dokter", "3. Liat Forum Konsultasi", "4. Exit", "", "Silahkan pilih keperluan anda.", "MAIN", false, akun, *T)
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		pasien(*akun, T, false)
	case 2:
		fmt.Println("SOON")
	case 3:
		fmt.Println("SOON")
	case 4:
		bye()
	default:
		menu("Masukkan angka sesuai option.", "1. Pasien", "2. Dokter", "3. Liat Forum Konsultasi", "4. Exit", "", "Silahkan pilih keperluan anda.", "MAIN", false, akun, *T)
		fmt.Scan(&pilihan)
	}
}

func pasien(akun acc, T *int, hide bool) {
	if !hide {
		menu("Selamat datang di menu pasien.", "1. Login Akun Pasien", "2. Buat Akun Baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &akun, *T)
	}
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		menu("", "Login Akun Pasien", "", "", "", "", "Silahkan masukkan username dan password anda.", "PASIEN", true, &akun, *T)
		loginAccount(akun, *T)
	case 2:
		menu("", "Buat Akun Baru Pasien.", "", "", "", "", "Silahkan masukkan username dan password anda.", "PASIEN", true, &akun, *T)
		makeAcc(&akun, T)
	case 3:
		menu("", "Recovery Akun Pasien.", "", "", "", "", "Masukkan username dan PIN anda.", "PASIEN", true, &akun, *T)
		recovery(&akun, *T)
	case 4:
		main_menu(&akun, T)
	default:
		menu("Masukkan angka sesuai option.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &akun, *T)
		fmt.Scan(&pilihan)
	}
}

func menu(header, kata1, kata2, kata3, kata4, kata5, pesan, menuType string, hide bool, akun *acc, T int) {
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

func loginAccount(A acc, T int) {
	var username, password string
	var indexUser int
	fmt.Printf("%3s %s", "", "Masukkan username Anda : ")
	fmt.Scan(&username)
	fmt.Printf("%3s %s", "", "Masukkan password Anda : ")
	fmt.Scan(&password)
	if exist(A, T, username){
		indexUser = findIndexUser(A, T, username)
		if A[indexUser].username == username && A[indexUser].password == password {
			//Login berhasil
		} else {
			menu("Password anda salah, gunakan PIN untuk recovery akun anda.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &A, T)
			pasien(A, &T, true)
		}
	}else{
		menu("Akun anda tidak ditemukan, silahkan buat akun baru.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &A, T)
		pasien(A, &T, true)
	}
}


func makeAcc(A *acc, T *int) {
	var nama, username, password string
	var pin int
	fmt.Printf("%3s %s", "", "Masukkan nama Anda : ")
	fmt.Scan(&nama)
	fmt.Printf("%3s %s", "", "Masukkan username Anda : ")
	fmt.Scan(&username)
	fmt.Printf("%3s %s", "", "Masukkan password Anda : ")
	fmt.Scan(&password)
	fmt.Printf("%3s %s", "", "Masukkan PIN Anda (4 Digit) : ")
	fmt.Scan(&pin)

	if exist(*A,*T, username) {
		menu("Username yang ada masukan telah di gunakan, silahkan gunakan username lain.", "Buat Akun Baru Pasien", "", "", "", "", "Silahkan masukkan username dan password anda.", "PASIEN", true, A, *T)
		makeAcc(A, T)
	} else {
		A[*T].nama = nama
		A[*T].username = username
		A[*T].password = password
		A[*T].PIN = pin
		*T += 1
		menu("Akun berhasil dibuat, Silahkan login menggunakan akun anda!", "1. Login Akun Pasien", "2. Buat Akun Baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, *T)
		pasien(*A, T, true)
	}
}

func exist(A acc, T int, usr string) bool {
	for i := 0; i < T; i++ {
		if usr == A[i].username {
			return true
		}
	}
	return false
}

func recovery(A *acc, T int){
	var username string
	var passwordBaru string
	var indexUser, pin int
	fmt.Printf("%3s %s", "", "Masukkan username Anda : ")
	fmt.Scan(&username)
	if exist(*A, T, username) {
		indexUser = findIndexUser(*A, T, username)
		fmt.Printf("%3s %s", "", "Masukkan PIN baru Anda : ")
		fmt.Scan(&pin)
		
		if A[indexUser].PIN == pin {
			fmt.Printf("%3s %s", "", "Masukkan password baru Anda : ")
			fmt.Scan(&passwordBaru)
			A[indexUser].password = passwordBaru
			menu("Akun anda berhasil di Recovery, silahkan login dengan password baru.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, T)
			pasien(*A, &T, true)
		}else{
			menu("Anda memasukan PIN yang salah.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, T)
			pasien(*A, &T, true)
		}
	}else{
		menu("Akun anda tidak ditemukan, silahkan buat akun baru.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, T)
		pasien(*A, &T, true)
	}
}

func findIndexUser(A acc,  T int, usr string) int{
	for i := 0; i < T; i++ {
		if usr == A[i].username {
			return i
		}
	}
	return 0
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
	fmt.Println("  ____    ___     ___    ____      ____   __   __  _____   _")
	fmt.Println(" / ___|  / _ \\   / _ \\  |  _ \\    | __ )  \\ \\ / / | ____| | |")
	fmt.Println("| |  _  | | | | | | | | | | | |   |  _ \\   \\ V /  |  _|   | |")
	fmt.Println("| |_| | | |_| | | |_| | | |_| |   | |_) |   | |   | |___  |_|")
	fmt.Println(" \\____|  \\___/   \\___/  |____/    |____/    |_|   |_____| (_)")
}
