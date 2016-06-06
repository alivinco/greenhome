package adapters

import (
	"github.com/olahol/melody"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/alivinco/iotmsglibgo"
)

type WsAdapter struct {
	mel *melody.Melody
	ginE *gin.Engine
	msgHandler MessageHandler
}

func NewWsAdapter(ginE *gin.Engine)(*WsAdapter){
	m := melody.New()
	wsa := WsAdapter{mel:m,ginE:ginE}

	ginE.GET("/greenhome/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("Client connected from ",s.Request.URL)
	})
	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("Client disconnected from ",s.Request.URL)
	})

	return &wsa
}

func (wsa *WsAdapter)SetMessageHandler(msgHandler MessageHandler){
	wsa.msgHandler = msgHandler
}

func (wsa *WsAdapter)Publish(topic string,iotMsg *iotmsglibgo.IotMsg , qos byte)(error){
	msg , err := iotmsglibgo.ConvertIotMsgToBytes(topic,iotMsg,nil)
	if err == nil{
		wsa.mel.Broadcast(msg)
	}
	return err
}