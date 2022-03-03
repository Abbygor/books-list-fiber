package controllers

import (
	"log"
	"strconv"

	"github.com/Abbygor/books-list-fiber/models"
	"github.com/Abbygor/books-list-fiber/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BookController struct{}

var bookService services.BookService

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (bc BookController) GetBooks(c *fiber.Ctx) error {
	log.Println("Controller GetBooks-Fiber")

	books, err := bookService.GetBooks()

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(books)

}

func (bc BookController) GetBook(c *fiber.Ctx) error {
	log.Println("Controller GetBook-Fiber")

	apiErr := validateVariable(c.Params("book_id"))

	if apiErr != nil {
		log.Println(apiErr.Error())
		return c.JSON(apiErr.Error())
	}

	book_id, _ := strconv.Atoi(c.Params("book_id"))

	book, apiErr := bookService.GetBook(book_id)

	if apiErr != nil {
		return c.JSON(apiErr)
	}

	return c.JSON(book)
}

func (bc BookController) AddBook(c *fiber.Ctx) error {
	log.Println("Controller AddBook-Fiber")

	new_book := new(models.Book)

	err := c.BodyParser(new_book)

	if err != nil {
		return err
	}

	err = validateStruct(new_book)

	if err != nil {
		return err.(validator.ValidationErrors)
	}

	book, err := bookService.AddBook(new_book)

	if err != nil {
		return err
	}

	return c.JSON(book)
}

func (bc BookController) PutBook(c *fiber.Ctx) error {
	log.Println("Controller PutBook-Fiber")

	update_book := new(models.Book)

	err := c.BodyParser(update_book)

	if err != nil {
		return err
	}

	err = validateStruct(update_book)

	if err != nil {
		return err.(validator.ValidationErrors)
	}

	book, err := bookService.UpdateBook(update_book)

	if err != nil {
		return err
	}

	return c.JSON(book)
}

func (bc BookController) DeleteBook(c *fiber.Ctx) error {
	log.Println("Controller DeleteBook-Fiber")

	apiErr := validateVariable(c.Params("book_id"))

	if apiErr != nil {
		log.Println(apiErr.Error())
		return c.JSON(apiErr.Error())
	}

	book_id, _ := strconv.Atoi(c.Params("book_id"))

	book_deleted, apiErr := bookService.DeleteBook(book_id)

	if apiErr != nil {
		return c.JSON(apiErr)
	}

	return c.JSON(book_deleted)
}

func validateStruct(object interface{}) error {
	log.Println("validateStruct")
	err := validate.Struct(object)

	if err != nil {
		return err
	}
	return nil
}

func validateVariable(variable interface{}) error {
	log.Println("validateVariable")
	err := validate.Var(variable, "required,number")

	log.Println(err)
	if err != nil {
		return err
	}
	return nil
}
