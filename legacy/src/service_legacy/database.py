import os
from functools import lru_cache
from sqlalchemy import create_engine
from sqlalchemy.engine import make_url
from sqlalchemy.orm import Session

# Credentials, host and database come exclusively from the connection.
@lru_cache(maxsize=1)
def get_engine():
    url = make_url(os.environ["DATABASE_URL"]).set(drivername="postgresql+psycopg")
    return create_engine(url, pool_pre_ping=True)


def get_session():
    return Session(get_engine())
