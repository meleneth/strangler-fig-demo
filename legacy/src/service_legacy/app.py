from contextlib import asynccontextmanager
from fastapi import FastAPI
from .routes import register
from .database import get_engine
from .telemetry import configure, instrument_app


def create_app(*, telemetry=True):
    provider = configure(get_engine()) if telemetry else None
    @asynccontextmanager
    async def lifespan(app):
        try:
            yield
        finally:
            if get_engine.cache_info().currsize:
                get_engine().dispose()
            if provider is not None:
                provider.shutdown()

    app = FastAPI(lifespan=lifespan)

    @app.get("/health")
    def health():
        return {"status": "ok"}

    register(app)
    if telemetry:
        instrument_app(app)
    return app
