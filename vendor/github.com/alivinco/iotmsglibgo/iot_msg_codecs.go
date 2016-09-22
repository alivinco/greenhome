package iotmsglibgo

import (
	"github.com/alivinco/iotmsglibgo/json-types"
	"encoding/json"
	//"strings"
	"time"
)

//type IotMsgToJsonIotMsgV0Codec struct {
//	msgTypeToStrMap  map[IotMsgType]string
//
//}
//
//func NewIotMsgToJsonIotMsgV0Codec ()(*IotMsgToJsonIotMsgV0Codec){
//	typeMap := map[IotMsgType]string {
//		MsgTypeCmd:"command",
//		MsgTypeEvt:"event",
//		MsgTypeGet:"get",
//	}
//	cod := IotMsgToJsonIotMsgV0Codec{typeMap}
//	return &cod
//}

// Converts IotMsg into it's V0-JSON byte array
func EncodeIotMsgToJsonStrV0 (msg *IotMsg,spid string)([]byte,error){
	var res interface{}
	switch msg.Type {
	case MsgTypeCmd:
		r := json_types.IotMsgCmdV0{}
		r.Command.Type = msg.Class
		r.Command.Subtype = msg.SubClass
		r.Command.Default.Value = msg.Default.Value
		r.Command.Default.Unit = msg.Default.Unit
		r.Command.Properties = msg.Properties
		r.UUID = msg.Uuid
		r.CreationTime = msg.Timestamp.Unix()
		r.Corid = msg.Corid
		r.Transport = msg.Transport
		r.Spid = spid
		res = r
	case MsgTypeEvt:
		r := json_types.IotMsgEvtV0{}
		r.Event.Default.Value = msg.Default.Value
		r.Event.Default.Unit = msg.Default.Unit
		r.Event.Properties = msg.Properties
		r.UUID = msg.Uuid
		r.CreationTime = msg.Timestamp.Unix()
		r.Corid = msg.Corid
		r.Transport = msg.Transport
		r.Spid = spid
		res = r
	}
	jsonBA,err := json.Marshal(res)
	return jsonBA,err
}


func DecodeIotMsgToJsonMsgStrV0 (msg []byte,topic string )(*IotMsg ,error){
	var iotMsg *IotMsg
	var msgType IotMsgType

	//if strings.Contains(topic,"commands"){
	//	msgType = MsgTypeCmd
	//} else {
	//	msgType = MsgTypeEvt
	//}
	v0msg := json_types.IotMsgV0{}
	err := json.Unmarshal(msg,&v0msg)
	if err != nil{
			return nil,err
	}
	if v0msg.Command.Type == "" {
		msgType = MsgTypeEvt
	}else{
		msgType = MsgTypeCmd
	}
	switch msgType {
	case MsgTypeCmd:
		iotMsg = NewIotMsg(MsgTypeCmd,v0msg.Command.Type,v0msg.Command.Subtype,nil)
		iotMsg.Default.Value = v0msg.Command.Default.Value
		iotMsg.Default.Unit = v0msg.Command.Default.Unit
		iotMsg.Properties = v0msg.Command.Properties
		iotMsg.Uuid = v0msg.UUID
		iotMsg.Transport = v0msg.Transport
	case MsgTypeEvt:
		iotMsg = NewIotMsg(MsgTypeEvt,v0msg.Event.Type,v0msg.Event.Subtype,nil)
		iotMsg.Default.Value = v0msg.Event.Default.Value
		iotMsg.Default.Unit = v0msg.Event.Default.Unit
		iotMsg.Properties = v0msg.Event.Properties
		iotMsg.Uuid = v0msg.UUID
		iotMsg.Transport = v0msg.Transport
	}
	iotMsg.Spid = v0msg.Spid
	return iotMsg , nil
}

func EncodeIotMsgToJsonStrV1 (msg *IotMsg)([]byte,error){
	msgTypeToStringMap := map[IotMsgType]string{
		MsgTypeCmd : "cmd",
		MsgTypeEvt : "evt",
		MsgTypeGet : "get",
	}

	r := json_types.IotMsgV1{}
	r.Type = msgTypeToStringMap[msg.Type]
	r.Cls = msg.Class
	r.Subcls = msg.SubClass
	r.Def.Value = msg.Default.Value
	r.Def.Unit = msg.Default.Unit
	r.Props = msg.Properties
	r.UUID = msg.Uuid
	r.Ctime = msg.Timestamp.Format(time.RFC3339)
	r.Corid = msg.Corid
	r.Ver = msg.Version
	//r.Topic = msg.Topic
	jsonBA,err := json.Marshal(r)
	return jsonBA,err
}

func DecodeIotMsgToJsonMsgStrV1 (msg []byte,topic string )(*IotMsg ,error){
	msgTypeStringToIntMap := map[string]IotMsgType{
		"cmd" : MsgTypeCmd ,
		"evt" : MsgTypeEvt ,
		"get" : MsgTypeGet ,
	}

	v1msg := json_types.IotMsgV1{}
	err := json.Unmarshal(msg,&v1msg)
	if err != nil{
		return nil,err
	}
	iotMsg := NewIotMsg(msgTypeStringToIntMap[v1msg.Type],v1msg.Cls,v1msg.Subcls,nil)
	iotMsg.Default.Value = v1msg.Def.Value
	iotMsg.Default.Unit = v1msg.Def.Unit
	iotMsg.Properties = v1msg.Props
	iotMsg.Uuid = v1msg.UUID
	iotMsg.Transport = v1msg.Transp
	//iotMsg.Timestamp = v1msg.Ctime
	iotMsg.Version = v1msg.Ver
	iotMsg.Topic = topic

	return iotMsg , nil
}
