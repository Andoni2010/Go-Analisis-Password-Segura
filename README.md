# **Proyecto: Análisis de Contraseñas Inseguras en Go**

Este proyecto tiene como objetivo analizar un archivo con contraseñas y compararlas con una lista de contraseñas comunes para identificar cuáles son inseguras. Está diseñado como una introducción práctica al manejo de archivos, mapas y estructuras comunes en el lenguaje **Go**.

---

## **Descripción del Proyecto**

El programa:

- Lee un archivo con contraseñas comunes (`top_password.txt`) y lo almacena en un mapa para búsquedas rápidas.
- Lee un archivo con contraseñas a analizar (`password.txt`).
- Compara cada contraseña del archivo con las contraseñas comunes para determinar si es insegura.
- Imprime en consola las contraseñas inseguras detectadas.

### **Este proyecto es útil para:**
- **Ciberseguridad**: Detectar contraseñas débiles o comprometidas.
- **Educación**: Aprender manejo de archivos, mapas y procesamiento de datos en Go.

---

## **Requisitos**

- **Go instalado**:  
  Asegúrate de tener una versión reciente instalada. Descárgala desde [golang.org](https://golang.org/).
  
- **Archivos necesarios**:  
  - `password.txt`: Archivo con las contraseñas a analizar.  
  - `top_password.txt`: Archivo con contraseñas comunes.

---

## **Archivos Necesarios**

### **Archivo de contraseñas a analizar (`password.txt`)**
Ejemplo de contenido:


```bash
password123
123456
mypassword
qwerty
supersecure
```

### **Archivo de contraseñas comunes (top_password.txt)**
Ejemplo de contenido:

```bash
123456
password
qwerty
abc123
admin
```
---
# Pasos para Ejecutar
### 1. Clonar el repositorio
Clona este proyecto en tu máquina local:

```bash
git clone https://github.com/usuario/proyecto-go-password-analyzer.git
```

###  2. Colocar los archivos necesarios
Asegúrate de que los archivos password.txt y top_password.txt estén en el mismo directorio que el archivo de código.

###  3. Ejecutar el programa
Compila y ejecuta el programa:

```bash
go run main.go
```
###  4. Revisar los resultados
El programa imprimirá las contraseñas inseguras en la consola. Ejemplo de salida:

```bash
Contraseñas inseguras encontradas:
123456
qwerty
```
---
## **Desglose del Código**
###  1. main
Esta función principal orquesta todo el programa:

- Declara los nombres de los archivos (password.txt y top_password.txt).
- Llama a loadCommonPassword para cargar las contraseñas comunes en un mapa.
- Llama a analyzePasswords para comparar las contraseñas del archivo de entrada con las contraseñas comunes.
```go
func main() {
	passwordFile := "password.txt"
	commonPasswordFile := "top_password.txt"

	commonPassword := loadCommonPassword(commonPasswordFile)
	analyzePasswords(passwordFile, commonPassword)
}
```
###  2. loadCommonPassword
Propósito: Leer un archivo con contraseñas comunes y almacenarlas en un mapa (map[string]bool) para búsquedas rápidas.
Cómo funciona:
- Abre el archivo con os.Open.
- Usa bufio.Scanner para leer línea por línea.
- Limpia cada línea con strings.TrimSpace.
- Almacena cada contraseña en un mapa donde la clave es la contraseña y el valor es true.
```go
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

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo de las contraseñas comunes:", err)
	}

	return commonPassword
}
```
###  3. analyzePasswords
Propósito: Comparar las contraseñas del archivo de entrada con el mapa de contraseñas comunes.
Cómo funciona:
- Abre el archivo con os.Open.
- Usa bufio.Scanner para leer línea por línea.
- Verifica si cada contraseña está en el mapa usando la función IsInsecurePassword.
- Almacena las contraseñas inseguras en un slice ([]string) y las imprime en consola.
```go
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

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo de las contraseñas a analizar:", err)
	}

	fmt.Println("Contraseñas inseguras encontradas:")
	for _, p := range insecurePasswords {
		fmt.Println(p)
	}
}
```
###  4. IsInsecurePassword
Propósito: Verificar si una contraseña está en el mapa de contraseñas comunes.
Cómo funciona:
- Toma una contraseña y verifica si está presente en el mapa.
- Retorna true si la contraseña es insegura, de lo contrario, retorna false.
```go
func IsInsecurePassword(password string, commonPassword map[string]bool) bool {
	return commonPassword[password]
}
```
---
## **Posibles Extensiones**
- Leer los nombres de los archivos (password.txt y top_password.txt) desde argumentos de línea de comandos.
- Exportar las contraseñas inseguras a un archivo CSV o JSON.
- Permitir analizar archivos con miles de contraseñas.
