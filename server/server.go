// this implementation is borrowed heavily from:
// https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go

package main

import (
	"context"
	"flag"
	"fmt"
	pb "http-user-agent-analyzer/httpanalyzer"
	"log"
	"net"

	"github.com/mileusna/useragent"
	"google.golang.org/grpc"
)

var (
	// TODO(cait) port number chosen randomly, pick a more informed port
	port = flag.Int("port", 26817, "The server port")
)

type server struct {
	pb.UnimplementedHttpAnalyzerServer
}

// TODO(cait) Allow(...) based on a http.Request.UserAgent()

func (s *server) Allow(ctx context.Context, in *pb.AllowRequest) (*pb.AllowReply, error) {
	input := in.GetUserAgentString()
	log.Printf("Received request: %v", input)
	ua := useragent.Parse(input)
	// because the current spec only allows for FireFox requests to be allowed,
	// if we don't match on Firefox then we don't allow. it's likely we'll include
	// more cases in the future but for now this is a simpler implementation
	return &pb.AllowReply{IsAllowed: ua.IsFirefox()}, nil
}

func main() {
	// for when we're parsing command line args
	// flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHttpAnalyzerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
