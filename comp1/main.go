package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	action "./action"
	pb "./comp1"
	"./engine"
	trigger "./trigger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:50000"
	port    = ":50051"
)

type server struct{}

func (c *server) Fetch(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	comps := []*pb.Component{{"t1", "trigger"}, {"a1", "action"}}
	return &pb.Response{comps}, nil
}

func (c *server) Execute(ctx context.Context, in *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	compName := in.ComponentName
	if compName == "t1" {
		var t trigger.Trigger
		t.Submit()
	} else if compName == "a1" {
		action.Execute()
	}
	return &pb.ExecuteResponse{}, nil
}

func doServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterModuleServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func doClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := engine.NewEngineClient(conn)

	// Contact the server and print out its response.
	r, err := c.Register(context.Background(), &engine.EngineRequest{"1", "localhost:50051"})
	if err != nil {
		log.Fatalf("could not fetch: %v", err)
	}
	log.Printf("Fetching: %s", r)
}

func main() {
	go doServer()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		doClient()
	}
}
