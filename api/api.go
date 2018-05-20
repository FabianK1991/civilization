package api

import (
	"github.com/gin-gonic/gin"
	"github.com/FabianK1991/civilization/simulation"
)

func StartApi (sim *simulation.Simulation) {
	r := gin.Default()

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