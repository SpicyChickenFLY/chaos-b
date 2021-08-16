package service

import (
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/installer"
	"github.com/romberli/log"
)

const (
	srcSQLFileDef   = "./static/mysql/mysql.tar.gz"
	srcCnfFileDef   = "./static/conf/my.cnf"
	servInstInfoDef = "root:123@localhost:3306|3307"
	mysqlPwdDef     = "123456"
)

// InstallCustomInstances install instance with custom mycnf file
func InstallCustomInstances(infoStr, srcSQLFile, srcCnfFile, mysqlPwd string) error {

	log.Info("Custom parameters:")
	log.Info(fmt.Sprintf("srcSQLFile: %s", srcSQLFile))
	log.Info(fmt.Sprintf("srcCnfFile: %s", srcCnfFile))
	log.Info(fmt.Sprintf("mysqlPwd: %s", mysqlPwd))
	log.Info(fmt.Sprintf("RunMode: %s", "custom"))
	// installer.InstallCustomInstance(infoStr, srcSQLFile, srcCnfFile)
	return nil
}

// InstallStandardInstances install instances with std mycnf file
func InstallStandardInstances(infoStr, mysqlPwd string) {
	log.Info("Custom parameters:")
	log.Info(fmt.Sprintf("srcSQLFile: %s", srcSQLFileDef))
	log.Info(fmt.Sprintf("srcCnfFile: %s", srcCnfFileDef))
	log.Info(fmt.Sprintf("mysqlPwd: %s", mysqlPwd))
	log.Info(fmt.Sprintf("RunMode: %s", "standard"))
	installer.InstallStandardMultiInstanceOnMultiServer(srcSQLFileDef, infoStr, mysqlPwd)
}
