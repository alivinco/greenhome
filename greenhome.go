package main

import (
	"os"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/base64"
	"github.com/alivinco/greenhome/gincontrib/jwt"
	"net/http"
	//"github.com/olahol/melody"
	"github.com/alivinco/greenhome/adapters"
	"github.com/alivinco/greenhome/routers"
	"gopkg.in/mgo.v2"
	"github.com/alivinco/greenhome/store"
)
var session *mgo.Session
var db *mgo.Database
var projectStore *store.ProjectStore
var mobileUiStore *store.MobileUiStore
var mqa *adapters.MqttAdapter

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
}

func Subscribe(){
	subs , _ := mobileUiStore.GetSubscriptions("")
	for _ , topic := range subs {
		mqa.Subscribe(topic,1)
	}
}
func Unsubscribe(){
	subs , _ := mobileUiStore.GetSubscriptions("")
	for _ , topic := range subs{
		mqa.Unsubscribe(topic)
	}
}

func RunHttpServer(bindAddress string,jwtSecret string) {
	decoded_secret, _ := base64.URLEncoding.DecodeString(jwtSecret)
	r := gin.Default()
	//m := melody.New()
	r.Static("/greenhome/static","./static")
	r.LoadHTMLGlob("templates/**/*")
	mobAppRoot := r.Group("/greenhome/ui/m")
	mobAppRoot.Use(jwt.Auth(string(decoded_secret)))
	mobAppRoot.GET("/home",func(c *gin.Context) {
			c.Get("UserData")
			//user,_:=c.Get("UserData")
			projectId := "57582d2a6dcdd112edb1278e"
			mobUi , _ := mobileUiStore.Get(projectId,"")
        		c.HTML(http.StatusOK, "start.html",mobUi[0])
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

	wsa := adapters.NewWsAdapter(r)
	mqa = adapters.NewMqttAdapter("tcp://localhost:1883","greenhome_test")
	mqa.Start()
	Subscribe()
	routers.NewMainRouter(mqa,wsa)
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
	InitDb()
	InitStores()
	RunHttpServer(bindAddress,jwtSecret)

}
