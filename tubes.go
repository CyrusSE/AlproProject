package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/common-nighthawk/go-figure"
)

var pilihan int

func main() {
	main_menu()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 3:
		bye()
	default:
		main()
	}
}

func main_menu() {
	clear()
	color("a")
	title("TUGAS BESAR ALGORITMA PEMPROGRAMAN")
	fmt.Println("    ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("    ║%73s║\n", "")
	fmt.Printf("    ║%48s %24s║\n", "Aplikasi TEL-U Library", "")
	fmt.Printf("    ║%73s║\n", "")
	fmt.Printf("    ║%55s %17s║\n", "Created by : • Faisal Ihsan Santoso", "")
	fmt.Printf("    ║%59s %13s║\n", "• Arie Farchan Fyrzatullah", "")
	fmt.Printf("    ║%73s║\n", "")
	fmt.Println("    ╠═════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("    ║%73s║\n", "")
	fmt.Printf("    ║%16s %56s║\n", "1. Mahasiswa", "")
	fmt.Printf("    ║%12s %60s║\n", "2. Admin", "")
	fmt.Printf("    ║%11s %61s║\n", "3. Exit", "")
	fmt.Printf("    ║%73s║\n", "")
	fmt.Printf("    ║%73s║\n", "")
	fmt.Printf("    ║%28s %44s║\n", "Silahkan pilih peran anda.", "")
	fmt.Printf("    ║%73s║\n", "")
	fmt.Printf("    ║%61s╔═══════════╣\n", "")
	fmt.Printf("    ║%72s %s║\n", "║ MAIN MENU", "")
	fmt.Println("    ╚═════════════════════════════════════════════════════════════╩═══════════╝")
	fmt.Print("    SELECT> ")
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
