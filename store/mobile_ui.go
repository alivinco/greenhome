package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/alivinco/greenhome/model"
	"fmt"
)

type MobileUiStore struct {
	session *mgo.Session
	db *mgo.Database
	mobileUiC *mgo.Collection
}

func NewMobileUiStore(session *mgo.Session,db *mgo.Database)(*MobileUiStore){
	imst := MobileUiStore{session:session,db:db}
	imst.mobileUiC = db.C("mobile_ui")
	return &imst
}

func (ps *MobileUiStore) Upsert(mobileUi *model.MobileUi) (string,error){
	var selector bson.M
	if len(mobileUi.Id)>0 {
		selector = bson.M{"_id":mobileUi.Id}
	}else{
		selector = bson.M{"_id":bson.NewObjectId()}
	}
	info , err := ps.mobileUiC.Upsert(selector,*mobileUi)
	if err == nil {
		if info.UpsertedId != nil {
			return info.UpsertedId.(bson.ObjectId).Hex(), err
		} else {
			return "", err
		}
	} else {
		return "" , err
	}
}
func (ms *MobileUiStore) Delete(ID string) error{
	return ms.mobileUiC.RemoveId(bson.ObjectIdHex(ID))
}

func (ms *MobileUiStore) GetSubscriptions(projectId string)([]string ,error){
	var results []model.MobileUi
	//projection := bson.M{"views.things.displayelementtopic":1}
	err := ms.mobileUiC.Find(nil).All(&results)
	if err != nil {
		return nil,err
	}
	subs := []string{}

	for _ ,mobUi := range results{
		for _,view := range mobUi.Views{
			for _,thing := range view.Things{
				subs = append(subs,thing.DisplayElementTopic)
			}
		}
	}
	return subs, nil
}

// GetList returns list of all apps
func (ms *MobileUiStore) Get(id string , project_id string) ([]model.MobileUi,error){
	var results []model.MobileUi
	selector := bson.M{}
	if len(id)>0 {
		selector = bson.M{"_id":bson.ObjectIdHex(id)}
	}else if project_id != "" {
		selector = bson.M{"project":bson.ObjectIdHex(project_id)}
	}
	err := ms.mobileUiC.Find(selector).All(&results)
	fmt.Println("Results All: ", results)
	return results,err
}

