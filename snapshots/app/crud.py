from motor.motor_asyncio import AsyncIOMotorDatabase, AsyncIOMotorCollection

from .models import SnapShot


async def create_snapshot(db: AsyncIOMotorDatabase, snapshot: SnapShot):
    snapshots: AsyncIOMotorCollection = db.get_collection("snapshots")
    snapshot = await snapshots.insert_one(snapshot.dict())
    new_snapshot = await snapshots.find_one({"_id": snapshot.inserted_id})
    return new_snapshot
