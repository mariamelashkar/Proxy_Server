package main

import (
    "flag"
    "log"
    "net/http"
    "os"

    "proxy/routers"
    "proxy/middlewares"
)

func main() {
    // Open log file
    logFile, err := os.OpenFile("proxy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()
    log.SetOutput(logFile)

    // Define command-line flag for the server address
    addr := flag.String("addr", "127.0.0.1:8080", "proxy address")
    flag.Parse()

    // Create the router
    r := routers.NewRouter()

    // Start the server with rate limiting middleware
    log.Println("Starting proxy server on", *addr)
    if err := http.ListenAndServe(*addr, middlewares.RateLimit(r)); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
