package main

import (
	//"bytes"
	//"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	//	"strings"
	"testing"
)

func init() {
}

const TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGkiOiJ0ZXN0IiwiZXhwIjoxNDE1NDE5NjA3fQ.HqblPguDy-5ElaI8rIn3BZVbeXfKQx1UUMS-qD-PMncs4DIeRkMR2AFV1NlTccIqDyMhnypXwIp_YQQxNzMt0Q35oMeb3CjHB46qtZz7xoIvJo8Is0mYnbSEuAM8t0kq7KKZ4W8NKSR-eqGT1-0w0AuZuoLxxCCjVzw4QtanhtUE_-AsfM-79D4yOcvDd_OrACZ1N9YMhc3RyB8mAfAc_HKrvfrTxA7USsjnc6i_J_81s9gMG8ol-oF2Dhd4hD7XoXuqskIsdgWfbydT8SCV_bhKgN2nr4NljvxEouiTj5uERhLdPqKSs59Zg_n9fkpRQkrUhst3uTwo6vw2CT-mNA"

func TestHandlerLoginReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/login?api=test", nil)
	response := httptest.NewRecorder()

	handleLogin(response, request)

	assert.Equal(t, response.Code, http.StatusOK, "Response body did not contain expected %v:\n\tbody: %v")
}

func TestHandleLoginReturnsJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/login?api=test", nil)

	response := httptest.NewRecorder()

	handleLogin(response, request)

	ct := response.HeaderMap["Content-Type"][0]
	assert.Equal(t, ct, "application/json", "Content-Type does not equal 'application/json'")
}

func TestAuthMiddlewareAccessDenied(t *testing.T) {
	request, _ := http.NewRequest("POST", "/api", nil)
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	n := negroni.New()
	n.Use(AuthMiddleware())
	n.ServeHTTP(response, request)

	assert.Equal(t, response.Code, 401, "Access wasn't denied")
}

func TestAuthMiddlewareAccessiAllowed(t *testing.T) {
	request, _ := http.NewRequest("POST", "/api", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "BEARER "+TOKEN)

	response := httptest.NewRecorder()

	n := negroni.New()
	n.Use(AuthMiddleware())
	n.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK, "Access wasn't allowed")
}

func TestHandleApiReturnsJSON(t *testing.T) {
	request, _ := http.NewRequest("POST", "/api", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "BEARER "+TOKEN)

	response := httptest.NewRecorder()

	handleApi(response, request)

	ct := response.HeaderMap["Content-Type"][0]
	assert.Equal(t, ct, "application/json", "Content-Type does not equal 'application/json'")
}
