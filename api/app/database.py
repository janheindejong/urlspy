from loguru import logger
from pymongo import MongoClient
from pymongo.database import Database

from .config import config


def get_db_uri():
    return "mongodb://{username}:{password}@{host}:{port}".format(  # pylint: disable=consider-using-f-string
        username=config.mongo_db_username,
        password=config.mongo_db_password,
        host=config.mongo_db_host,
        port=config.mongo_db_port,
    )


class DbConnHandler:
    client: MongoClient


db = DbConnHandler()


def get_application_database() -> Database:
    return db.client[config.database_name]


def connect_to_database() -> None:
    db.client = MongoClient(get_db_uri(), tz_aware=True)
    logger.info("Successfully connected DB")


def disconnect_from_database() -> None:
    db.client.close()
    logger.info("Successfully disconnected DB")
