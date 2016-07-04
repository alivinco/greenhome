package controller

import (
	"github.com/alivinco/greenhome/store"
	"github.com/gin-gonic/gin"
	"errors"
	"github.com/alivinco/greenhome/gincontrib/utils"
	"net/http"
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
			c.AbortWithError(401, errors.New("You are not allowed to access the app"))
		}

	} else {
		c.JSON(http.StatusNotFound,gin.H{"error": err.Error()})
	}

}