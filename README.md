# URLStalker 

[![Code style: black](https://img.shields.io/badge/code%20style-black-000000.svg)](https://github.com/psf/black)

App to scrape URLs at regular interval

## Architecture 

* Resources: API for storing and retrieving URLs that need to be tracked, including any relevant metadata
* Snapshots: API for NoSQL DB with snapshots of URLs 
* Scraper: Go application for scraping the URLs 

![Architecture](docs/architecture.svg)

## Development 

For the development, I wanted to experiment with the use of VSCode's development containers. If you have the plugin enabled in VSCode, and have Docker running on your machine, VSCode should automatically launch in the development container for the application you want to work on. Next to the development container, it also starts containers for all the other apps in the stack. 

The environments of the three apps are completely separate and decoupled. The idea is that you can optimize your development environment to the app at hand. 

## To do

- [x] Create Snapshot app 
- [x] Create docker compose file for deployment 
- [x] Add persistent storage to snapshot app 
- [x] Add Scraper app 
- [x] Add storing functionality to scraper app 
- [x] Add delete endpoint
- [ ] Improve devcontainer workflow to allow to work on multiple parts at same time
- [ ] Add automatic test data to development PostgressDB
- [ ] Unittests (yes, I've been a bad boi)
- [X] Add Debug configurations for VSCode
- [ ] Add read end-points to snapshot app
- [ ] Add e-mail functionality  
- [ ] Add deployment pipeline 

## Techniques

This application uses the following techniques: 

- Docker and Docker compose 
- Development containers 
- Python and Go
- PostgreSQL and MongoDB 
- Alembic for SQL database migration
- SQLAlchemy ORM
- Python autoformatting using `black`, `isort` and `autoflake`
- Concurrent programming using `async` in Python, and Go-routines in Go
- FastAPI 
- Dependency management in Python with `poetry`
