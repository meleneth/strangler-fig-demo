package app

import (
    "context"
    "net/http"
    "time"
    "os"
    "github.com/jackc/pgx/v5/pgxpool"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "github.com/exaring/otelpgx"
    "log"
)

type Dependencies struct {
    HTTPClient *http.Client
    DB *pgxpool.Pool
}

func New(ctx context.Context) (http.Handler, func(), error) {
    shutdown, err := configureTelemetry(ctx)
    if err != nil { return nil, nil, err }
    var pool *pgxpool.Pool
    cleanup := func() {
        if pool != nil { pool.Close() }
        shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        if err := shutdown(shutdownCtx); err != nil { log.Printf("telemetry shutdown: %v", err) }
    }
    cfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
    if err != nil { cleanup(); return nil, nil, err }
    cfg.ConnConfig.Tracer = otelpgx.NewTracer()
    pool, err = pgxpool.NewWithConfig(ctx, cfg)
    if err != nil { cleanup(); return nil, nil, err }
    pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
    err = pool.Ping(pingCtx)
    cancel()
    if err != nil { cleanup(); return nil, nil, err }
    deps := Dependencies{HTTPClient: &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}, DB: pool}
    deps.HTTPClient.Timeout = 10*time.Second
    return Handler(deps), cleanup, nil
}

// Handler constructs HTTP behavior without opening connections or exporters.
func Handler(deps Dependencies) http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _, _ = w.Write([]byte(`{"status":"ok"}`))
    })
    register(mux, deps)
    return otelhttp.NewHandler(mux, "http")
}
