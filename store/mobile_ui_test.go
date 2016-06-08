package store

import (
	"github.com/alivinco/greenhome/model"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestMobileUiStore_Upsert(t *testing.T) {
	prStore := NewMobileUiStore(session, db)
	mobileUi := model.MobileUi{
		Id:      bson.ObjectIdHex("57582d2a6dcdd112edb1278e"),
		Project: bson.ObjectIdHex("57573834554efc2c77b59f97"),
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
					DisplayElementTopic: "jim1/evt/ta/zw/3/bin_switch/1",
					ControlElementTopic: "jim1/cmd/ta/zw/3/bin_switch/1",
					UiElement:           "binary_switch",
				}, model.Thing{
					Name:                "Temp",
					Type:                "sensor.temperature",
					DisplayElementTopic: "/dev/zw/99/sen_temp/1/events",
					ControlElementTopic: "",
					UiElement:           "sensor",
					Value:               "12.3",
					Unit:                "C",
				},
			},
			},
		},
	}

	// Update
	//MobileUi := model.MobileUi{Id:bson.ObjectIdHex("57573834554efc2c77b59f97"),Name:"StavangerHome",Domain:"livincovi",GeoLocation:model.GeoLocation{Lat:58.955755,Long:5.691449}}

	prStore.Upsert(&mobileUi)
}

func TestMobileUiStore_GetSubscriptions(t *testing.T) {
	prStore := NewMobileUiStore(session, db)
	r , err :=prStore.GetSubscriptions("")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}
func TestMobileUiStore_Get(t *testing.T) {
	prStore := NewMobileUiStore(session, db)
	pr, err := prStore.Get("57582d2a6dcdd112edb1278e", "")
	if err == nil {
		t.Log(pr[0].Views[0].Name)
	} else {
		t.Error(err)
	}

}

//func TestMain(m *testing.M) {
//	var err error
//
//	session ,err = mgo.Dial("localhost")
//	if err == nil {
//		session.SetMode(mgo.Monotonic, true)
//		db = session.DB("greenhome")
//	}
//
//	os.Exit(m.Run())
//
//	session.Close()
//}
