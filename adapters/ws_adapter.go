package adapters

import (
	"github.com/olahol/melody"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/alivinco/iotmsglibgo"
	"encoding/json"
	"github.com/golang/glog"
	"strings"
	"github.com/alivinco/greenhome/model"
)

type WsMessage struct {
	Topic string `json:"topic"`
	Payload string `json:"payload"`
}

type WsAdapter struct {
	mel *melody.Melody
	ginE *gin.RouterGroup
	msgHandler MessageHandler
}

func NewWsAdapter(ginE *gin.RouterGroup )(*WsAdapter){
	m := melody.New()
	wsa := WsAdapter{mel:m,ginE:ginE}

	ginE.GET("", func(c *gin.Context) {
		reqDomain := c.Query("domain")
		authRequest ,exists := c.Get("AuthRequest")
		if exists{
			if authRequest.(model.AuthRequest).ValidateDomain(reqDomain){
				m.HandleRequest(c.Writer, c.Request)
				return
			}
		}
		fmt.Print("Request can't be authenticated")
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
	wsMsg.Topic = strings.Replace(wsMsg.Topic," ","",-1)
	domain := s.Request.URL.Query().Get("domain")
	fmt.Println("New WS message from topic = %v for domain =",wsMsg.Topic,domain)
	fmt.Println(wsMsg.Payload)
	iotMsg ,err := iotmsglibgo.ConvertBytesToIotMsg(wsMsg.Topic,[]byte(wsMsg.Payload),map[string]string{"override_payload_type":"jim1"})
	ctx := model.Context{Domain:domain}
	if err == nil {
		wsa.msgHandler("ws",wsMsg.Topic,iotMsg,&ctx)
	} else {
		glog.Error(err)

	}
}
// Publish messages over websocket
func (wsa *WsAdapter)Publish(topic string,iotMsg *iotmsglibgo.IotMsg , qos byte)(error){
	msg , err := iotmsglibgo.ConvertIotMsgToBytes(topic,iotMsg,map[string]string{"override_payload_type":"jim1"})
	wsMsg := WsMessage{Topic:topic,Payload:string(msg)}
	if err == nil{
		msg , _ := json.Marshal(wsMsg)
		wsa.mel.Broadcast(msg)
	}
	return err
}