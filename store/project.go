package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/alivinco/greenhome/model"
	"fmt"
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

