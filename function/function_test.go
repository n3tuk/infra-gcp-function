package function_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	f "github.com/n3tuk/infra-gcp-function/function"

	"github.com/stretchr/testify/assert"
)

var contentType = "application/json; charset=utf-8"

func TestRoot(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	f.NewServer().ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code, "The server should be responding with the status code %d, but received %d", http.StatusOK, w.Code)
	assert.Equal(t, contentType, w.Header().Get("Content-Type"), "The server should be responding with the Content-Type as '%s', but found '%s'", contentType, w.Header().Get("Content-Type"))
	assert.Equal(t, "{}", strings.TrimSuffix(string(b), "\n"), "The server should responding with an empty response, but found %s", strings.TrimSuffix(string(b), "\n"))
}

func TestNotFound(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/not-found", nil)
	w := httptest.NewRecorder()
	f.NewServer().ServeHTTP(w, r)

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
	assert.NotEmpty(t, e.Run, "The response should be set with the Run UUID value, but it was empty")
}

func TestAlive(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/alive", nil)
	w := httptest.NewRecorder()
	f.NewServer().ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	a := &f.AliveResponse{}
	err = json.Unmarshal(b, a)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code, "The server should be responding with the status code %d, but received %d", http.StatusOK, w.Code)
	assert.Equal(t, contentType, w.Header().Get("Content-Type"), "The server should be responding with the Content-Type as '%s', but found '%s'", contentType, w.Header().Get("Content-Type"))
	assert.Equal(t, "alive", a.Type, "The response should be set with the Type of 'alive', but found '%v'", a.Type)
	assert.True(t, a.Alive, "The response should be set with Alive as true, but found %v", a.Alive)
	assert.NotEmpty(t, a.Date, "The response should be set with the Date/Time, but it was empty")
	assert.NotEmpty(t, a.Run, "The response should be set with the Run UUID value, but it was empty")
}
