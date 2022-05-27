# URLStalker 

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
- [ ] Add Debug configurations for VSCode
- [x] Add persistent storage to snapshot app 
- [ ] Add read end-points to snapshot app
- [x] Add Scraper app 
- [ ] Add storing functionality to scraper app 
- [ ] Add e-mail functionality  
- [ ] Add deployment pipeline 



