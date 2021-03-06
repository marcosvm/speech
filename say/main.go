package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/marcosvm/speech/api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav file")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("usage:\n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	text := pb.Text{Text: os.Args[1]}
	res, err := client.Say(context.Background(), &text)
	if err != nil {
		log.Fatalf("could not say %s: %v", text, err)
	}
	err = ioutil.WriteFile(*output, res.Audio, 0666)
	if err != nil {
		log.Fatalf("could not write to %s: %v", *output, err)
	}
}
