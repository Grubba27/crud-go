package pet

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Controller(r *gin.RouterGroup) {
	service := Service()

	r.GET("/pets/:id", func(c *gin.Context) {
		id := c.Param("id")
		res := service.getById(id)
		c.JSON(http.StatusOK, res)
	})

	r.PUT("/pets/:id", func(c *gin.Context) {
		id := c.Param("id")
		var pet Pet

		if err := c.ShouldBind(&pet); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}

		res := service.updatePet(id, pet)
		c.JSON(http.StatusOK, res)
	})

	r.POST("/pets", func(c *gin.Context) {
		var pet Pet
		if err := c.ShouldBind(&pet); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		res := service.createPet(pet)
		c.JSON(http.StatusOK, res)
	})

	r.DELETE("/pets/:id", func(c *gin.Context) {
		id := c.Param("id")
		res := service.deleteById(id)
		c.JSON(http.StatusOK, res)
	})
}
