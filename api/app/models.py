from datetime import datetime
from typing import Optional

from bson import ObjectId
from pydantic import AnyHttpUrl, BaseModel, EmailStr, Field


class PydanticObjectId(ObjectId):
    @classmethod
    def __get_validators__(cls):
        yield cls.validate

    @classmethod
    def validate(cls, value):
        if not ObjectId.is_valid(value):
            raise ValueError("Invalid objectid")
        return ObjectId(value)

    @classmethod
    def __modify_schema__(cls, field_schema):
        field_schema.update(type="string")


class DBModelMixin(BaseModel):
    id: PydanticObjectId = Field(..., alias="_id")

    class Config:
        allow_population_by_field_name = True
        json_encoders = {ObjectId: str}
        arbitrary_types_allowed = True


class Resource(BaseModel):
    url: AnyHttpUrl
    email: Optional[EmailStr] = None


class ResourceInDB(Resource, DBModelMixin):
    ...


class Snapshot(BaseModel):
    datetime: datetime
    status_code: int
    body: str


class SnapshotInDB(Snapshot, DBModelMixin):
    resource_id: PydanticObjectId = Field(..., alias="_id")
