package iotmsglibgo

import (
	"testing"
	//"fmt"
)

func TestDetectPayloadType(t *testing.T) {
	topic := "/ta/zw/8/bin_switch/1/commands"
	typ , err := DetectPayloadType(topic)
	t.Log(PayloadTypeToStrMap[typ])
	if err != nil {
		t.Error(err)
	}
	if typ != PayloadTypeJsonIotMsgV0 {
		t.Fail()
	}
	topic = "jim1/evt/ta/zw/1/bin_switch/1"
	typ , err = DetectPayloadType(topic)
	t.Log(PayloadTypeToStrMap[typ])
	if err != nil {
		t.Error(err)
	}
	if typ != PayloadTypeJsonIotMsgV1 {
		t.Fail()
	}

}

func TestConvertBytesToIotMsgV0(t *testing.T) {
	jsonStr := `{   "origin": {
			  "vendor": "Sensio",
			  "@id": "smartly_ios",
			  "@type": "app"
			 },
			 "uuid": "86f93c70-6c47-4cfc-91eb-30d3b069e417",
			 "creation_time": 1385815582,
			 "command": {
			  "default": {
			   "unit": "C",
			   "value": 23.1
			  },
			  "subtype": "thermostat",
			  "target": "/dev/zw/16/lvl_thermostat/1/commands",
			  "@type": "level",
			  "properties": {
			   "setpoint_type": "heating"
			  }
			 },
			 "spid": "S-451-12",
			 "@context": "http://smartly.no/context"
			}`
	byt := []byte(jsonStr)
	iotMsg,err := ConvertBytesToIotMsg("/dev/zw/16/lvl_thermostat/1/commands",byt,nil)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if iotMsg.Type != MsgTypeCmd || iotMsg.Class != "level" {
		t.Error("Wrong msg type of class ")
	}
	t.Log(iotMsg.String())
}
func TestConvertIotMsgToBytesV0(t *testing.T) {
	config := map[string]string{"SPID":"JR-1243"}
	iotMsg := NewIotMsg(MsgTypeCmd,"level","thermostat",nil)
	iotMsg.SetDefaultFloat(20.5,"")
	iotMsg.SetStrProperty("setpoint_type","heating")
	byt ,err := ConvertIotMsgToBytes("/dev/zw/16/lvl_thermostat/1/commands",iotMsg,config)
	if err != nil {
		t.Error(err)
	}
	t.Log(" \n Type autodetect \n")
	t.Log(string(byt))
}

func TestConvertIotMsgToBytesV1_overrideType(t *testing.T) {
	config := map[string]string{"SPID":"JR-1243","override_payload_type":"jim1"}
	iotMsg := NewIotMsg(MsgTypeCmd,"level","thermostat",nil)
	iotMsg.SetDefaultFloat(20.5,"")
	iotMsg.SetStrProperty("setpoint_type","heating")
	byt ,err := ConvertIotMsgToBytes("/dev/zw/16/lvl_thermostat/1/commands",iotMsg,config)
	if err != nil {
		t.Error(err)
	}
	t.Log(" \n Overriding type to jim1 \n")
	t.Log(string(byt))
}

func TestConvertBytesToIotMsgV1(t *testing.T) {
	t.Log("TestConvertBytesToIotMsgV1")
	jsonStr :=  `{"type":"evt","cls": "binary","subcls": "switch","def": {"value": true},"props": {"p1": 165}, "ctime": "2016-05-29T15:28:26.013751", "uuid": "e48fbe58-3aaf-442d-b769-7a24aed8b716"}`
	byt := []byte(jsonStr)
	iotMsg,err := ConvertBytesToIotMsg("jim1/evt/ta/zw/1/bin_switch/1",byt,nil)
	if err != nil {
		t.Log(err)
		t.Fail()
	}else {
		t.Log(iotMsg.String())
	}

}

func TestConvertIotMsgToBytesV1(t *testing.T) {
	iotMsg := NewIotMsg(MsgTypeCmd,"level","thermostat",nil)
	iotMsg.SetDefaultFloat(20.5,"")
	iotMsg.SetStrProperty("setpoint_type","heating")
	iotMsg.Topic = "jim1/cmd/ta/zw/1/lvl_thermostat/1"
	byt ,err := ConvertIotMsgToBytes("jim1/cmd/ta/zw/1/lvl_thermostat/1",iotMsg,nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(byt))
}
