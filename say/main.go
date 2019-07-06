package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"

	pb "say-grpc/backend/api"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address a say backend")
	output := flag.String("o", "output.wav", "WAV output file name")

	flag.Parse()

	conn, err := grpc.Dial(*backend)

	if err != nil {
		log.Fatal("couldn't connect to %s:%v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)

	res, err := client.Say(context.Background(), &pb.Text{Text: "Hello World"})

	if err != nil{
		log.Fatalf("couldn't get back say words %v", err)
	}


	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("couldn't write WAV file %s:%v", *output, err)
	}

}
