package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vadim-shalnev/GOlibrary/config"
	"github.com/vadim-shalnev/GOlibrary/internal/router"
	"github.com/vadim-shalnev/GOlibrary/run"
	"log"

	"net/http"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Создаем конфигурацию приложения
	conf := config.NewAppConf()

	CO := run.Boostrup(conf)

	r := router.New_Router(CO)

	http.ListenAndServe(":8080", r)
}
