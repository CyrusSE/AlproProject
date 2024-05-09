package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const NMAX int = 1000

type question struct {
	tag        string
	pertanyaan string
	tanggapan  string
	date       string
	author     string
}

type dataAccount struct {
	nama        string
	username    string
	password    string
	PIN         int
	jumlah_post int
}

type acc [NMAX]dataAccount
type que [NMAX]question

var pilihan int

func main() {
	var account acc
	var question que
	var totalAccount, totalPertanyaan int
	menu("Selamat datang di aplikasi konsultasi online.", "1. Pasien", "2. Dokter", "3. Liat Forum Konsultasi", "4. Exit", "", "Silahkan pilih keperluan anda.", "MAIN", false, &account, totalAccount, totalPertanyaan, &question)
	main_menu(&account, &totalAccount, &totalPertanyaan, &question)
	fmt.Print(account[0].nama)
}

func main_menu(akun *acc, T *int, P *int, question *que) {
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		menu("Selamat datang di menu pasien.", "1. Login Akun Pasien", "2. Buat Akun Baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, akun, *T, *P, question)
		pasien(*akun, T, P, *question)
	case 2:
		fmt.Println("SOON")
	case 3:
		fmt.Println("SOON")
	case 4:
		bye()
	case 9:
		load(akun, T, question)
		menu("Selamat datang di aplikasi konsultasi online.", "1. Pasien", "2. Dokter", "3. Liat Forum Konsultasi", "4. Exit", "", "Silahkan pilih keperluan anda.", "MAIN", false, akun, *T, *P, question)
		main_menu(akun, T, P, question)
	default:
		menu("Masukkan angka sesuai option.", "1. Pasien", "2. Dokter", "3. Liat Forum Konsultasi", "4. Exit", "", "Silahkan pilih keperluan anda.", "MAIN", false, akun, *T, *P, question)
		main_menu(akun, T, P, question)
	}
}

func pasien(akun acc, T *int, P *int, question que) {
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		menu("", "Login Akun Pasien", "", "", "", "", "Silahkan masukkan username dan password anda.", "PASIEN", true, &akun, *T, *P, &question)
		loginAccount(akun, *T, *P, question)
	case 2:
		menu("", "Buat Akun Baru Pasien.", "", "", "", "", "Silahkan masukkan username dan password anda.", "PASIEN", true, &akun, *T, *P, &question)
		makeAcc(&akun, T, P, &question)
	case 3:
		menu("", "Recovery Akun Pasien.", "", "", "", "", "Masukkan username dan PIN anda.", "PASIEN", true, &akun, *T, *P, &question)
		recovery(&akun, *T, *P, &question)
	case 4:
		main_menu(&akun, T, P, &question)
	default:
		menu("Masukkan angka sesuai option.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &akun, *T, *P, &question)
		pasien(akun, T, P, question)
	}
}

