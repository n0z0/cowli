package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/masterzen/winrm"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Host Target: ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)

	fmt.Print("Masukkan Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukkan Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Println("\n--- Menghubungkan ke", host, "---")

	endpoint := winrm.NewEndpoint(
		host,
		5985,
		false, // https
		true,  // insecure (tetap true untuk menghindari bug)
		nil,
		nil,
		nil,
		time.Second*60,
	)

	// Gunakan NewClient yang lebih sederhana
	client, err := winrm.NewClient(endpoint, username, password)
	if err != nil {
		panic(err)
	}

	// --- PERBAIKAN UTAMA ---
	// Bungkus perintah PowerShell di dalam pemanggilan powershell.exe
	// Perintah ini akan dieksekusi oleh shell default (cmd.exe), yang kemudian memanggil PowerShell
	psCommand := `dir`

	fmt.Printf("Menjalankan perintah: %s\n", psCommand)

	// Jalankan perintah dan tangkap outputnya
	stdout, stderr, exitCode, err := client.RunWithString(psCommand, "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// --- Tampilkan Hasil ---
	fmt.Printf("\nExit Code: %d\n", exitCode)
	fmt.Println("--- STDOUT ---")
	fmt.Println(stdout)
	if stderr != "" {
		fmt.Println("--- STDERR ---")
		fmt.Println(stderr)
	}
}
