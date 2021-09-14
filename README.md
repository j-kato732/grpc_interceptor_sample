# grpc_postgres_sample
```
$ git clone git@github.com:j-kato732/grpc_gateway_sample.git
$ cd grpc_gateway_sample
$ docker-compose up -d
$ docker exec -ti grpc_gateway_sample bash
```
# proto変更した場合
```
// serverの生成
$ protoc -I .:${GOPATH}/src --go_out ./ --go_opt paths=source_relative     --go-grpc_out ./ --go-grpc_opt paths=source_relative proto/aimo.proto
```
```
// gatewayの生成
$ protoc -I .:${GOPATH}/src --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative proto/aimo.proto
```

# How to connect
```
$ sqlite3 test.db 
```
# show table
```
$ .table
```

# sqlite web
```
$ sqlite_web test.db
```
