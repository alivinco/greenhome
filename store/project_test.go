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

func  TestProjectStore_Upsert(t *testing.T) {
	prStore := NewProjectStore(session,db)
	// Insert
	//project := model.Project{Name:"StavangerHome3",Domain:"livincovi",GeoLocation:model.GeoLocation{Lat:58.955755,Long:5.691449}}
	// Update
	project := model.Project{Id:bson.ObjectIdHex("57573834554efc2c77b59f97"),Name:"StavangerHome",Domain:"livincovi",GeoLocation:model.GeoLocation{Lat:58.955755,Long:5.691449}}

	prStore.Upsert(&project)
}

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

