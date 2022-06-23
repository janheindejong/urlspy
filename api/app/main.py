from fastapi import FastAPI

from .api import router
from .database import connect_to_database, disconnect_from_database

app = FastAPI()

app.include_router(router)

app.add_event_handler("startup", connect_to_database)
app.add_event_handler("shutdown", disconnect_from_database)
