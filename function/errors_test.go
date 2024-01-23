package function_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	f "github.com/n3tuk/infra-gcp-function/function"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/not-found", nil)
	w := httptest.NewRecorder()
	f.NewRouter().ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	e := &f.ErrorResponse{}
	err = json.Unmarshal(b, e)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, w.Code, "The server should be responding with the status code %d, but received %d", http.StatusNotFound, w.Code)
	assert.Equal(t, contentType, w.Header().Get("Content-Type"), "The server should be responding with the Content-Type as '%s', but found '%s'", contentType, w.Header().Get("Content-Type"))
	assert.Equal(t, "error", e.Type, "The response should set with the Type of 'error', but found '%v'", e.Type)
	assert.Equal(t, http.StatusNotFound, e.Status, "The response should be set with the Status as %d, but found %d", http.StatusNotFound, e.Status)

	assert.NotEmpty(t, w.Header().Get("X-Server-ID"), "The response should have the header X-Server-ID set with the UUID for the Server runtime")
	assert.NotEmpty(t, w.Header().Get("X-Request-ID"), "The response should have the header X-Request-ID set with the UUID for the Request")
}
