package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
	"net/http"
)

var Pattern = [20][20]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

var x = 9
var y = 0
var Position = Pattern[x][y]
var Win = Pattern[19][15]

type Cmd struct {
	Command string `json:"command"`
}

//POST/labirint
func Movement(c *gin.Context) {
	var input Cmd
	var pos models.Coordinates

	//if err := c.ShouldBindJSON(&input); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	if input.Command == "/move_top" {
		if Pattern[x][y-1] == 1 {
			Position = Pattern[x][y-1]
			c.JSON(http.StatusOK, gin.H{"message": "Команда успешно выполнена!"})
		} else if Pattern[x][y-1] <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Прохода нет!"})
		} else if Pattern[x][y-1] == 2 {
			c.JSON(http.StatusOK, gin.H{"message": "Вы нашли выход!"})
		}

	}
	if input.Command == "/move_bot" {
		if Pattern[x][y+1] == 1 {
			Position = Pattern[x][y+1]
			c.JSON(http.StatusOK, gin.H{"message": "Команда успешно выполнена!"})
		} else if Pattern[x][y+1] <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Прохода нет!"})
		} else if Pattern[x][y+1] == 2 {
			c.JSON(http.StatusOK, gin.H{"message": "Вы нашли выход!"})
		}
	}
	if input.Command == "/move_left" {
		if Pattern[x+1][y] == 1 {
			Position = Pattern[x+1][y]
			c.JSON(http.StatusOK, gin.H{"message": "Команда успешно выполнена!"})
		} else if Pattern[x+1][y] <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Прохода нет!"})
		} else if Pattern[x+1][y] == 2 {
			c.JSON(http.StatusOK, gin.H{"message": "Вы нашли выход!"})
		}
	}
	if input.Command == "/move_right" {
		if Pattern[x-1][y] == 1 {
			Position = Pattern[x-1][y]
			c.JSON(http.StatusOK, gin.H{"message": "Команда успешно выполнена!"})
		} else if Pattern[x-1][y] <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Прохода нет!"})
		} else if Pattern[x-1][y] == 2 {
			c.JSON(http.StatusOK, gin.H{"message": "Вы нашли выход!"})
		}
	}
	if input.Command == "/check" {
		c.IndentedJSON(http.StatusOK, pos)
	}
}

