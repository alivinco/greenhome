package auth

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/alivinco/greenhome/model"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"golang.org/x/oauth2"
	log "github.com/Sirupsen/logrus"
)


func AuthMiddleware(store *sessions.CookieStore ) gin.HandlerFunc {
	return func(c *gin.Context) {
		authRequest := model.AuthRequest{}
		session , err := store.Get(c.Request,"gh_user")
		//authorize_url := "http://localhost:5010/greenhome/login"
		authorize_url := "/greenhome/login"
		if session.IsNew {
//			c.AbortWithError(401, err)
			c.Redirect(303,authorize_url)
			fmt.Println(err)
			authRequest.IsAuthenticated = false
			authRequest.Error = err
		}else{
			if session.Values["domain_id"].(string) != ""{
				authRequest.IsAuthenticated = true
				authRequest.Username = session.Values["username"].(string)
				authRequest.Email = session.Values["email"].(string)
				authRequest.DomainId = session.Values["domain_id"].(string)

				fmt.Println("Request is authenticated as ",session.Values)
			}else{
				fmt.Println("Request is not authenticated as ",session.Values)
				authRequest.IsAuthenticated = false
				c.Redirect(303,authorize_url)
			}
		}
		c.Set("AuthRequest",authRequest)
	}
}

func OAuth2CallbackHandler(store *sessions.CookieStore,c *gin.Context ,config *model.AppConfigs) {

	domain := "zmarlin.eu.auth0.com"

	redirectUrl := config.AppRootUrl+"/greenhome/ui/m/home"
	log.Debug("Redirect URL : ",redirectUrl)
	conf := &oauth2.Config{
		ClientID:     config.AuthClientId,
		ClientSecret: config.AuthClientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       []string{"openid", "name", "email", "nickname"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	// Getting the Code that we got from Auth0
	code := c.Query("code")
	fmt.Println("Login callback with code = ",code)
	if code != "" {
		// Exchanging the code for a token
		token, err := conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Error("Can't exchange the code for a token")
			c.AbortWithError(http.StatusInternalServerError,err)
			return
		}

		// Getting now the User information
		client := conf.Client(oauth2.NoContext, token)
		resp, err := client.Get("https://" + domain + "/userinfo")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError,err)
			return
		}

		// Reading the body
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError,err)
			return
		}

		// Unmarshalling the JSON of the Profile

		userInfo := Auth0UserInfo{}
		if err :=json.Unmarshal(body,&userInfo);err != nil{
			c.AbortWithError(http.StatusInternalServerError,err)
			return
		}
		session , _ := store.Get(c.Request,"gh_user")
		session.Values["username"] = userInfo.Name
		session.Values["email"] = userInfo.Email
		session.Values["domain_id"] = userInfo.DomainID
		session.Values["project_id"] = "57573834554efc2c77b59f97"
		if userInfo.DomainID == ""{
			fmt.Println("Error getting User info from Auth0")
			c.AbortWithError(http.StatusInternalServerError,err)
		}
		session.Save(c.Request,c.Writer)
		c.Redirect(303,redirectUrl)
	}else {
		fmt.Println("Something went wrong , token is empty.")
		c.HTML(http.StatusOK, "login.html",config)
	}


}

func Logout(store *sessions.CookieStore,c *gin.Context,config *model.AppConfigs){
	session , _ := store.Get(c.Request,"gh_user")
	session.Options.MaxAge = -1
	session.Save(c.Request,c.Writer)
	logout_url := "https://zmarlin.eu.auth0.com/logout?client_id="+config.AuthClientId
	c.Redirect(303,logout_url)

}



