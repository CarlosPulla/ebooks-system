package main

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	Year       int
	Genre      string
	Available  bool
	BorrowedBy string // userID si está prestado, "" si no
}

func (b Book) String() string {
	status := "Disponible ✅"
	if !b.Available {
		status = fmt.Sprintf("Prestado ❌ (Usuario: %s)", b.BorrowedBy)
	}
	return fmt.Sprintf("[%s] %s - %s (%d) | Género: %s | %s",
		b.ID, b.Title, b.Author, b.Year, b.Genre, status)
}
