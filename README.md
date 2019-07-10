# Postgres connection ping segfault

This code will cause a panic/segfault when trying to ping a database that is down or doesn't exist.

Tested with:
- go version: 1.12.7
- lib/pq: 1.1.1


### Steps to reproduce

1. `git clone git@github.com:jc21/golang-sql-pq-test.git`
2. `cd golang-sql-pq-test.git`
3. `go build -o pq-test`
4. `./pq-test`


#### Expected

The Ping method should return a valid error


#### Actual

```
PostgresError: dial tcp: lookup doesnt-exist.example.com on 10.0.1.1:53: no such hostpanic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4c16da]

goroutine 1 [running]:
database/sql.(*DB).conn(0x0, 0x6b6a60, 0xc00009e010, 0x6ae701, 0x618020, 0x6ae730, 0x618020)
        /usr/lib/go-1.12/src/database/sql/sql.go:1080 +0x3a
database/sql.(*DB).PingContext(0x0, 0x6b6a60, 0xc00009e010, 0x5d956c, 0x666e28)
        /usr/lib/go-1.12/src/database/sql/sql.go:730 +0x90
database/sql.(*DB).Ping(...)
        /usr/lib/go-1.12/src/database/sql/sql.go:748
main.main()
        /home/jamiec/web/golang/pq-test/main.go:12 +0x45
```

