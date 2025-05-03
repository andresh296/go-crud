package middleware

import (
    "bytes"
    "net/http"
    "net/http/httptest"

    "github.com/gin-gonic/gin"
)

func createMockContext(method, path string, body []byte) (*gin.Context, *httptest.ResponseRecorder) {
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    req, _ := http.NewRequest(method, path, bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    c.Request = req

    return c, w
}