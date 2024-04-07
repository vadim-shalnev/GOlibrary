package Repository

import "github.com/vadim-shalnev/GOlibrary/internal/entity"

func (r *LibraryRepository) GetUsers() ([]entity.UserBooks, error) {
	rows, err := r.DB.Query(`
		SELECT u.first_name, u.last_name, u.email, b.title, a.full_name
		FROM user_books ub
		JOIN users u ON ub.user_id = u.id
		JOIN books b ON ub.book_id = b.id
		JOIN authors a ON b.author_id = a.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersBooks []entity.UserBooks
	var currentUser entity.User
	var currentBooks []entity.Books

	for rows.Next() {
		var user entity.User
		var book entity.Books
		var author entity.Author

		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &book.Title, &author.FullName)
		if err != nil {
			return nil, err
		}

		// Проверяем, изменился ли пользователь
		if len(usersBooks) == 0 || currentUser.Email != user.Email {
			// Если это новый пользователь, добавляем предыдущего пользователя и его книги в список
			if len(usersBooks) > 0 {
				usersBooks = append(usersBooks, entity.UserBooks{Users: currentUser, Books: currentBooks})
			}
			// Начинаем новый список книг для текущего пользователя
			currentBooks = []entity.Books{book}
			currentUser = user
		} else {
			// Если это тот же пользователь, добавляем книгу в список
			currentBooks = append(currentBooks, book)
		}
	}

	// Добавляем последнего пользователя и его книги в список
	if len(currentBooks) > 0 {
		usersBooks = append(usersBooks, entity.UserBooks{Users: currentUser, Books: currentBooks})
	}

	return usersBooks, nil
}

func (r *LibraryRepository) GetAuthors() ([]entity.Author, error) {
	rows, err := r.DB.Query("SELECT full_name FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []entity.Author
	for rows.Next() {
		var author entity.Author
		err := rows.Scan(&author.FullName)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *LibraryRepository) GetBooks() ([]entity.Books, error) {
	rows, err := r.DB.Query(`
		SELECT b.title, a.full_name
		FROM books b
		JOIN authors a ON b.author_id = a.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.Books
	for rows.Next() {
		var book entity.Books
		var author entity.Author
		err := rows.Scan(&book.Title, &author.FullName)
		if err != nil {
			return nil, err
		}
		book.Author = author
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *LibraryRepository) AddAuthor(author entity.Author) error {
	// Проверяем, существует ли уже автор
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM authors WHERE full_name = $1)", author.FullName).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		// Добавляем нового автора
		_, err := r.DB.Exec("INSERT INTO authors (full_name) VALUES ($1)", author.FullName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *LibraryRepository) AddBook(book entity.Books) error {
	// Добавляем автора, если он ещё не существует
	err := r.AddAuthor(book.Author)
	if err != nil {
		return err
	}

	// Получаем ID автора
	var authorID int
	err = r.DB.QueryRow("SELECT id FROM authors WHERE full_name = $1", book.Author.FullName).Scan(&authorID)
	if err != nil {
		return err
	}

	// Добавляем книгу
	_, err = r.DB.Exec("INSERT INTO books (title, author_id) VALUES ($1, $2)", book.Title, authorID)
	if err != nil {
		return err
	}

	return nil
}
