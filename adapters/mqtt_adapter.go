package adapters

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/alivinco/iotmsglibgo"
	"strings"
	"github.com/alivinco/greenhome/model"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

type MqttAdapter struct {
	client     MQTT.Client
	msgHandler MessageHandler
}

type MessageHandler func (adapter string,topic string,iotMsg *iotmsglibgo.IotMsg , ctx *model.Context)

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

func (mh *MqttAdapter)SetMessageHandler(msgHandler MessageHandler){
	mh.msgHandler = msgHandler
}

func (mh *MqttAdapter)Start()(error){
	log.Info("Connecting to MQTT broker ")
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
	log.Debug("Subscribing to topic:",topic)
	if token := mh.client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		log.Info(token.Error())
		return token.Error()
	}
	return nil
}


func (mh *MqttAdapter)Unsubscribe(topic string)(error){
	  log.Debug("Unsubscribing from topic:",topic)
	  if token := mh.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
	    return token.Error()
	  }
	  return nil
}

// The method should be invoked whenever topics are modified in data model
func (mh *MqttAdapter)TopicChangeHandler(topics []string,isSub bool,ctx *model.Context){
	for _,topic := range topics{
		if isSub {
			mh.Subscribe(AddDomainToTopic(ctx.Domain,topic),1)
		}else{
			if topic != ""{
				mh.Unsubscribe(AddDomainToTopic(ctx.Domain,topic))
			}else{
				log.Debug("Topic is empty , nothing to unsubscribe.")
			}

		}
	}
}



//define a function for the default message handler
func (mh *MqttAdapter) onMessage(client MQTT.Client, msg MQTT.Message) {
	log.Info("TOPIC: %s\n", msg.Topic())
	log.Info("MSG: %s\n", msg.Payload())
	fmt.Println("New msg from topic:",msg.Topic())
	domain , topic := DetachDomainFromTopic(msg.Topic())
	iotMsg ,err := iotmsglibgo.ConvertBytesToIotMsg(topic,msg.Payload(),nil)
	ctx := model.Context{Domain:domain}
	if err == nil {
		mh.msgHandler("mqtt",topic,iotMsg , &ctx)
	} else {
		log.Error(err)

	}
}

func (mh *MqttAdapter)Publish(topic string,iotMsg *iotmsglibgo.IotMsg , qos byte,ctx *model.Context)(error){
	topic = AddDomainToTopic(ctx.Domain,topic)
	bytm , err := iotmsglibgo.ConvertIotMsgToBytes(topic,iotMsg,nil)
	if err == nil {
		fmt.Println("Publishing msg to topic:",topic)
		mh.client.Publish(topic,qos,false,bytm)
		return nil
	}else{
		return err
	}

}

func AddDomainToTopic(domain string , topic string )string {
	// Check if topic is already prefixed with  "/" if yes then concat without adding "/"
	// 47 is code of "/"
	if topic[0] == 47 {
		return domain+topic
	}
	return domain+"/"+topic
}
func DetachDomainFromTopic(topic string ) (string , string) {
	spt := strings.Split(topic , "/")
	// spt[0] - domain
	var top string
	if strings.Contains(spt[1],"jim"){
		top = strings.Replace(topic,spt[0]+"/","",1)
	}else{
		top = strings.Replace(topic,spt[0],"",1)
	}
	// returns domain , topic
	return spt[0] , top

}


