package main

import (
	"os"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	//"encoding/base64"
	"net/http"
	//"github.com/olahol/melody"
	"github.com/alivinco/greenhome/adapters"
	"github.com/alivinco/greenhome/routers"
	"gopkg.in/mgo.v2"
	"github.com/alivinco/greenhome/store"
	"github.com/alivinco/greenhome/model"
	"github.com/alivinco/greenhome/gincontrib/auth"
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	"io/ioutil"
)
var session *mgo.Session
var db *mgo.Database
var projectStore *store.ProjectStore
var mobileUiStore *store.MobileUiStore
var mqa *adapters.MqttAdapter
var thingsCacheStore *store.ThingsCacheStore
var secretFileName string
var sessionStore *sessions.CookieStore
var configs *model.AppConfigs

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
	mobileUiStore = store.NewMobileUiStore(session,db)
	mobileUiStore.SetProjectStore(projectStore)
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

func Subscribe(){
	subs , _ := mobileUiStore.GetSubscriptions("",true)
	for _ , topic := range subs {
		mqa.Subscribe(topic,1)
	}
}
func Unsubscribe(){
	subs , _ := mobileUiStore.GetSubscriptions("",true)
	for _ , topic := range subs{
		mqa.Unsubscribe(topic)
	}
}

func RunHttpServer(bindAddress string,jwtSecret string) {
	//decoded_secret, _ := base64.URLEncoding.DecodeString(jwtSecret)
	r := gin.Default()
	//m := melody.New()
	r.Static("/greenhome/static","./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/greenhome/login",func(c *gin.Context) {
			auth.OAuth2CallbackHandler(sessionStore ,c,configs)

		})
	r.GET("/greenhome/logout",func(c *gin.Context) {
			auth.Logout(sessionStore ,c)
			c.Redirect(303,"http://localhost:5010/greenhome/ui/m/home",)
		})
	mobAppRoot := r.Group("/greenhome/ui/m")
	//mobAppRoot.Use(auth.Auth(string(decoded_secret)))
	mobAppRoot.Use(auth.AuthMiddleware(sessionStore))
	mobAppRoot.GET("/home",func(c *gin.Context) {
			c.Get("UserData")
			//user,_:=c.Get("UserData")
			projectId := "57582d2a6dcdd112edb1278e"
			mobUi , _ := mobileUiStore.GetMobileUi(projectId,"")
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
	wsGroup := r.Group("/greenhome/ws")
	wsGroup.Use(auth.AuthMiddleware(sessionStore))
	wsa := adapters.NewWsAdapter(wsGroup)
	mqa = adapters.NewMqttAdapter("tcp://localhost:1883","greenhome_test")
	mqa.Start()
	Subscribe()
	routers.NewMainRouter(mqa,wsa,thingsCacheStore)
	r.Run(bindAddress)
}

func main() {
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
	fmt.Println("addr:",bindAddress)
	fmt.Println("jwt_secret:",jwtSecret)
	defer func(){
		Unsubscribe()
		session.Close()
	}()
	configs = &model.AppConfigs{}
	InitDb()
	InitStores()
	RunHttpServer(bindAddress,jwtSecret)

}
