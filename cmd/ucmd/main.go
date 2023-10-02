package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*type SingleUser struct {
	Data struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}*/

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type KE struct {
	ID          string `json:"id"`
	Name        string `json:"Наименование"`
	System      string `json:"ИТ система"`
	Env         string `json:"Среда"`
	MainIP      string `json:"Основной IP адрес"`
	IPs         string `json:"IP адреса"`
	AltName     string `json:"Альтернативное имя"`
	Domain      string `json:"ci_domain"`
	Segment     string `json:"Сетевой сегмент"`
	OS          string `json:"Семейство ОС"`
	NameOS      string `json:"Имя ОС"`
	VersionOS   string `json:"Версия ОС"`
	MainOwner   string `json:"Основной владелец"`
	SecondOwner string `json:"Замещающий владелец"`
	CMDBID      string `json:"CMDB_ID"`
	Status      string `json:"Статус"`
	InfGroup    string `json:"Инфраструктурная группа"`
	Comment     string `json:"Комментарии"`
}

type Answer struct {
	CurrentUser string `json:"current_user"`
	TotalSize   int    `json:"total_size"`
	RowsInPage  int    `json:"rows_in_page"`
	Pagination  struct {
		CurrentPageNum  int    `json:"current_page_num"`
		CurrentPageSize int    `json:"current_page_size"`
		Previous        string `json:"previous"`
		Next            string `json:"next"`
	} `json:"pagination"`
	Data []KE `json:"data"`
}

func main() {
	var loginPrm, passPrm, urlTokenPrm, urlPrm string
	flag.StringVar(&loginPrm, "login", "user", "as string")
	flag.StringVar(&passPrm, "password", "pwd", "as string")
	flag.StringVar(&urlTokenPrm, "urlToken", "https://localhost/", "as string")
	flag.StringVar(&urlPrm, "url", "https://localhost/", "as string")
	flag.Parse()

	data := []byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, loginPrm, passPrm))
	r := bytes.NewReader(data)
	response, apiError := http.Post(urlTokenPrm, "application/json", r)

	if apiError != nil {
		log.Fatal(apiError)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		panic(err.Error())
	}

	client := &http.Client{}
	req, err2 := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("token", token.AccessToken)

	if err2 != nil {
		panic(err2.Error())
	}
	resp, err3 := client.Do(req)
	if err3 != nil {
		log.Fatal(err3.Error())
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(response.Body)
	var answer Answer
	err = json.Unmarshal(body, &answer)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Results: %v\n", token)
	fmt.Println("login:", loginPrm)
	fmt.Println("password:", passPrm)
	fmt.Println("url token:", urlTokenPrm)
	fmt.Println("url:", urlPrm)

	fmt.Println(token.TokenType)
	fmt.Println(token.AccessToken)

	fmt.Println(answer.CurrentUser)
	fmt.Println(answer.TotalSize)
	fmt.Println(answer.RowsInPage)
}
