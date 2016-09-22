package iotmsglibgo

import (
	"testing"
)

func TestNewIotMsg(t *testing.T) {
	msg := NewIotMsg(MsgTypeCmd, "binary", "switch", nil)
	msg.SetDefaultStr("test value", "")
	msg.SetStrProperty("prop1","value1")
	t.Log(msg.String())
	if r := msg.GetDefaultStr(); r != "test value" {
		t.Failed()
	}
	if msg.GetStrProperty("prop1") != "value1" {
		t.Failed()
	}
}

func TestIotMsg_GetDefaultStr(t *testing.T) {
	msg := NewIotMsg(MsgTypeCmd, "binary", "test", nil)

	msg.Default.Value = "test"
	msgStr :=  msg.GetDefaultStr()
	if msgStr != "test"{
		t.Error("Wrong default value")
	}

	msg.Default.Value = 3
	msgStr =  msg.GetDefaultStr()
	if msgStr != "3"{
		t.Error("Wrong default value")
	}

	msg.Default.Value = true
	msgStr =  msg.GetDefaultStr()
	if msgStr != "true"{
		t.Error("Wrong default value")
	}

	msg.Default.Value = map[string]string{"el1":"test"}
	msgStr =  msg.GetDefaultStr()
	t.Log(msgStr)
	if msgStr != "map[el1:test]"{
		t.Error("Wrong default value")
	}
}

func TestIotMsg_GetDefaultInt(t *testing.T) {
	msg := NewIotMsg(MsgTypeCmd, "binary", "test", nil)
	msg.Default.Value = "3"
	if val , _ := msg.GetDefaultInt(); val != 3 {
		t.Error("Wrong default value = ")
	}
	msg.Default.Value = 2
	if val , _ := msg.GetDefaultInt(); val != 2 {
		t.Error("Wrong default value = ")
	}
	msg.Default.Value = 4.5
	if val , _ := msg.GetDefaultInt(); val != 4 {
		t.Error("Wrong default value = ")
	}
}