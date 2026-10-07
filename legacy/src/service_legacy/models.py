from datetime import date, datetime
from decimal import Decimal
from uuid import UUID
import sqlalchemy as sa
from sqlalchemy import ForeignKey
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column


class Base(DeclarativeBase):
    pass


class Item(Base):
    __tablename__ = "items"
    id: Mapped[int] = mapped_column(sa.Integer, primary_key=True, nullable=False, unique=False)
    name: Mapped[str] = mapped_column(sa.String(80), primary_key=False, nullable=False, unique=True)
