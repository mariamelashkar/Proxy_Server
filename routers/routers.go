package routers

import (
    "net/http"
    "github.com/gorilla/mux"
    "proxy/handlers"
)

// NewRouter creates a new Gorilla Mux router
func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/", handlers.ServeIndex).Methods("GET")
    router.HandleFunc("/proxy", handlers.HandleRequestAndRedirect).Methods("GET")
    router.HandleFunc("/logs", handlers.HandleLogs).Methods("GET")

    // Serve static files
    fileServer := http.FileServer(http.Dir("./static"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

    return router
}
