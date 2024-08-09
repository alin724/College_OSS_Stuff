package main

import (
	"context"
	"log"
	"net"
        "google.golang.org/grpc"
        pb "grpc-server-op/pb/protos_grpc_op"
)

type grpc_server struct {
        pb.UnimplementedGrpcOpServer
}

func (s *grpc_server) GrpcOp_RPC(ctx context.Context, data_in *pb.GrpcOpReq) (*pb.GrpcOpResp, error) {
	log.Printf("Received: %v", data_in)
        return &pb.GrpcOpResp{RespVal: 2}, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:20001")
	if err != nil {
	log.Fatal(err)
	}
	server := grpc.NewServer()
        pb.RegisterGrpcOpServer(s, &grpc_server{})
	log.Printf("Server is listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
