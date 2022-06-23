from pymongo.database import Database

from app.models import Resource, ResourceInDB

resource_collection_name = "resources"


def create_resource(resource: Resource, db: Database):
    db[resource_collection_name].insert_one(resource.dict())


def read_resources(db: Database):
    return [ResourceInDB(**row) for row in db[resource_collection_name].find()]
