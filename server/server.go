package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/hamadakafu/grpc-sample/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// sleepして3秒毎に挨拶するようにする

type GreetServer struct{}

func (s *GreetServer) UnaryGreet(ctx context.Context, in *pb.RequestGreet) (*pb.ResponseGreet, error) {
	res := pb.ResponseGreet{
		Message: "hello" + in.Message,
	}
	return &res, nil
}

// ServerStreamGreetはリクエストが来ると5秒毎にレスポンスを返すようにする
func (s *GreetServer) ServerStreamGreet(in *pb.RequestGreet, srv pb.Greet_ServerStreamGreetServer) error {
	mes := in.Message
	log.Printf("message: %s", mes)
	for {
		if err := srv.Send(&pb.ResponseGreet{
			Message: "hello" + mes,
		}); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}
func (s *GreetServer) BidirectionalStreamGreet(srv pb.Greet_BidirectionalStreamGreetServer) error {
	for {
		req, err := srv.Recv()
		if err != nil {
			return err
		}

		mes := req.Message
		log.Printf("message: %s", mes)

		if err := srv.Send(&pb.ResponseGreet{
			Message: "Hello" + mes,
		}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGreetServer(s, &GreetServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
