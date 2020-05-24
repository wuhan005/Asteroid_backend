package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

var token, port string
var hub *Hub
var teams []string

func init() {
	gin.SetMode(gin.ReleaseMode)

	hub = NewHub()
	flag.StringVar(&token, "token", randomString(32), "Authorization Token")
	flag.StringVar(&port, "port", "12345", "HTTP Listening Port")

	// Load team file
	teamData, err := ioutil.ReadFile("team.txt")
	if err != nil {
		log.Panicln(err)
	}
	teams = strings.Split(string(teamData), "\n")
}

func main() {
	flag.Parse()

	fmt.Println("===== Teams =====")
	for k, v := range teams {
		fmt.Printf("%2d - %s\n", k, v)
	}
	fmt.Printf("\ntoken: %s\n\n", token)

	r := gin.Default()
	r.GET("/websocket", func(c *gin.Context) {
		ServeWebSocket(hub, c)
	})
	auth := r.Use(func(c *gin.Context) {
		if c.GetHeader("Authorization") != token {
			c.JSON(makeErrJSON(403, 40300, "unauthorized"))
			c.Abort()
			return
		}
		c.Next()
	})

	auth.POST("/attack", attackHandler)
	auth.POST("/rank", rankHandler)
	auth.POST("/status", statusHandler)
	auth.POST("/round", roundHandler)
	auth.GET("/easterEgg", eggHandler)
	auth.POST("/time", timeHandler)
	auth.POST("/clear", clearHandler)
	auth.POST("/clearAll", clearAllHandler)

	go hub.Run()
	log.Printf("Listening and serving HTTP on :%s\n", port)
	log.Panicln(r.Run(":" + port))
}
