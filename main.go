package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Load team file
	teamData, err := ioutil.ReadFile("team.txt")
	if err != nil {
		log.Panicln(err)
	}

	teams := strings.Split(string(teamData), "\n")
	fmt.Println(teams)

	r := gin.Default()
	hub := NewHub()
	go hub.Run()

	r.GET("/websocket", func(c *gin.Context) {
		ServeWebSocket(hub, c)
	})
	r.Run(":12345")
}
