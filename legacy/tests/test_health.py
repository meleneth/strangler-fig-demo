from fastapi.testclient import TestClient
from service_legacy.app import create_app


def test_health():
    with TestClient(create_app(telemetry=False)) as client:
        response = client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "ok"}
    assert response.headers["content-type"] == "application/json"
