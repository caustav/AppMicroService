package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"sync"

	pb "./comp1/comp1"
	engine "./engine"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50000"
)

type componentHolder struct {
	component    *pb.Component
	moduleClient pb.ModuleClient
}

type connectionHolder struct {
	url        string
	clientConn *grpc.ClientConn
}

var connHolderArray []connectionHolder
var componentMap map[string]*componentHolder
var connHolderMap map[string]connectionHolder

type server struct{}

func doServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	engine.RegisterEngineServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (c *server) Register(ctx context.Context, in *engine.EngineRequest) (*engine.EngineResponse, error) {
	out := new(engine.EngineResponse)
	conn, err := grpc.Dial(in.Url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	connHolderMap[in.UniqueNumber] = connectionHolder{in.Url, conn}
	moduleClient := pb.NewModuleClient(connHolderMap[in.UniqueNumber].clientConn)
	readBundle(moduleClient)
	return out, nil
}

func initEngine() {
	connHolderMap = make(map[string]connectionHolder)
	componentMap = make(map[string]*componentHolder)
}

func readBundle(moduleClient pb.ModuleClient) {
	r, err := moduleClient.Fetch(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("could not fetch: %v", err)
	}
	for j := range r.Component {
		componentMap[r.Component[j].Name] = &componentHolder{r.Component[j], moduleClient}
	}
}

func doLoadComponents() {
	urls := [...]string{"localhost:50051", "localhost:50052"}
	connHolderArray = make([]connectionHolder, len(urls))
	for i := range connHolderArray {
		connHolderArray[i].url = urls[i]
	}
	componentMap = make(map[string]*componentHolder)
	for i := range connHolderArray {
		conn, err := grpc.Dial(connHolderArray[i].url, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		connHolderArray[i].clientConn = conn

		c := pb.NewModuleClient(conn)

		r, err := c.Fetch(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatalf("could not fetch: %v", err)
		}
		for j := range r.Component {
			componentMap[r.Component[j].Name] = new(componentHolder)
			componentMap[r.Component[j].Name].component = r.Component[j]
			componentMap[r.Component[j].Name].moduleClient = c
		}
		log.Printf("Fetching: %s", r.Component)
		fmt.Println(r.Component[i].Type)
	}
}

func executeChannel(channels [2]string, wg *sync.WaitGroup) {
	// channels := [...]string{"t2", "a2"}
	for i := range channels {
		compHolder := componentMap[channels[i]]
		log.Printf("Fetching: %s", componentMap[channels[i]])
		c := compHolder.moduleClient
		c.Execute(context.Background(), &pb.ExecuteRequest{ComponentName: compHolder.component.Name, Type: compHolder.component.Type})
	}
	wg.Done()
}

func closeConnection() {
	for _, connHolder := range connHolderMap {
		connHolder.clientConn.Close()
	}
}

func main() {
	initEngine()
	go doServer()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Press any key")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		runChannels()
	}
}

func runChannels() {
	var wg sync.WaitGroup
	wg.Add(1)
	channels1 := [...]string{"t2", "a2"}
	//channels2 := [...]string{"t1", "a1"}
	//channels3 := [...]string{"t2", "a1"}
	// go executeChannel(channels1, &wg)
	go executeChannel(channels1, &wg)
	// go executeChannel(channels3, &wg)
	wg.Wait()
	//closeConnection()
}
