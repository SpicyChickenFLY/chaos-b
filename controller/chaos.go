package controller

import (
	"net/http"

	"github.com/SpicyChickenFLY/chaos-b/service"
	"github.com/gin-gonic/gin"
)

// ListChaosTest find all

// CreateChaosTest call chaos service
func CreateChaosTest(c *gin.Context) {
	serverAddr := c.Param("addr")
	cmd := c.Param("cmd")

	resp := service.CreateChaosTest(serverAddr, cmd)

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
