# Motis

Motis is a transit engine (alternative to OTP)

## First time local setup

Download and extract the latest motis release binary:

```
wget https://github.com/motis-project/motis/releases/latest/download/motis-linux-amd64.tar.bz2
tar xf motis-linux-amd64.tar.bz2
```

This should unzip a motis binary, tiles-profiles/, and ui/ folders

```
./motis config bolivia-260415.osm.pbf a_gtfs_zip_or_a_folder_with_gtfs_zips
./motis import
./motis server
```

The `config` command should create a default `config.yml` file

## Current setup

Motis is a full trip planner engine, it could replace multiple components of the HSL Trip Planner appplication:

- Routing engine (OTP is kept)
  The HSL FE appplication is highly attached to the OTP graphql query format. It was decided to keep using OTP.

- Tiles (hsl-map-server is kept)
  Though Motis UI renders tiles, the format of the HSL is not compatible with this, HSL FE requests png tiles, while motis returns pbf tiles. Also FE expects some geojson based responses from the tiles service. It was decided to keep using hsl-map-server.

- Geocoding (Pelias was replaced by motis)
  Since setting up a Pelias service is highly more harder than Motis, we are using Motis as a geocoding service, however a pelias-to-motis adapter service is required for FE

```
docker build . -t motis
docker run --name motis -d -p 8083:8083 motis
```

## Tests

```
 curl 'http://localhost:8083/api/v1/geocode?text=multicine'
```

## Notes:

- it looks like anytime you make a change on the `config.yml` file, you need to re-run `./motis import`.

