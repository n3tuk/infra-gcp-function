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

func TestAlive(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/alive", nil)
	w := httptest.NewRecorder()
	f.NewRouter().ServeHTTP(w, r)

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

	assert.NotEmpty(t, w.Header().Get("X-Server-ID"), "The response should have the header X-Server-ID set with the UUID for the Server runtime")
	assert.NotEmpty(t, w.Header().Get("X-Request-ID"), "The response should have the header X-Request-ID set with the UUID for the Request")
}
