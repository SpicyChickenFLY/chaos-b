package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SpicyChickenFLY/chaos-b/controller"
	"github.com/SpicyChickenFLY/chaos-b/pkgs/middleware"
	"github.com/gin-gonic/gin"

	"github.com/romberli/log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	logFileName = "/tmp/run.log"
)

const ( // GIN CONFIG
	ginPort = ":8080"
)

func main() {
	// Init logger
	_, _, err := log.InitLoggerWithDefaultConfig(logFileName)
	if err != nil {
		fmt.Printf("Init logger failed: %s\n", err.Error())
		panic(err)
	}
	fmt.Println("Init logger succeed")

	// Init router
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Static("/static", "./static")

	// Group: Todo List
	groupAPI := router.Group("/api")
	{
		groupVersion1 := groupAPI.Group("/v1")
		{
			groupMysqlInstaller := groupVersion1.Group("/auto-mysql")
			{
				groupMysqlInstaller.POST("/standard", controller.InstallStandardInstances)
				groupMysqlInstaller.POST("/custom", controller.InstallCustomInstances)
				groupMysqlInstaller.DELETE("/instance", controller.RemoveInstances)
			}
			groupCnfManager := groupVersion1.Group("/auto-mycnf")
			{
				groupCnfManager.GET("/cnf", controller.GetCnfTemplateFile)
				groupCnfManager.POST("/cnf", controller.AddNewCnfFile)
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
			fmt.Println("server encount error while listen and serve:", err)
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
		log.Fatalf("Error occurs when server shutdown: %s", err.Error())
	}
	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		log.Info("timeout of 1 seconds.")
	}
	log.Info("Server exiting")
}
