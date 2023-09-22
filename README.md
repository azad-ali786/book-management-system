# Book Management System Backend

Welcome to the Book Management System backend documentation. This README provides an overview of the backend code for managing books using the Go programming language. This system is designed to handle various operations related to books, including creating, retrieving, updating, and deleting book records.

## Table of Contents
- [Introduction](#introduction)
- [Routes](#routes)
- [Installation](#installation)

## Introduction

The Book Management System backend is built using Go and leverages the Gorilla Mux router and MYSQL to define and handle API routes. It provides a set of routes and controllers for managing books. These routes allow users to perform actions such as creating new books, retrieving book details, updating book information, and deleting books.

## Routes

The backend defines the following routes:

- `POST /book`: Create a new book.
- `GET /book`: Retrieve a list of all books.
- `GET /book/{bookId}`: Retrieve a specific book by its unique identifier.
- `PUT /book/{bookId}`: Update the information of a specific book.
- `DELETE /book/{bookId}`: Delete a specific book by its unique identifier.

## Installation

Before running the Book Management System backend, ensure you have Go installed on your system. You can follow the official Go installation guide at [https://golang.org/doc/install](https://golang.org/doc/install) if you haven't already.

To install and run the backend:

1. Clone the repository:

   ```bash
   git clone https://github.com/azad-ali786/bookmanagementsystem.git
   cd bookmanagementsystem
   ```
2. Install the required dependencies using go get:
   eg.
   ```bash
   go get github.com/gorilla/mux
   ```
3. Build and run the application
   ```bash
   go run main.go
   ```
The backend will be accessible at http://localhost:8080.



