package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Pattern = [20][20]int{{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

var x = 0
var y = 0
var pos Coordinates

type Cmd struct {
	Command string `json:"command"`
}
type Coordinates struct {
	Top   int `json:"top"`
	Bot   int `json:"bot"`
	Left  int `json:"left"`
	Right int `json:"right"`
}

func Movement(c *gin.Context) {

	var input Cmd

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Command == "/move_top" {
		if Pattern[x][y-1] == 1 {
			y -= 1
			c.JSON(http.StatusOK, gin.H{"message": "Команда успешно выполнена!"})
		} else if Pattern[x][y-1] <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Прохода нет!"})
		} else if Pattern[x][y-1] == 2 {
			c.JSON(http.StatusOK, gin.H{"message": "Вы нашли выход!"})
		}

	} else if input.Command == "/move_bot" {
		if Pattern[x][y+1] == 1 {
			y += 1
			c.JSON(http.StatusOK, gin.H{"message": "Команда успешно выполнена!"})
		} else if Pattern[x][y+1] <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Прохода нет!"})
		} else if Pattern[x][y+1] == 2 {
			c.JSON(http.StatusOK, gin.H{"message": "Вы нашли выход!"})
		}
	} else if input.Command == "/move_left" {
		if Pattern[x-1][y] == 1 {
			x -= 1
			c.JSON(http.StatusOK, gin.H{"message": "Команда успешно выполнена!"})
		} else if Pattern[x+1][y] <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Прохода нет!"})
		} else if Pattern[x+1][y] == 2 {
			c.JSON(http.StatusOK, gin.H{"message": "Вы нашли выход!"})
		}
	} else if input.Command == "/move_right" {
		if Pattern[x+1][y] == 1 {
			x += 1
			c.JSON(http.StatusOK, gin.H{"message": "Команда успешно выполнена!"})
		} else if Pattern[x-1][y] <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Прохода нет!"})
		} else if Pattern[x-1][y] == 2 {
			c.JSON(http.StatusOK, gin.H{"message": "Вы нашли выход!"})
		}
	} else if input.Command == "/check" {

		pos.Top = Pattern[x][y-1]
		pos.Bot = Pattern[x][y+1]
		pos.Left = Pattern[x-1][y]
		pos.Right = Pattern[x+1][y]

		c.IndentedJSON(http.StatusOK, pos)

	} else if input.Command == "/init" {

		x = 0
		y = 0

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "400"})
	}

}
