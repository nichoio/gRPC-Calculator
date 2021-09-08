A simple calculator, implemented with Go and gRPC.

Setup: docker-compose up

Run client, example: docker exec client /bin/bash -c 'go run calculate/calculate_client/client.go -a 2 -b 4 -method add -prec 2'

Show client help: docker exec client /bin/bash -c 'go run calculate/calculate_client/client.go --help'

Usage of client:
  -a float
        First Number (default 1)
  -b float
        Second Number (default 1)
  -method string
        Operator to use (add, sub, mult, div, sqd, root) (default "add")
  -prec uint
        Precision for rounding (round to nth digit) (default 3)


Alternatively, run shell session within client container and execute client from within there:
docker exec -it client /bin/bash

Test: docker exec server /bin/bash -c 'go test /app/...'
