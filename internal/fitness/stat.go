package health

import "github.com/jinzhu/gorm"

type Stat struct {
  gorm.Model
  VideoID uint
  Username string
  Type string
}

func AddStat(VID uint, username string, viewType string) Stat{
  st := Stat{
    VideoID:  VID,
    Username: username,
    Type:     viewType,
  }
  Db.Create(&st)
  return st
}

func ListStats() []Stat {
  var stats []Stat
  Db.Find(&stats)
  return stats
}

