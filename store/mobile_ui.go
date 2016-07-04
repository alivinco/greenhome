package store
//
//import (
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
//	"github.com/alivinco/greenhome/model"
//	"fmt"
//	"github.com/alivinco/greenhome/adapters"
//)
//
//type MobileUiStore struct {
//	session *mgo.Session
//	db *mgo.Database
//	mobileUiC *mgo.Collection
//	projectStore *ProjectStore
//}
//
//func NewMobileUiStore(session *mgo.Session,db *mgo.Database)(*MobileUiStore){
//	imst := MobileUiStore{session:session,db:db}
//	imst.mobileUiC = db.C("mobile_ui")
//	return &imst
//}
//
//func (ms *MobileUiStore) SetProjectStore(projectStore *ProjectStore){
//	ms.projectStore = projectStore
//}
//
//func (ps *MobileUiStore) Upsert(mobileUi *model.MobileUi) (string,error){
//	var selector bson.M
//	if len(mobileUi.Id)>0 {
//		selector = bson.M{"_id":mobileUi.Id}
//	}else{
//		selector = bson.M{"_id":bson.NewObjectId()}
//	}
//	info , err := ps.mobileUiC.Upsert(selector,*mobileUi)
//	if err == nil {
//		if info.UpsertedId != nil {
//			return info.UpsertedId.(bson.ObjectId).Hex(), err
//		} else {
//			return "", err
//		}
//	} else {
//		return "" , err
//	}
//}
//func (ms *MobileUiStore) Delete(ID string) error{
//	return ms.mobileUiC.RemoveId(bson.ObjectIdHex(ID))
//}
//
//func (ms *MobileUiStore) UpdateThingValue(topic string , value string){
//
//
//}
//
//func (ms *MobileUiStore) GetSubscriptions(projectId string , global bool)([]string ,error){
//	var results []model.MobileUi
//	//projection := bson.M{"views.things.displayelementtopic":1}
//	err := ms.mobileUiC.Find(nil).All(&results)
//	if err != nil {
//		return nil,err
//	}
//	subs := []string{}
//
//	for _ ,mobUi := range results{
//		project , _ := ms.projectStore.GetById(mobUi.Project.Hex())
//		for _,view := range mobUi.Views{
//			for _,thing := range view.Things{
//				if global{
//					subs = append(subs,adapters.AddDomainToTopic(project.Domain,thing.DisplayElementTopic))
//				}else {
//					subs = append(subs,thing.DisplayElementTopic)
//				}
//
//			}
//		}
//	}
//	return subs, nil
//}
//
//// GetList returns list of all apps
//func (ms *MobileUiStore) GetMobileUi(id string , projectId string) (*model.MobileUi,error){
//	var results model.MobileUi
//	selector := bson.M{}
//	if len(id)>0 {
//		selector = bson.M{"_id":bson.ObjectIdHex(id)}
//	}else if projectId != "" {
//		selector = bson.M{"project":bson.ObjectIdHex(projectId)}
//	}
//	err := ms.mobileUiC.Find(selector).One(&results)
//	fmt.Println("Results All: ", results)
//	return &results,err
//}
//
////func ExtendMobileUiWithValue(cache *ThingsCacheStore , mobUi *model.MobileUi , ctx *model.Context ){
////	fmt.Println("Extending mobUi for domain",ctx.Domain)
////	for view_i,view := range mobUi.Views{
////			for thing_i,thing := range view.Things{
////				value  := cache.Get(thing.DisplayElementTopic,ctx)
////				if value != nil {
////					mobUi.Views[view_i].Things[thing_i].Value = value.Default.Value
////					fmt.Println("Value from cache=",thing.Value)
////
////				}else{
////					fmt.Println("No entry in cache")
////				}
////
////			}
////		}
////}
//
