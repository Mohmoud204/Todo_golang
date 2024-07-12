package controller

import(
   "github.com/gin-gonic/gin"
  "net/http"
   "example.com/todo/models"
    "gorm.io/gorm"
   
  )
  
  var db *gorm.DB 
  func SetDB(database *gorm.DB) {
    db = database
}
  func CreateTodo(c *gin.Context) {
    var todo models.Todo 
    
    if err := c.BindJSON(&todo); err != nil {
    
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Create(&todo).Error; err != nil {
      
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, todo)
  }
func GetAllTodos(c *gin.Context){
  var todos []models.Todo
  
    if err := db.Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, todos)
} 

func DeleteTodo(c *gin.Context){
  id := c.Param("id")
  var todo models.Todo 
  db.First(&todo,id) 
  if todo.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"message":"id not found"})
        return
  }
  db.Unscoped().Delete(&todo) 
  c.JSON(http.StatusOK, gin.H{"message":"delete success"})
}

func ShowTodoById (c *gin.Context){
  id := c.Param("id")
  var todo models.Todo 
  db.First(&todo,id) 
  if todo.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"message":"id not found"})
        return
  }
  c.JSON(http.StatusOK,todo)
  return
}



func UpdateTodo(c *gin.Context) {
  type NewTodo struct {
    Title  string ` json:"title"`
    Done   bool   `json:"done"`
  }
  
  id := c.Param("id")
  var todo models.Todo 
  db.First(&todo,id) 
  if todo.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"message":"id not found"})
        return
  } 
  
  
  var newTodo NewTodo
  if err := c.BindJSON(&newTodo); err != nil {
    
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
  if newTodo.Title != "" {
    todo.Title = newTodo.Title
  }
  if newTodo.Done != todo.Done  {
    todo.Done = newTodo.Done
  }

  db.Save(&todo)
    c.JSON(http.StatusOK, todo)
  return
}