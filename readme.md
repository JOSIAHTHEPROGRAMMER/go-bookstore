# Go Bookstore

[![Go](https://img.shields.io/badge/Go-1.21-blue)](https://golang.org/)
[![License](https://img.shields.io/github/license/JOSIAHTHEPROGRAMMER/go-bookstore)](https://github.com/JOSIAHTHEPROGRAMMER/go-bookstore/blob/main/LICENSE)


---

## Overview

Go Bookstore is a RESTful API built with Go, Gorilla Mux, and GORM, designed to manage books, authors, publishers, and publications. It connects to an SQL Server database and provides endpoints for creating, updating, retrieving, and deleting records.  

---

## Features

- CRUD operations for Books, Authors, Publishers, and Publications
- Validates relationships between books, authors, and publishers before creating publications
- Prevents duplicates at the API level
- Search functionality for books and publications

---

## Tech Stack

- Go (Golang)
- Gorilla Mux (HTTP router)
- GORM (ORM for database)
- SQL Server (Database)

---

## Installation

1. **Clone the repository**
```bash
git clone https://github.com/JOSIAHTHEPROGRAMMER/go-bookstore.git
cd go-bookstore
```

2. **Set up environment variables**
Create a .env file at the root with the following content:
```ini
CONNECTION_STRING=your_sql_server_connection_string
```

3. **Run the server**
```bash
go run main.go
```

The server will start on http://localhost:8080.

## API Endpoints

### Books

- GET /books - Get all books

- GET /books/{id} - Get a book by ID

- POST /books - Create a book

- PUT /books/{id} - Update a book

- DELETE /books/{id} - Delete a book

- GET /books/search?title={title} - Search books by title

### Authors

- GET /authors - Get all authors

- GET /authors/{id} - Get an author by ID

- POST /authors - Create an author

- PUT /authors/{id} - Update an author

- DELETE /authors/{id} - Delete an author

### Publishers

- GET /publishers - Get all publishers

- GET /publishers/{id} - Get a publisher by ID

- POST /publishers - Create a publisher

- PUT /publishers/{id} - Update a publisher

- DELETE /publishers/{id} - Delete a publisher

### Publications

- GET /publications - Get all publications

- GET /publications/{id} - Get a publication by ID

- POST /publications - Create a publication (requires existing book, author, publisher)

- PUT /publications/{id} - Update a publication

- DELETE /publications/{id} - Delete a publication

- GET /publications/search?title={title}&author={author} - Search publications



## License

This project is licensed under the MIT License. See the LICENSE
 file for details.