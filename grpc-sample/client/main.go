package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-client-op/pb/protos_grpc_op"
)

func main() {
	val_set_ip := os.Getenv("SET_IPADDR") + ":20001"
	dial_rpc, err := grpc.Dial(val_set_ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer dial_rpc.Close()

	req1 := &pb.GrpcOpReq {
		ReqVal: 32,
	}

	client := pb.NewGrpcOpClient(dial_rpc)
	req_data, err := client.GrpcOp_RPC(context.Background(), req1)
        if err != nil {
                log.Fatalf("failed to send request: %v", err)
        }
        log.Printf("Response: %v", req_data)
}


 
