package samplegrpc

import (
	"context"
	"time"
	"log"
	"google.golang.org/grpc"
	pb "github.com/harisaginting/ginting/model/proto"
)

const addr string = "localhost:50051"
var c pb.SampleClient

func conn() *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = pb.NewSampleClient(conn)
	return conn
}

func Sample(name string)(res string){
	conn := conn()
	defer conn.Close()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	res = r.GetMessage()
	return
}