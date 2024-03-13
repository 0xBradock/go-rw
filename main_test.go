package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponse(t *testing.T, expected, actual int) {
	require.Equal(t, expected, actual)
}

func TestHealthCheck(t *testing.T) {
	s := NewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/health", nil)
	response := executeRequest(req, s)

	checkResponse(t, http.StatusOK, response.Code)
}
