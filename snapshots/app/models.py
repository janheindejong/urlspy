from datetime import datetime

from bson import ObjectId
from pydantic import AnyHttpUrl, BaseModel, Field


class ObjectIdField(ObjectId):
    @classmethod
    def __get_validators__(cls):
        yield cls.validate

    @classmethod
    def validate(cls, v):
        if not ObjectId.is_valid(v):
            raise ValueError("Invalid objectid")
        return ObjectId(v)

    @classmethod
    def __modify_schema__(cls, field_schema):
        field_schema.update(type="string")


class CreateSnapShot(BaseModel):
    url: AnyHttpUrl
    datetime: datetime
    response: int
    body: str


class SnapShot(CreateSnapShot):
    id: ObjectIdField = Field(default_factory=ObjectIdField, alias="_id")

    class Config:
        allow_population_by_field_name = True
        arbitrary_types_allowed = True
        json_encoders = {ObjectId: str}
