package main

import (
	"flag"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	log "unknwon.dev/clog/v2"
)

var token, port string
var hub *Hub
var teams []team

func init() {
	_ = log.NewConsole()
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	// Parse CLI.
	flag.StringVar(&token, "token", randomString(32), "Authorization Token")
	flag.StringVar(&port, "port", "12345", "HTTP Listening Port")
	flag.Parse()

	// Load team file.
	teamData, err := ioutil.ReadFile("team.txt")
	if err != nil {
		log.Fatal("Failed to read file team.txt: %v", err)
	}

	teamsName := strings.Split(string(teamData), "\n")
	for index, t := range teamsName {
		teams = append(teams, team{
			Id:    index,
			Name:  t,
			Rank:  index,
			Score: 1000,
		})
	}

	hub = NewHub()

	log.Info("===== Teams =====")
	for k, v := range teams {
		log.Trace("%2d - %s", k, v.Name)
	}

	log.Info("Authorization token: %s\n", token)

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

	log.Info("Listening and serving HTTP on :%s\n", port)
	log.Fatal("Failed to start web server: %v", r.Run(":"+port))
}
