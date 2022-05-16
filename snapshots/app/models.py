from datetime import datetime

from pydantic import AnyHttpUrl, BaseModel


class SnapShot(BaseModel):
    url: AnyHttpUrl
    datetime: datetime
    response: int
    body: str
