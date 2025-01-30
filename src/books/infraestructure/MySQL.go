package infraestructure

import (
	"database/sql"
	"demo/src/books/domain"
	"fmt"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{
		db: db,
	}
}

func (m *MySQL) ValidateAuthor(authorId int32) (bool, error) {
	var exists bool
	err := m.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", authorId).Scan(&exists)
	if (err != nil) {
		return false, err
	}
	return exists, nil
}

func (m *MySQL) Save(book domain.Book) (domain.Book, error) {
	// Primero validamos que exista el autor
	exists, err := m.ValidateAuthor(book.Author)
	if err != nil {
		return domain.Book{}, err
	}
	if !exists {
		return domain.Book{}, fmt.Errorf("el autor con ID %d no existe", book.Author)
	}

	// Procedemos con el guardado
	result, err := m.db.Exec("INSERT INTO books (title, author, description) VALUES (?, ?, ?)", book.Title, book.Author, book.Description)
	if err != nil {
		return domain.Book{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Book{}, err
	}
	book.Id = int32(id)
	book = domain.Book{Id: int32(id), Title: book.Title, Author: book.Author, Description: book.Description}
	return book, nil
}

func (m *MySQL) GetAll() ([]domain.Book, error) {
	rows, err := m.db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Description)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books, nil
}

func (m *MySQL) GetByID(id int32) (domain.Book, error) {
	row := m.db.QueryRow("SELECT * FROM books WHERE id = ?", id)
	var book domain.Book
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Description)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (m *MySQL) Update(id int32, book domain.Book) (domain.Book, error) {
	_, err := m.db.Exec("UPDATE books SET title = ?, author = ?, description = ? WHERE id = ?", book.Title, book.Author, book.Description, id)
	if err != nil {
		return domain.Book{}, err
	}
	book = domain.Book{Id: id, Title: book.Title, Author: book.Author, Description: book.Description}
	return book, nil
}

func (m *MySQL) Delete(id int32) (int64, error) {
	result, err := m.db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
