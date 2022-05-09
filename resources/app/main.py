"""Hello, world"""

from fastapi import FastAPI

from . import models, schemas, database, crud
from sqlalchemy.orm import Session

from fastapi import Depends

app = FastAPI()


@app.get("/")
def root():
    """Return home"""
    return "Hello, world!"


def get_db():
    db = database.SessionLocal()
    try:
        yield db
    finally:
        db.close()


@app.get("/resource", response_model=list[schemas.Resource])
def get_resource(db: Session = Depends(get_db)) -> list[models.Resource]:
    resources = crud.get_resources(db)
    return resources


@app.post("/resource", response_model=schemas.Resource)
def post_resource(
    resource: schemas.ResourceBase, db: Session = Depends(get_db)
) -> models.Resource:
    return crud.create_new_resource(db, resource)
