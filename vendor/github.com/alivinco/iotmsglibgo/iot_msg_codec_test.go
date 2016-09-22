package iotmsglibgo

import (
	"testing"
)

func TestEncodeIotMsgToJsonMsgV0(t *testing.T) {
	msg := NewIotMsg(PayloadTypeJsonIotMsgV1, "binary", "switch", nil)
	msg.SetDefaultStr("test value", "")
	msg.SetStrProperty("prop1","value1")
	byt , err := EncodeIotMsgToJsonStrV0(msg,"SPId1")
	if err != nil{
		t.Error(err)
	}
 	t.Log(string(byt))


}
func TestDecodedeIotMsgToJsonMsgV0(t *testing.T) {
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
	iotMsg ,err := DecodeIotMsgToJsonMsgStrV0(byt,"/dev/zw/16/lvl_thermostat/1/commands")
	if err != nil {
		t.Error(err)
	}
	t.Log(iotMsg.String())
	if iotMsg.Default.Value != 23.1 {
		t.Fail()
	}
	if iotMsg.GetDefaultFloat() != 23.1 {
		t.Fail()
	}
}