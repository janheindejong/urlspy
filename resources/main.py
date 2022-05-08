"""Hello, world"""

from fastapi import FastAPI

app = FastAPI()


@app.get("/")
def root():
    """Return home"""
    return "Hello, world!"


print("Hello, world!")
