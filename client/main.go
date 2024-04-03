package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "go_chat/internal/delivery/proto/protocol_buffer"

	"google.golang.org/grpc"
)

func main() {
	// dial server
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	defer conn.Close()
	// create stream
	client := pb.NewServicesClient(conn)
	in := &pb.StreamServerRequest{Id: 1}
	stream, err := client.StreamServer(context.Background(), in)
	if err == io.EOF {
		fmt.Println("completely send data from server")
		return
	}
	if err != nil && err != io.EOF {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Msg)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}
