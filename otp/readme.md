# OTP

This uses a Digitransit's OpenTripPlanner fork. It's the main route calculations engine for the FE application.

On the `data/` folder you'll need an `osm.pbf` and `gtfs.zip` files like these:

```
LaPazBusGTFS_v1.zip
MiTelefericoGTFS.zip
bolivia-260415.osm.pbf
```

The `entrypoint.sh` file is used to hijack the default container entrypoint

## Manual run using hsl image:

```
docker run \
  --name otp \
  -v ${PWD}/data:/var/otp/v3-prod/waltti \
  -v ${PWD}/entrypoint.sh:/entrypoint.sh \
  -v ${PWD}/build-config.json:/var/otp/v3-prod/waltti/build-config.json \
  -p 8081:8080 \
  -d \
  -e OTP_GRAPH_DIR=v3-prod/waltti \
  --entrypoint /entrypoint.sh \
  hsldevcom/opentripplanner:v2-prod-waltti
```

## Image run:

```
docker image build . -t otp
docker run -d \
  --name otp \
  -p 8081:8080 \
  -v ${PWD}/data:/var/otp/v3-prod/waltti \
  otp:latest
```

## OTP debug tools

```
http://localhost:8081
http://localhost:8081/graphiql
```

