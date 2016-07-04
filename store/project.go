package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/alivinco/greenhome/model"
	"fmt"
	"github.com/alivinco/greenhome/adapters"
	log "github.com/Sirupsen/logrus"
)

type ProjectStore struct {
	session *mgo.Session
	db *mgo.Database
	projectC *mgo.Collection
}

func NewProjectStore(session *mgo.Session,db *mgo.Database)(*ProjectStore){
	imst := ProjectStore{session:session,db:db}
	imst.projectC = db.C("projects")
	return &imst
}

func (ps *ProjectStore) Upsert(project *model.Project) (string,error){
	var selector bson.M
	if len(project.Id)>0 {
		selector = bson.M{"_id":project.Id}
	}else{
		selector = bson.M{"_id":bson.NewObjectId()}
	}
	info , err := ps.projectC.Upsert(selector,*project)
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
func (ms *ProjectStore) Delete(ID string) error{
	return ms.projectC.RemoveId(bson.ObjectIdHex(ID))
}

// GetList returns list of all apps
func (ms *ProjectStore) Get(filter *model.Project) ([]model.Project,error){
	var results []model.Project
	selector := bson.M{}
	if len(filter.Id)>0 {
		selector = bson.M{"_id":filter.Id}
	}else if filter.Domain != "" {
		selector = bson.M{"domain":filter.Domain}
	}
	err := ms.projectC.Find(selector).All(&results)
	fmt.Println("Results All: ", results)
	return results,err
}
// GetList returns list of all apps
func (ms *ProjectStore) GetById(id string) (*model.Project,error){
	result := model.Project{}
	err := ms.projectC.FindId(bson.ObjectIdHex(id)).One(&result)
	fmt.Println("Results All: ", result)
	if err == nil{
		return &result,nil
	}else {
		return nil,err
	}

}

func (ms *ProjectStore) GetSubscriptions(projectId string , global bool)([]string ,error){
	var results []model.Project
	//projection := bson.M{"views.things.displayelementtopic":1}
	err := ms.projectC.Find(nil).All(&results)
	if err != nil {
		return nil,err
	}
	subs := []string{}

	for _ ,project := range results{
		for _,view := range project.Views{
			for _,thing := range view.Things{
				if global{
					subs = append(subs,adapters.AddDomainToTopic(project.Domain,thing.DisplayElementTopic))
				}else {
					subs = append(subs,thing.DisplayElementTopic)
				}

			}
		}
	}
	return subs, nil
}

func ExtendMobileUiWithValue(cache *ThingsCacheStore , mobUi *model.Project , ctx *model.Context ){
	log.Debug("Extending mobUi for domain =",ctx.Domain)
	for view_i,view := range mobUi.Views{
			for thing_i,thing := range view.Things{
				value  := cache.Get(thing.DisplayElementTopic,ctx)
				if value != nil {
					mobUi.Views[view_i].Things[thing_i].Value = value.Default.Value
					log.Debug("Value from cache=",thing.Value)

				}else{
					log.Debug("No entry in cache")
				}

			}
		}
}
