package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

var hub *Hub
var teams []string

func init() {
	gin.SetMode(gin.ReleaseMode)

	hub = NewHub()

	// Load team file
	teamData, err := ioutil.ReadFile("team.txt")
	if err != nil {
		log.Panicln(err)
	}
	teams = strings.Split(string(teamData), "\n")
}

func main() {
	token := flag.String("token", randomString(32), "Authorization Token")
	port := flag.String("port", "12345", "HTTP Listening Port")
	fmt.Println("===== Teams =====")
	for k, v := range teams {
		fmt.Printf("%2d - %s\n", k, v)
	}
	fmt.Printf("\ntoken: %s\n\n", *token)

	r := gin.Default()
	r.GET("/websocket", func(c *gin.Context) {
		ServeWebSocket(hub, c)
	})
	auth := r.Use(func(c *gin.Context) {
		if c.GetHeader("Authorization") != *token {
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
	log.Printf("Listening and serving HTTP on :%s\n", *port)
	log.Panicln(r.Run(":" + *port))
}
