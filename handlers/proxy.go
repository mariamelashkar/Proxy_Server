package handlers

import (
    "io"
    "net/http"
    "net/url"
)

// HandleRequestAndRedirect handles incoming proxy requests
func HandleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
    LogRequest(r)

    requestURL := r.URL.Query().Get("url")
    if requestURL == "" {
        ServeCustomErrorPage(w, r, http.StatusBadRequest)
        return
    }

    proxyURL, err := url.Parse(requestURL)
    if err != nil {
        ServeCustomErrorPage(w, r, http.StatusBadRequest)
        return
    }

    proxyReq, err := http.NewRequest(r.Method, proxyURL.String(), r.Body)
    if err != nil {
        ServeCustomErrorPage(w, r, http.StatusInternalServerError)
        return
    }

    // Remove headers that could reveal client information
    proxyReq.Header.Del("User-Agent")
    proxyReq.Header.Del("Referer")
    proxyReq.Header.Del("X-Forwarded-For")
    proxyReq.Header.Del("X-Real-IP")

    for name, values := range r.Header {
        if name != "User-Agent" && name != "Referer" && name != "X-Forwarded-For" && name != "X-Real-IP" {
            for _, value := range values {
                proxyReq.Header.Add(name, value)
            }
        }
    }

    // Optionally, set your own IP address or a placeholder in X-Forwarded-For to hide client's IP
    proxyReq.Header.Set("X-Forwarded-For", "127.0.0.1")

    client := &http.Client{}
    resp, err := client.Do(proxyReq)
    if err != nil {
        ServeCustomErrorPage(w, r, http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)

    LogResponse(resp.StatusCode)
}
