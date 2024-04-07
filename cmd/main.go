package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	//err := godotenv.Load()
	//conf := config.NewAppConf()

	book := gofakeit.Book()
	fmt.Println(book)
}
