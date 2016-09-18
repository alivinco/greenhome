package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/alivinco/greenhome/model"
	"fmt"
	"github.com/alivinco/greenhome/adapters"
	log "github.com/Sirupsen/logrus"
	"github.com/alivinco/iotmsglibgo"
	"errors"
)

type ProjectStore struct {
	session *mgo.Session
	db *mgo.Database
	projectC *mgo.Collection
	topicChangeHandler TopicChangeHandler
}
// callback function is invoked whenever topic is updated
// isSub - indicates if handlers should subscribe or unsubscibe from topics listed in first param
type TopicChangeHandler func (topics []string ,isSub bool, ctx *model.Context)

func NewProjectStore(session *mgo.Session,db *mgo.Database)(*ProjectStore){
	imst := ProjectStore{session:session,db:db}
	imst.projectC = db.C("projects")
	return &imst
}

func (ps *ProjectStore) SetTopicChangeHandler(topicChangeHandler TopicChangeHandler){
	ps.topicChangeHandler = topicChangeHandler
}

func (ps *ProjectStore) Upsert(project *model.Project) (string,error){
	var selector bson.M
	var oldProject *model.Project
	if len(project.Id)>0 {
		selector = bson.M{"_id":project.Id}
		err := ps.projectC.FindId(project.Id).One(&oldProject)
		log.Error(err)
	}else{
		selector = bson.M{"_id":bson.NewObjectId()}
		oldProject = nil
	}
	for vi,view := range project.Views {
		if view.Id == "" {
			project.Views[vi].Id = bson.NewObjectId()
		}
		for thi,thing := range view.Things {
			if thing.Id == "" {
				project.Views[vi].Things[thi].Id = bson.NewObjectId()
			}
		}
	}

	info , err := ps.projectC.Upsert(selector,*project)
	if err == nil {
		subT , unsubT := GetUpdatedTopics(project,oldProject)
		log.Info("Topics for sub:",subT)
		log.Info("Topics to unsub:",unsubT)
		ctx := model.Context{Domain:project.Domain}
		if ps.topicChangeHandler != nil{
			ps.topicChangeHandler(subT,true,&ctx)
			ps.topicChangeHandler(unsubT,false,&ctx)
		}
		if info.UpsertedId != nil {
			return info.UpsertedId.(bson.ObjectId).Hex(), err
		} else {
			return "", err
		}
	} else {
		return "" , err
	}
}
func (ps *ProjectStore) Delete(ID string) error{
	ctx := model.Context{ps.GetDomainByProjectId(ID)}
	if ctx.Domain == "" {
		return errors.New("Invalid project id ")
	}
	unsubTopics , _ := ps.GetSubscriptions(ID,true)
	if ps.topicChangeHandler != nil{
			ps.topicChangeHandler(unsubTopics,false,&ctx)
	}
	return ps.projectC.RemoveId(bson.ObjectIdHex(ID))
}

// GetList returns list of all apps
func (ms *ProjectStore) Get(filter *model.Project) ([]model.Project,error){
	results := []model.Project{}
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
	if id == "" {
		return nil , errors.New("Id can't be empty.")
	}
	result := model.Project{}
	err := ms.projectC.FindId(bson.ObjectIdHex(id)).One(&result)
	fmt.Println("Results All: ", result)
	if err == nil{
		return &result,nil
	}else {
		return nil,err
	}

}

// GetList returns list of all apps
func (ms *ProjectStore) GetDomainByProjectId(id string) (string){
	result := model.Project{}
	//var domain string
	err := ms.projectC.FindId(bson.ObjectIdHex(id)).Select(bson.M{"domain":1}).One(&result)
	if err == nil{
		return result.Domain
	}else {
		log.Error("Error while domain lookup",err)
		return ""
	}

}

func (ms *ProjectStore) GetSubscriptions(projectId string , global bool)([]string ,error){
	var results []model.Project
	//projection := bson.M{"views.things.displayelementtopic":1}
	var err error
	if projectId != "" {
		err = ms.projectC.FindId(bson.ObjectIdHex(projectId)).All(&results)
	}else {
		err = ms.projectC.Find(nil).All(&results)
	}

	if err != nil {
		return nil,err
	}
	subs := []string{}

	for _ ,project := range results{
		for _,view := range project.Views{
			for _,thing := range view.Things{
				if thing.DisplayElementTopic != ""{
					if global{
						subs = append(subs,adapters.AddDomainToTopic(project.Domain,thing.DisplayElementTopic))
					}else {
						subs = append(subs,thing.DisplayElementTopic)
					}
				}


			}
		}
	}
	return subs, nil
}

// Return array of modified topics .
func GetUpdatedTopics(newProject *model.Project , oldProject *model.Project)(subTopics,unsubTopics []string){
	for _,viewN := range newProject.Views {
		for _, thingN := range viewN.Things {
			if oldProject != nil {
				for _,viewO := range oldProject.Views {
					if viewN.Id == viewO.Id{
						for _, thingO := range viewO.Things {
							if thingN.Id == thingO.Id && thingN.DisplayElementTopic != thingO.DisplayElementTopic {
								subTopics = append(subTopics,thingN.DisplayElementTopic)
								unsubTopics = append(unsubTopics,thingO.DisplayElementTopic)
							}
						}
					}
				}
			}else {
				subTopics = append(subTopics,thingN.DisplayElementTopic)
			}
		}
	}
	return
}

func ExtendThingsWithValues(cache *ThingsCacheStore , project *model.Project , ctx *model.Context ){
	log.Debug("Extending project with values for domain =",ctx.Domain)
	var value *iotmsglibgo.IotMsg
	for view_i,view := range project.Views{
			for thing_i,thing := range view.Things{

				if thing.DisplayElementTopic != ""{
					value  = cache.Get(thing.DisplayElementTopic,ctx)
				} else {
					value  = cache.Get(thing.ControlElementTopic,ctx)
				}

				if value != nil {
					project.Views[view_i].Things[thing_i].Value = value.Default.Value
					log.Debug("Value from cache=",thing.Value)

				}else{
					log.Debug("No entry in cache")
				}

			}
		}
}



