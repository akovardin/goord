package ord

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDictionaryService_GetKKTUCodes(t *testing.T) {
	// Тестовые данные
	testResponse := KKTUResponse{
		TotalItemsCount: 2,
		Limit:           10,
		Items: []KKTUItem{
			{
				Code: "01.01.01",
				Name: "Рекламные услуги",
			},
			{
				Code: "01.01.02",
				Name: "Маркетинговые услуги",
			},
		},
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		// Проверяем путь запроса
		assert.Equal(t, "/v1/dict/kktu", r.URL.Path, "Expected path /v1/dict/kktu")

		// Проверяем параметры запроса
		search := r.URL.Query().Get("search")
		lang := r.URL.Query().Get("lang")
		offset := r.URL.Query().Get("offset")
		limit := r.URL.Query().Get("limit")
		codes := r.URL.Query().Get("codes")

		assert.Equal(t, "test", search, "Expected search=test")
		assert.Equal(t, "ru", lang, "Expected lang=ru")
		assert.Equal(t, "10", offset, "Expected offset=10")
		assert.Equal(t, "20", limit, "Expected limit=20")
		assert.Equal(t, "[code1 code2]", codes, "Expected codes=[code1 code2]")

		// Возвращаем тестовый ответ
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Создаем сервис словаря
	dictService := &DictionaryService{client: client}

	// Выполняем запрос
	result, err := dictService.GetKKTUCodes(context.Background(), "test", "ru", 10, 20, []string{"code1", "code2"})
	require.NoError(t, err, "GetKKTUCodes should not return an error")

	// Проверяем результат
	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, testResponse.Limit, result.Limit, "Limit should match")
	assert.Equal(t, len(testResponse.Items), len(result.Items), "Items count should match")

	for i, item := range testResponse.Items {
		assert.Equal(t, item.Code, result.Items[i].Code, "Item code should match")
		assert.Equal(t, item.Name, result.Items[i].Name, "Item name should match")
	}
}

func TestDictionaryService_GetKKTUCodes_Error(t *testing.T) {
	// Создаем тестовый сервер, который возвращает ошибку
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Создаем сервис словаря
	dictService := &DictionaryService{client: client}

	// Выполняем запрос
	_, err := dictService.GetKKTUCodes(context.Background(), "", "", 0, 0, nil)
	require.Error(t, err, "GetKKTUCodes should return an error")
}

func TestDictionaryService_GetERIRMessage(t *testing.T) {
	// Тестовые данные
	testResponse := ERIRMessageResponse{
		Items: []ERIRMessageItem{
			{
				Message: "SUCCESS",
				Name:    "Успешно",
			},
		},
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		// Проверяем путь запроса
		assert.Equal(t, "/v1/dict/erir_message", r.URL.Path, "Expected path /v1/dict/erir_message")

		// Проверяем параметры запроса
		lang := r.URL.Query().Get("lang")
		message := r.URL.Query().Get("message")

		assert.Equal(t, "ru", lang, "Expected lang=ru")
		assert.Equal(t, "SUCCESS", message, "Expected message=SUCCESS")

		// Возвращаем тестовый ответ
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Создаем сервис словаря
	dictService := &DictionaryService{client: client}

	// Выполняем запрос
	result, err := dictService.GetERIRMessage(context.Background(), "ru", "SUCCESS")
	require.NoError(t, err, "GetERIRMessage should not return an error")

	// Проверяем результат
	assert.Equal(t, len(testResponse.Items), len(result.Items), "Items count should match")
	assert.Equal(t, testResponse.Items[0].Message, result.Items[0].Message, "Message should match")
	assert.Equal(t, testResponse.Items[0].Name, result.Items[0].Name, "Name should match")
}

func TestDictionaryService_GetERIRMessage_Error(t *testing.T) {
	// Создаем тестовый сервер, который возвращает ошибку
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Создаем сервис словаря
	dictService := &DictionaryService{client: client}

	// Выполняем запрос
	_, err := dictService.GetERIRMessage(context.Background(), "", "")
	require.Error(t, err, "GetERIRMessage should return an error")
}

func TestDictionaryService_PostERIRMessages(t *testing.T) {
	// Тестовые данные
	testResponse := ERIRMessageResponse{
		Items: []ERIRMessageItem{
			{
				Message: "SUCCESS",
				Name:    "Успешно",
			},
			{
				Message: "ERROR",
				Name:    "Ошибка",
			},
		},
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "POST", r.Method, "Expected POST request")

		// Проверяем путь запроса
		assert.Equal(t, "/v1/dict/erir_message", r.URL.Path, "Expected path /v1/dict/erir_message")

		// Проверяем заголовки
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"), "Expected Content-Type application/json")

		// Проверяем тело запроса
		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		assert.NoError(t, err, "Should be able to decode request body")

		// Проверяем данные в теле запроса
		assert.Equal(t, "ru", reqBody["lang"], "Lang should match")
		messages, ok := reqBody["messages"].([]interface{})
		assert.True(t, ok, "Messages should be an array")
		assert.Equal(t, 2, len(messages), "Should have 2 messages")
		assert.Equal(t, "SUCCESS", messages[0], "First message should match")
		assert.Equal(t, "ERROR", messages[1], "Second message should match")

		// Возвращаем тестовый ответ
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Создаем сервис словаря
	dictService := &DictionaryService{client: client}

	// Выполняем запрос
	result, err := dictService.PostERIRMessages(context.Background(), "ru", []string{"SUCCESS", "ERROR"})
	require.NoError(t, err, "PostERIRMessages should not return an error")

	// Проверяем результат
	assert.Equal(t, len(testResponse.Items), len(result.Items), "Items count should match")

	for i, item := range testResponse.Items {
		assert.Equal(t, item.Message, result.Items[i].Message, "Item message should match")
		assert.Equal(t, item.Name, result.Items[i].Name, "Item name should match")
	}
}

func TestDictionaryService_PostERIRMessages_Error(t *testing.T) {
	// Создаем тестовый сервер, который возвращает ошибку
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Создаем сервис словаря
	dictService := &DictionaryService{client: client}

	// Выполняем запрос
	_, err := dictService.PostERIRMessages(context.Background(), "", []string{})
	require.Error(t, err, "PostERIRMessages should return an error")
}
