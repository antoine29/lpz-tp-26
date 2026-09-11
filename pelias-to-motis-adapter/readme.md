## Digitransit UI - Motis  API adapter

This is an adapter service to translate Digitransit UI API requests into requests that can be served by Motis

```
docker build . -t adapter
docker run -d \
  --name adapter \
  -p 8084:8084 \
  -e PORT=8084 \
  -e MOTIS_HOST=http://172.17.0.3:8083 \
  adapter
``` 

## Test

```
curl 'http://localhost:8084/v1/search?text=multicine'
curl 'http://localhost:8084/v1/reverse?point.lat=-16.5105435&point.lon=-68.1220433'


curl 'http://localhost:8084/v1/reverse?\
  &point.lat=60.18601826737799\
  &point.lon=24.953212738037113\
  &boundary.circle.radius=0.1\
  &lang=en\
  &size=1\
  &layers=address\
  &zones=1' \
  -H 'accept: application/json'
```

