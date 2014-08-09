# Gopher Endpoint

A [Wombat](wombat.co) endpoint that retruns gophers.

### Request

```
curl -i -X POST -d "{ \"request_id\": \"$RANDOM\" }" http://gopher-endpoint.url
```

###  Response

```
HTTP/1.1 200 OK
Date: Sat, 09 Aug 2014 18:23:22 GMT
Content-Length: 87
Content-Type: text/plain; charset=utf-8

{"request_id":"4676","gophers":[{"id":2101718032230089003},{"id":1966644620900390021}]}
```