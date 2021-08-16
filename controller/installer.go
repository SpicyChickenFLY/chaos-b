package controller

import (
	"net/http"

	"github.com/SpicyChickenFLY/chaos-b/service"
	"github.com/gin-gonic/gin"
)

// ShowMysqlInstallerUI show GUI for user
func ShowMysqlInstallerUI(c *gin.Context) {
	c.HTML(http.StatusOK, "installer.tmpl", gin.H{})
}

// InstallStandardInstances install instances with std mycnf file
func InstallStandardInstances(c *gin.Context) {
	// service.InstallStandardInstances()
}

// InstallCustomInstances install instance with custom mycnf file
func InstallCustomInstances(c *gin.Context) {
	data := &struct {
		InfoStr    string `json:"InfoStr"`
		SrcSQLFile string `json:"SrcSQLFile"`
		SrcCnfFile string `json:"SrcCnfFile"`
		MysqlPwd   string `json:"MysqlPwd"`
	}{}
	c.BindJSON(&data)
	service.InstallStandardInstances(data.InfoStr, data.MysqlPwd)
	c.JSON(http.StatusOK, gin.H{})
}

// RemoveInstances remove instance
func RemoveInstances(c *gin.Context) {
	// param := c.Param("param")
	c.String(http.StatusOK, "not supported now!")
}
