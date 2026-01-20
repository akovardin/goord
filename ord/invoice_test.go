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

func TestClient_GetInvoices(t *testing.T) {
	// Тестовые данные
	testResponse := InvoiceListResponse{
		ExternalIDs:     []string{"id1", "id2", "id3"},
		TotalItemsCount: 3,
		Limit:           10,
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		// Проверяем путь запроса
		assert.Equal(t, "/v1/invoice", r.URL.Path, "Expected path /v1/invoice")

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
	result, err := client.GetInvoices(context.Background(), 0, 10)
	require.NoError(t, err, "GetInvoices should not return an error")

	// Проверяем результат
	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, testResponse.Limit, result.Limit, "Limit should match")
	assert.Equal(t, testResponse.ExternalIDs, result.ExternalIDs, "ExternalIDs should match")
}

func TestClient_GetInvoices_Error(t *testing.T) {
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

	// Выполняем запрос
	_, err := client.GetInvoices(context.Background(), 0, 10)
	require.Error(t, err, "GetInvoices should return an error")
}

func TestClient_GetInvoice(t *testing.T) {
	// Тестовые данные
	testResponse := Invoice{
		ContractExternalID: "test-contract-id",
		Date:               "2023-01-01",
		DateStart:          "2023-01-01",
		DateEnd:            "2023-01-31",
		Amount: InvoiceAmount{
			Services: InvoiceAmountGroup{
				ExcludingVat: "1000.00",
				VatRate:      "20",
				Vat:          "200.00",
				IncludingVat: "1200.00",
			},
		},
		ClientRole:     "advertiser",
		ContractorRole: "publisher",
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		// Проверяем путь запроса
		assert.Equal(t, "/v4/invoice/test-invoice-id", r.URL.Path, "Expected path /v4/invoice/test-invoice-id")

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
	result, err := client.GetInvoice(context.Background(), "test-invoice-id")
	require.NoError(t, err, "GetInvoice should not return an error")

	// Проверяем результат
	assert.Equal(t, testResponse.ContractExternalID, result.ContractExternalID, "ContractExternalID should match")
	assert.Equal(t, testResponse.Date, result.Date, "Date should match")
	assert.Equal(t, testResponse.DateStart, result.DateStart, "DateStart should match")
	assert.Equal(t, testResponse.DateEnd, result.DateEnd, "DateEnd should match")
	assert.Equal(t, testResponse.ClientRole, result.ClientRole, "ClientRole should match")
	assert.Equal(t, testResponse.ContractorRole, result.ContractorRole, "ContractorRole should match")
}

func TestClient_GetInvoice_Error(t *testing.T) {
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

	// Выполняем запрос
	_, err := client.GetInvoice(context.Background(), "test-invoice-id")
	require.Error(t, err, "GetInvoice should return an error")
}

func TestClient_CreateInvoiceHeader(t *testing.T) {
	// Тестовые данные
	invoice := Invoice{
		ContractExternalID: "test-contract-id",
		Date:               "2023-01-01",
		DateStart:          "2023-01-01",
		DateEnd:            "2023-01-31",
		Amount: InvoiceAmount{
			Services: InvoiceAmountGroup{
				ExcludingVat: "1000.00",
				VatRate:      "20",
				Vat:          "200.00",
				IncludingVat: "1200.00",
			},
		},
		ClientRole:     "advertiser",
		ContractorRole: "publisher",
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "PUT", r.Method, "Expected PUT request")

		// Проверяем путь запроса
		assert.Equal(t, "/v4/invoice/test-invoice-id/header", r.URL.Path, "Expected path /v4/invoice/test-invoice-id/header")

		// Проверяем тело запроса
		var reqBody Invoice
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		assert.NoError(t, err, "Should be able to decode request body")

		// Проверяем данные в теле запроса
		assert.Equal(t, invoice.ContractExternalID, reqBody.ContractExternalID, "ContractExternalID should match")
		assert.Equal(t, invoice.Date, reqBody.Date, "Date should match")
		assert.Equal(t, invoice.DateStart, reqBody.DateStart, "DateStart should match")
		assert.Equal(t, invoice.DateEnd, reqBody.DateEnd, "DateEnd should match")
		assert.Equal(t, invoice.ClientRole, reqBody.ClientRole, "ClientRole should match")
		assert.Equal(t, invoice.ContractorRole, reqBody.ContractorRole, "ContractorRole should match")

		// Возвращаем пустой ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Выполняем запрос
	err := client.CreateInvoiceHeader(context.Background(), "test-invoice-id", invoice)
	require.NoError(t, err, "CreateInvoiceHeader should not return an error")
}

func TestClient_CreateInvoiceHeader_Error(t *testing.T) {
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

	// Выполняем запрос
	err := client.CreateInvoiceHeader(context.Background(), "test-invoice-id", Invoice{})
	require.Error(t, err, "CreateInvoiceHeader should return an error")
}

func TestClient_AddContractsToInvoice(t *testing.T) {
	// Тестовые данные
	items := []InvoiceItem{
		{
			ContractExternalID: stringPtr("test-contract-id"),
			Amount: InvoiceAmountGroup{
				ExcludingVat: "1000.00",
				VatRate:      "20",
				Vat:          "200.00",
				IncludingVat: "1200.00",
			},
		},
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "PATCH", r.Method, "Expected PATCH request")

		// Проверяем путь запроса
		assert.Equal(t, "/v4/invoice/test-invoice-id/items", r.URL.Path, "Expected path /v4/invoice/test-invoice-id/items")

		// Проверяем тело запроса
		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		assert.NoError(t, err, "Should be able to decode request body")

		// Проверяем данные в теле запроса
		reqItems, ok := reqBody["items"].([]interface{})
		assert.True(t, ok, "Items should be an array")
		assert.Equal(t, len(items), len(reqItems), "Should have same number of items")

		// Возвращаем пустой ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Выполняем запрос
	err := client.AddContractsToInvoice(context.Background(), "test-invoice-id", items)
	require.NoError(t, err, "AddContractsToInvoice should not return an error")
}

func TestClient_AddContractsToInvoice_Error(t *testing.T) {
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

	// Выполняем запрос
	err := client.AddContractsToInvoice(context.Background(), "test-invoice-id", []InvoiceItem{})
	require.Error(t, err, "AddContractsToInvoice should return an error")
}

func TestClient_DeleteInvoice(t *testing.T) {
	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "DELETE", r.Method, "Expected DELETE request")

		// Проверяем путь запроса
		assert.Equal(t, "/v4/invoice/test-invoice-id", r.URL.Path, "Expected path /v4/invoice/test-invoice-id")

		// Возвращаем пустой ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Выполняем запрос
	err := client.DeleteInvoice(context.Background(), "test-invoice-id")
	require.NoError(t, err, "DeleteInvoice should not return an error")
}

