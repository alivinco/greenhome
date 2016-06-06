package jwt

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/alivinco/greenhome/model"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt_lib.ParseFromRequest(c.Request, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})
		authRequest := model.AuthRequest{}

		if err != nil {
//			c.AbortWithError(401, err)
			fmt.Println(err)
			authRequest.IsAuthenticated = false
			authRequest.Error = err
		}else{
			authRequest.IsAuthenticated = true
			authRequest.Email,_ = token.Claims["email"].(string)
			authRequest.Username,_ = token.Claims["nickname"].(string)
			meta,_ := token.Claims["app_metadata"].(map[string]interface{})
			authRequest.DomainName = meta["domain_name"].(string)
			authRequest.DomainId = meta["domain_id"].(string)
			fmt.Println("Domain name = ",authRequest.DomainName)
		}
		c.Set("AuthRequest",authRequest)
	}
}

