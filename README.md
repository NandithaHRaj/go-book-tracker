📚 Go Book Tracker

A simple CRUD backend application built with Go to track books, reading status, and ratings.

This project is part of my journey to strengthen Go fundamentals before moving into gRPC and distributed systems.

🚀 Features

- Add a new book
- List all books
- Get a book by ID
- Update book details
- Delete a book

🛠 Tech Stack

- Go (net/http)
- In-memory storage (Phase 1)
- PostgreSQL (Phase 2 – planned)

📁 Project Structure
go-book-tracker/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── model/
│   ├── repository/
│   └── handler/
│
├── go.mod