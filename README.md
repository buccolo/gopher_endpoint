# Gopher Endpoint [![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy?template=https://github.com/buccolo/gopher_endpoint)

A [Wombat](https://wombat.co) endpoint that returns gophers.

### Start the endpoint

```
PORT=5000 go run gopher_endpoint.go
```

### Request

```
curl -i -X POST -d "{ \"request_id\": \"$RANDOM\" }" http://0.0.0.0:5000
```

###  Response

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 09 Aug 2014 18:23:22 GMT
Content-Length: 87

{"request_id":"4676","gophers":[{"id":2101718032230089003},{"id":1966644620900390021}]}
```
