"""SQL database definition"""
import os

from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker


def get_url():
    return "postgresql://%s:%s@%s" % (
        os.getenv("APP_CONFIG_DB_USER", "postgres"),
        os.getenv("APP_CONFIG_DB_PASSWORD", "postgres"),
        os.getenv("APP_CONFIG_DB_HOST", "postgres:5432"),
    )


engine = create_engine(get_url())

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
