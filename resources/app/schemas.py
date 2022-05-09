from pydantic import AnyHttpUrl, BaseModel


class ResourceBase(BaseModel):

    path: AnyHttpUrl


class Resource(ResourceBase):
    id: int

    class Config:
        orm_mode = True
