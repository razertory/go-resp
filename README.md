# go-resp

A RESP(Redis Serialization Protocol) DB written in go.

## Usage

- `go run cmd/main.go`
- redis-cli -p 50000
samples:

```shell
➜  ~ redis-cli -p 50000
127.0.0.1:50000> hello
(error) ERR command not support
127.0.0.1:50000> ping
PONG
127.0.0.1:50000> set db1 redis
OK
127.0.0.1:50000> get db1
"redis"
127.0.0.1:50000> get db
(nil)
127.0.0.1:50000> mset db1 mysql db2 redis db3 pg
OK
127.0.0.1:50000> mget db1 db3
1) "mysql"
2) "pg"
````

