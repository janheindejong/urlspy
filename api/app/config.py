from pydantic import BaseSettings


class Config(BaseSettings):

    mongo_db_name: str = "urlstalker"
    mongo_db_host: str = "mongo"
    mongo_db_port: int = 27017
    mongo_db_username: str = "mongodb"
    mongo_db_password: str = "mongodb"

    class Config:
        env_prefix = "app_"


config = Config()
