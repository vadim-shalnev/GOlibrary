package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	resp, err := client.Get("http://localhost:8080/api/library/users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var users []UserBooks
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
	//
	resp, err = client.Get("http://localhost:8080/api/library/books")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var books []Books
	err = json.NewDecoder(resp.Body).Decode(&books)
	if err != nil {
		fmt.Println(err)
	}
	for _, book := range books {
		fmt.Printf("%+v\n", book)
	}
}
