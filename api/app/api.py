from fastapi import APIRouter, Depends
from pymongo.database import Database

from app.database import get_application_database

from .crud import (
    create_resource,
    create_snapshot,
    read_resources,
    read_single_resource,
    read_snapshots,
)
from .models import Resource, ResourceInDB, Snapshot, SnapshotInDB

router = APIRouter()


@router.get("/")
def root():
    return "Hello, world!"


@router.get("/resource", response_model=list[ResourceInDB])
def get_resources(db: Database = Depends(get_application_database)):
    return read_resources(db)


@router.post("/resource", response_model=ResourceInDB)
def post_resource(resource: Resource, db: Database = Depends(get_application_database)):
    return create_resource(resource, db)


@router.get("/resource/{resource_id}", response_model=ResourceInDB)
def get_resource(resource_id: str, db: Database = Depends(get_application_database)):
    return read_single_resource(resource_id, db)


@router.get("/resource/{resource_id}/snapshots", response_model=list[SnapshotInDB])
def get_resource_snapshots(
    resource_id: str, db: Database = Depends(get_application_database)
):
    return read_snapshots(resource_id, db)


@router.post("/resource/{resource_id}/snapshots", response_model=SnapshotInDB)
def post_resource_snapshots(
    resource_id: str,
    snapshot: Snapshot,
    db: Database = Depends(get_application_database),
):
    return create_snapshot(resource_id, snapshot, db)
