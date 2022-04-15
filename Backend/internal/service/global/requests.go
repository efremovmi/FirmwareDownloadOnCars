package global

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type TypeData struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Lastupdate string `json:"lastupdate"`
	File       string `json:"file"`
	Status     string `json:"status"`
	Usergroup  string `json:"usergroup"`
}

type AllData struct {
	Status bool       `json:"status"`
	Data   []TypeData `json:"data"`
}

func AskAll() (error, *AllData) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParam("key", "getCars").
		SetQueryParam("apikey", "LQRYA65qLUoiQj1nN9caeO4CxAE9f8Oi").
		Get("https://hackmpei.ru/api/getData.php")

	if err != nil {
		return err, nil
	}

	var parsingStruct AllData

	err = json.Unmarshal(resp.Body(), &parsingStruct)

	if err != nil {
		return err, nil
	}

	return nil, &parsingStruct
}

func GetFile(name, filename string) error {
	client := resty.New()

	_, err := client.R().
		SetQueryParam("key", "dowloadFile").
		SetQueryParam("apikey", "LQRYA65qLUoiQj1nN9caeO4CxAE9f8Oi").
		SetQueryParam("name", filename).
		SetOutput(fmt.Sprintf("%s_%s", name, filename)).
		Get("https://hackmpei.ru/api/download.php")

	if err != nil {
		return err
	}

	return nil
}
