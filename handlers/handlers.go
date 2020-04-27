package handlers

import (
  "github.com/gin-gonic/gin"
  health "health/internal/fitness"
  "path/filepath"
  "strconv"
)

func GetVideosListHandler(c *gin.Context) {
  videos := health.GetVideosList()
  c.JSON(200, videos)
}
func GetSheduleHandler(c *gin.Context) {

}
func UploadVideo(c *gin.Context) {
  file, err := c.FormFile("file")
  if err != nil {
    c.JSON(500, err)
    return
  }
  vid, err := health.AddVideo()
  if err != nil {
    c.JSON(500, err)
    return
  }
  if err := c.SaveUploadedFile(file, "upload/" + strconv.Itoa(int(vid.ID)) + filepath.Ext(file.Filename)); err != nil {
    c.JSON(500, err)
    return
  }

}

func DeleteVideoHandler(c *gin.Context) {
  idStr := c.Param("id")
  id, err := strconv.Atoi(idStr)
  if err != nil {
    c.JSON(500, err)
    return
  }
  health.DeleteVideo(uint(id))
}

func AddSheduleHandler(c *gin.Context) {
  dayStr := c.Param("day")
  day, err := strconv.Atoi(dayStr)
  if err != nil {
    c.JSON(500, err)
    return
  }
  shed := health.AddShedule(uint8(day))
  c.JSON(200, shed)
}

func ListSheduleHandler(c *gin.Context) {
  dayStr := c.Param("day")
  day, err := strconv.Atoi(dayStr)
  if err != nil {
    c.JSON(500, err)
    return
  }
  c.JSON(200, health.ListShedule(uint8(day)))
}

func UpdSheduleHandler(c *gin.Context) {
  var shed health.Shedule
  err := c.BindJSON(&shed)
  if err != nil {
    c.JSON(500, err)
    return
  }
  shed.Update()
}

func DeleteSheduleHandler(c *gin.Context) {
  idStr := c.Param("id")
  id, err := strconv.Atoi(idStr)
  if err != nil {
    c.JSON(500, err)
    return
  }
  health.DeleteShedule(uint(id))
}
func ToggleVideoHandler(c *gin.Context) {
  idStr := c.Param("id")
  id, err := strconv.Atoi(idStr)
  if err != nil {
    c.JSON(500, err)
    return
  }
  err = health.ToggleVid(uint(id))
  if err != nil {
    c.JSON(500, err)
    return
  }
}
func UpdVideoHandler(c *gin.Context)  {
  idStr := c.Param("id")
  id, err := strconv.Atoi(idStr)
  if err != nil {
    c.JSON(500, err)
    return
  }
  var req struct{
    Name string
  }
  err = c.BindJSON(&req)
  if err != nil {
    c.JSON(500, err)
    return
  }
  vid, err := health.UpdVid(uint(id), req.Name)
  if err != nil {
    c.JSON(500, err)
    return
  }
  c.JSON(200, vid)
}

func AddStatHandler(c *gin.Context) {
  var req struct{
    VID uint
    Name string
    Type string
  }
  err := c.BindJSON(&req)
  if err != nil {
    c.JSON(400, err)
    return
  }

  health.AddStat(req.VID, req.Name, req.Type)
}

func ListStatsHandler(c *gin.Context) {
  c.JSON(200, health.ListStats())
}