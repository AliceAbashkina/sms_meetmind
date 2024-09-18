package noti_send_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my-go-project/universal_code"
	"net/http"
)

const (
	apiKey = "de5f0d28d53fd2f3b528b61bdd4d5d00" // API ключ (специфичный для NotiSend)
	apiUrl = "https://sms.notisend.ru/api"      // Базовый URL API
)

// Структура для данных запроса (специфична для этого API)
type SmsRequestData struct {
	Project    string `json:"project"`
	Recipients string `json:"recipients"`
	Message    string `json:"message"`
	ApiKey     string `json:"apikey"`
}

// Создание тела запроса с использованием JSON данных
func createRequestBody(data SmsRequestData) (*bytes.Buffer, string, error) {
	// Преобразуем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, "", fmt.Errorf("не удалось преобразовать данные в JSON: %w", err)
	}

	// Создаем буфер с JSON данными
	requestBody := bytes.NewBuffer(jsonData)
	return requestBody, "application/json", nil
}

// Создание HTTP запроса
func createHttpRequest(method, url string, body *bytes.Buffer, contentType string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать запрос: %w", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// Выполнение HTTP запроса
func executeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("запрос не удался: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неожиданный код ответа: %d", resp.StatusCode)
	}
	return resp, nil
}

// Обработка ответа от API
func processApiResponse(resp *http.Response) (*universal_code.ApiResponse, error) {
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать ответ: %w", err)
	}

	var response universal_code.ApiResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("не удалось распарсить тело ответа: %w", err)
	}

	return &response, nil
}

// Основная функция для отправки SMS (специфична для API NotiSend)
func SendSms(message string, recipient string) (*universal_code.ApiResponse, error) {
	data := SmsRequestData{
		Project:    "sms_meetmind",
		Recipients: recipient,
		Message:    message,
		ApiKey:     apiKey,
	}

	// Создаем тело запроса
	requestBody, contentType, err := createRequestBody(data)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании тела запроса: %w", err)
	}

	// Формируем запрос
	req, err := createHttpRequest("POST", apiUrl+"/message/send", requestBody, contentType)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать запрос: %w", err)
	}

	// Выполняем запрос
	resp, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	// Обрабатываем ответ от API
	return processApiResponse(resp)
}
