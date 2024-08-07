package handlers

import (
    "html/template"
    "net/http"
    "os"
    "sync"
    "log"
)

var logMutex sync.Mutex
var templates = template.Must(template.ParseFiles("templates/base.html", "templates/logs.html"))

// ServeIndex serves the index.html file
func ServeIndex(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
}

// HandleLogs serves the proxy log file
func HandleLogs(w http.ResponseWriter, r *http.Request) {
    logMutex.Lock()
    defer logMutex.Unlock()
    logs, err := os.ReadFile("proxy.log")
    if err != nil {
        ServeCustomErrorPage(w, r, http.StatusInternalServerError)
        return
    }
    templates.ExecuteTemplate(w, "logs.html", string(logs))
}

// ServeCustomErrorPage serves custom error pages
func ServeCustomErrorPage(w http.ResponseWriter, r *http.Request, statusCode int) {
    w.WriteHeader(statusCode)
    http.ServeFile(w, r, "static/"+http.StatusText(statusCode)+".html")
}

// LogRequest logs incoming requests
func LogRequest(r *http.Request) {
    logMutex.Lock()
    defer logMutex.Unlock()
    log.Printf("Received request: %s %s from %s\n", r.Method, r.URL.String(), r.RemoteAddr)
}

// LogResponse logs response status codes
func LogResponse(statusCode int) {
    logMutex.Lock()
    defer logMutex.Unlock()
    log.Printf("Response status: %d\n", statusCode)
}
