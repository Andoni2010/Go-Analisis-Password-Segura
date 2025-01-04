package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	passwordFile := "password.txt"
	commonPasswordFile := "top_password.txt"

	commonPassword := loadCommonPassword(commonPasswordFile)
	analyzePasswords(passwordFile, commonPassword)
}

func loadCommonPassword(filePath string) map[string]bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo de las contraseñas comunes:", err)
		os.Exit(1)
	}
	defer file.Close()

	commonPassword := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commonPassword[strings.TrimSpace(scanner.Text())] = true
	}

	if err := scanner.Err(); err != nil { // Corregido el operador de asignación ":="
		fmt.Println("Error al leer el archivo de las contraseñas comunes:", err)
	}

	return commonPassword
}

func analyzePasswords(filePath string, commonPassword map[string]bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo de las contraseñas a analizar:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var insecurePasswords []string

	for scanner.Scan() {
		password := scanner.Text()
		if IsInsecurePassword(password, commonPassword) {
			insecurePasswords = append(insecurePasswords, password)
		}
	}

	if err := scanner.Err(); err != nil { // Corregido el operador de asignación ":="
		fmt.Println("Error al leer el archivo de las contraseñas a analizar:", err)
	}

	fmt.Println("Contraseñas inseguras encontradas:")
	for _, p := range insecurePasswords {
		fmt.Println(p)
	}
}

// Función IsInsecurePassword verifica si una contraseña está en el mapa de contraseñas comunes
func IsInsecurePassword(password string, commonPassword map[string]bool) bool {
	return commonPassword[password]
}
