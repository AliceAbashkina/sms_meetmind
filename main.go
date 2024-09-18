package main

import (
	"fmt"
	"log"
	"my-go-project/noti_send_sdk"
)

func main() {
	response, err := noti_send_sdk.SendSms("Привет!", "79533912220")
	if err != nil {
		log.Fatalf("Ошибка при отправке SMS: %v", err)
	}
	fmt.Printf("Ответ от API: %+v\n", response)
}
