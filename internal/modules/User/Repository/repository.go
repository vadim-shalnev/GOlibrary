package Repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/vadim-shalnev/GOlibrary/internal/entity"
)

func (r *AuthRepository) RegisterUser(regData entity.User) (string, error) {
	_, err := r.DB.Exec("INSERT INTO users(first_name, last_name, Email, Password) VALUES (?, ?, ?, ?)", regData.FirstName, regData.LastName, regData.Email, regData.Password)
	if err != nil {
		return "", err

	}
	return "success", nil
}

func (r *AuthRepository) CheckUserInfo(loginData entity.LoginData) (string, error) {
	var user entity.User
	err := r.DB.QueryRow("SELECT * FROM users WHERE Email = ? AND Password = ?", loginData.Email, loginData.Password).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return "", errors.New("неправильный логин или пароль")
	}
	if user.Email != loginData.Email || user.Password != loginData.Password {
		return "", errors.New("неправильный логин или пароль")
	}
	return "success", nil
}

func (r *AuthRepository) RentBook(book entity.Books, login string) (entity.Books, error) {
	var userID, bookID int

	// Проверка наличия пользователя
	err := r.DB.QueryRow("SELECT id FROM users WHERE email = $1", login).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("пользователь с логином %s не найден", login)
		}
		return book, err
	}

	// Проверка наличия книги
	err = r.DB.QueryRow("SELECT id FROM books WHERE title = $1", book.Title).Scan(&bookID)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("книга с названием %s не найдена", book.Title)
		}
		return book, err
	}

	// Проверка наличия записи о взятии книги пользователем
	var exists bool
	err = r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_books WHERE book_id = $1)", bookID).Scan(&exists)
	if err != nil {
		return book, err
	}
	if exists {
		return book, fmt.Errorf("пользователь %s уже взял книгу %s", login, book.Title)
	}

	// Добавление записи о взятии книги пользователем
	_, err = r.DB.Exec("INSERT INTO user_books (user_id, book_id) VALUES ($1, $2)", userID, bookID)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *AuthRepository) ReturnBook(book entity.Books, login string) (string, error) {
	var userID, bookID int
	var exists bool

	// Получение ID пользователя по логину
	err := r.DB.QueryRow("SELECT id FROM users WHERE email = $1", login).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("пользователь с логином %s не найден", login)
		}
		return "", err
	}

	// Получение ID книги по названию
	err = r.DB.QueryRow("SELECT id FROM books WHERE title = $1", book.Title).Scan(&bookID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("книга с названием %s не найдена", book.Title)
		}
		return "", err
	}

	// Проверка наличия записи о взятии книги пользователем
	err = r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_books WHERE user_id = $1 AND book_id = $2)", userID, bookID).Scan(&exists)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("пользователь %s не взял книгу %s", login, book.Title)
	}

	// Удаление записи о взятии книги пользователем
	_, err = r.DB.Exec("DELETE FROM user_books WHERE user_id = $1 AND book_id = $2", userID, bookID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Книга %s успешно возвращена пользователем %s", book.Title, login), nil
}
