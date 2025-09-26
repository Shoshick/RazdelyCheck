package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	apiURL := "https://proverkacheka.com/api/v1/check/get"

	// твой токен из личного кабинета
	token := "35571.3OFIStTeEEbgjlH9N"

	// сырая строка из QR-кода
	qrraw := "t=20250920T1536&s=10673.80&fn=7380440802042622&i=9675&fp=1637223767&n=1"

	// собираем данные формы
	data := url.Values{}
	data.Set("token", token)
	data.Set("qrraw", qrraw)

	// создаём POST-запрос
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
