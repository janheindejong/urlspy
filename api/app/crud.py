from motor.motor_asyncio import AsyncIOMotorCollection, AsyncIOMotorDatabase

from .models import CreateSnapShot, SnapShotQuery, SnapShot


async def create_snapshot(db: AsyncIOMotorDatabase, snapshot: CreateSnapShot):
    snapshots: AsyncIOMotorCollection = db.get_collection("snapshots")
    snapshot = await snapshots.insert_one(snapshot.dict())
    new_snapshot = await snapshots.find_one({"_id": snapshot.inserted_id})
    return new_snapshot


async def read_snapshot(db: AsyncIOMotorDatabase, snapshot_query: SnapShotQuery) -> list[SnapShot]: 
    result = []
    snapshots: AsyncIOMotorCollection = db.get_collection("snapshots")
    cursor = snapshots.find({"url": snapshot_query.url})
    cursor.sort("datetime", -1).limit(snapshot_query.limit)
    async for document in cursor: 
        result.append(document)
    return result