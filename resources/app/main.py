"""Hello, world"""

from fastapi import Depends, FastAPI, Response, status
from sqlalchemy.orm import Session

from . import crud, database, models, schemas

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


@app.delete("/resource")
def delete_resource(resource_id: int, db: Session = Depends(get_db)) -> models.Resource:
    try:
        crud.delete_resource(db, resource_id)
        return Response(status_code=status.HTTP_204_NO_CONTENT)
    except crud.ResourceNotFound:
        return Response(status_code=status.HTTP_404_NOT_FOUND)
