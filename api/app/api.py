from fastapi import APIRouter, Depends
from pymongo.database import Database

from app.database import get_application_database

from .crud import create_resource, read_resources
from .models import Resource, ResourceInDB

router = APIRouter()


@router.get("/")
def root():
    return "Hello, world!"


@router.get("/resource", response_model=list[ResourceInDB])
def get_resource(db: Database = Depends(get_application_database)):
    return read_resources(db)


@router.post("/resource")
def post_resource(resource: Resource, db: Database = Depends(get_application_database)):
    create_resource(resource, db)
