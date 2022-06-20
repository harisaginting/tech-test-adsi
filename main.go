package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"os"
	"fmt"

	"github.com/harisaginting/ginting/pkg/tracer"
	database "github.com/harisaginting/ginting/db"
	router "github.com/harisaginting/ginting/api"
	"github.com/harisaginting/ginting/pkg/log"
	"github.com/harisaginting/ginting/pkg/utils/helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	ctx, stop 	:= signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	tracer.InitTracer()

	// DB CONNECTION
	db := database.Connection()
	database.Migration(db)

	port 	:= os.Getenv("PORT")
	app 	:= gin.New()
	

	// get default url request
	app.UseRawPath 		   = true
	app.UnescapePathValues = true
	// cors configuration
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization", "x-source")
	config.AllowAllOrigins 	= true
	config.AllowMethods 	= []string{"OPTIONS", "PUT", "POST", "GET", "DELETE"}
	app.Use(cors.New(config))

	// error recorvery
	app.Use(gin.CustomRecovery(panicHandler))
	app.Use(otelgin.Middleware("ginting-server"))

	// route
	app.GET("/ping", ping)
	app.NoRoute(lostInSpce)
	// API
	api := app.Group("api")
	router.RestV1(api, db)
	
	// handling server gracefully shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}
	// Initializing the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info(ctx, fmt.Sprintf("listen: %s", port))
		}
	}()
	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Warn(ctx,"shutting down gracefully, press Ctrl+C again to force 🔴")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Warn(ctx,"Server forced to shutdown 🔴: ", err)
	}
	log.Warn(ctx,"Server shutdown 🔴")
}

func lostInSpce(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":        404,
		"data":          nil,
		"error_message": "No Route Found",
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"service_name": os.Getenv("APP_NAME"),
	})
	return
}

// Custom Recovery Panic Error
func panicHandler(c *gin.Context, err interface{}) {
	ctx 	:= c.Request.Context()
	newerr 	:= helper.ForceError(err)
	log.Error(ctx, newerr, "Panic Error 🔴")
    c.JSON(500, gin.H{
    	"status" : 500,
        "error_message":   err,
    })
}