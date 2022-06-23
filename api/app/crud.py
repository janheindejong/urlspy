from bson import ObjectId
from pymongo.database import Database

from app.models import Resource, ResourceInDB, Snapshot, SnapshotInDB

resource_collection_name = "resources"
snapshot_collection_name = "snapshots"


def create_resource(resource: Resource, db: Database) -> ResourceInDB:
    new_resource = db[resource_collection_name].insert_one(resource.dict())
    created_resource = db[resource_collection_name].find_one(
        {"_id": new_resource.inserted_id}
    )
    return created_resource


def read_resources(db: Database) -> list[ResourceInDB]:
    return [ResourceInDB(**row) for row in db[resource_collection_name].find()]


def read_single_resource(resource_id: str, db: Database) -> Resource:
    return db[resource_collection_name].find_one({"_id": ObjectId(resource_id)})


def create_snapshot(resource_id: str, snapshot: Snapshot, db: Database) -> SnapshotInDB:
    new_snapshot = db[snapshot_collection_name].insert_one(
        {**snapshot.dict(), **{"resource_id": resource_id}}
    )
    created_snapshot = db[snapshot_collection_name].find_one(
        {"_id": new_snapshot.inserted_id}
    )
    return created_snapshot


def read_snapshots(resource_id: str, db: Database) -> list[SnapshotInDB]:
    return [
        SnapshotInDB(**row)
        for row in db[snapshot_collection_name].find({"resource_id": resource_id})
    ]
