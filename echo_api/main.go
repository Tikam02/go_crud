package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Player struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Online   bool   `json:"online"`
}

type Players []Player

var players = Players{Player{generateId(), "DarthVader", true}, Player{generateId(), "Batman", true}}

func generateId() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10000)
}

func getPlayers(c echo.Context) error {
	return c.JSON(http.StatusOK, players)
}

func postPlayer(c echo.Context) error {
	player := Player{}
	err := c.Bind(&player)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	player.Id = generateId()
	players = append(players, player)
	return c.JSON(http.StatusCreated, players)
}

func getPlayer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, player := range players {
		if player.Id == id {
			return c.JSON(http.StatusOK, player)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func putPlayer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range players {
		if players[i].Id == id {
			players[i].Online = false
			return c.JSON(http.StatusOK, players)
		}
	}
	return c.JSON(http.StatusOK, players)
}

func deletePlayer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range players {
		if players[i].Id == id {
			players = append(players[:i], players[i+1:]...)
			return c.JSON(http.StatusOK, players)
		}
	}
	return c.JSON(http.StatusOK, players)
}

func main() {

	fmt.Println("Running Server....")
	e := echo.New()
	e.GET("/players", getPlayers)
	e.POST("/players", postPlayer)
	e.GET("/players/:id", getPlayer)
	e.PUT("/players/:id", putPlayer)
	e.DELETE("/players/:id", deletePlayer)
	e.Start(":9000")

}