func TestClient_DeleteInvoice_Error(t *testing.T) {
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

	// Выполняем запрос
	err := client.DeleteInvoice(context.Background(), "test-invoice-id")
	require.Error(t, err, "DeleteInvoice should return an error")
}

func TestClient_SendInvoiceToErir(t *testing.T) {
	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "POST", r.Method, "Expected POST request")

		// Проверяем путь запроса
		assert.Equal(t, "/v4/invoice/test-invoice-id/ready", r.URL.Path, "Expected path /v4/invoice/test-invoice-id/ready")

		// Возвращаем пустой ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Выполняем запрос
	err := client.SendInvoiceToErir(context.Background(), "test-invoice-id")
	require.NoError(t, err, "SendInvoiceToErir should not return an error")
}

func TestClient_SendInvoiceToErir_Error(t *testing.T) {
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

	// Выполняем запрос
	err := client.SendInvoiceToErir(context.Background(), "test-invoice-id")
	require.Error(t, err, "SendInvoiceToErir should return an error")
}

func TestClient_DeleteContractsFromInvoice(t *testing.T) {
	// Тестовые данные
	deleteInfo := map[string]interface{}{
		"items": []map[string]interface{}{
			{
				"contract_external_id": "test-contract-id",
			},
		},
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "POST", r.Method, "Expected POST request")

		// Проверяем путь запроса
		assert.Equal(t, "/v4/invoice/test-invoice-id/delete", r.URL.Path, "Expected path /v4/invoice/test-invoice-id/delete")

		// Проверяем тело запроса
		var reqBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		assert.NoError(t, err, "Should be able to decode request body")

		// Возвращаем пустой ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Выполняем запрос
	err := client.DeleteContractsFromInvoice(context.Background(), "test-invoice-id", deleteInfo)
	require.NoError(t, err, "DeleteContractsFromInvoice should not return an error")
}

func TestClient_DeleteContractsFromInvoice_Error(t *testing.T) {
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

	// Выполняем запрос
	err := client.DeleteContractsFromInvoice(context.Background(), "test-invoice-id", map[string]interface{}{})
	require.Error(t, err, "DeleteContractsFromInvoice should return an error")
}

func TestClient_CreateWholeInvoice(t *testing.T) {
	// Тестовые данные
	invoice := Invoice{
		ContractExternalID: "test-contract-id",
		Date:               "2023-01-01",
		DateStart:          "2023-01-01",
		DateEnd:            "2023-01-31",
		Amount: InvoiceAmount{
			Services: InvoiceAmountGroup{
				ExcludingVat: "1000.00",
				VatRate:      "20",
				Vat:          "200.00",
				IncludingVat: "1200.00",
			},
		},
		ClientRole:     "advertiser",
		ContractorRole: "publisher",
	}

	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		assert.Equal(t, "PUT", r.Method, "Expected PUT request")

		// Проверяем путь запроса
		assert.Equal(t, "/v4/invoice/test-invoice-id", r.URL.Path, "Expected path /v4/invoice/test-invoice-id")

		// Проверяем тело запроса
		var reqBody Invoice
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		assert.NoError(t, err, "Should be able to decode request body")

		// Проверяем данные в теле запроса
		assert.Equal(t, invoice.ContractExternalID, reqBody.ContractExternalID, "ContractExternalID should match")
		assert.Equal(t, invoice.Date, reqBody.Date, "Date should match")
		assert.Equal(t, invoice.DateStart, reqBody.DateStart, "DateStart should match")
		assert.Equal(t, invoice.DateEnd, reqBody.DateEnd, "DateEnd should match")
		assert.Equal(t, invoice.ClientRole, reqBody.ClientRole, "ClientRole should match")
		assert.Equal(t, invoice.ContractorRole, reqBody.ContractorRole, "ContractorRole should match")

		// Возвращаем пустой ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаем клиент с тестовым сервером
	client := NewClient(Config{
		BaseURL: server.URL,
		Token:   "test-token",
	})

	// Выполняем запрос
	err := client.CreateWholeInvoice(context.Background(), "test-invoice-id", invoice)
	require.NoError(t, err, "CreateWholeInvoice should not return an error")
}

func TestClient_CreateWholeInvoice_Error(t *testing.T) {
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

	// Выполняем запрос
	err := client.CreateWholeInvoice(context.Background(), "test-invoice-id", Invoice{})
	require.Error(t, err, "CreateWholeInvoice should return an error")
}
