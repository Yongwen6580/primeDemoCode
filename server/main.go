package main

import (
	"context"
	"log"
	"math/big"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "primeDemo/prime"
)

type server struct {
	pb.UnimplementedPrimeCheckerServer
}

func (s *server) mustEmbedUnimplementedPrimeCheckerServer() {
	//TODO implement me
	panic("implement me")
}

func (s *server) IsPrime(ctx context.Context, req *pb.PrimeRequest) (*pb.PrimeResponse, error) {
	n := big.NewInt(int64(req.Number))
	if n.ProbablyPrime(20) {
		return &pb.PrimeResponse{IsPrime: true}, nil
	}
	return &pb.PrimeResponse{IsPrime: false}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPrimeCheckerServer(s, &server{})
	reflection.Register(s)

	log.Println("starting server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
