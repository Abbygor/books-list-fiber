package repositories

import (
	"log"

	"github.com/Abbygor/books-list-fiber/config"
	"github.com/Abbygor/books-list-fiber/models"
	"github.com/subosito/gotenv"
)

type BookRepository struct{}

func init() {
	gotenv.Load()
}

func (br BookRepository) GetBooks() ([]models.Book, error) {
	log.Println("Repository GetBooks-Fiber")

	db := config.ConnectDB()

	books := []models.Book{}
	book := models.Book{}

	rows, err := db.Query("select * from books")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	return books, err

}

func (br BookRepository) GetBook(book_id int) (models.Book, error) {
	log.Printf("Repository GetBook %v -Fiber", book_id)

	db := config.ConnectDB()

	book := models.Book{}

	row := db.QueryRow("select * from books where id = $1", book_id)

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	return book, err
}

func (br BookRepository) AddBook(new_book *models.Book) (*models.Book, error) {
	log.Println("Repository AddBook-Fiber")

	db := config.ConnectDB()

	err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id;",
		&new_book.Title, &new_book.Author, &new_book.Year).Scan(&new_book.ID)

	return new_book, err

}

func (br BookRepository) UpdateBook(update_book *models.Book) (*models.Book, error) {
	log.Println("Repository UpdateBook-Fiber")

	db := config.ConnectDB()

	result, err := db.Exec("update books set title = $1, author = $2, year = $3 where id = $4 RETURNING id;",
		&update_book.Title, &update_book.Author, &update_book.Year, &update_book.ID)

	if err != nil {
		return new(models.Book), err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil || rowsUpdated == 0 {
		return new(models.Book), err
	}

	return update_book, nil
}

func (br BookRepository) DeleteBook(book_id int) (int, error) {
	log.Println("Repository DeleteBook-Fiber")

	db := config.ConnectDB()

	result, err := db.Exec("delete from books where id = $1;", book_id)

	if err != nil {
		return 0, err
	}

	rowsDel, err := result.RowsAffected()

	if err != nil || rowsDel == 0 {
		return 0, err
	}

	return int(rowsDel), nil
}
