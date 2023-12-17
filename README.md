# Bookstore Management API using GoFr Framework

The "Bookstore Management API" is a simple yet powerful application designed to handle the basic operations of a bookstore, such as creating, retrieving, listing, updating, and deleting books. The API is built using the Go programming language and leverages the GoFr framework, which simplifies the development of RESTful APIs.

Language/Framework: Go (Golang), GoFr Framework
Functionality: CRUD operations for managing books in a bookstore.
Key Components: API Server, BookstoreService, User Interface (HTTP requests)
Data Storage: In-memory map

### Install required packages:
go get -u gofr.dev/pkg/gofr
go get -u gofr.dev/pkg/log
go get -u gofr.dev/pkg/service

### Run the project:
go run main.go

### Test the API:
Use Postman to interact with API.
Create a new book using a POST request to http://localhost:9000/books with a JSON body like:
{
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald"
}
