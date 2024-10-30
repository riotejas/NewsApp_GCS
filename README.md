
# News API Application

Simple API to query News APIs services 

# Build
Build a docker image
`make docker-build`

# Run
Run locally
`make docker-run`

# Query
## Scan articles
```shell
curl --request GET \
  --url 'http://localhost:8080/v1/scan'
```

## Doc
Get OpenAPI doc
```shell
curl --request GET \
  --url http://localhost:8080/doc
```

## Health
Get Application Health
```shell
curl --request GET \
  --url http://localhost:8080/health
```

## Simulate an error response
Error response for users to test their error handling
```shell
curl --request GET \
  --url http://localhost:8080/error
```