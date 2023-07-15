package database

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type PData struct {
	FirstPl  string
	SecondPl string
	ThirdPl  string
	MyTitle  string
	MyText1  string
	MyText2  string
	Email    string
}

var host string = "http://localhost:8080"

func Check(email, pass string) int {
	// Создаем JSON-объект с данными email и pass
	data := map[string]string{"email": email, "pass": pass}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Отправляем POST-запрос на сервер
	resp, err := http.Post(host+"/check", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Получаем ответ от сервера
	var result int
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Append(email, pass, name string) {
	// Создаем JSON-объект с данными email, pass и name
	data := map[string]string{"email": email, "pass": pass, "name": name}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Отправляем POST-запрос на сервер
	resp, err := http.Post(host+"/append", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func PageData(email string) PData {
	// Создаем JSON-объект с данными email
	data := map[string]string{"email": email}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Отправляем POST-запрос на сервер
	resp, err := http.Post(host+"/pagedata", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Получаем ответ от сервера
	var result PData
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func WData(email string) []string {
	// Создаем JSON-объект с данными email
	data := map[string]string{"email": email}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Отправляем POST-запрос на сервер
	resp, err := http.Post(host+"/wdata", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Получаем ответ от сервера
	var result []string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
