# OTP GQL examples:

Following queries should work on the OTP gql interface:

## Get all routes

```
query srcSearchRoutesQuery($feeds: [String!]!, $modes: [Mode]) {
  viewer {
    routes(feeds: $feeds, transportModes: $modes) {
      gtfsId
      agency {
        name
        id
      }
      type
      shortName
      mode
      color
      longName
      patterns {
        code
        id
      }
      id
    }
  }
}

{
  "feeds": [
    "LPB",
    "MT"
  ],
  "modes": null
}
```


