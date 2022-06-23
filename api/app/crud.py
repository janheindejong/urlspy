from bson import ObjectId
from fastapi import HTTPException
from pymongo.database import Database

from app.models import Resource, ResourceInDB, Snapshot, SnapshotInDB

from .config import config


def create_new_resource(resource: Resource, db: Database) -> ResourceInDB:
    new_resource = db[config.resources_collection_name].insert_one(resource.dict())
    created_resource = db[config.resources_collection_name].find_one(
        {"_id": new_resource.inserted_id}
    )
    assert isinstance(created_resource, dict), "Resource not successfully created"
    return ResourceInDB(**created_resource)


def read_many_resources(db: Database) -> list[ResourceInDB]:
    return [ResourceInDB(**row) for row in db[config.resources_collection_name].find()]


def read_single_resource(resource_id: str, db: Database) -> ResourceInDB:
    document = _read_single_resource_raw(resource_id, db)
    return ResourceInDB(**document)


def create_new_snapshot(
    resource_id: str, snapshot: Snapshot, db: Database
) -> SnapshotInDB:
    # Verify if resource exists
    _read_single_resource_raw(resource_id, db)

    # Create new snapshot
    new_snapshot = db[config.snapshots_collection_name].insert_one(
        {**snapshot.dict(), **{"resource_id": resource_id}}
    )
    created_snapshot = db[config.snapshots_collection_name].find_one(
        {"_id": new_snapshot.inserted_id}
    )
    assert isinstance(created_snapshot, dict), "Snapshot not successfully created"

    # Set latest_snapshot field in resource document to new snapshot (embedded pattern)
    db[config.resources_collection_name].update_one(
        {"_id": ObjectId(resource_id)}, {"$set": {"latest_snapshot": created_snapshot}}
    )
    return SnapshotInDB(**created_snapshot)


def read_many_snapshots(resource_id: str, db: Database) -> list[SnapshotInDB]:
    return [
        SnapshotInDB(**row)
        for row in db[config.snapshots_collection_name].find(
            {"resource_id": resource_id}
        )
    ]


def _read_single_resource_raw(resource_id: str, db: Database) -> dict:
    document = db[config.resources_collection_name].find_one(
        {"_id": ObjectId(resource_id)}
    )
    if not document:
        raise HTTPException(status_code=404, detail=f"Resource {resource_id} not found")
    return document
