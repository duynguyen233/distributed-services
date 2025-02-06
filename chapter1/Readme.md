# Chapter 1: Let's Go
In this chapter, I tried to build a simple Commit Log when the app receive the HTTP request.
I built a simple JSON/HTTP commit log service that accepts and responds with JSON and stores the records into memory log which is slice of record struct

```
go run chapter1/server/main.go

curl -X POST localhost:8080 -d \
    '{"record": {"value": "TGV0J3MgR28gIzEK"}}'

curl -X POST localhost:8080 -d \
    '{"record": {"value": "TGV0J3MgR28gIzEK"}}'

curl -X POST localhost:8080 -d \
    '{"record": {"value": "TGV0J3MgR28gIzEK"}}'

curl -X GET localhost:8080 -d '{"offset": 0}'

curl -X GET localhost:8080 -d '{"offset": 0}'

curl -X GET localhost:8080 -d '{"offset": 0}'
```