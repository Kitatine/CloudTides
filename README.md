# CloudTides

## Dependencies

- All tutorials require [Docker](https://docs.docker.com/docker-for-mac/install/)

## Setting up the app

Checkout the branch for a given tutorial, and run `docker-compose build` or `docker-compose up --build`

## Running the app on local

Run `docker-compose up` to see messages in the terminal. Run `docker-compose start` to run the app in the background.

## Issues
If encounter timeout error when setting up *docker-compose* for frontend then use the following in terminal before running *docker-compose*
```
export DOCKER_CLIENT_TIMEOUT=120
export COMPOSE_HTTP_TIMEOUT=120
```
