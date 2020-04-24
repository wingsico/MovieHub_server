package tests

import (
	"github.com/go-playground/assert/v2"
	"github.com/wingsico/movie_server/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovieQuery(t *testing.T) {
	router := routes.InitRouter()
	w := httptest.NewRecorder() // httptest 记录器
	req, _ := http.NewRequest(http.MethodGet, "/api/movie/222", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

