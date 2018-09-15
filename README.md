# Distributed Log Querier

A Golang distributed log querier that can query distributed log files on multiple machines, from any of those machines.

## To Run:
Run server on all machines:
```
go run server.go
```

On any one of the machine, start the client to query log files:
```
go build client.go
./client.go [query] 3000 [log file name]
```
### To Run Unit Test:
Generate all unit tests:
```
go run generate_testfiles.go
```
Run server on all machines and on any one machine, run:
```
go test -v client_test.go

```

### Delete Unit Test Files
```
go run delete_testfiles.go
```

