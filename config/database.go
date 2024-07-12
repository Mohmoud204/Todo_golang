package config 
import (
//   "example.com/todo/models"
  "log"
  "gorm.io/driver/mysql"
 "gorm.io/gorm"
   "gorm.io/gorm/logger"
   "os"
   "time"
"github.com/joho/godotenv"
  )

func ConnectDatabase() *gorm.DB {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second,   // Slow SQL threshold
            LogLevel:      logger.Info,   // Log level
            Colorful:      true,          // Disable color
        },
    )
   dsn := os.Getenv("DATA_BASE_UPLOAD")
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
      Logger: newLogger,
    })
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    return db
}