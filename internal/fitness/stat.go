package health

import "github.com/jinzhu/gorm"

type Stat struct {
  gorm.Model
  VideoID uint
  Username string
  Type string
}