func menu(header, kata1, kata2, kata3, kata4, kata5, pesan, menuType string, hide bool, akun *acc, T int, P int, question *que) {
	clear()
	color("a")
	title("TUGAS BESAR ALGORITMA PEMPROGRAMAN")
	fmt.Printf("%25s╔═════════════════════════════════════════════════════════════════════════╗\n", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	fmt.Printf("%25s║%53s%-20s║\n", "", "Aplikasi Konsultasi Kesehatan", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	fmt.Printf("%25s║%55s %17s║\n", "", "Created by : • Faisal Ihsan Santoso", "")
	fmt.Printf("%25s║%59s %13s║\n", "", "• Arie Farchan Fyrzatullah", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	fmt.Printf("%25s╠═════════════════════════════════════════════════════════════════════════╣\n", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	if header != "" {
		fmt.Printf("%25s║  %-71s║\n", "", header)
		fmt.Printf("%25s║%73s║\n", "", "")
	}
	if kata1 != "" {
		fmt.Printf("%25s║%4s%-69s║\n", "", "", kata1)
	}
	if kata2 != "" {
		fmt.Printf("%25s║%4s%-69s║\n", "", "", kata2)
	}
	if kata3 != "" {
		fmt.Printf("%25s║%4s%-69s║\n", "", "", kata3)
	}
	if kata4 != "" {
		fmt.Printf("%25s║%4s%-69s║\n", "", "", kata4)
	}
	if kata5 != "" {
		fmt.Printf("%25s║%4s%-69s║\n", "", "", kata5)
	}
	fmt.Printf("%25s║%73s║\n", "", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	fmt.Printf("%25s║  %-71s║\n", "", pesan)
	fmt.Printf("%25s║%73s║\n", "", "")
	if menuType == "PASIEN" {
		fmt.Printf("%25s║%59s╔═════════════╣\n", "", "")
		fmt.Printf("%25s║%73s║\n", "", "║ MENU PASIEN ")
		fmt.Printf("%25s╚═══════════════════════════════════════════════════════════╩═════════════╝\n", "")
	} else if menuType == "DOKTER" {
		fmt.Printf("%25s║%56s╔════════════════╣\n", "", "")
		fmt.Printf("%25s║%73s║\n", "", "║ MENU DOKTER ")
		fmt.Printf("%25s╚════════════════════════════════════════════════════════╩════════════════╝\n", "")
	} else {
		fmt.Printf("%25s║%61s╔═══════════╣\n", "", "")
		fmt.Printf("%25s║%73s║\n", "", "║ MAIN MENU ")
		fmt.Printf("%25s╚═════════════════════════════════════════════════════════════╩═══════════╝\n", "")
	}
	if !hide {
		fmt.Printf("%25sSELECT> ", "")
	}
}

func main_pasien(A acc, T int, P int, username string, question que) {
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		menu("", "Posting keluhan anda yang ingin anda tanyakan pada dokter.", "", "", "", "", "Silahkan masukkan Judul Pertanyaan dan Pertanyaan anda.", "PASIEN", true, &A, T, P, &question)
		PostQuestion(&A, T, P, username, &question)
	case 2:
		list1 := "1. " + question[0].tag + "   " + question[0].date + "     " + question[0].author
		menu("Selamat datang di Forum Konsultasi.", list1, "", "", "", "", "Silahkan pilih pertanyaan yang ingin anda lihat.", "PASIEN", false, &A, T, P, &question)
	default:
		menu("Masukkan angka sesuai option.", "1. Konsultasi", "2. Liat Forum Konsultasi", "3. Logout", "", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &A, T, P, &question)
		fmt.Scanln(&pilihan)
	}
}

func PostQuestion(A *acc, T int, P int, username string, Q *que) {
	Index := findIndexUser(*A, T, username)
	var judul, pertanyaan string
	fmt.Printf("%24s %s", "", "Masukkan Judul Pertanyaan anda : ")
	fmt.Scanln(&judul)
	fmt.Printf("%24s %s", "", "Masukkan Pertanyaan Anda : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pertanyaan = scanner.Text()
	Q[P].tag = judul
	Q[P].pertanyaan = pertanyaan
	A[Index].jumlah_post += 1
	currentTime := time.Now()
	Q[P].date = string(currentTime.Format("2 January 2006 at 15:04"))
	Q[P].author = A[Index].nama
	menu("Pertanyaan anda berhasil di posting!", "1. Konsultasi", "2. Liat Forum Konsultasi", "3. Logout", "", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, T, P, Q)
	main_pasien(*A, T, P, username, *Q)
}

func loginAccount(A acc, T int, P int, question que) {
	var username, password string
	var indexUser int
	fmt.Printf("%24s %s", "", "Masukkan username Anda : ")
	fmt.Scanln(&username)
	fmt.Printf("%24s %s", "", "Masukkan password Anda : ")
	fmt.Scanln(&password)
	if exist(A, T, username) {
		indexUser = findIndexUser(A, T, username)
		if A[indexUser].username == username && A[indexUser].password == password {
			welcome_text := "Selamat datang " + A[indexUser].nama + ", di menu pasien."
			menu(welcome_text, "1. Konsultasi", "2. Liat Forum Konsultasi", "3. Logout", "", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &A, T, P, &question)
			main_pasien(A, T, P, username, question)
			//Login Berhasil
		} else {
			menu("Password anda salah, gunakan PIN untuk recovery akun anda.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &A, T, P, &question)
			pasien(A, &T, &P, question)
			//Login Gagal
		}
	} else {
		menu("Akun anda tidak ditemukan, silahkan buat akun baru.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, &A, T, P, &question)
		pasien(A, &T, &P, question)
		//Akun tidak ditemukan
	}
}

func makeAcc(A *acc, T *int, P *int, question *que) {
	var nama, username, password string
	var pin int
	fmt.Printf("%24s %s", "", "Masukkan nama Anda : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama = scanner.Text()
	fmt.Printf("%24s %s", "", "Masukkan username Anda : ")
	fmt.Scanln(&username)
	fmt.Printf("%24s %s", "", "Masukkan password Anda : ")
	fmt.Scanln(&password)
	fmt.Printf("%24s %s", "", "Masukkan PIN Anda (4 Digit) : ")
	fmt.Scanln(&pin)

	if exist(*A, *T, username) {
		menu("Username telah di gunakan, silahkan gunakan username lain.", "Buat Akun Baru Pasien", "", "", "", "", "Silahkan masukkan username dan password anda.", "PASIEN", true, A, *T, *P, question)
		makeAcc(A, T, P, question)
	} else {
		A[*T].nama = nama
		A[*T].username = username
		A[*T].password = password
		A[*T].PIN = pin
		*T += 1
		menu("Akun berhasil dibuat, Silahkan login menggunakan akun anda!", "1. Login Akun Pasien", "2. Buat Akun Baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, *T, *P, question)
		pasien(*A, T, P, *question)
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

func recovery(A *acc, T int, P int, question *que) {
	var username string
	var passwordBaru string
	var indexUser, pin int
	fmt.Printf("%24s %s", "", "Masukkan username Anda : ")
	fmt.Scanln(&username)
	if exist(*A, T, username) {
		indexUser = findIndexUser(*A, T, username)
		fmt.Printf("%24s %s", "", "Masukkan PIN Anda : ")
		fmt.Scanln(&pin)
		if A[indexUser].PIN == pin {
			fmt.Printf("%24s %s", "", "Masukkan password baru Anda : ")
			fmt.Scanln(&passwordBaru)
			A[indexUser].password = passwordBaru
			menu("Akun anda berhasil di Recovery, silahkan login dengan password baru.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, T, P, question)
			pasien(*A, &T, &P, *question)
		} else {
			menu("Anda memasukan PIN yang salah.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, T, P, question)
			pasien(*A, &T, &P, *question)
		}
	} else {
		menu("Akun anda tidak ditemukan, silahkan buat akun baru.", "1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back", "", "Silahkan pilih keperluan anda.", "PASIEN", false, A, T, P, question)
		pasien(*A, &T, &P, *question)
	}
}

func findIndexUser(A acc, T int, usr string) int {
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

func load(A *acc, T *int, Q *que) {
	A[0].nama = "FAISAL IHSAN SANTOSO"
	A[0].username = "faisal"
	A[0].password = "123"
	A[0].PIN = 1234
	*T += 1
}
