# INCROWD TEST TASK

This repository contains Sports News test task. Tech stack:
- Golang,
- MongoDb,
- Docker Compose,

## Project structure
There are the following folders and files in the project directory:
- `config` - yaml-config for the web-server,
- `db` - initial scripts for creating databases, collections, indexes etc.,
- `server` - source files for the web-server and Dockerfile for the image,
- `.env` - environment variables for services,
- `docker-compose.yaml` - file with compose instructions for starting the infrastructure for the application,
- `README.MD` - this file,


## Run the project
```sh
docker compose up
```

This command runs 2 services:
- Web-server at `localhost:8083`,
- MongoDB database at `localhost:27017`,

## Features

After the project is started, a web-server launches at `localhost:8083` (by default). It provides 2 endpoints:
- `/articles` - to list all the articles from the database. Supported query parameters:
    - `pageSize` - number of articles to retrieve (10 by default),
    - `pageNum` - number of the page from which we want to display articles (0 by default, every page contains `pageSize` records, list starts from 0),
- `/articles/{id}` - to list the article with given `id` from the database,

## Description

After the application is started, it reads the config. Config file contains the following sections:
- `storage` - database connection details,
- `http_server` - web-server connection details and inner params,
- `ecb_api` - Ecb API connection details,

Path to the config can be set with flag `-config` on startup, or read from `DEFAULT_API_CONFIG_PATH` environment variable (`/etc/incrowd-api/conf.yaml` by default).

After setting the connection to MongoDB database, application inits the caching layer. In the current version the cache is just a map with mutex and default ttl = 1 hour. In the next versions it can be replaced with services like Redis. The main requirement is that the caching layer should implement the Cache interface from file `articlecache.go`.

Then, application inits the poller that will poll the external API (`ecb` enpoints in this task). `ecbPoll` polls the endpoints in a goroutine, and as the next steps we can add as many pollers as we need (one poller per one external API).

For this task the strategy for "maintaining" the most actual (the last) 100 articles from the external API was chosen:
- at first, we retrieve lightweight data for the last 2 * 100 (last 100 articles with the margin) from the db and then populate cache:
    - key = article.id from ecb cms,
    - value = "composed" key from last modified timestamp for the article in our database and article.id (uuid) from our database, splitted by `:` symbol,
- then we poll the external api in the loop to get data in chunks divided by `pageSize`, and compare ids and lastModifiedDates of api articles with our data:
    - if the article is found in cache:
        - if its last modification time is the same, we can just skip this article and not update it in the database,
        - if the last modification time from the api for the article is newer than ours, we replace the old article in the database with the new one,
    - if the article is not found in cache, we will try to find it in the database and skip or update it if it is necessary,

This api polling operation repeats every 60 minutes (`ecb_api.period` in config).

It is important to mention that in the current version the most of the fields of the articles in database is filled by the dummy data (method `convertArticle` does the work), because there aren't any instructions on matching the fields between database and api fields at this time. Also, we don't use the `references` field now, but one of the next versions could implement this feature.

So, we are using the strategy to maintain the latest articles now, but it can be replaced with another strategy when we maintain the articles before some date, or for current sport season only, or for this year only etc.

Finally, web-server launches, so the application provides 2 endpoints for its users described in `Features` section above.

If there are some problems with the external api it won't lead to the application crash, because we still can provide the data from our database and poll the api later.

Graceful shutdown mechanism also implemented to prevent memory leaks after the application stopped.

## Database model

After the database started, the following entities created:
- new database (`incrowd` by default, set by `MONGO_INITDB_DATABASE` environment variable),
- new user (`incrowd` by default, set by `MONGO_USER` environment variable) with password (`password` by default, set by `MONGO_PASSWORD` environment variable),
- new collection `articles`,
- indexes for `id` and `source.sourceSystem` fields,

Important:
- the value for `id` field generates during the convertation of article model from api to database model. There is a non-zero probability that its value is not unique, but there aren't any additional checks for this at this time,
- there is a `source` object for every article in the database - it serves as a link between api model and database model, but we don't return it for our users:
    - `source.sourceSystem` - a name of the external api (`ecb`),
    - `source.sourceId` - id of the article in the ecb model,

## Tests

There are a couple of simple unit tests:
- `server/internal/validators/funcs_test.go` - table unit test for query parameter validator for positive ints (`ValidatePositiveInt`),
- `server/internal/articleserver/articleserver_test.go` - test for response status for `/article/{id}` endpoint with storage mock,

To run all tests:
```
go test ./... -v
```