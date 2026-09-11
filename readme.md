# LPZ trip planner application

This is a web based trip planner application to provide information about the public transport system on La Paz city. It uses and adapts DigiTransit services and components.

## Components:

- otp

This the main routing service. It uses a DigiTransit OpenTripPlanner fork container image.

it uses GTFS datasets along osm files

- hsl-map-server

Service to provide map tiles to be used on the FE application. It is a DigiTransit forked repo.

it uses mbtiles and geojson as data sources

Forked repo at `https://github.com/antoine29/hsl-map-server`

- motis

Used as geocoding service to be used on the FE application. It has to be used along an adapter service

- pelias-to-motis-adapter

Service used to adapt FE requests to pelias, to be answered by Motis

- digitransit-ui

Main FE web application, it uses above otp, map-server, and motis adapter as BE dependencies.

## Local development setup

Refer to each folder `readme.md` file to setup each project locally

Forked repo at `https://github.com/antoine29/digitransit-ui`

## Production deployment

ToDo: write K8S manifests

