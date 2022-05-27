from sqlalchemy.orm import Session

from . import models, schemas


class ResourceNotFound(Exception):
    ...


def get_resources(db: Session) -> list[models.Resource]:
    return db.query(models.Resource).all()


def create_new_resource(db: Session, obj_in: schemas.ResourceBase) -> models.Resource:
    resource = models.Resource(**obj_in.dict())
    db.add(resource)
    db.commit()
    db.refresh(resource)
    return resource
