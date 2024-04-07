package DBemptyChecker

import (
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"log"
)

func CreateTable(db *sql.DB) {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(255),
			last_name VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255)
    );`

	authors := `
	CREATE TABLE IF NOT EXISTS authors (
	    	id SERIAL PRIMARY KEY,
			full_name VARCHAR(255)
	);`

	books := `
	CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255),
			author_id INTEGER REFERENCES authors(id)
	);`
	userBooks := `
		CREATE TABLE IF NOT EXISTS user_books (
			user_id INTEGER REFERENCES users(id),
			book_id INTEGER REFERENCES books(id),
			PRIMARY KEY (user_id, book_id)
);`
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to start transaction:", err)
	}

	_, err = tx.Exec(createUsersTable)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to create users table:", err)
	}

	_, err = tx.Exec(authors)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to create geocodes table:", err)
	}

	_, err = tx.Exec(books)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to create response_history table:", err)
	}
	_, err = tx.Exec(userBooks)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to create user_books table:", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Failed to commit transaction:", err)
	}

	log.Println("Tables created successfully")
}

// Заполняем таблицы фековыми данными если пустые
func checkAndFillTables(db *sql.DB) error {
	// Проверка и заполнение таблицы users
	err := checkAndFillTable(db, "users", 51, func(db *sql.DB) error {
		_, err := db.Exec(`
			INSERT INTO users (first_name, last_name, email, password)
			VALUES ($1, $2, $3, $4)
		`, gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email(), gofakeit.Password(true, true, true, true, true, 10))
		return err
	})
	if err != nil {
		return err
	}

	// Проверка и заполнение таблицы authors
	err = checkAndFillTable(db, "authors", 10, func(db *sql.DB) error {
		_, err = db.Exec(`
			INSERT INTO authors (full_name)
			VALUES ($1)
		`, gofakeit.BookAuthor())
		return err
	})
	if err != nil {
		return err
	}

	// Проверка и заполнение таблицы books
	err = checkAndFillTable(db, "books", 100, func(db *sql.DB) error {
		_, err = db.Exec(`
			INSERT INTO books (title, author_id)
			VALUES ($1, $2)
		`, gofakeit.BookTitle(), gofakeit.Number(1, 10))
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func checkAndFillTable(db *sql.DB, tableName string, expectedCount int, insertFunc func(*sql.DB) error) error {
	var count int
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		for i := 0; i < expectedCount; i++ {
			err := insertFunc(db)
			if err != nil {
				return err
			}
		}
		log.Printf("Добавлено %d записей в таблицу %s\n", expectedCount, tableName)
	} else {
		log.Printf("Таблица %s не пуста, содержит %d записей\n", tableName, count)
	}

	return nil
}
