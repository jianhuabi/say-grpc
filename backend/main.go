package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"net"
	"os/exec"

	pb "github.com/jianhuabi/say-grpc/backend/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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


func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {

	f, err := ioutil.TempFile("", "")

	if err != nil {
		return nil, fmt.Errorf("Fatal: couldn't create temp file %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close %s: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}


// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")

// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr

// if err := cmd.Run(); err != nil {
// 	log.Fatal(err)
// }
