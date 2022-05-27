import os

import motor.motor_asyncio


def get_url():
    return "mongodb://%s:%s@%s" % (
        os.getenv("APP_CONFIG_DB_USER", "mongodb"),
        os.getenv("APP_CONFIG_DB_PASSWORD", "mongodb"),
        os.getenv("APP_CONFIG_DB_HOST", "mongo:27017"),
    )


client = motor.motor_asyncio.AsyncIOMotorClient(get_url())

database: motor.motor_asyncio.AsyncIOMotorDatabase = client.snapshots
