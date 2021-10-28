package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	// "reflect"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errdetails "grpc_gateway_sample/errors"
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
			log.Println(errors.Is(err, errdetails.InvalidParam))
			log.Printf("Response: %v", &pb.GetUserInfoResponse{})
			log.Printf("Error: %v", err.Error())
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}

		// メソッド呼び出し後の処理
		fmt.Printf("%+v\n", resp)

		return resp, nil
	}
}

func main() {
	// prepare zap logger
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Println(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v¥n", err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			SampleInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_validator.UnaryServerInterceptor(),
		)),
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
