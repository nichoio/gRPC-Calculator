A simple calculator, implemented with Go and gRPC.

### Requirements ###

Docker 18.06.0+ and Docker Compose.

### Setup ###

Run `docker-compose up`

### Run ###

Example: `docker exec client /bin/bash -c 'go run calculate/calculate_client/client.go -a 2 -b 4 -method add -prec 2'`

Show help: `docker exec client /bin/bash -c 'go run calculate/calculate_client/client.go --help'`  
The help command generates the following:

```
Usage of client:
  -a float
        First Number (default 1)
  -b float
        Second Number (default 1)
  -method string
        Operator to use (add, sub, mult, div, sqd, root) (default "add")
  -prec uint
        Precision for rounding (round to nth digit) (default 3)
```

Alternatively, run shell session within client container and execute client from within there:  
`docker exec -it client /bin/bash`, then e.g. `go run calculate/calculate_client/client.go -a 8 -b 4 -method mult`

### Test ###
Run unit tests: `docker exec server /bin/bash -c 'go test /app/...'`
