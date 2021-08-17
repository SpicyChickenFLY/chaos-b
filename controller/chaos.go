package controller

import (
	"net/http"

	"github.com/SpicyChickenFLY/chaos-b/service"
	"github.com/gin-gonic/gin"
)

// StartChaosTest call chaos service
func StartChaosTest(c *gin.Context) {
	serverAddr := c.Param("addr")
	cmd := c.Param("cmd")
	resp := service.StartChaosTest(serverAddr, cmd)
	c.JSON(http.StatusOK, gin.H{"data": resp})
}
