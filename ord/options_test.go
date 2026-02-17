//nolint:errcheck
package ord

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithToken(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err, "NewClient should not return an error")

	token := "test-token"
	option := WithToken(token)
	err = option(client)
	require.NoError(t, err, "WithToken should not return an error")

	assert.Equal(t, token, client.token, "Client token should be set")
}

func TestWithHttpClient(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err, "NewClient should not return an error")

	httpClient := &http.Client{}
	option := WithHttpClient(httpClient)
	err = option(client)
	require.NoError(t, err, "WithHttpClient should not return an error")

	assert.Equal(t, httpClient, client.http, "HTTP client should be set")
}

func TestWithBase(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err, "NewClient should not return an error")

	base := "https://test.api.com"
	option := WithBase(base)
	err = option(client)
	require.NoError(t, err, "WithBase should not return an error")

	assert.Equal(t, base, client.base, "Base URL should be set")
}
