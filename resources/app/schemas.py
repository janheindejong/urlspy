from typing import Optional

from pydantic import AnyHttpUrl, BaseModel, EmailStr


class ResourceBase(BaseModel):

    url: AnyHttpUrl
    email_address: Optional[EmailStr]


class Resource(ResourceBase):
    id: int

    class Config:
        orm_mode = True
