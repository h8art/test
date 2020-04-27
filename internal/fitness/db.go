package health

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/joho/godotenv"
  "log"
  "os"
)

var Db *gorm.DB

func init() {

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal(err)
  }

  username := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASS")
  dbName := os.Getenv("DB_NAME")
  dbHost := os.Getenv("DB_HOST")

  dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
  fmt.Println(dbURI)

  Conn, err := gorm.Open("postgres", dbURI)
  if err != nil {
    log.Println(err)
  }

  Db = Conn
  Db.Debug().AutoMigrate(&Video{}, &Stat{}, &Shedule{})

}

