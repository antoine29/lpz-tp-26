# geocoding request adapter

## pelias request
the adapter service will receive this request and re-use some of the params to make a motis request. the motis response fields will be used to make this request response

- add the /v1/search endpoint to the adapter server
- it should receive and parse the 4 listed url params (text, lang, sources, and layers)
- it should reuse the text url param to make a motis request (line :94)
- the endpoint should return the json response in the format shown from line :16

curl 'http://localhost:8084/v1/search?text=airport&lang=en&sources=oa%2Cosm%2Cnlsfi%2CgtfsHSL%2CgtfsHSLlautta&layers=station%2Cvenue%2Caddress%2Cstreet' \
  -H 'accept: application/json'

{
  "geocoding": { # fixed
    "version": "0.1", # fixed
    "attribution": "http://some-attribution", # fixed
    "query": {
      "text": "airport", # same as url param
      "size": 2, # length of features 
      "lang": "en", # same as url param
      "layers": [ # same as url param
        "station",
        "venue",
        "address",
        "street"
      ],
      "sources": [ # same as url params
        "openaddresses",
        "openstreetmap",
        "nlsfi",
        "gtfshsl",
        "gtfshsllautta"
      ],
      "private": false, # fixed
      "querySize": 100, # fixed
      "parsed_text": { # fixed
        "neighbourhood": "airport", # same as text url param
        "name": "airport" # same as text url param
      }
    },
    "warnings": [ # fixed
      "Invalid Parameter: digitransit-subscription-key" # fixed
    ],
    "engine": { # fixed (the whole object)
      "name": "Pelias",
      "author": "Mapzen",
      "version": "1.0"
    },
    "timestamp": 1777988168749 # current timestamp
  },
  "type": "FeatureCollection", # fixed
  "features": [
    {
      "type": "Feature", # fixed
      "geometry": { # fixed
        "type": "Point", # fixed
        "coordinates": [ # fixed
          24.966793,  # motis[].lat
          60.316841   # motis[].long
        ]
      },
      "properties": { # fixed
        "id": "GTFS:HSL:4000215", # motis[].id
        "layer": "station", # motis[].type
        "source": "gtfshsl", # fixed
        "name": "Airport", # motis[].name
        "label": "Airport, Lentokenttä, Vanda" # motis[].name
      }
    },
    {
      "type": "Feature", # fixed
      "geometry": { # fixed
        "type": "Point", # fixed
        "coordinates": [ # fixed
          24.968296, # motis[].lat
          60.318933  # motis[].long
        ]
      },
      "properties": { # fixed
        "id": "node:26608365", # motis[].id
        "layer": "venue", # fixed
        "source": "openstreetmap", # fixed
        "name": "Helsingfors-Vanda flygplats", # motis[].name
        "label": "Helsingfors-Vanda flygplats (Airport), Flygstationsvägen 1, Vanda" # motis[].name
      }
    }
  ],
}


# motis request
the adapter service will make make/send this request using the params it received on a /v1/search request

- it should reuse the text url param. language and mode are constants/fixed
- the response of this motis request should be used to build the /v1/search response

curl 'http://localhost:8083/api/v1/geocode?\
&text=airport\
&language=en\
&mode=AIRPLANE,HIGHSPEED_RAIL,LONG_DISTANCE,NIGHT_RAIL,COACH,RIDE_SHARING,REGIONAL_RAIL,SUBURBAN,SUBWAY,TRAM,BUS,FERRY,ODM,FUNICULAR,AERIAL_LIFT,OTHER' \
  -H 'Accept: */*' \
  -H 'Connection: keep-alive' \
  -b 'lang=en'

[
  {
    "type": "PLACE",
    "category": "post_box_12",
    "name": "Calle 10",
    "id": "node/[6066402902]",
    "lat": -16.4925075,
    "lon": -68.1081368,
    "country": "BO",
    "tz": "America/La_Paz",
    "score": -19.54684066772461
  },
  {
    "type": "PLACE",
    "category": "post_box_12",
    "name": "Calle 6",
    "id": "node/[6066506697]",
    "lat": -16.492236,
    "lon": -68.110287,
    "country": "BO",
    "tz": "America/La_Paz",
    "score": -19.54684066772461
  }
]


