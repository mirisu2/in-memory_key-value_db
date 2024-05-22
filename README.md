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

### Config
```
engine:
  # mandatory
  type: "memory"
network:
  # optional; default=127.0.0.1:5555
  address: "127.0.0.1:5555"
  # optional;  default=100
  max_connections: 2
logging:
  # optional; available values: debug, warn, error, info; default=info 
  level: "info"
  
  # optional; available values: console or file; default=console
  output: "file"
  # optional, it matters if "output" is set; default=app.log
  log_file_name: "server.log"
  
  # optional; available values: json or text; default=text
  format: "json"
  
  # optional; available values: true or false; default=true
  source: "true"
```