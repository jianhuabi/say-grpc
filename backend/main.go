package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "say-grpc/backend/api"
)

func main() {
	port := flag.Int("p", 8080, "say grpc port to listen to")
	flag.Parse()

	sock, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		logrus.Fatalf("Fatal: couldn't listen to port %d: %v", *port, err)
	}

	logrus.Infof("Server is now listening on port %d", *port)

	s := grpc.NewServer()


	pb.RegisterTextToSpeechServer(s, server{})

	err = s.Serve(sock)

	if err != nil {
		logrus.Fatalf("Fatal: can't server on port %d", *port)
	}
}


type server struct{}


func (server) Say(context.Context, *pb.Text) (*pb.Speech, error) {
	return nil, fmt.Errorf("not implemented")
}


// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")

// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr

// if err := cmd.Run(); err != nil {
// 	log.Fatal(err)
// }
