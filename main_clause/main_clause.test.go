package main_clause

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBeefSummary(t *testing.T) {
	router := gin.Default()
	router.GET("/food/mainclause", getDataBeef)

	req, _ := http.NewRequest("GET", "/food/mainclause", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
