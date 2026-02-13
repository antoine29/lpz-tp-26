# Pelias

## Setting up the ES instance

- run es
docker compose up -d elasticsearch
curl -X GET "http://localhost:9200/_cluster/health?pretty"

- create the es schema (wait for some minutes before running since es service can take some minutes to sping up)
docker compose run schema

## Importing ES data

- docker compose run openstreetmap

- docker compose run whosonfirst

## Importing other data 

- docker compose run placeholder_importer


## Running the main api

- docker compose up -d api

curl "http://localhost:3100/v1/search?text=yanacocha&boundary.country=BO"


todo:
if boundary.country=BO is not passed on req, then hsl api will close the search to FIN, thus no pelias match will be returned. we need to either fork hsl api to change the default country, or to fork ui to use the req parameter
