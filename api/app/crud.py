from bson import ObjectId
from pymongo.database import Database

from app.models import Resource, ResourceInDB, Snapshot, SnapshotInDB

from .config import config


def create_new_resource(resource: Resource, db: Database) -> ResourceInDB:
    new_resource = db[config.resources_collection_name].insert_one(resource.dict())
    created_resource = db[config.resources_collection_name].find_one(
        {"_id": new_resource.inserted_id}
    )
    return created_resource


def read_many_resources(db: Database) -> list[ResourceInDB]:
    return [ResourceInDB(**row) for row in db[config.resources_collection_name].find()]


def read_single_resource(resource_id: str, db: Database) -> Resource:
    return db[config.resources_collection_name].find_one({"_id": ObjectId(resource_id)})


def create_new_snapshot(
    resource_id: str, snapshot: Snapshot, db: Database
) -> SnapshotInDB:
    # Create new snapshot
    new_snapshot = db[config.snapshots_collection_name].insert_one(
        {**snapshot.dict(), **{"resource_id": resource_id}}
    )
    created_snapshot = db[config.snapshots_collection_name].find_one(
        {"_id": new_snapshot.inserted_id}
    )

    # Set latest_snapshot field in resource document to new snapshot (embedded pattern)
    db[config.resources_collection_name].update_one(
        {"_id": ObjectId(resource_id)}, {"$set": {"latest_snapshot": created_snapshot}}
    )
    return created_snapshot


def read_many_snapshots(resource_id: str, db: Database) -> list[SnapshotInDB]:
    return [
        SnapshotInDB(**row)
        for row in db[config.snapshots_collection_name].find(
            {"resource_id": resource_id}
        )
    ]
