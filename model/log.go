package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type EventLogConfig struct {
	Project bson.ObjectId `json:"project_id"`
	Subsriptions []string `json:"subscriptions"`
}

type EventLogEntry struct {
	Project bson.ObjectId `json:"project_id"`
	Topic string  `json:"topic"`
	StringValue string `json:"string_value"`
	NumericValue float32 `json:"numeric_value"`
	BoolValue bool `json:"bool_value"`
}

type UserActivityLogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Type string `json:"type"`
	Action string `json:"action"`
}
