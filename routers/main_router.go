package routers

import (
	IotMsg "github.com/alivinco/iotmsglibgo"
	"github.com/alivinco/greenhome/adapters"
	"fmt"
)

type Router interface {
	OnMessage(adapter string,topic string,iotMsg *IotMsg.IotMsg)
}

type MainRouter struct {
	mqttAdapter *adapters.MqttAdapter
	wsAdapter *adapters.WsAdapter
}

func NewMainRouter(mqttAdapter *adapters.MqttAdapter , wsAdapter *adapters.WsAdapter)(*MainRouter){
	mr := MainRouter{mqttAdapter,wsAdapter}
	mr.mqttAdapter.SeMessageHandler(mr.onMqttMessage)
	//mqttAdapter.Subscribe("jim1/cmd/test/grhome",1)
	mr.wsAdapter.SetMessageHandler(mr.onWsMessage)
	return &mr
}

func (mr *MainRouter)onMqttMessage(adapter string,topic string,iotMsg *IotMsg.IotMsg){
	fmt.Println(iotMsg.String())
	mr.wsAdapter.Publish("jim1"+topic,iotMsg,1)
}
func (mr *MainRouter)onWsMessage(adapter string,topic string,iotMsg *IotMsg.IotMsg){
	fmt.Println(iotMsg.String())
	mr.mqttAdapter.Publish(topic,iotMsg,1)
}
