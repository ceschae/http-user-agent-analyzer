package client

import (
	"context"
	"fmt"
	pb "http-user-agent-analyzer/httpanalyzer"
	"http-user-agent-analyzer/server"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", server.DefaultPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v%v", err, conn)
	}
	defer conn.Close()
	c := pb.NewHttpAnalyzerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Allow(ctx, &pb.AllowRequest{})
	if err != nil {
		log.Fatalf("could not process user agent string")
	}
	log.Printf("Allowed? %t", r.IsAllowed)
}
