package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Ahora la UI trabaja con una INTERFAZ (Unidad 3).
	var lib LibraryManager = NewLibrary()
	seedDemoData(lib)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n==============================")
		fmt.Println("      EBOOKS SYSTEM (Go)")
		fmt.Println("==============================")
		fmt.Println("1) Listar libros")
		fmt.Println("2) Agregar libro")
		fmt.Println("3) Eliminar libro")
		fmt.Println("4) Buscar libro por título")
		fmt.Println("5) Listar usuarios")
		fmt.Println("6) Agregar usuario")
		fmt.Println("7) Eliminar usuario")
		fmt.Println("8) Prestar libro")
		fmt.Println("9) Devolver libro")
		fmt.Println("0) Salir")
		fmt.Print("Elige una opción: ")

		opt := readLine(reader)

		switch opt {
		case "1":
			listBooks(lib)
		case "2":
			addBook(lib, reader)
		case "3":
			removeBook(lib, reader)
		case "4":
			searchBook(lib, reader)
		case "5":
			listUsers(lib)
		case "6":
			addUser(lib, reader)
		case "7":
			removeUser(lib, reader)
		case "8":
			borrowBook(lib, reader)
		case "9":
			returnBook(lib, reader)
		case "0":
			fmt.Println("Saliendo... ")
			return
		default:
			fmt.Println("Opción inválida ")
		}
	}
}

// ------------------ Helpers UI ------------------

func readLine(r *bufio.Reader) string {
	text, _ := r.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt(r *bufio.Reader) int {
	for {
		txt := readLine(r)
		n, err := strconv.Atoi(txt)
		if err == nil {
			return n
		}
		fmt.Print("Ingresa un número válido: ")
	}
}

// ------------------ Menú actions ------------------

func listBooks(lib LibraryManager) {
	books := lib.ListBooks()
	if len(books) == 0 {
		fmt.Println("No hay libros registrados.")
		return
	}
	fmt.Println("\n--- LIBROS ---")
	for _, b := range books {
		fmt.Println(b.String())
	}
}

func addBook(lib LibraryManager, r *bufio.Reader) {
	fmt.Println("\n--- AGREGAR LIBRO ---")
	fmt.Print("ID: ")
	id := readLine(r)

	fmt.Print("Título: ")
	title := readLine(r)

	fmt.Print("Autor: ")
	author := readLine(r)

	fmt.Print("Año: ")
	year := readInt(r)

	fmt.Print("Género: ")
	genre := readLine(r)

	b, err := NewBook(id, title, author, year, genre)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := lib.AddBook(b); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Libro agregado ✅")
}

func removeBook(lib LibraryManager, r *bufio.Reader) {
	fmt.Println("\n--- ELIMINAR LIBRO ---")
	fmt.Print("ID del libro: ")
	id := readLine(r)

	if err := lib.RemoveBook(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Libro eliminado ✅")
}

func searchBook(lib LibraryManager, r *bufio.Reader) {
	fmt.Println("\n--- BUSCAR LIBRO ---")
	fmt.Print("Buscar por título: ")
	q := readLine(r)

	res := lib.FindBooksByTitle(q)
	if len(res) == 0 {
		fmt.Println("No se encontraron libros.")
		return
	}
	for _, b := range res {
		fmt.Println(b.String())
	}
}

func listUsers(lib LibraryManager) {
	users := lib.ListUsers()
	if len(users) == 0 {
		fmt.Println("No hay usuarios registrados.")
		return
	}
	fmt.Println("\n--- USUARIOS ---")
	for _, u := range users {
		fmt.Println(u.String())
	}
}

func addUser(lib LibraryManager, r *bufio.Reader) {
	fmt.Println("\n--- AGREGAR USUARIO ---")
	fmt.Print("ID: ")
	id := readLine(r)

	fmt.Print("Nombre: ")
	name := readLine(r)

	fmt.Print("Email: ")
	email := readLine(r)

	u, err := NewUser(id, name, email)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := lib.AddUser(u); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Usuario agregado ✅")
}

func removeUser(lib LibraryManager, r *bufio.Reader) {
	fmt.Println("\n--- ELIMINAR USUARIO ---")
	fmt.Print("ID del usuario: ")
	id := readLine(r)

	if err := lib.RemoveUser(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Usuario eliminado ✅")
}

func borrowBook(lib LibraryManager, r *bufio.Reader) {
	fmt.Println("\n--- PRESTAR LIBRO ---")
	fmt.Print("ID del libro: ")
	bookID := readLine(r)

	fmt.Print("ID del usuario: ")
	userID := readLine(r)

	if err := lib.BorrowBook(bookID, userID); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Libro prestado ✅")
}

func returnBook(lib LibraryManager, r *bufio.Reader) {
	fmt.Println("\n--- DEVOLVER LIBRO ---")
	fmt.Print("ID del libro: ")
	bookID := readLine(r)

	if err := lib.ReturnBook(bookID); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Libro devuelto ✅")
}

// ------------------ Demo data ------------------

func seedDemoData(lib LibraryManager) {
	u1, _ := NewUser("U001", "Carlos", "carlos@example.com")
	u2, _ := NewUser("U002", "Ana", "ana@example.com")
	_ = lib.AddUser(u1)
	_ = lib.AddUser(u2)

	b1, _ := NewBook("B001", "Clean Code", "Robert C. Martin", 2008, "Software")
	b2, _ := NewBook("B002", "The Go Programming Language", "Alan A. A. Donovan", 2015, "Programación")
	_ = lib.AddBook(b1)
	_ = lib.AddBook(b2)
}
