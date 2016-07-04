package store

import (
	"testing"
	"os"
	"gopkg.in/mgo.v2"
	"github.com/alivinco/greenhome/model"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session
var db *mgo.Database

func TestProjectStore_Upsert(t *testing.T) {
	prStore := NewProjectStore(session,db)
	project := model.Project{
		Id:bson.ObjectIdHex("57573834554efc2c77b59f97"),
		Name:"StavangerHome",
		Domain:"livincovi",
		GeoLocation:model.GeoLocation{Lat:58.955755,Long:5.691449},
		Views: []model.View{
			model.View{Name: "Living room", Room: "Living", Floor: 1, Things: []model.Thing{
				model.Thing{
					Name:                "switch2",
					Type:                "bianry.switch",
					DisplayElementTopic: "jim1/evt/ta/zw/3/bin_switch/1",
					ControlElementTopic: "jim1/cmd/ta/zw/3/bin_switch/1",
					UiElement:           "binary_switch"},
			},
			},
			model.View{Name: "Home", Room: "Living", Floor: 1, Things: []model.Thing{
				model.Thing{
					Name:                "switch1",
					Type:                "bianry.switch",
					DisplayElementTopic: "jim1/evt/ta/zw/2/bin_switch/1",
					ControlElementTopic: "jim1/cmd/ta/zw/2/bin_switch/1",
					UiElement:           "binary_switch",
				}, model.Thing{
					Name:                "switch12",
					Type:                "bianry.switch",
					DisplayElementTopic: "/dev/zw/2/bin_switch/1/events",
					ControlElementTopic: "/dev/zw/2/bin_switch/1/commands",
					UiElement:           "binary_switch",
				}, model.Thing{
					Name:                "Temp",
					Type:                "sensor.temperature",
					DisplayElementTopic: "/dev/zw/99/sen_temp/1/events",
					ControlElementTopic: "",
					UiElement:           "sensor",
					Value:               "12.3",
					Unit:                "C",
				}, model.Thing{
					Name:                "Temp Living",
					Type:                "sensor.temperature",
					DisplayElementTopic: "/dev/zw/99/sen_temp/2/events",
					ControlElementTopic: "",
					UiElement:           "sensor",
					Value:               "12.3",
					Unit:                "C",
				},
			},
			},
		},
	}
	prStore.Upsert(&project)
	// Update
	//MobileUi := model.MobileUi{Id:bson.ObjectIdHex("57573834554efc2c77b59f97"),Name:"StavangerHome",Domain:"livincovi",GeoLocation:model.GeoLocation{Lat:58.955755,Long:5.691449}}

}

func TestProjectsStore_GetSubscriptions(t *testing.T) {
	prStore := NewProjectStore(session, db)
	r , err :=prStore.GetSubscriptions("",true)
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

//func  TestProjectStore_Upsert(t *testing.T) {
//	prStore := NewProjectStore(session,db)
//	// Insert
//	//project := model.Project{Name:"StavangerHome3",Domain:"livincovi",GeoLocation:model.GeoLocation{Lat:58.955755,Long:5.691449}}
//	// Update
//	project := model.Project{Id:bson.ObjectIdHex("57573834554efc2c77b59f97"),Name:"StavangerHome",Domain:"livincovi",GeoLocation:model.GeoLocation{Lat:58.955755,Long:5.691449}}
//
//	prStore.Upsert(&project)
//}

func  TestProjectStore_Get(t *testing.T) {
	prStore := NewProjectStore(session,db)
	project := model.Project{Id:bson.ObjectIdHex("57573834554efc2c77b59f97")}
	pr ,err := prStore.Get(&project)
	if err == nil{
		t.Log(pr[0].Name)
	}else{
		t.Error(err)
	}
}

func  TestProjectStore_GetById(t *testing.T) {
	prStore := NewProjectStore(session,db)
	pr ,err := prStore.GetById("57573834554efc2c77b59f97")
	if err == nil{
		t.Log(pr.Name)
	}else{
		t.Error(err)
	}

}

func TestMain(m *testing.M) {
	var err error

	session ,err = mgo.Dial("localhost")
	if err == nil {
		session.SetMode(mgo.Monotonic, true)
		db = session.DB("greenhome")
	}

	os.Exit(m.Run())

	session.Close()
}

