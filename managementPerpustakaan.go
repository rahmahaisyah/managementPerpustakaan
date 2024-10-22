package main

import (
    "fmt"
)
type Book struct {
    ID    int
    Title string
    IsBorrowed bool
}

// Member represents a library member
type Member struct {
    ID   int
    Name string
}

// Global data
var books []Book
var members []Member

// AddBook menambahkan buku baru ke perpustakaan
func AddBook(id int, title string) {
    book := Book{ID: id, Title: title, IsBorrowed: false}
    books = append(books, book)
    fmt.Println("Buku berhasil ditambahkan:", book)
}

// AddMember menambahkan anggota baru
func AddMember(id int, name string) {
    member := Member{ID: id, Name: name}
    members = append(members, member)
    fmt.Println("Anggota berhasil ditambahkan:", member)
}

// ListBooks menampilkan semua buku
func ListBooks() {
    fmt.Println("Daftar Buku:")
    for _, book := range books {
        status := "Tersedia"
        if book.IsBorrowed {
            status = "Dipinjam"
        }
        fmt.Printf("ID: %d, Judul: %s, Status: %s\n", book.ID, book.Title, status)
    }
}

// ListMembers menampilkan semua anggota
func ListMembers() {
    fmt.Println("Daftar Anggota:")
    for _, member := range members {
        fmt.Printf("ID: %d, Nama: %s\n", member.ID, member.Name)
    }
}

// BorrowBook meminjam buku berdasarkan ID
func BorrowBook(bookID int) {
    for i, book := range books {
        if book.ID == bookID {
            if !book.IsBorrowed {
                books[i].IsBorrowed = true
                fmt.Printf("Buku '%s' berhasil dipinjam\n", book.Title)
            } else {
                fmt.Println("Buku sudah dipinjam!")
            }
            return
        }
    }
    fmt.Println("Buku tidak ditemukan")
}

// ReturnBook mengembalikan buku berdasarkan ID
func ReturnBook(bookID int) {
    for i, book := range books {
        if book.ID == bookID {
            if book.IsBorrowed {
                books[i].IsBorrowed = false
                fmt.Printf("Buku '%s' berhasil dikembalikan\n", book.Title)
            } else {
                fmt.Println("Buku belum dipinjam!")
            }
            return
        }
    }
    fmt.Println("Buku tidak ditemukan")
}
func main() {
    // Menambahkan beberapa data contoh
    AddBook(1, "Pemrograman Go")
    AddBook(2, "Belajar Algoritma")
    AddMember(1, "Alice")
    AddMember(2, "Bob")

    // Menu sederhana
    for {
        fmt.Println("\n--- Manajemen Perpustakaan ---")
        fmt.Println("1. Daftar Buku")
        fmt.Println("2. Daftar Anggota")
        fmt.Println("3. Tambah Buku")
        fmt.Println("4. Tambah Anggota")
        fmt.Println("5. Pinjam Buku")
        fmt.Println("6. Kembalikan Buku")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih opsi: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            ListBooks()
        case 2:
            ListMembers()
        case 3:
            var id int
            var title string
            fmt.Print("Masukkan ID Buku: ")
            fmt.Scan(&id)
            fmt.Print("Masukkan Judul Buku: ")
            fmt.Scan(&title)
            AddBook(id, title)
        case 4:
            var id int
            var name string
            fmt.Print("Masukkan ID Anggota: ")
            fmt.Scan(&id)
            fmt.Print("Masukkan Nama Anggota: ")
            fmt.Scan(&name)
            AddMember(id, name)
        case 5:
            var id int
            fmt.Print("Masukkan ID Buku yang ingin dipinjam: ")
            fmt.Scan(&id)
            BorrowBook(id)
        case 6:
            var id int
            fmt.Print("Masukkan ID Buku yang ingin dikembalikan: ")
            fmt.Scan(&id)
            ReturnBook(id)
        case 0:
            fmt.Println("Keluar dari sistem...")
            return
        default:
            fmt.Println("Opsi tidak valid, coba lagi!")
        }
    }
}