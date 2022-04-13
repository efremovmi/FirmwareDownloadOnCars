package usecase

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

var (
	IpMaps = map[string]string{
		"carAr_1": "http://192.168.1.169/mycar",
		"carAr_2": "http://192.168.1.248/mycar",
		"carAr_3": "http://192.168.1.169/mycar",
		"carAr_4": "http://192.168.1.169/mycar",
	}
)

func (uc *Handler) UpdateArduino(name string, file string) error {
	if _, ok := IpMaps[name]; !ok {
		return fmt.Errorf("I have no idea what is it")
	}

	fmt.Println("Upload on", name)

	client := resty.New().SetTimeout(time.Second * 30)
	fmt.Println(file)

	resp, err := client.R().
		SetFile("firmware", file).
		Post(IpMaps[name])

	fmt.Println(err.Error())

	fmt.Println(resp.Status())

	return err
}
