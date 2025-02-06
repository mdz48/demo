package infraestructure

import (
	"database/sql"
	"demo/src/books/domain"
	"fmt"
	"strings"
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

func (m *MySQL) GetAll() ([]domain.BookWithAuthor, error) {
    query := `
        SELECT b.id, b.title, b.description, a.id as author_id, a.name as author_name 
        FROM books b 
        INNER JOIN users a ON b.author = a.id
    `
    rows, err := m.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []domain.BookWithAuthor
    for rows.Next() {
        var book domain.BookWithAuthor
        err := rows.Scan(&book.Id, &book.Title, &book.Description, &book.AuthorId, &book.AuthorName)
        if err != nil {
            return nil, err
        }
        books = append(books, book)
    }

    if err = rows.Err(); err != nil {
        return nil, err
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

func (m *MySQL) GetBooksByAuthor(authorId int32) ([]domain.BookWithAuthor, error) {
    query := `
        SELECT b.id, b.title, b.description, a.id as author_id, a.name as author_name 
        FROM books b 
        INNER JOIN users a ON b.author = a.id 
        WHERE b.author = ?
    `
    rows, err := m.db.Query(query, authorId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []domain.BookWithAuthor
    for rows.Next() {
        var book domain.BookWithAuthor
        err := rows.Scan(&book.Id, &book.Title, &book.Description, &book.AuthorId, &book.AuthorName)
        if err != nil {
            return nil, err
        }
        books = append(books, book)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return books, nil
}

func (m *MySQL) AddFavoriteBook(userId int32, bookId int32) (int64, error) {
    result, err := m.db.Exec("INSERT INTO favorite_books (user_id, book_id) VALUES (?, ?)", userId, bookId)
    if err != nil {
        if strings.Contains(err.Error(), "Duplicate entry") {
            return 0, fmt.Errorf("el libro ya está en favoritos")
        }
        return 0, err
    }
    return result.RowsAffected()
}
