package utils
import (
	"github.com/gin-gonic/gin"
	"github.com/alivinco/blackflowhub/model"
)

func GetAuthRequest(c *gin.Context) ( *model.AuthRequest){
	authRequestUntyped , _ := c.Get("AuthRequest")
	authRequest,_ := authRequestUntyped.(model.AuthRequest)
	return  &authRequest
}
