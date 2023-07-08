// this implementation is borrowed heavily from:
// https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go

package server

import (
	"flag"
	"fmt"
	pb "http-user-agent-analyzer/httpanalyzer"
	"log"
	"net"

	"google.golang.org/grpc"
)

const DefaultPort = 26817

var (
	// TODO(cait) port number chosen randomly, pick a more informed port
	port = flag.Int("port", DefaultPort, "The server port")
)

type server struct {
	pb.UnimplementedHttpAnalyzerServer
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
