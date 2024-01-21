package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	f "github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/google/uuid"
)

type (
	AliveResponse struct {
		Type  string    `json:"type"`
		Alive bool      `json:"alive"`
		Date  string    `json:"date,omitempty"`
		Run   uuid.UUID `json:"run,omitempty"`
	}

	ErrorResponse struct {
		Type   string    `json:"type"`
		Status int       `json:"status"`
		Error  string    `json:"error,omitempty"`
		Run    uuid.UUID `json:"run,omitempty"`
	}
)

var (
	server *http.ServeMux
	id     uuid.UUID
)

func init() {
	id, _ = uuid.NewRandom()
	server = NewServer()

	f.HTTP("server", request)
}

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/alive", alive)

	return mux
}

func request(w http.ResponseWriter, r *http.Request) {
	server.ServeHTTP(w, r)
}

// error provides for a standard formatted response to the endpoint based on the
// status code and error message given.
func newError(w http.ResponseWriter, e interface{}, c int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(c)

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		// Fallback Message if JSON encoding fails
		fmt.Fprintln(w, "{\"type\":\"error\",\"status\":500,\"error\":\"unknown\"}")
	}
}

// errorNotFound provides a 404 response with a JSON body reporting an error
// where the page could not be found, either intensionally or otherwise.
func errorNotFound(w http.ResponseWriter) {
	e := ErrorResponse{
		Type:   "error",
		Status: http.StatusNotFound,
		Error:  "File Not Found",
		Run:    id,
	}

	newError(w, e, http.StatusNotFound)
}

// errorNotFound provides a 404 response with a JSON body reporting an error
// where the page could not be found, either intensionally or otherwise.
func errorInternalServer(w http.ResponseWriter, err error) {
	e := ErrorResponse{
		Type:   "error",
		Status: http.StatusInternalServerError,
		Error:  fmt.Sprintf("Internal Server Error: %s", err),
		Run:    id,
	}

	newError(w, e, http.StatusNotFound)
}

// root provides a base handler which handles requests to the root of the
// endpoint, or for any URL which is unknown, so it should handle 404 responses
// to invalid requests, as well as its own.
func root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorNotFound(w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintln(w, "{}")
}

// alive provides a simple response to verify that the endpoint is alive and
// service responses back to the client, including a date/time to validate the
// request is not cached
func alive(w http.ResponseWriter, r *http.Request) {
	a := AliveResponse{
		Type:  "alive",
		Alive: true,
		Date:  time.Now().Format(time.RFC3339),
		Run:   id,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		errorInternalServer(w, err)
	}
}
