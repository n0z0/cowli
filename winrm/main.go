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

	// --- Konfigurasi Awal (Hanya Sekali) ---
	fmt.Print("Masukkan Host Target: ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)

	fmt.Print("Masukkan Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukkan Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Buat koneksi endpoint
	endpoint := winrm.NewEndpoint(
		host,
		5985,
		false, // https
		true,  // insecure
		nil,
		nil,
		nil,
		time.Second*60,
	)

	// Buat klien WinRM
	client, err := winrm.NewClient(endpoint, username, password)
	if err != nil {
		panic("Gagal membuat klien WinRM: " + err.Error())
	}

	fmt.Printf("\n--- Berhasil terhubung ke %s ---\n", host)
	fmt.Println("--- Ketik 'exit' untuk keluar ---\n")

	// --- Loop Interaktif untuk Perintah ---
	for {
		fmt.Print("winrm> ") // Prompt seperti di CMD

		// Baca perintah dari user
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		// Cek jika user ingin keluar
		if strings.ToLower(command) == "exit" {
			fmt.Println("Keluar dari program.")
			break
		}

		// Jika perintah kosong, lanjut ke iterasi berikutnya
		if command == "" {
			continue
		}

		// Jalankan perintah di server
		fmt.Printf("[Menjalankan: %s]\n", command)
		stdout, stderr, exitCode, err := client.RunWithString(command, "")

		// Tampilkan hasil
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Exit Code: %d\n", exitCode)
			if stdout != "" {
				fmt.Println("--- STDOUT ---")
				fmt.Println(stdout)
			}
			if stderr != "" {
				fmt.Println("--- STDERR ---")
				fmt.Println(stderr)
			}
		}
		fmt.Println("---------------------------------") // Pemisah
	}
}
