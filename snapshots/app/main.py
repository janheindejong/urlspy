from fastapi import FastAPI

from .models import SnapShot
from .crud import create_snapshot
from .database import database

app = FastAPI()


@app.get("/")
def hello_world():
    return "Hello, world!"


@app.post("/snapshot")
async def post_snapshot(snapshot: SnapShot):
    new_snapshot = await create_snapshot(database, snapshot)
    return new_snapshot
