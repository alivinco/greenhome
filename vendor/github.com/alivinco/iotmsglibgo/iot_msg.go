package iotmsglibgo

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

const (
	MsgTypeCmd = 1
	MsgTypeEvt = 2
	MsgTypeGet = 3

	PayloadTypeJsonIotMsgV0   = 0
	PayloadTypeJsonIotMsgV1   = 1
	PayloadTypeJsonOpaque     = 2
	PayloadTypeBinaryIotMsgV1 = 3
	PayloadTypeBinaryValue    = 4
	PayloadTypeBinaryOpaque   = 5
)

var PayloadStrToTypeMap = map[string]IotMsgPayloadType{
	"jim0": PayloadTypeJsonIotMsgV0,
	"jim1": PayloadTypeJsonIotMsgV1,
	"jopq": PayloadTypeJsonOpaque,
	"bim1": PayloadTypeBinaryIotMsgV1,
	"bval": PayloadTypeBinaryValue,
	"bopq": PayloadTypeBinaryOpaque,
}

var PayloadTypeToStrMap = map[IotMsgPayloadType]string{
	PayloadTypeJsonIotMsgV0:   "jim0",
	PayloadTypeJsonIotMsgV1:   "jim1",
	PayloadTypeJsonOpaque:     "jopq",
	PayloadTypeBinaryIotMsgV1: "bim1",
	PayloadTypeBinaryValue:    "bval",
	PayloadTypeBinaryOpaque:   "bopq",
}

type IotMsgDefault struct {
	Value interface{}
	Unit  string
	Type  string
}
type IotMsgProperties map[string]interface{}
type IotMsgType uint8
type IotMsgPayloadType uint8

type IotMsg struct {
	Origin     string
	Type       IotMsgType
	Class      string
	SubClass   string
	Default    IotMsgDefault
	Properties IotMsgProperties
	Topic      string
	Timestamp  time.Time
	Uuid       string
	Corid      string
	Transport  string
	Version    float32
	Spid	   string
}

func NewIotMsg(msgType IotMsgType, msgClass string, msgSubClass string, reqMsg *IotMsg) *IotMsg {
	iotMsg := IotMsg{Type: msgType,
		Class:     msgClass,
		SubClass:  msgSubClass,
		Timestamp: time.Now(),
		Uuid:      uuid.NewV4().String(),
	}
	if reqMsg != nil {
		iotMsg.Corid = reqMsg.Uuid
	}
	return &iotMsg
}

func (msg *IotMsg) SetDefaultStr(value string, unit string) {
	msg.Default = IotMsgDefault{Value: value, Unit: unit, Type: "string"}
}
func (msg *IotMsg) GetDefaultStr() (string) {
	return fmt.Sprintf("%v",msg.Default.Value)
}
func (msg *IotMsg) SetDefaultBool(value bool, unit string) {
	msg.Default = IotMsgDefault{Value: value, Unit: unit, Type: "bool"}
}
func (msg *IotMsg) GetDefaultBool() bool {
	return msg.Default.Value.(bool)
}
func (msg *IotMsg) SetDefaultInt(value int, unit string) {
	msg.Default = IotMsgDefault{Value: value, Unit: unit, Type: "int"}
}
func (msg *IotMsg) GetDefaultInt() (int, error) {
	switch v := msg.Default.Value.(type) {
	case float64 , float32:
		return int(v.(float64)), nil
	case string:
		return strconv.Atoi(v)
	case int:
		return v, nil
	case bool:
		if v {
			return 1 , nil
		} else {
			return 0 , nil
		}
	default:
		return 0, errors.New("Variable can't be converted into int")
	}
}
func (msg *IotMsg) SetDefaultFloat(value float64, unit string) {
	msg.Default = IotMsgDefault{Value: value, Unit: unit, Type: "float"}
}
func (msg *IotMsg) GetDefaultFloat() float64 {
	return msg.Default.Value.(float64)
}
func (msg *IotMsg) SetDefault(value interface{}, unit string) {
	msg.Default = IotMsgDefault{Value: value, Unit: unit, Type: ""}
}
func (msg *IotMsg) GetDefault() interface{} {
	return msg.Default.Value
}
func (msg *IotMsg) GetProperties() *IotMsgProperties {
	return &msg.Properties
}
func (msg *IotMsg) SetProperties(properties IotMsgProperties) {
	msg.Properties = properties
}
func (msg *IotMsg) GetStrProperty(key string) string {
	return msg.Properties[key].(string)
}
func (msg *IotMsg) SetStrProperty(key string, value string) {
	if msg.Properties == nil {
		msg.Properties = IotMsgProperties{}
	}
	msg.Properties[key] = value
}
func (msg *IotMsg) String() string {
	return fmt.Sprintf("Type = %v , Class = %v , SubClass = %v \n Default = %v \n Properties = %v \n UUID = %v \n Timestamp = %v \n",
		msg.Type, msg.Class, msg.SubClass, msg.Default, msg.Properties, msg.Uuid, msg.Timestamp.Format(time.RFC3339))
}
