from fastapi import FastAPI, status

from .crud import create_snapshot, read_snapshot
from .database import database
from .models import CreateSnapShot, SnapShot, SnapShotQuery

app = FastAPI()


@app.get("/")
def hello_world():
    return "Hello, world!"


@app.get("/snapshot", response_model=list[SnapShot])
async def get_snapshot(snapshot_query: SnapShotQuery):
    return await read_snapshot(database, snapshot_query)



@app.post("/snapshot", response_model=SnapShot, status_code=status.HTTP_201_CREATED)
async def post_snapshot(snapshot: CreateSnapShot):
    new_snapshot = await create_snapshot(database, snapshot)
    return new_snapshot
