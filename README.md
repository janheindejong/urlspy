# URL Stalker 

[![Code style: black](https://img.shields.io/badge/code%20style-black-000000.svg)](https://github.com/psf/black)

App to scrape URLs at regular interval

## Tools & techniques

This application uses the following tools & techniques: 

- Micro-services architecture with **Docker**, **Docker Compose** & **Docker Secrets**
- Optimized and controllable development environment using **Visual Studio Code devcontainers** 
- Persistent storage using **MongoDB**
- A REST API for the database using **Python** and **FastAPI** 
- A scraper, built in **Go**
- E-mail functionality, using **SMTP**, and Go's **`net/smtp`** package
- Dependency management for Python with **Poetry**

## Architecture 

The app consists of two parts: 

* A MongoDB, with a Python based API 
* A Go-based scraper, that runs periodically

![Architecture](docs/architecture.svg)

Initially, I envisioned a setup with both an SQL DB, and a MongoDB. THat was a bit overkill, so now the app has one database within a MongoDB - **`urlstalker`** - with two collections: 

- **`resources`** - contains resources that need to be tracked, including metadata and latest snapshot
- **`snapshots`** - contains full list of all snapshots, including link to resource id

The resource model has the following format: 

```json
{
    "_id": 123456789,
    "url": "https://example.com", 
    "email": "john@doe.com", 
    "latest_snapshot": null  // One snapshot, optional
} 
```

Snapshots have the following shape: 

```json 
{
    "_id": 123456789, 
    "datetime": "2022-01-20T00:00:00.0+00:00", 
    "status_code": 200, 
    "body": "<html><body>Hello, world!<body/><html/>",
    "resource_id": 123456789
}
```

## Development 

For the development, I wanted to experiment with the use of VSCode's development containers. If you have the plugin enabled in VSCode, and have Docker running on your machine, VSCode should automatically launch in the development container for the application you want to work on. Next to the development container, it also starts containers for all the other apps in the stack. 

The setup is such that if you open any of the application folders (i.e. `./api` or `./scraper`), it spins up a development container named `<appname>-dev` with all the right tooling, and mounts the full project to `/workspace` in the container. Additionally, it launches (or re-uses already running) services from the stack it needs (e.g. the database). Pretty neat. 

For more reading on why development containers are promising, read [this](https://www.infoq.com/articles/devcontainers/). 

## Connecting to Gmail 

Currently, I'm using SMTP for connection to gmail. To be able to connect to Gmail, I've had to enable 2FA on the account, and create an app password. In future, I might want to use the Gmail REST-API. For now, the account name and app password are stored in files pointed to by the docker compose files: 

* Development enviroment: 
    - Email account: `./secrets/email_account.txt` 
    - Email password: `./secrects/email_password.txt` 
* For production environment 
    - Email account: `/opt/urlstalker/email_account.txt`
    - Email password: `/opt/urlstalker/email_password.txt` 

## To do

- [x] Create docker compose file for deployment 
- [x] Add Scraper app 
- [x] Add storing functionality to scraper app 
- [x] Improve devcontainer workflow to allow to work on multiple parts at same time
- [x] Add debug configurations for VSCode
- [x] Add e-mail functionality  
- [ ] Improve Go project structure; read [here](https://tutorialedge.net/golang/go-project-structure-best-practices/)
- [ ] Add DELETE endpoint to database API
- [ ] Unittests (yes, I've been a bad boi)
- [ ] Add deployment pipeline 
- [ ] Add user-defineable tags to user stories, to make nicer e-mail subjects
