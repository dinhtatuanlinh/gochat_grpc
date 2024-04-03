package delivery

import (
	"fmt"
	protocolBuffer "go_chat/internal/delivery/proto/protocol_buffer"
	"io"
)

func (delivery *GrpcDelivery) SendMessageToChannel(msgStream protocolBuffer.Services_SendMessageToChannelServer) error {

	go func() {
		for {
			select {
			case <-msgStream.Context().Done():
			//	msgStream.
			case receivingMsg := <-delivery.message:
				numberOfWorker := 5
				jobs := make(chan ClientMsg, numberOfWorker)
				for i := 0; i < numberOfWorker; i++ {
					go worker(jobs, receivingMsg)
				}
				for _, clientMsg := range delivery.clientConnections[receivingMsg.Channel.Name] {
					if clientMsg.sender != receivingMsg.Channel.SendersName {
						fmt.Printf("username: %s sent message to channel: %s \n", receivingMsg.Channel.SendersName, receivingMsg.Channel.Name)
						jobs <- clientMsg
					}
				}
			}
		}
	}()
	for {
		msg, err := msgStream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil && err != io.EOF {
			return err
		}
		clientMsg := ClientMsg{
			sender: msg.Channel.SendersName,
			conn:   msgStream,
		}

		_, ok := delivery.connections[fmt.Sprintf("%s-%s", msg.Channel.Name, msg.Channel.SendersName)]
		if !ok {
			delivery.connections[fmt.Sprintf("%s-%s", msg.Channel.Name, msg.Channel.SendersName)] = 1
			delivery.clientConnections[msg.Channel.Name] = append(delivery.clientConnections[msg.Channel.Name], clientMsg)
		}

		delivery.message <- msg

	}

}

func worker(jobs <-chan ClientMsg, msg *protocolBuffer.Message) {
	for job := range jobs {
		err := job.conn.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
