package main

import (
    "flag"
    "log/slog"
    "net/http"
    "os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {
    addr := flag.String("addr", ":4000", "HTTP network address")
    flag.Parse()

    mux := http.NewServeMux()

    mux.HandleFunc("GET /{$}", home)
    mux.HandleFunc("GET /dreams/{dbName}/{description}", dreamsView)

    logger.Info("starting server", "addr", *addr)

    err := http.ListenAndServe(*addr, mux)
    logger.Error(err.Error())
    os.Exit(1)
}
