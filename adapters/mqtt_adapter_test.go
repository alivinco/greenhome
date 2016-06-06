package adapters

import "testing"
import (IotMsg "github.com/alivinco/iotmsglibgo")
var syncChan = make(chan string)


func  OnMessage(adapter string,topic string,iotMsg *IotMsg.IotMsg){
	//tr.t.Log(topic,iotMsg.String())
	syncChan <- iotMsg.String()
}

func TestNewMqttAdapter(t *testing.T) {
	ad := NewMqttAdapter("tcp://localhost:1883","greenhome_test")
	ad.SeMessageHandler(OnMessage)
	ad.Start()
	ad.Subscribe("jim1/cmd/test/grhome",1)
	iotMsg := IotMsg.NewIotMsg(IotMsg.MsgTypeCmd,"binary","switch",nil)
	iotMsg.SetDefaultBool(true,"")
	ad.Publish("jim1/cmd/test/grhome",iotMsg,1)
	chmsg := <-syncChan
	t.Log(chmsg)
	ad.Unsubscribe("jim1/cmd/test/grhome")
	ad.Stop()
}
