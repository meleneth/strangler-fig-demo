package main

import (
    "context"
    "errors"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "mobilis.local/candidate/internal/app"
)

func run() error {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()
    handler, cleanup, err := app.New(ctx)
    if err != nil { return err }
    defer cleanup()
    port := os.Getenv("PORT")
    if port == "" { port = "8080" }
    server := &http.Server{Addr: ":" + port, Handler: handler, ReadHeaderTimeout: 5*time.Second}
    result := make(chan error, 1)
    go func() { result <- server.ListenAndServe() }()
    select {
    case err := <-result:
        if !errors.Is(err, http.ErrServerClosed) { return err }
    case <-ctx.Done():
        shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        if err := server.Shutdown(shutdownCtx); err != nil { _ = server.Close(); return err }
        if err := <-result; !errors.Is(err, http.ErrServerClosed) { return err }
    }
    return nil
}

func main() {
    if err := run(); err != nil { log.Fatal(err) }
}
