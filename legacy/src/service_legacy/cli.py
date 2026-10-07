import os
import uvicorn
from .app import create_app


def main():
    uvicorn.run(create_app(), host="0.0.0.0", port=int(os.environ.get("PORT", "8000")))
