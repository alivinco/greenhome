package routers

import (
	IotMsg "github.com/alivinco/iotmsglibgo"
	"github.com/alivinco/greenhome/adapters"
	"github.com/alivinco/greenhome/store"
	"github.com/alivinco/greenhome/model"
	log "github.com/Sirupsen/logrus"
)

type Router interface {
	OnMessage(adapter string,topic string,iotMsg *IotMsg.IotMsg)
}


type MainRouter struct {
	mqttAdapter *adapters.MqttAdapter
	wsAdapter *adapters.WsAdapter
	thingsCache *store.ThingsCacheStore
}

func NewMainRouter(mqttAdapter *adapters.MqttAdapter , wsAdapter *adapters.WsAdapter , thingsCache *store.ThingsCacheStore)(*MainRouter){
	mr := MainRouter{mqttAdapter,wsAdapter,thingsCache}
	mr.mqttAdapter.SetMessageHandler(mr.onMqttMessage)
	//mqttAdapter.Subscribe("jim1/cmd/test/grhome",1)
	mr.wsAdapter.SetMessageHandler(mr.onWsMessage)
	mr.thingsCache = thingsCache
	return &mr
}

func (mr *MainRouter)onMqttMessage(adapter string,topic string,iotMsg *IotMsg.IotMsg ,ctx *model.Context){
	log.Debug(iotMsg.String())
	mr.thingsCache.Set(topic,*iotMsg,ctx)
	mr.wsAdapter.Publish(topic,iotMsg,1)
}
func (mr *MainRouter)onWsMessage(adapter string,topic string,iotMsg *IotMsg.IotMsg ,ctx *model.Context){
	log.Debug(iotMsg.String())
	mr.thingsCache.Set(topic,*iotMsg,ctx)
	mr.mqttAdapter.Publish(topic,iotMsg,1,ctx)
}
