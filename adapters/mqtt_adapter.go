package adapters

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/glog"
	"github.com/alivinco/iotmsglibgo"
)

type MqttAdapter struct {
	client     MQTT.Client
	msgHandler MessageHandler
}

type MessageHandler func (adapter string,topic string,iotMsg *iotmsglibgo.IotMsg)

//serverUri="tcp://localhost:1883"
func NewMqttAdapter(serverUri string ,clientId string)(*MqttAdapter) {
	mh := MqttAdapter{}
	opts := MQTT.NewClientOptions().AddBroker(serverUri)
	opts.SetClientID(clientId)
	opts.SetDefaultPublishHandler(mh.onMessage)
	opts.SetCleanSession(true)

	//create and start a client using the above ClientOptions
	mh.client = MQTT.NewClient(opts)
	return &mh
}

func (mh *MqttAdapter)SeMessageHandler(msgHandler MessageHandler){
	mh.msgHandler = msgHandler
}

func (mh *MqttAdapter)Start()(error){
	if token := mh.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
func (mh *MqttAdapter)Stop(){
	mh.client.Disconnect(250)
}

func (mh *MqttAdapter)Subscribe(topic string,qos byte)(error){
	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := mh.client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		glog.Info(token.Error())
		return token.Error()
	}
	return nil
}

func (mh *MqttAdapter)Unsubscribe(topic string)(error){
	//unsubscribe from /go-mqtt/sample
	  if token := mh.client.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
	    return token.Error()
	  }
	  return nil
}

//define a function for the default message handler
func (mh *MqttAdapter) onMessage(client MQTT.Client, msg MQTT.Message) {
	glog.Info("TOPIC: %s\n", msg.Topic())
	glog.Info("MSG: %s\n", msg.Payload())
	iotMsg ,err := iotmsglibgo.ConvertBytesToIotMsg(msg.Topic(),msg.Payload(),nil)

	if err == nil {
		mh.msgHandler("mqtt",msg.Topic(),iotMsg)
	} else {
		glog.Error(err)

	}
}

func (mh *MqttAdapter)Publish(topic string,iotMsg *iotmsglibgo.IotMsg , qos byte)(error){
	bytm , err := iotmsglibgo.ConvertIotMsgToBytes(topic,iotMsg,nil)
	if err == nil {
		mh.client.Publish(topic,qos,false,bytm)
		return nil
	}else{
		return err
	}

}

