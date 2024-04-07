package Service

import "github.com/vadim-shalnev/GOlibrary/internal/entity"

func (s *AuthService) Register(regData entity.User) (string, error) {
	return s.Repository.RegisterUser(regData)
}
func (s *AuthService) Login(loginData entity.LoginData) (string, error) {
	return s.Repository.CheckUserInfo(loginData)
}
func (s *AuthService) RentBook(book entity.Books, login string) (entity.Books, error) {
	return s.Repository.RentBook(book, login)
}
func (s *AuthService) ReturnBook(book entity.Books, login string) (string, error) {
	return s.Repository.ReturnBook(book, login)
}
