package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// "reflect"

	"google.golang.org/grpc"

	// db "grpc_gateway_sample/db"

	pb "grpc_gateway_sample/proto"
)

const (
	port = ":8080"
)

func SampleInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// メソッド呼び出し前の処理
		fmt.Printf("%+v\n", req)

		// リクエストされたgrpcメソッドを実行
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		// メソッド呼び出し後の処理
		fmt.Printf("%+v\n", resp)

		return resp, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v¥n", err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(SampleInterceptor()),
	)
	// 実行したい実処理をserverに登録する
	// periodService := &getPeriodService{}
	// pb.RegisterAimoServer(server, periodService)
	service := &GetPeriodService{}
	pb.RegisterAimoServer(server, service)
	log.Printf("server listening at %v\n", lis.Addr())
	if err != nil {
		log.Fatalf("faild to serve: %v\n", err)
	}
	server.Serve(lis)
}
