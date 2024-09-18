package universal_code

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Структура для универсального API-ответа
type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Создание HTTP-запроса (универсальная функция)
func CreateHttpRequest(method, url string, body *bytes.Buffer, contentType string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать запрос: %w", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// Выполнение HTTP-запроса (универсальная функция)
func ExecuteRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("запрос не удался: %w", err)
	}

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неожиданный код ответа: %d", resp.StatusCode)
	}

	return resp, nil
}

// Обработка API-ответа (универсальная функция)
func ProcessApiResponse(resp *http.Response) (*ApiResponse, error) {
	defer resp.Body.Close()

	// Читаем тело ответа
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать ответ: %w", err)
	}

	// Парсим тело ответа
	var response ApiResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("не удалось распарсить тело ответа: %w", err)
	}

	return &response, nil
}
