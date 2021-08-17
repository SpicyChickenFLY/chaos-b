package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/SpicyChickenFLY/chaos-b/controller"
	"github.com/SpicyChickenFLY/chaos-b/pkgs/middleware"
	"github.com/gin-gonic/gin"
	"spicychicken.top/chaos-b/pkgs/mysql"

	"github.com/romberli/log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

const (
	defaultLogFileRelPath  = "log/never-todo.log"
	defaultConfFileRelPath = "static/init.cfg"
)

const ( // GIN CONFIG
	ginPort = ":8080"
)

func main() {
	// get cwd
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("get current Directory failed: %s\n", err.Error())
		os.Exit(1)
	}

	logPath := path.Join(currDir, defaultLogFileRelPath)
	confPath := path.Join(currDir, defaultConfFileRelPath)

	sysType := runtime.GOOS
	switch sysType {
	case "linux":
		logPath = strings.ReplaceAll(logPath, "\\", "/")
		confPath = strings.ReplaceAll(confPath, "\\", "/")
	case "windows":
		logPath = strings.ReplaceAll(logPath, "/", "\\")
		confPath = strings.ReplaceAll(confPath, "/", "\\")
	}

	// Init logger
	if _, _, err := log.InitFileLoggerWithDefault(logPath); err != nil {
		fmt.Printf("Init logger failed: %s\n", err.Error())
		os.Exit(1)
	}
	log.Info("=============================")
	log.Info("Program Started")

	cfg, err := ini.Load(confPath)
	if err != nil {
		log.Error(err.Error())
		log.Info("Program Terminated")
		log.Info("=============================")
		os.Exit(1)
	}
	// get mysql root@localhost password
	dbType := cfg.Section("db").Key("type").String()
	if dbType == "mysql" {
		serverHost := cfg.Section("db").Key("server_host").String()
		serverPort := cfg.Section("db").Key("server_port").String()
		userName := cfg.Section("db").Key("user_name").String()
		userPwd := cfg.Section("db").Key("user_pwd").String()
		dbName := cfg.Section("db").Key("db_name").String()
		dbCharset := cfg.Section("db").Key("db_charset").String()
		// Initialize MySQL connection
		if err := mysql.CreateGormConn(
			userName, userPwd,
			serverHost, serverPort,
			dbName, dbCharset); err != nil {
			handleError(err)
		}
		log.Info("mysql initialization compelete")
	}

	// Init router
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Static("/static", "./static")

	// Group: Todo List
	groupAPI := router.Group("/api")
	{
		groupVersion1 := groupAPI.Group("/v1")
		{
			groupChaos := groupVersion1.Group("/auto-mysql")
			{
				groupChaos.GET("/test/:id", controller.ListChaosTest)
				groupChaos.GET("/test/:id", controller.FindChaosTestByID)
				groupChaos.POST("/test/:id", controller.CreateChaosTest)
				groupChaos.DELETE("/test/:id", controller.DeleteChaosTest)
			}
			groupInstance := groupVersion1.Group("/auto-mycnf")
			{
				groupInstance.GET("/cnf", controller.GetCnfTemplateFile)
				groupInstance.POST("/cnf", controller.AddNewCnfFile)
			}
		}
	}

	server := &http.Server{
		Addr:    ginPort,
		Handler: router,
	}

	go func() {
		// service connections
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			handleError(err)
		}
	}()

	// Wait for interrupt signal to shutdown  with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (default SIGTERM) (-2 SIGINT) (-9 SIGKILL cant be catched)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		handleError(err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		log.Info("timeout of 1 seconds.")
	}
	log.Info("Server exiting")
}

func handleError(err error) {
	log.Error(err.Error())
	log.Info("Program Terminated")
	log.Info("=============================")
	os.Exit(1)
}
