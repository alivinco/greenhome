package adapters

import (
	"github.com/olahol/melody"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/alivinco/iotmsglibgo"
	"encoding/json"
	"github.com/golang/glog"
)

type WsMessage struct {
	Topic string `json:"topic"`
	Payload string `json:"payload"`
}

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
	m.HandleMessage(wsa.OnMessage)
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
func (wsa *WsAdapter)OnMessage(s *melody.Session, msg []byte){
	//wsa.mel.Broadcast(msg)
	wsMsg := WsMessage{}
	err := json.Unmarshal(msg,&wsMsg)
	fmt.Println(wsMsg)
	fmt.Println("New message from topic = ",wsMsg.Topic)
	fmt.Println(wsMsg.Payload)
	iotMsg ,err := iotmsglibgo.ConvertBytesToIotMsg(wsMsg.Topic,[]byte(wsMsg.Payload),nil)
	if err == nil {
		wsa.msgHandler("ws",wsMsg.Topic,iotMsg)
	} else {
		glog.Error(err)

	}
}

func (wsa *WsAdapter)Publish(topic string,iotMsg *iotmsglibgo.IotMsg , qos byte)(error){
	msg , err := iotmsglibgo.ConvertIotMsgToBytes(topic,iotMsg,nil)
	wsMsg := WsMessage{Topic:topic,Payload:string(msg)}
	if err == nil{
		msg , _ := json.Marshal(wsMsg)
		wsa.mel.Broadcast(msg)
	}
	return err
}