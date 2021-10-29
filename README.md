# middlewareのインストール

grpc-middlewareのインストールが済んでいない場合はこれから

```go
go get github.com/grpc-ecosystem/go-grpc-middleware
```

# validatorのインストール

```go
$ go get https://github.com/mwitkow/go-proto-validators.git
```


grpc-middlewareにはRequestのValidateを行うミドルウェアが存在する

使用するにはprotoにバリデート方法を記述する

```go
syntax = "proto3";
package grpc_sample;

import "go-proto-validators/validator.proto";

message getUserInfoRequest {
    int32  user_id = 1; 
    string  period = 2 [(validator.field) = {string_not_empty : true}];
}
```

# protoc実行(コード生成)
goファイルを生成するコマンドに`—govalidators_out`を追加する

—govalidators_optで各種設定も行える

```go
protoc -I .:${GOPATH}/src:${GOPATH}/pkg/mod --go_out ./  --go_opt paths=source_relative  --govalidators_out ./ --govalidators_opt paths=source_relative   --go-grpc_out ./ --go-grpc_opt paths=source_relative proto/aimo.proto
```

バリデートファイルの出力先をオプションで追加することで、

指定したディレクトリ配下にhello.validator.pb.goというファイルができる。

今回の場合は実行時ディレクトリ配下

ソースの生成にインポート先のものがないとおこられるので、

`git clone [https://github.com/mwitkow/go-proto-validators.git](https://github.com/mwitkow/go-proto-validators.git) ${GOPATH}/src/go-proto-validators`

を実行してゴリ押しで.protoファイルを落としてくる

# gRPCサーバにvalidatorsをセット
最後にgRPCサーバにinterceptorをセットする

```go
import "github.com/grpc-ecosystem/go-grpc-middleware/validator"
```

```go
server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			SampleInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
```
# 参考

[GitHub - mwitkow/go-proto-validators: Generate message validators from .proto annotations.](https://github.com/mwitkow/go-proto-validators)
[go-grpc-middlewareを一通り試してみる - Qiita](https://qiita.com/Morix1500/items/7a20d76a931af68d860d#grpc_validator)


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
```
// validatorの生成
protoc -I .:${GOPATH}/src --go_out ./  --go_opt paths=source_relative  --govalidators_out ./ --govalidators_opt paths=source_relative   --go-grpc_out ./ --go-grpc_opt paths=source_relative proto/aimo.proto
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
