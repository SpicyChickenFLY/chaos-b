package controller

import (
	"html/template"
	"net/http"

	"github.com/SpicyChickenFLY/chaos-b/service"
	"github.com/gin-gonic/gin"
)

const (
	cnfTemplateFilePath = "static/mycnf.template"
)

// ShowCnfManagerUI show GUI for user
func ShowCnfManagerUI(c *gin.Context) {

}

// GetCnfTemplateFile get Cnf template file
func GetCnfTemplateFile(c *gin.Context) {
	if result, err := service.GetCnfPara(cnfTemplateFilePath); err != nil {
		c.HTML(http.StatusOK, "error.tmpl",
			gin.H{"error": err})
	} else {
		c.HTML(http.StatusOK, "index.tmpl",
			gin.H{"html": template.HTML(result)})
	}
}

// AddNewCnfFile is a
func AddNewCnfFile(c *gin.Context) {

}
