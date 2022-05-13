"""SQLAlchemy models"""

from __future__ import annotations

from typing import Any

from sqlalchemy import Column, Integer, String
from sqlalchemy.orm import declarative_base

Base: Any = declarative_base()


class Resource(Base):
    __tablename__ = "resource"

    id = Column(Integer, primary_key=True)
    url = Column(String)
