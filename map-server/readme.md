# map server

- you need an osm.pbf file. get it from:
https://download.geofabrik.de/south-america/bolivia.html

- extract LPZ tiles with:
docker run -it -v $(pwd)/data:/data --entrypoint /bin/sh ghcr.io/systemed/tilemaker:master
# ./tilemaker /data/bolivia-latest.osm.pbf --bbox=-68.5,-17,-68,-16.3 --output lpz.mbtiles

you can see/edit the bbox on https://bboxfinder.com/#-17.239555,-69.206561,-15.751656,-67.281207

- test the mbtiles:
docker run -it -v $(pwd)/bolivia-shortbread-1.0.mbtiles:/data/bol.mbtiles -p 8081:8080 maptiler/tileserver-gl:latest

- run the hsl map server
-- take this repo as reference: https://github.com/HSLdevcom/hsl-map-server/tree/master
-- replace the config.js file
-- build the docker image 
-- rename the mbtiles file to finland.mbtiles TODO: fix this:

docker run -d -p 8080:8080 -v /home/anthony/lpz-tp/map-server/hsl-map-server/data/:/opt/hsl-map-server/data map-server:latest

following urls should work:
http://localhost:8080/map/v3/hsl-map/15/10181/17906.png
http://localhost:8080/map/v3/hsl-map/17/40733/71628.png

* for some reason when stoppig and restarting the map-server container, map-tiles service stops working well, you need to run a new service container
- additional notes:
- to download lpz mbtiles https://www.maptiler.com/on-prem-datasets/dataset/osm/south-america/bolivia/la-paz/#9.26/-16.4295/-68.116 (it requires registration. not tested)

- https://openmaptiles.org/docs/generate/create-custom-extract/
- https://download.geofabrik.de/south-america/bolivia.html

