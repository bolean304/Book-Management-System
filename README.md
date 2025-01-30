# Book Management System (Backend)

## Overview
This is a simple backend system for a "Book Management" application built using Golang and the Gin framework. The application provides user authentication, book management, and borrowing/returning functionality while ensuring security best practices such as password hashing and JWT-based authentication.

## Features
### **User Management**
- User registration with username and password (passwords are securely hashed using bcrypt).
- JWT-based authentication for secure API access.

### **Book Management**
- Add a new book (Title, Author, Genre, Published Year).
- View all books .
- Search for books by title or author.
- Update book details.
- Delete a book.

### **Borrow/Return Books**
- Borrow multiple books by their IDs.
- Return borrowed books.

## Tech Stack
- **Language**: Golang
- **Framework**: Gin (for REST API handling)
- **Database**: MySQL (with GORM ORM)
- **Authentication**: JWT (JSON Web Token)
- **Password Security**: bcrypt hashing
- **Environment Management**: godotenv

## Setup and Installation
### **Prerequisites**
- Go installed (version 1.18 or later)
- PostgreSQL installed and running
- Git installed

### **Clone the Repository**
```sh
$ git clone https://github.com/yourusername/book-management-system.git
$ cd book-management-system
```

### **Set Up Environment Variables**
Create a `.env` file in the project root and configure the following:
```env
DB_USERNAME=root
DB_PASSWORD=root123
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=test_db
SECRET_KEY="my-secret-key"
PORT=8080

```

### **Install Dependencies**
```sh
$ go mod tidy
```



### **Run the Application**
```sh
$ go run main.go
```
The server will start on `http://localhost:8080`

## API Endpoints
### **User Management**
| Method | Endpoint         | Description |
|--------|-----------------|-------------|
| POST   | `/register`      | Register a new user |
| POST   | `/login`         | User login and JWT generation |

### **Book Management**
| Method | Endpoint       | Description |
|--------|---------------|-------------|
| POST   | `/add-book`       | Add a new book |
| GET    | `/fetch-books`       | Get all books  |
| GET    | `/books/search?query=title_or_author` | Search for books by title or author |
| PUT    | `/books/:id`   | Update book details |
| DELETE | `/books/:id`   | Delete a book |

### **Borrow/Return Books**
| Method | Endpoint       | Description |
|--------|---------------|-------------|
| POST   | `/book-borrow`      | Borrow books |
| GET    | `/etch-borrowed-books/:user_id` | Get all books borrowed by a user |

## Security and Cryptography
- **Password Hashing**: User passwords are hashed using `bcrypt` before storage.
- **JWT Authentication**: API access is secured using JWT tokens.
- **Input Validation & SQL Injection Prevention**: Using Gin’s request binding and GORM to prevent raw SQL injections.
- **Data Encryption Utility**: Basic encryption and decryption for user notes using AES.


## Author
Developed by Ankit Chauhan.


