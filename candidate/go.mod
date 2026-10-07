module mobilis.local/candidate

go 1.25.0

require (
	github.com/jackc/pgx/v5 v5.7.6
	go.opentelemetry.io/otel v1.38.0
	go.opentelemetry.io/otel/sdk v1.38.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.38.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.63.0
	github.com/exaring/otelpgx v0.9.3
)
