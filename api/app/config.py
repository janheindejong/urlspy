from pydantic import BaseSettings


class Config(BaseSettings):

    # MongoDB address and credentials
    mongo_db_host: str = "mongo"
    mongo_db_port: int = 27017
    mongo_db_username: str = "mongodb"
    mongo_db_password: str = "mongodb"

    # MongoDB database names
    database_name = "urlspy"
    resources_collection_name = "resources"
    snapshots_collection_name = "snapshots"

    class Config:  # pylint: disable=too-few-public-methods
        env_prefix = "app_"


config = Config()
