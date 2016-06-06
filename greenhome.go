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
)

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
        		c.HTML(http.StatusOK, "start.html",gin.H{})
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

	//r.GET("/greenhome/ws", func(c *gin.Context) {
	//	m.HandleRequest(c.Writer, c.Request)
	//})
	//
	//m.HandleMessage(func(s *melody.Session, msg []byte) {
	//	m.Broadcast(msg)
	//})
	//m.HandleConnect(func(s *melody.Session) {
	//	fmt.Println("Client connected from ",s.Request.URL)
	//})
	//m.HandleDisconnect(func(s *melody.Session) {
	//	fmt.Println("Client disconnected from ",s.Request.URL)
	//})
	wsa := adapters.NewWsAdapter(r)
	mqa := adapters.NewMqttAdapter("tcp://localhost:1883","greenhome_test")
	mqa.Start()
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
	RunHttpServer(bindAddress,jwtSecret)
}
