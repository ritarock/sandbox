package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	hellopb "smaple-grpc/pkg/grpc"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func (s *myServer) HelloServerStream(req *hellopb.HelloRequest, stream hellopb.GreetingService_HelloServerStreamServer) error {
	resCount := 5
	for i := 0; i < resCount; i++ {
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func (s *myServer) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)

	for {
		// リクエスト受信
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			// リクエストをすべて受け取った後の処理
			message := fmt.Sprintf("Hello, %v!", nameList)
			// SendAndClose メソッドを呼ぶことでレスポンスを返す
			return stream.SendAndClose(&hellopb.HelloResponse{
				Message: message,
			})
		}
		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}

func (s *myServer) HelloBiStreams(stream hellopb.GreetingService_HelloBiStreamsServer) error {
	for {
		// リクエスト受信
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		message := fmt.Sprintf("Hello, %v!", req.GetName())
		if err := stream.Send(&hellopb.HelloResponse{
			Message: message,
		}); err != nil {
			return err
		}
	}
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	// listener を作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// gRPC サーバを作成
	s := grpc.NewServer()

	// gRPC サーバに GreetingService を登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// サーバリフレクションの設定
	// シリアライズのルールを知らない grpcurl がサーバから proto ファイルの情報を取得する
	// これによって grpcurl が使えるようになる
	reflection.Register(s)

	// gRPC サーバを稼働
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// Ctrl + C が入力されたら shutdown する
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
