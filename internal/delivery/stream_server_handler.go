package delivery

import (
	"fmt"
	protocolBuffer "go_chat/internal/delivery/proto/protocol_buffer"
	"log"
	"time"
)

func (delivery *GrpcDelivery) StreamServer(in *protocolBuffer.StreamServerRequest, srv protocolBuffer.Services_StreamServerServer) error {
	log.Printf("fetch response for id : %d", in.Id)

	//use wait group to allow process to be concurrent
	count := 0
	for {
		//time sleep to simulate server process time
		time.Sleep(time.Duration(count) * time.Second)
		resp := protocolBuffer.StreamServerResponse{Msg: fmt.Sprintf("Request #%d For Id:%d", count, in.Id)}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("finishing request number : %d", count)
		count++
		if count == 1000 {
			break
		}
	}

	return nil
}
