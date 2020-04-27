package health

import "github.com/jinzhu/gorm"

type Shedule struct {
  gorm.Model
  DayOfWeek uint8
  Time string
  VideoID uint
}

func AddShedule(day uint8) *Shedule{
  shed := Shedule{
    DayOfWeek: day,
  }
  Db.Create(&shed)
  return &shed
}
func ListShedule(day uint8) []Shedule {
  var shedules []Shedule
  Db.Where("day_of_week = ?", day).Find(&shedules)
  return shedules
}

func (s *Shedule) Update() {
  Db.Save(&s)
}
func DeleteShedule(id uint) {
  Db.Where("id = ?", id).Delete(&Shedule{})
}