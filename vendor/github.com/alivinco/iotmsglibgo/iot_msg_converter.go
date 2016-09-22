package iotmsglibgo

import (
	"strings"
	"errors"
)

func DetectPayloadType(topic string)(IotMsgPayloadType,error){
	tsplit := strings.Split(topic,"/")
	if len(tsplit)>0{
		if tsplit[0] != ""{
			return PayloadStrToTypeMap[tsplit[0]] , nil
		}else{
			return PayloadTypeJsonIotMsgV0 , nil
		}
	}else {
		return 0,errors.New("Type can't be detected.")
	}
}
// Converts bytes slice into IotMsg
func ConvertBytesToIotMsg(topic string,byteMsg []byte,configs map[string]string)(*IotMsg,error)  {
	ptype,s := configs["override_payload_type"]
	var payloadType IotMsgPayloadType
	var err error
	if s {
		payloadType,err = DetectPayloadType(ptype)
	}else {
		payloadType, err = DetectPayloadType(topic)
	}
	if err != nil {
		return nil,err
	}

	switch payloadType {
	case PayloadTypeJsonIotMsgV0:
		return DecodeIotMsgToJsonMsgStrV0(byteMsg,topic)
	case PayloadTypeJsonIotMsgV1:
		return DecodeIotMsgToJsonMsgStrV1(byteMsg,topic)
	default:
		return nil , errors.New("Unknown payload type")
	}

}

func ConvertIotMsgToBytes(topic string ,iotMsg *IotMsg , configs map[string]string)([]byte,error){
	ptype,s := configs["override_payload_type"]
	var payloadType IotMsgPayloadType
	var err error
	if s {
		payloadType,err = DetectPayloadType(ptype)
	}else {
		payloadType, err = DetectPayloadType(topic)
	}
	if err != nil {
		return nil,err
	}
	spid ,_ := configs["SPID"]
	switch payloadType {
	case PayloadTypeJsonIotMsgV0:
		return EncodeIotMsgToJsonStrV0(iotMsg,spid)
	case PayloadTypeJsonIotMsgV1:
		return EncodeIotMsgToJsonStrV1(iotMsg)
	default:
		return nil , errors.New("Unknown payload type")
	}
}