### Run server
```
make server SERVER_ARGS="-config=server.yaml"
OR
bin/server -config=server.yaml
```
### Run client
```
make client CLIENT_ARGS="-address=localhost:5555"
OR
bin/client -address=localhost:5555
```
### Commands
```
SET foo bar
GET foo
DELETE foo
```