package api

import (
	"github.com/gin-gonic/gin"
	"github.com/FabianK1991/civilization/simulation"
	"io/ioutil"
	"net/http"
	"github.com/FabianK1991/civilization/simulation/world"
	"log"
	"fmt"
)

func StartApi (sim *simulation.Simulation) {
	r := gin.Default()

	r.POST("/updateTile" , func(c *gin.Context) {
		var updatedTile world.Tile

		err := c.BindJSON(&updatedTile)

		if(err != nil) {
			log.Fatal(err)
		} else {
			fmt.Println(updatedTile)
			sim.UpdateTile(&updatedTile)
			c.Status(200)
		}
	})

	r.GET("/", func(c *gin.Context) {
		data, _ := ioutil.ReadFile("client/index.html")

		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	r.GET("/getSimulation", func(c *gin.Context) {
		c.String(200, "%s", sim.GetJSON())
	})

	r.GET("/stop", func(c *gin.Context) {
		sim.Stop()
		c.Status(200)
	})

	r.GET("/start", func(c *gin.Context) {
		go sim.Start()
		c.Status(200)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}