# LPZ trip planner application

This is a web based trip planner application to provide information about the public transport system on La Paz city. It uses and adapts DigiTransit services and components.

## Components:

- otp

This the main routing service. It uses a DigiTransit OpenTripPlanner fork container image.

- hsl-map-server

Service to provide map tiles to be used on the FE application. It is a DigiTransit forked repo.

Repo at `https://github.com/antoine29/hsl-map-server`

- pelias (deprecated)

DigiTransit geocoding service to used on the FE application. Deprecated in favour of geocoder service since pelias requires more resources and it is harder to setup (harder than geocoder motis)

- geocoder

Geocoding service to be used on the FE application. It uses a Motis instance and an adapter service to translate DigiTransit pelias requests into Motis

- ui

Main FE web application, it uses above otp, map-server, and geocoder as BE dependencies.

## Local development setup

Refer to each folder readme.md file to setup each project locally

## Production deployment

ToDo: write K8S manifests

