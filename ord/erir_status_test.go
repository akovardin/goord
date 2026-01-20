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

func TestClient_GetObjectProcessingStatus(t *testing.T) {
	// Тестовые данные
	testResponse := ErirStatusEntity{
		ErirStatus:      "verified",
		UpdatedByUserTs: "2023-05-25T12:17:26Z",
		FinalizedTs:     stringPtr("2023-05-28T12:17:26Z"),
		Messages:        []string(nil),
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		// Проверяем путь запроса
		assert.Equal(t, "/v1/person/id1/erir_status", r.URL.Path, "Expected path /v1/person/id1/erir_status")

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
	result, err := client.GetObjectProcessingStatus(context.Background(), "person", "id1")
	require.NoError(t, err, "GetObjectProcessingStatus should not return an error")

	// Проверяем результат
	assert.Equal(t, testResponse.ErirStatus, result.ErirStatus, "ErirStatus should match")
	assert.Equal(t, testResponse.UpdatedByUserTs, result.UpdatedByUserTs, "UpdatedByUserTs should match")
	assert.Equal(t, testResponse.FinalizedTs, result.FinalizedTs, "FinalizedTs should match")
	assert.Equal(t, testResponse.Messages, result.Messages, "Messages should match")
}

func TestClient_GetObjectProcessingStatus_Error(t *testing.T) {
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
	_, err := client.GetObjectProcessingStatus(context.Background(), "person", "id1")
	require.Error(t, err, "GetObjectProcessingStatus should return an error")

	// Проверяем, что ошибка содержит ожидаемый текст
	assert.Contains(t, err.Error(), "failed to get object processing status", "Error message should contain expected text")
}

func TestClient_GetAdObjectProcessingStatus(t *testing.T) {
	// Тестовые данные
	testResponse := ErirStatusEntities{
		TotalItemsCount: 1,
		Limit:           10,
		LimitPerEntity:  5,
		Items: []ErirStatusEntityItem{
			{
				DataType:        "person",
				ExternalID:      "id1",
				Name:            "Test Person",
				ErirTaxStatus:   "no_tax",
				ErirStatus:      "verified",
				UpdatedByUserTs: "2023-05-25T12:17:26Z",
				FinalizedTs:     stringPtr("2023-05-28T12:17:26Z"),
				Messages:        []string(nil),
			},
		},
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		// Проверяем путь запроса
		assert.Equal(t, "/v1/erir_statuses", r.URL.Path, "Expected path /v1/erir_statuses")

		// Проверяем параметры запроса
		dataType := r.URL.Query().Get("data_type")
		erirStatus := r.URL.Query().Get("erir_status")
		offset := r.URL.Query().Get("offset")
		limit := r.URL.Query().Get("limit")
		limitPerEntity := r.URL.Query().Get("limit_per_entity")
		externalID := r.URL.Query()["external_id"]

		assert.Equal(t, "person", dataType, "Expected data_type=person")
		assert.Equal(t, "verified", erirStatus, "Expected erir_status=verified")
		assert.Equal(t, "0", offset, "Expected offset=0")
		assert.Equal(t, "10", limit, "Expected limit=10")
		assert.Equal(t, "5", limitPerEntity, "Expected limit_per_entity=5")
		assert.Equal(t, []string{"id1", "id2"}, externalID, "Expected external_id=[id1,id2]")

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
	result, err := client.GetAdObjectProcessingStatus(context.Background(), "person", "verified", 0, 10, 5, []string{"id1", "id2"})
	require.NoError(t, err, "GetAdObjectProcessingStatus should not return an error")

	// Проверяем результат
	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, testResponse.Limit, result.Limit, "Limit should match")
	assert.Equal(t, testResponse.LimitPerEntity, result.LimitPerEntity, "LimitPerEntity should match")

	require.Len(t, result.Items, 1, "Should have one item")
	item := result.Items[0]
	expectedItem := testResponse.Items[0]
	assert.Equal(t, expectedItem.DataType, item.DataType, "DataType should match")
	assert.Equal(t, expectedItem.ExternalID, item.ExternalID, "ExternalID should match")
	assert.Equal(t, expectedItem.Name, item.Name, "Name should match")
	assert.Equal(t, expectedItem.ErirTaxStatus, item.ErirTaxStatus, "ErirTaxStatus should match")
	assert.Equal(t, expectedItem.ErirStatus, item.ErirStatus, "ErirStatus should match")
	assert.Equal(t, expectedItem.UpdatedByUserTs, item.UpdatedByUserTs, "UpdatedByUserTs should match")
	assert.Equal(t, expectedItem.FinalizedTs, item.FinalizedTs, "FinalizedTs should match")
	assert.Equal(t, expectedItem.Messages, item.Messages, "Messages should match")
}

func TestClient_GetAdObjectProcessingStatus_Error(t *testing.T) {
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
	_, err := client.GetAdObjectProcessingStatus(context.Background(), "person", "verified", 0, 10, 5, []string{"id1"})
	require.Error(t, err, "GetAdObjectProcessingStatus should return an error")

	// Проверяем, что ошибка содержит ожидаемый текст
	assert.Contains(t, err.Error(), "failed to get ad object processing statuses", "Error message should contain expected text")
}

func TestClient_PostAdObjectProcessingStatus(t *testing.T) {
	// Тестовые данные
	testResponse := ErirStatusEntities{
		TotalItemsCount: 1,
		Limit:           10,
		LimitPerEntity:  5,
		Items: []ErirStatusEntityItem{
			{
				DataType:        "person",
				ExternalID:      "id1",
				Name:            "Test Person",
				ErirTaxStatus:   "no_tax",
				ErirStatus:      "verified",
				UpdatedByUserTs: "2023-05-25T12:17:26Z",
				FinalizedTs:     stringPtr("2023-05-28T12:17:26Z"),
				Messages:        []string(nil),
			},
		},
	}

	testRequest := PostAdObjectProcessingStatusRequest{
		DataType:       "person",
		ErirStatus:     "verified",
		ExternalID:     []string{"id1", "id2"},
		Offset:         0,
		Limit:          10,
		LimitPerEntity: 5,
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "POST", r.Method, "Expected POST request")

		// Проверяем путь запроса
		assert.Equal(t, "/v1/erir_statuses", r.URL.Path, "Expected path /v1/erir_statuses")

		// Проверяем заголовки
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"), "Expected Content-Type application/json")

		// Читаем и проверяем тело запроса
		var request PostAdObjectProcessingStatusRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		require.NoError(t, err, "Should be able to decode request body")

		assert.Equal(t, testRequest.DataType, request.DataType, "DataType should match")
		assert.Equal(t, testRequest.ErirStatus, request.ErirStatus, "ErirStatus should match")
		assert.Equal(t, testRequest.ExternalID, request.ExternalID, "ExternalID should match")
		assert.Equal(t, testRequest.Offset, request.Offset, "Offset should match")
		assert.Equal(t, testRequest.Limit, request.Limit, "Limit should match")
		assert.Equal(t, testRequest.LimitPerEntity, request.LimitPerEntity, "LimitPerEntity should match")

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
	result, err := client.PostAdObjectProcessingStatus(context.Background(), testRequest)
	require.NoError(t, err, "PostAdObjectProcessingStatus should not return an error")

	// Проверяем результат
	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, testResponse.Limit, result.Limit, "Limit should match")
	assert.Equal(t, testResponse.LimitPerEntity, result.LimitPerEntity, "LimitPerEntity should match")

	require.Len(t, result.Items, 1, "Should have one item")
	item := result.Items[0]
	expectedItem := testResponse.Items[0]
	assert.Equal(t, expectedItem.DataType, item.DataType, "DataType should match")
	assert.Equal(t, expectedItem.ExternalID, item.ExternalID, "ExternalID should match")
	assert.Equal(t, expectedItem.Name, item.Name, "Name should match")
	assert.Equal(t, expectedItem.ErirTaxStatus, item.ErirTaxStatus, "ErirTaxStatus should match")
	assert.Equal(t, expectedItem.ErirStatus, item.ErirStatus, "ErirStatus should match")
	assert.Equal(t, expectedItem.UpdatedByUserTs, item.UpdatedByUserTs, "UpdatedByUserTs should match")
	assert.Equal(t, expectedItem.FinalizedTs, item.FinalizedTs, "FinalizedTs should match")
	assert.Equal(t, expectedItem.Messages, item.Messages, "Messages should match")
}

func TestClient_PostAdObjectProcessingStatus_Error(t *testing.T) {
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
	testRequest := PostAdObjectProcessingStatusRequest{
		DataType: "person",
	}

	_, err := client.PostAdObjectProcessingStatus(context.Background(), testRequest)
	require.Error(t, err, "PostAdObjectProcessingStatus should return an error")

	// Проверяем, что ошибка содержит ожидаемый текст
	assert.Contains(t, err.Error(), "failed to post ad object processing statuses", "Error message should contain expected text")
}
