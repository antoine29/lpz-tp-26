# OTP

This uses a Digitransit's OpenTripPlanner fork. It's the main route calculations engine for the FE application.

On the `data/` folder you'll need an `osm.pbf` and `gtfs.zip` files like these to build the otp graph file:

```
LaPazBusGTFS_v1.zip
MiTelefericoGTFS.zip
bolivia-260415.osm.pbf
```

The `entrypoint.sh` file is used to hijack the default container entrypoint

## Image build:

```
docker build . -t otp
```

## Build the otp graph:

```
docker run --rm \
  --name otp-builder \
  -v ${PWD}/data:/var/otp/v3-prod/waltti \
  otp:latest build
```

## Run the server:

```
docker run -d \
  --name otp \
  -p 8081:8080 \
  -v ${PWD}/data:/var/otp/v3-prod/waltti \
  otp:latest server
```

## OTP debug tools

```
http://localhost:8081
http://localhost:8081/graphiql
```

