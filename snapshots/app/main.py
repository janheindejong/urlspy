from fastapi import FastAPI, status

from .crud import create_snapshot
from .database import database
from .models import CreateSnapShot, SnapShot

app = FastAPI()


@app.get("/")
def hello_world():
    return "Hello, world!"


@app.post("/snapshot", response_model=SnapShot, status_code=status.HTTP_201_CREATED)
async def post_snapshot(snapshot: CreateSnapShot):
    new_snapshot = await create_snapshot(database, snapshot)
    return new_snapshot
