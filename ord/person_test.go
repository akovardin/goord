package ord

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetPersons(t *testing.T) {
	// Тестовые данные
	testResponse := PersonListResponse{
		ExternalIDs:     []string{"id1", "id2", "id3"},
		TotalItemsCount: 3,
		Limit:           10,
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		// Проверяем путь запроса
		assert.Equal(t, "/v1/person", r.URL.Path, "Expected path /v1/person")

		// Проверяем параметры запроса
		offset := r.URL.Query().Get("offset")
		limit := r.URL.Query().Get("limit")
		assert.Equal(t, "0", offset, "Expected offset=0")
		assert.Equal(t, "10", limit, "Expected limit=10")

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

	// Выполняем запрос
	result, err := client.GetPersons(context.Background(), 0, 10)
	require.NoError(t, err, "GetPersons should not return an error")

	// Проверяем результат
	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, testResponse.Limit, result.Limit, "Limit should match")
	assert.Equal(t, testResponse.ExternalIDs, result.ExternalIDs, "ExternalIDs should match")
}

func TestClient_GetPersons_Error(t *testing.T) {
	// Создаем тестовый сервер, который возвращает ошибку
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Выполняем запрос
	_, err := client.GetPersons(context.Background(), 0, 10)
	require.Error(t, err, "GetPersons should return an error")

	// Проверяем, что ошибка содержит ожидаемый текст
	assert.Contains(t, err.Error(), "failed to get persons", "Error message should contain expected text")
}
