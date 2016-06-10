package store
import (
	iotMsg "github.com/alivinco/iotmsglibgo"
	"github.com/alivinco/greenhome/model"
)

type ThingsCacheStore struct{
	store map [string]iotMsg.IotMsg
}

func NewThingsCacheStore()(*ThingsCacheStore){
	cache := ThingsCacheStore{map[string]iotMsg.IotMsg{}}
	return &cache
}

func (ch *ThingsCacheStore) Set(topic string,msg iotMsg.IotMsg,ctx *model.Context){
	ch.store[ctx.Domain+"|"+topic] = msg
}

func (ch *ThingsCacheStore) Get(topic string,ctx *model.Context)(*iotMsg.IotMsg){
	tp , ok := ch.store[ctx.Domain+"|"+topic]
	if ok {
		return &tp
	}
	return nil

}

