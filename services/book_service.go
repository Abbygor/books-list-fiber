package services

import (
	"log"

	"github.com/Abbygor/books-list-fiber/models"
	"github.com/Abbygor/books-list-fiber/repositories"
)

type BookService struct{}

var bookRepository repositories.BookRepository

func (bs BookService) GetBooks() ([]models.Book, error) {
	log.Println("Service GetBooks-Fiber")

	return bookRepository.GetBooks()

}

func (bs BookService) GetBook(book_id int) (models.Book, error) {
	log.Printf("Service GetBook %v -Fiber", book_id)

	return bookRepository.GetBook(book_id)
}

func (bs BookService) AddBook(new_book *models.Book) (*models.Book, error) {
	log.Println("Service AddBook-Fiber")

	return bookRepository.AddBook(new_book)
}

func (bs BookService) UpdateBook(update_book *models.Book) (*models.Book, error) {
	log.Println("Service UpdateBook-Fiber")

	return bookRepository.UpdateBook(update_book)
}

func (bs BookService) DeleteBook(book_id int) (int, error) {
	log.Printf("Service DeleteBook %v -Fiber", book_id)

	return bookRepository.DeleteBook(book_id)

}
