package client

import (
	"context"
	"flag"
	pb "http-user-agent-analyzer/httpanalyzer"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// TODO(cait) pick a better default; rework description
	userAgent = flag.String("user-agent-header", "default-header", "description of flag")
)

func Allow(userAgentString string) error {
	allow(userAgentString)
	return nil
}

func main() {
	flag.Parse()
	// TODO(cait) this port appears twice it'd be nice to DRY this up
	allow(*userAgent)
}

func allow(userAgentString string) {
	conn, err := grpc.Dial("localhost:26817", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v%v", err, conn)
	}
	defer conn.Close()
	c := pb.NewHttpAnalyzerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Allow(ctx, &pb.AllowRequest{UserAgentString: userAgentString})
	if err != nil {
		log.Fatalf("could not process user agent string: %s", userAgentString)
	}
	log.Printf("Allowed? %t", r.IsAllowed)
}
