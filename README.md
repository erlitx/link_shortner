# URL Shortener

A simple URL shortening service written in Go using the Chi router and PostgreSQL.  
It provides API endpoints to create and resolve short links, all containerized with Docker.

---

## Features

- Generate short URLs from long ones
- Redirect short links to original URLs
- RESTful API with JSON
- In-memory LRU cache for performance
- PostgreSQL for persistent storage
- Health and profiling endpoints via `pprof`

---

## Tech Stack

- Go 1.23
- Chi router
- PostgreSQL
- Migrate CLI 

---

## API Endpoints

### Create a short link

**POST** `/api/v1/create_shortlink`

#### Request Body:

```json
{
  "rawUrl": "https://example.com"
}
```
## Start an App
Provide .env file with following env variables:

```
POSTGRES_USER
POSTGRES_PASSWORD
POSTGRES_PORT
POSTGRES_HOST
POSTGRES_DB_NAME

APP_NAME = LINK-SHORTNER
APP_VERSION = 0.3.0
```

Use `make` tool to run an app and all other comands.

If not yet installed:

 `sudo apt update
sudo apt install build-essential
`

### Run services in the docker

`make up`

### Install migrate tool and make first migration
`make migrate-install`

`make migrate-up`


### Run the app

`make run`