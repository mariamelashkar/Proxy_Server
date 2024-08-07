package middlewares

import (
    "net/http"
    "time"
)

var rateLimiter = time.Tick(time.Second / 10) // 10 requests per second

// rateLimit is middleware that limits the rate of incoming requests
func RateLimit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        <-rateLimiter
        next.ServeHTTP(w, r)
    })
}
