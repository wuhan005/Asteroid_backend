package main

import "github.com/gin-gonic/gin"

const (
	INIT      = "init"
	ATTACK    = "attack"
	RANK      = "rank"
	STATUS    = "status"
	ROUND     = "round"
	EGG       = "easterEgg"
	TIME      = "time"
	CLEAR     = "clear"
	CLEAR_ALL = "clearAll"
)

func attackHandler(c *gin.Context) {
	var attackData attack
	if err := c.BindJSON(&attackData); err != nil {
		c.JSON(makeErrJSON(400, 40000, "payload error"))
		return
	}

	sendMessage(ATTACK, attackData)
	c.JSON(makeSuccessJSON("success"))
}

func rankHandler(c *gin.Context) {
	var rankData rank
	if err := c.BindJSON(&rankData); err != nil {
		c.JSON(makeErrJSON(400, 40000, "payload error"))
		return
	}

	sendMessage(RANK, rankData)
	c.JSON(makeSuccessJSON("success"))
}

func statusHandler(c *gin.Context) {
	var statusData status
	if err := c.BindJSON(&statusData); err != nil {
		c.JSON(makeErrJSON(400, 40000, "payload error"))
		return
	}

	sendMessage(STATUS, statusData)
	c.JSON(makeSuccessJSON("success"))
}

func roundHandler(c *gin.Context) {
	var roundData round
	if err := c.BindJSON(&roundData); err != nil {
		c.JSON(makeErrJSON(400, 40000, "payload error"))
		return
	}

	sendMessage(ROUND, roundData)
	c.JSON(makeSuccessJSON("success"))
}

func eggHandler(c *gin.Context) {
	sendMessage(EGG, nil)
	c.JSON(makeSuccessJSON("success"))
}

func timeHandler(c *gin.Context) {
	var timeData clock
	if err := c.BindJSON(&timeData); err != nil {
		c.JSON(makeErrJSON(400, 40000, "payload error"))
		return
	}

	sendMessage(TIME, timeData)
	c.JSON(makeSuccessJSON("success"))
}

func clearHandler(c *gin.Context) {
	var clearData clearStatus
	if err := c.BindJSON(&clearData); err != nil {
		c.JSON(makeErrJSON(400, 40000, "payload error"))
		return
	}

	sendMessage(CLEAR, clearData)
	c.JSON(makeSuccessJSON("success"))
}

func clearAllHandler(c *gin.Context) {
	sendMessage(CLEAR_ALL, nil)
	c.JSON(makeSuccessJSON("success"))
}
