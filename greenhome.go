package main

import (
	"os"
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
	session ,err = mgo.Dial("localhost")
	if err == nil {
		session.SetMode(mgo.Monotonic, true)
		db = session.DB("greenhome")
	}
}
func InitStores(){
	projectStore = store.NewProjectStore(session,db)
	thingsCacheStore = store.NewThingsCacheStore()
	secretFileName = "./sessionsecret.db"
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
		mqa = adapters.NewMqttAdapter("tcp://localhost:1883","greenhome_test")
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

func InitHttpServer(bindAddress string,jwtSecret string)(*gin.Engine) {
	//decoded_secret, _ := base64.URLEncoding.DecodeString(jwtSecret)
	r := gin.Default()
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
			c.Get("UserData")
			//user,_:=c.Get("UserData")
			projectId := "57573834554efc2c77b59f97"
			mobUi , _ := projectStore.GetById(projectId)
			session , _ := sessionStore.Get(c.Request,"gh_user")
			domain := session.Values["domain_id"].(string)
			ctx := model.Context{domain}
			store.ExtendMobileUiWithValue(thingsCacheStore,mobUi,&ctx)
        		c.HTML(http.StatusOK, "start.html",gin.H{"mobUi":mobUi,"domain":domain})
		})
	mobAppRoot.GET("/security",func(c *gin.Context) {
			c.Get("UserData")
			//user,_:=c.Get("UserData")
        		c.HTML(http.StatusOK, "security.html",gin.H{})
		})
	mobAppRoot.GET("/rooms",func(c *gin.Context) {
			c.Get("UserData")
			//user,_:=c.Get("UserData")
        		c.HTML(http.StatusOK, "rooms.html",gin.H{})
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

	wsGroup = r.Group("/greenhome/ws")
	wsGroup.Use(auth.AuthMiddleware(sessionStore))
	return r
}

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp:true,ForceColors:true})
	log.SetLevel(log.DebugLevel)
	bindAddress := ":6000"
	jwtSecret := ""
    // Load configs from env variable or from command line .
	bindAddress = os.Getenv("BFH_BIND_ADDR")
	if bindAddress != "" {
		jwtSecret = os.Getenv("BFH_JWT_SECRET")
	}else{
		flag.StringVar(&bindAddress,"addr",":5010","Server bind address")
		flag.StringVar(&jwtSecret,"jwt_secret","","Jwt secret")
	}

	flag.Parse()
	log.Info("addr:",bindAddress)
	log.Info("jwt_secret:",jwtSecret)
	defer func(){
		UnsubscribeMqttTopics()
		session.Close()
	}()
	configs = &model.AppConfigs{}
	configs.AuthClientId = "njwDYXaCFOS2TzTHGQaBUTk8GiXNgLti"
	configs.AuthClientSecret = "T2kdCk2kTrbprreq2Dlc-qm5klDTjd5UAzHASWFPlehO4yAwoxfilnUgLoGMmR1p"
	configs.AppRootUrl = "http://192.168.80.237:5010"
	InitDb()
	InitStores()
	r := InitHttpServer(bindAddress,jwtSecret)
	InitAdaptersAndMainRouter()
	r.Run(bindAddress)

}
