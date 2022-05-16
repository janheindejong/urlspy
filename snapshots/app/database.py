import motor.motor_asyncio

MONGO_DETAILS = "mongodb://mongodb:mongodb@mongo:27017"

client = motor.motor_asyncio.AsyncIOMotorClient(MONGO_DETAILS)

database: motor.motor_asyncio.AsyncIOMotorDatabase = client.snapshots
