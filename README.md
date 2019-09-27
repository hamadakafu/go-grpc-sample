# gRPC

```go
./protoc-gen.sh
go run server/server.go
```

## proto -> go
protoディレクトリにあるsample.protoを使用して，
`./pb`の場所にsample.pd.goを作成する
```bash
protoc --proto_path proto --go_out=plugins=grpc:./pb sample.proto
```

## Server側が一方向Streamをする場合
実装すべき関数の型はUnaryの場合と異なる

Requestの構造体とStreamServerを引数に取る関数

## Bidirectional Streamをする場合
BidirectionalStreamServerを引数に取る関数
あるRPCがBidirectional Streamだった場合，専用のサーバが作成される模様

Server側でタイマーを用意するのはうまいやり方が分からなかった．
Client側でタイマーを用意して，終了するのが良さそう

