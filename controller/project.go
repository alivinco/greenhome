package controller

import (
	"github.com/alivinco/greenhome/store"
	"github.com/gin-gonic/gin"
	"errors"
	"github.com/alivinco/greenhome/gincontrib/utils"
	"net/http"
	"github.com/alivinco/greenhome/model"
	log "github.com/Sirupsen/logrus"
)

type ProjectRestController struct {
	ProjectStore *store.ProjectStore
}

func (mst *ProjectRestController) GetProject(c *gin.Context){
	projectId := c.Param("project_id")
	result,err := mst.ProjectStore.GetById(projectId)
	auth := utils.GetAuthRequest(c)
	if err == nil {
		if auth.DomainId == result.Domain || auth.IsAuthenticated {
			c.JSON(http.StatusOK, *result)
		}else{
			log.Warn("Unauthorized request")
			c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized request"))
		}

	} else {
		log.Info(err.Error())
		c.JSON(http.StatusNotFound,gin.H{"error": err.Error()})
	}
}

func (mst *ProjectRestController) GetProjects(c *gin.Context){
	auth := utils.GetAuthRequest(c)
	if auth.IsAuthenticated {
		filter := model.Project{Domain:auth.DomainId}
		result,err := mst.ProjectStore.Get(&filter)
		if err == nil {
			c.JSON(http.StatusOK, result)
		} else {
			log.Info(err.Error())
			c.JSON(http.StatusNotFound,gin.H{"error": err.Error()})
		}
	}else {
		log.Warn("Unauthorized request")
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized request"))
	}

}

func (mst *ProjectRestController) PostProject(c *gin.Context){
	auth := utils.GetAuthRequest(c)
	if auth.IsAuthenticated {
		var project model.Project
		if c.BindJSON(&project) == nil {
			project.Domain = auth.DomainId
			mst.ProjectStore.Upsert(&project)
		}else {
			log.Warn("Can't bind model")
			c.AbortWithError(http.StatusInternalServerError, errors.New("Can't bind model"))
		}

	}else{
		log.Warn("Unauthorized request")
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized request"))
	}

}

func (mst *ProjectRestController) DeleteProject(c *gin.Context){
	projectId := c.Param("project_id")
	auth := utils.GetAuthRequest(c)
	if auth.IsAuthenticated && auth.DomainId == mst.ProjectStore.GetDomainByProjectId(projectId){
		mst.ProjectStore.Delete(projectId)
	}else{
		log.Warn("Unauthorized request")
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized request"))
	}

}