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
	// Buat reader untuk membaca input dari keyboard
	reader := bufio.NewReader(os.Stdin)

	// --- Minta Input dari User ---
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

	// --- Konfigurasi Koneksi ---
	endpoint := winrm.NewEndpoint(
		host,           // Gunakan host dari input user
		5985,           // Port HTTP untuk WinRM
		false,          // Tidak menggunakan HTTPS
		true,           // Tetap true untuk menghindari bug
		nil,            // caCert ([]byte)
		nil,            // cert ([]byte)
		nil,            // key ([]byte)
		time.Second*60, // timeout (time.Duration)
	)

	// --- Kredensial ---
	// PERUBAHAN: Gunakan NewClient yang lebih sederhana
	// Kita tidak lagi menggunakan 'params' dan 'TransportDecorator'
	client, err := winrm.NewClient(
		endpoint,
		username, // Gunakan username dari input user
		password, // Gunakan password dari input user
	)
	if err != nil {
		panic(err)
	}

	// --- Eksekusi Perintah PowerShell ---
	psCommand := "Get-Process | Select-Object -First 5 | ConvertTo-Json"

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
