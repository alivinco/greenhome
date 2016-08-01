package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alivinco/greenhome/adapters"
	"github.com/alivinco/greenhome/routers"
	"gopkg.in/mgo.v2"
	"github.com/alivinco/greenhome/store"
	"github.com/alivinco/greenhome/model"
	"github.com/alivinco/greenhome/gincontrib/auth"
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"github.com/alivinco/greenhome/controller"
	"github.com/caarlos0/env"
	"github.com/BurntSushi/toml"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)
var session *mgo.Session
var db *mgo.Database
var projectStore *store.ProjectStore
var mqa *adapters.MqttAdapter
var wsa *adapters.WsAdapter
var thingsCacheStore *store.ThingsCacheStore
var secretFileName string
var sessionStore *sessions.CookieStore
var configs *model.AppConfigs
var wsGroup *gin.RouterGroup

func InitDb(){
	var err error
	session ,err = mgo.Dial(configs.MongoConnUri)
	if err == nil {
		session.SetMode(mgo.Monotonic, true)
		db = session.DB("greenhome")
	}
}
func InitStores(){
	projectStore = store.NewProjectStore(session,db)
	thingsCacheStore = store.NewThingsCacheStore()
	secretFileName = configs.SessionStoreFile
	//secretFileName = "./sessionsecret.db"
	var err error
	sessionSecret , err := ioutil.ReadFile(secretFileName)
	if err != nil {
		sessionSecret = securecookie.GenerateRandomKey(24)
		ioutil.WriteFile(secretFileName,sessionSecret,0777)
	}
	securecookie.GenerateRandomKey(24)
	//sessionStore = sessions.NewFilesystemStore("./",sessionSecret)
	sessionStore = sessions.NewCookieStore(sessionSecret)
}

func InitAdaptersAndMainRouter(){
	if wsGroup != nil{
		wsa = adapters.NewWsAdapter(wsGroup)
		mqa = adapters.NewMqttAdapter(configs.MqttBrokerUri,configs.MqttConnClientId,configs.MqttConnUsername,configs.MqttConnPassword)
		//mqa = adapters.NewMqttAdapter("tcp://localhost:1883","greenhome_test")
		projectStore.SetTopicChangeHandler(mqa.TopicChangeHandler)
		err := mqa.Start()
		if err !=nil {
			log.Fatal("Can't connect to mqtt broker. ",err)

			panic(err)
		}
		SubscribeMqttTopics()
		routers.NewMainRouter(mqa,wsa,thingsCacheStore)
	}else{
		log.Fatal("Ws Group is not initialized. Initialize it first.")
	}

}

func SubscribeMqttTopics(){
	subs , _ := projectStore.GetSubscriptions("",true)
	for _ , topic := range subs {
		mqa.Subscribe(topic,1)
	}
}
func UnsubscribeMqttTopics(){
	subs , _ := projectStore.GetSubscriptions("",true)
	for _ , topic := range subs{
		mqa.Unsubscribe(topic)
	}
}

func GetProject(c *gin.Context)(project *model.Project , domain string){
			session , _ := sessionStore.Get(c.Request,"gh_user")
			project , _ = projectStore.GetById(session.Values["project_id"].(string))
			domain = session.Values["domain_id"].(string)
			ctx := model.Context{domain}
			store.ExtendThingsWithValues(thingsCacheStore,project,&ctx)
			return
	}

func InitHttpServer(bindAddress string,jwtSecret string)(*gin.Engine) {
	//decoded_secret, _ := base64.URLEncoding.DecodeString(jwtSecret)
	r := gin.Default()
	//projectId := "57573834554efc2c77b59f97"

	//m := melody.New()
	r.Static("/greenhome/static","./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/greenhome/login",func(c *gin.Context) {
			auth.OAuth2CallbackHandler(sessionStore ,c,configs)

		})
	r.GET("/greenhome/logout",func(c *gin.Context) {
			auth.Logout(sessionStore ,c,configs)
			c.Redirect(303, configs.AppRootUrl+"/greenhome/ui/m/home",)
		})
	mobAppRoot := r.Group("/greenhome/ui/m")
	//mobAppRoot.Use(auth.Auth(string(decoded_secret)))
	mobAppRoot.Use(auth.AuthMiddleware(sessionStore))
	mobAppRoot.GET("/home",func(c *gin.Context) {
			project , domain := GetProject(c)
        		c.HTML(http.StatusOK, "start.html",gin.H{"project":project,"domain":domain})
		})
	mobAppRoot.GET("/view/:view_id",func(c *gin.Context) {
			project , domain := GetProject(c)
			viewId , _ := c.Params.Get("view_id")
			var view *model.View
			for _ , v :=  range project.Views{
				if v.Id == bson.ObjectIdHex(viewId){
					view = &v
					break
				}
			}
        		c.HTML(http.StatusOK, "view.html",gin.H{"view":view,"domain":domain,"view_id":viewId})
		})
	mobAppRoot.GET("/security",func(c *gin.Context) {
			project , domain := GetProject(c)
			c.HTML(http.StatusOK, "security.html",gin.H{"project":project,"domain":domain})
		})
	mobAppRoot.GET("/rooms",func(c *gin.Context) {
			c.Get("UserData")
			//user,_:=c.Get("UserData")
			project , domain := GetProject(c)
        		c.HTML(http.StatusOK, "rooms.html",gin.H{"project":project,"domain":domain})
		})
	mobAppRoot.GET("/logs",func(c *gin.Context) {
			c.Get("UserData")
			//user,_:=c.Get("UserData")
        		c.HTML(http.StatusOK, "logs.html",gin.H{})
		})
	// ADMIN UI
	adminAppRoot := r.Group("/greenhome/ui/adm")
	adminAppRoot.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html",gin.H{})
	})
	// REST API
	projectController := controller.ProjectRestController{projectStore}
	apiAppRoot := r.Group("/greenhome/api")
	apiAppRoot.Use(auth.AuthMiddleware(sessionStore))
	apiAppRoot.GET("/project/:project_id",projectController.GetProject)
	apiAppRoot.DELETE("/project/:project_id",projectController.DeleteProject)
	apiAppRoot.GET("/projects",projectController.GetProjects)
	apiAppRoot.POST("/project",projectController.PostProject)
	// WS Endpoint
	wsGroup = r.Group("/greenhome/ws")
	wsGroup.Use(auth.AuthMiddleware(sessionStore))
	return r
}

func LoadConfigs(){
	configs = &model.AppConfigs{}
    	var configFile string
	flag.StringVar(&configFile,"c","","Config file")
	flag.Parse()
	if configFile != "" {
		fmt.Println("Loading configs from file ",configFile)
		if _, err := toml.DecodeFile(configFile,configs);err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Loading configs from ENV .")
		if err := env.Parse(configs);err != nil {
			panic(err)
		}
	}
	fmt.Println("Starting GreenHome with paramters")
	fmt.Printf("%+v\n", configs)
}

func main() {
	LoadConfigs()
	log.SetFormatter(&log.TextFormatter{FullTimestamp:true,ForceColors:true})
	log.SetLevel(log.DebugLevel)
	defer func(){
		UnsubscribeMqttTopics()
		session.Close()
	}()
	InitDb()
	InitStores()
	r := InitHttpServer(configs.BindAddress,configs.JwtSecret)
	InitAdaptersAndMainRouter()
	r.Run(configs.BindAddress)

}
