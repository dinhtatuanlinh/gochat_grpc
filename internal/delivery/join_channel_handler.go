package delivery

//
//
//import (
//	"fmt"
//	protocolBuffer "go_chat/internal/delivery/proto/protocol_buffer"
//)
//
//func (delivery *GrpcDelivery) JoinChannel(ch *protocolBuffer.Channel, msgStream protocolBuffer.Services_JoinChannelServer) error {
//
//	//msgChannel := make(chan *protocolBuffer.Message)
//	//delivery.channel = make(map[string][]chan *protocolBuffer.Message)
//	//participate := protocolBuffer.Message{
//	//	Channel: &protocolBuffer.Channel{
//	//		Name:        ch.Name,
//	//		SendersName: ch.SendersName,
//	//	},
//	//	Message: "",
//	//}
//	//msgChannel <- &participate
//
//	//delivery.channel[ch.Name] = append(delivery.channel[ch.Name], msgChannel)
//	channel := ch.Name
//	sender := ch.SendersName
//	// doing this never closes the stream
//	for {
//		select {
//		case <-msgStream.Context().Done():
//			return nil
//		case msg := <-Message:
//			fmt.Printf("GO ROUTINE (got conn): %v \n", msg)
//			if channel == msg.Channel.Name && sender != msg.Channel.SendersName {
//				msgStream.Send(msg)
//			}
//
//		}
//	}
//}
