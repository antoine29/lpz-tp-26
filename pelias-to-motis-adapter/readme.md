## Digitransit UI - Motis  API adapter

This is an adapter service to translate Digitransit UI API requests into requests that can be served by Motis

```
docker build . -t adapter
docker run -d \
  --name adapter \
  -p 8084:8084 \
  -e PORT=8084 \
  -e MOTIS_HOST=http://172.17.0.4:8083 \
  adapter
``` 

## Test

```
curl 'http://localhost:8084/v1/search?text=multicine'
```

