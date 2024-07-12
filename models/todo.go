package models  
import (
  "gorm.io/gorm"

  )
type Todo struct {
    gorm.Model
	 Title  string `gorm:"not null" json:"title" binding:"required"`
    Done   bool   `gorm:"default:false" json:"done"`
}


func AutoMigrate(db *gorm.DB) {
    db.AutoMigrate(&Todo{})
}