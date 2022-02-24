package app

import "github.com/Abbygor/books-list-fiber/controllers"

var bookController controllers.BookController

func MapUrls() {
	app.Get("/books", bookController.GetBooks)
	app.Get("/books/:book_id", bookController.GetBook)
	app.Post("/books", bookController.AddBook)
	app.Put("/books", bookController.PutBook)
	app.Delete("/books/:book_id", bookController.DeleteBook)
}
