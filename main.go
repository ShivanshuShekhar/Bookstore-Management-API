package main

import (
	"encoding/json"
	"sync"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/log"
	"gofr.dev/pkg/service"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookstoreService struct {
	mu    sync.Mutex
	books map[string]Book
}

func NewBookstoreService() *BookstoreService {
	return &BookstoreService{
		books: make(map[string]Book),
	}
}

func (bs *BookstoreService) CreateBook(ctx *gofr.Context) (interface{}, error) {
	var book Book
	if err := json.NewDecoder(ctx.Request.Body).Decode(&book); err != nil {
		return nil, err
	}

	bs.mu.Lock()
	defer bs.mu.Unlock()

	book.ID = generateID()
	bs.books[book.ID] = book

	return book, nil
}

func (bs *BookstoreService) GetBook(ctx *gofr.Context) (interface{}, error) {
	bookID := ctx.Param("id")

	bs.mu.Lock()
	defer bs.mu.Unlock()

	book, found := bs.books[bookID]
	if !found {
		return nil, service.EntityNotFound{Entity: "book", ID: bookID}
	}

	return book, nil
}

func (bs *BookstoreService) ListBooks(ctx *gofr.Context) (interface{}, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	var books []Book
	for _, book := range bs.books {
		books = append(books, book)
	}

	return books, nil
}

func (bs *BookstoreService) UpdateBook(ctx *gofr.Context) (interface{}, error) {
	bookID := ctx.Param("id")

	var updatedBook Book
	if err := json.NewDecoder(ctx.Request.Body).Decode(&updatedBook); err != nil {
		return nil, err
	}

	bs.mu.Lock()
	defer bs.mu.Unlock()

	_, found := bs.books[bookID]
	if !found {
		return nil, service.EntityNotFound{Entity: "book", ID: bookID}
	}

	bs.books[bookID] = updatedBook

	return updatedBook, nil
}

func (bs *BookstoreService) DeleteBook(ctx *gofr.Context) (interface{}, error) {
	bookID := ctx.Param("id")

	bs.mu.Lock()
	defer bs.mu.Unlock()

	_, found := bs.books[bookID]
	if !found {
		return nil, service.EntityNotFound{Entity: "book", ID: bookID}
	}

	delete(bs.books, bookID)

	return nil, nil
}

func generateID() string {
	return "book_" + gofr.GenerateUUID()
}

func main() {

	k := gofr.New()

	bookstoreService := NewBookstoreService()

	k.POST("/books", bookstoreService.CreateBook)
	k.GET("/books/:id", bookstoreService.GetBook)
	k.GET("/books", bookstoreService.ListBooks)
	k.PUT("/books/:id", bookstoreService.UpdateBook)
	k.DELETE("/books/:id", bookstoreService.DeleteBook)

	k.SetLogLevel(log.LevelInfo)

	k.Start()
}
