package api

import (
  "fmt"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "health/handlers"
  "github.com/gin-contrib/static"
  "os"
)


func StartApi() {
  r := gin.Default()
  config := cors.DefaultConfig()
  config.AllowHeaders = []string{"Origin", "Authorization", "Content-Length", "Content-Type"}
  config.AllowAllOrigins = true
  config.ExposeHeaders = []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin"}
  config.AllowCredentials = true
  r.Use(cors.New(config))
  Logfile, _ := os.Create("request.log")
  Errorlogfile, _ := os.Create("error.log")

  gin.DefaultWriter = Logfile
  gin.DefaultErrorWriter = Errorlogfile

  r.MaxMultipartMemory = 24000 << 20 // 24 MiB

  r.Use(static.Serve("/upload", static.LocalFile("./upload", true)))

  router := r.Group("/api/v1")

  router.GET("/videos", handlers.GetVideosListHandler)
  router.DELETE("/videos/:id", handlers.DeleteVideoHandler)
  router.PUT("/videos/:id", handlers.UpdVideoHandler)
  router.GET("/videos/:id", handlers.ToggleVideoHandler)
  router.GET("/shedule", handlers.GetSheduleHandler)
  router.POST("/shedule/:day", handlers.AddSheduleHandler)
  router.GET("/shedule/:day", handlers.ListSheduleHandler)
  router.DELETE("/shedule/:id", handlers.DeleteSheduleHandler)
  router.PUT("/shedule", handlers.UpdSheduleHandler)
  router.POST("/videos", handlers.UploadVideo)
  err := r.Run(":8083")
  if err != nil {
    fmt.Println(err)
  }
}
